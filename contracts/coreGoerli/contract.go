// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coreGoerli

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

// CoreGoerliMetaData contains all meta data concerning the CoreGoerli contract.
var CoreGoerliMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValueAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotInSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagAllowAllSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"FeatureFlagDeniersReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagDenyAllSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getDeniers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagDenyAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeatureAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"setDeniers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagAllowAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagDenyAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"FeatureUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"InvalidAccountId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"InvalidPermission\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"}],\"name\":\"OnlyAccountTokenProxy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"PermissionNotGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionOutOfBounds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedAccountId\",\"type\":\"uint128\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountLastInteraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountPermissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"permissions\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIAccountModule.AccountPermissions[]\",\"name\":\"accountPerms\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"notifyAccountTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"renouncePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revokePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyDistribution\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateralRatio\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"MarketNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"NotFundedByPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint128ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint128\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"updatedDebt\",\"type\":\"int256\"}],\"name\":\"DebtAssociated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"associateDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InvalidMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotCcipRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNetwork\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCcipClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCcipClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredTime\",\"type\":\"uint256\"}],\"name\":\"AccountActivityTimeoutPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"CollateralDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InsufficientAccountCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existing\",\"type\":\"uint256\"}],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parameter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"PrecisionLost\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"CollateralLockCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"CollateralLockExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"cleanExpiredLocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cleared\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"createLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getAccountAvailableCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getAccountCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDeposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAssigned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalLocked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getLocks\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amountD18\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"lockExpirationTime\",\"type\":\"uint64\"}],\"internalType\":\"structCollateralLock.Data[]\",\"name\":\"locks\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"CollateralConfigured\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"configureCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getCollateralConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"hideDisabled\",\"type\":\"bool\"}],\"name\":\"getCollateralConfigurations\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getCollateralPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientCcipFee\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TransferCrossChainInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCrossChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gasTokenUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentDebt\",\"type\":\"int256\"}],\"name\":\"InsufficientDebt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"PoolNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"IssuanceFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UsdBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UsdMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnUsd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintUsd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CannotScaleEmptyMapping\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"currentCRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cratio\",\"type\":\"uint256\"}],\"name\":\"IneligibleForLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientMappedAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeVaultLiquidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt128ToUint128\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Liquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"VaultLiquidation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"isPositionLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"isVaultLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"}],\"name\":\"liquidate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxUsd\",\"type\":\"uint256\"}],\"name\":\"liquidateVault\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountToDeposit\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarketCollateralDepositable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarketCollateralWithdrawable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MarketCollateralDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MarketCollateralWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"systemAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MaximumMarketCollateralConfigured\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"configureMaximumMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"depositMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getMarketCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralAmountD18\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getMaximumMarketCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"IncorrectMarketInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NotEnoughLiquidity\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MarketRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"MarketSystemFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketUsdDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketUsdWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"SetMarketMinLiquidityRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"minDelegateTime\",\"type\":\"uint32\"}],\"name\":\"SetMinDelegateTime\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxIter\",\"type\":\"uint256\"}],\"name\":\"distributeDebtToPools\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketDebtPerShare\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getMarketFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFeeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketMinDelegateTime\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketNetIssuance\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getMarketPoolDebtDistribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesD18\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"totalSharesD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"valuePerShareD27\",\"type\":\"int128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketPools\",\"outputs\":[{\"internalType\":\"uint128[]\",\"name\":\"inRangePoolIds\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"outRangePoolIds\",\"type\":\"uint128[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketReportedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketTotalDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMinLiquidityRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleManager\",\"outputs\":[{\"internalType\":\"contractIOracleManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsdToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getWithdrawableMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"isMarketCapacityLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"registerMarket\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"minDelegateTime\",\"type\":\"uint32\"}],\"name\":\"setMarketMinDelegateTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"setMinLiquidityRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RecursiveMulticall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PoolApprovedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PoolApprovedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PreferredPoolSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"addApprovedPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApprovedPools\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreferredPool\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"removeApprovedPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"setPreferredPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"CapacityLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"timeRemaining\",\"type\":\"uint32\"}],\"name\":\"MinDelegationTimeoutPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"PoolAlreadyExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"indexed\":false,\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"markets\",\"type\":\"tuple[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolConfigurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolNameUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolNominationRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolNominationRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolOwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolOwnershipAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"SetMinLiquidityRatio\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"acceptPoolOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedPoolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinLiquidityRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getNominatedPoolOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"poolName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"nominatePoolOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"optionalCollateralType\",\"type\":\"address\"}],\"name\":\"rebalancePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"renouncePoolNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"revokePoolNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"setMinLiquidityRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"newMarketConfigurations\",\"type\":\"tuple[]\"}],\"name\":\"setPoolConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setPoolName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint32\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint32ToInt32\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint64ToInt64\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardDistributorNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardUnavailable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"RewardsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardsDistributorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardsDistributorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"}],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"getRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"registerRewardsDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"removeRewardsDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"updateRewards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newChainId\",\"type\":\"uint64\"}],\"name\":\"NewSupportedCrossChainNetwork\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ccipRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ccipTokenPool\",\"type\":\"address\"}],\"name\":\"configureChainlinkCrossChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleManagerAddress\",\"type\":\"address\"}],\"name\":\"configureOracleManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"v\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfigAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfigUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"v\",\"type\":\"bytes32\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"supportedNetworks\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"ccipSelectors\",\"type\":\"uint64[]\"}],\"name\":\"setSupportedCrossChainNetworks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numRegistered\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minDelegation\",\"type\":\"uint256\"}],\"name\":\"InsufficientDelegation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCollateralAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"}],\"name\":\"InvalidLeverage\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"DelegationUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralAmountD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"}],\"name\":\"delegateCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"collateralizationRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionCollateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultCollateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CoreGoerliABI is the input ABI used to generate the binding from.
// Deprecated: Use CoreGoerliMetaData.ABI instead.
var CoreGoerliABI = CoreGoerliMetaData.ABI

// CoreGoerli is an auto generated Go binding around an Ethereum contract.
type CoreGoerli struct {
	CoreGoerliCaller     // Read-only binding to the contract
	CoreGoerliTransactor // Write-only binding to the contract
	CoreGoerliFilterer   // Log filterer for contract events
}

// CoreGoerliCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoreGoerliCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreGoerliTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoreGoerliTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreGoerliFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoreGoerliFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreGoerliSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoreGoerliSession struct {
	Contract     *CoreGoerli       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreGoerliCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoreGoerliCallerSession struct {
	Contract *CoreGoerliCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CoreGoerliTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoreGoerliTransactorSession struct {
	Contract     *CoreGoerliTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CoreGoerliRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoreGoerliRaw struct {
	Contract *CoreGoerli // Generic contract binding to access the raw methods on
}

// CoreGoerliCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoreGoerliCallerRaw struct {
	Contract *CoreGoerliCaller // Generic read-only contract binding to access the raw methods on
}

// CoreGoerliTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoreGoerliTransactorRaw struct {
	Contract *CoreGoerliTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoreGoerli creates a new instance of CoreGoerli, bound to a specific deployed contract.
func NewCoreGoerli(address common.Address, backend bind.ContractBackend) (*CoreGoerli, error) {
	contract, err := bindCoreGoerli(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoreGoerli{CoreGoerliCaller: CoreGoerliCaller{contract: contract}, CoreGoerliTransactor: CoreGoerliTransactor{contract: contract}, CoreGoerliFilterer: CoreGoerliFilterer{contract: contract}}, nil
}

// NewCoreGoerliCaller creates a new read-only instance of CoreGoerli, bound to a specific deployed contract.
func NewCoreGoerliCaller(address common.Address, caller bind.ContractCaller) (*CoreGoerliCaller, error) {
	contract, err := bindCoreGoerli(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliCaller{contract: contract}, nil
}

// NewCoreGoerliTransactor creates a new write-only instance of CoreGoerli, bound to a specific deployed contract.
func NewCoreGoerliTransactor(address common.Address, transactor bind.ContractTransactor) (*CoreGoerliTransactor, error) {
	contract, err := bindCoreGoerli(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliTransactor{contract: contract}, nil
}

// NewCoreGoerliFilterer creates a new log filterer instance of CoreGoerli, bound to a specific deployed contract.
func NewCoreGoerliFilterer(address common.Address, filterer bind.ContractFilterer) (*CoreGoerliFilterer, error) {
	contract, err := bindCoreGoerli(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliFilterer{contract: contract}, nil
}

// bindCoreGoerli binds a generic wrapper to an already deployed contract.
func bindCoreGoerli(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoreGoerliMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoreGoerli *CoreGoerliRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoreGoerli.Contract.CoreGoerliCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoreGoerli *CoreGoerliRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CoreGoerliTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoreGoerli *CoreGoerliRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CoreGoerliTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoreGoerli *CoreGoerliCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoreGoerli.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoreGoerli *CoreGoerliTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreGoerli.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoreGoerli *CoreGoerliTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoreGoerli.Contract.contract.Transact(opts, method, params...)
}

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_CoreGoerli *CoreGoerliCaller) GetAccountAvailableCollateral(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getAccountAvailableCollateral", accountId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetAccountAvailableCollateral(accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreGoerli.Contract.GetAccountAvailableCollateral(&_CoreGoerli.CallOpts, accountId, collateralType)
}

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_CoreGoerli *CoreGoerliCallerSession) GetAccountAvailableCollateral(accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreGoerli.Contract.GetAccountAvailableCollateral(&_CoreGoerli.CallOpts, accountId, collateralType)
}

// GetAccountCollateral is a free data retrieval call binding the contract method 0xef45148e.
//
// Solidity: function getAccountCollateral(uint128 accountId, address collateralType) view returns(uint256 totalDeposited, uint256 totalAssigned, uint256 totalLocked)
func (_CoreGoerli *CoreGoerliCaller) GetAccountCollateral(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getAccountCollateral", accountId, collateralType)

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
func (_CoreGoerli *CoreGoerliSession) GetAccountCollateral(accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	return _CoreGoerli.Contract.GetAccountCollateral(&_CoreGoerli.CallOpts, accountId, collateralType)
}

// GetAccountCollateral is a free data retrieval call binding the contract method 0xef45148e.
//
// Solidity: function getAccountCollateral(uint128 accountId, address collateralType) view returns(uint256 totalDeposited, uint256 totalAssigned, uint256 totalLocked)
func (_CoreGoerli *CoreGoerliCallerSession) GetAccountCollateral(accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	return _CoreGoerli.Contract.GetAccountCollateral(&_CoreGoerli.CallOpts, accountId, collateralType)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCaller) GetAccountLastInteraction(opts *bind.CallOpts, accountId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getAccountLastInteraction", accountId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetAccountLastInteraction(&_CoreGoerli.CallOpts, accountId)
}

// GetAccountLastInteraction is a free data retrieval call binding the contract method 0x1b5dccdb.
//
// Solidity: function getAccountLastInteraction(uint128 accountId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCallerSession) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetAccountLastInteraction(&_CoreGoerli.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_CoreGoerli *CoreGoerliCaller) GetAccountOwner(opts *bind.CallOpts, accountId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getAccountOwner", accountId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_CoreGoerli *CoreGoerliSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _CoreGoerli.Contract.GetAccountOwner(&_CoreGoerli.CallOpts, accountId)
}

// GetAccountOwner is a free data retrieval call binding the contract method 0xbf60c31d.
//
// Solidity: function getAccountOwner(uint128 accountId) view returns(address)
func (_CoreGoerli *CoreGoerliCallerSession) GetAccountOwner(accountId *big.Int) (common.Address, error) {
	return _CoreGoerli.Contract.GetAccountOwner(&_CoreGoerli.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_CoreGoerli *CoreGoerliCaller) GetAccountPermissions(opts *bind.CallOpts, accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getAccountPermissions", accountId)

	if err != nil {
		return *new([]IAccountModuleAccountPermissions), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAccountModuleAccountPermissions)).(*[]IAccountModuleAccountPermissions)

	return out0, err

}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_CoreGoerli *CoreGoerliSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _CoreGoerli.Contract.GetAccountPermissions(&_CoreGoerli.CallOpts, accountId)
}

// GetAccountPermissions is a free data retrieval call binding the contract method 0xa796fecd.
//
// Solidity: function getAccountPermissions(uint128 accountId) view returns((address,bytes32[])[] accountPerms)
func (_CoreGoerli *CoreGoerliCallerSession) GetAccountPermissions(accountId *big.Int) ([]IAccountModuleAccountPermissions, error) {
	return _CoreGoerli.Contract.GetAccountPermissions(&_CoreGoerli.CallOpts, accountId)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_CoreGoerli *CoreGoerliCaller) GetAccountTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getAccountTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_CoreGoerli *CoreGoerliSession) GetAccountTokenAddress() (common.Address, error) {
	return _CoreGoerli.Contract.GetAccountTokenAddress(&_CoreGoerli.CallOpts)
}

// GetAccountTokenAddress is a free data retrieval call binding the contract method 0xa148bf10.
//
// Solidity: function getAccountTokenAddress() view returns(address)
func (_CoreGoerli *CoreGoerliCallerSession) GetAccountTokenAddress() (common.Address, error) {
	return _CoreGoerli.Contract.GetAccountTokenAddress(&_CoreGoerli.CallOpts)
}

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_CoreGoerli *CoreGoerliCaller) GetApprovedPools(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getApprovedPools")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_CoreGoerli *CoreGoerliSession) GetApprovedPools() ([]*big.Int, error) {
	return _CoreGoerli.Contract.GetApprovedPools(&_CoreGoerli.CallOpts)
}

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_CoreGoerli *CoreGoerliCallerSession) GetApprovedPools() ([]*big.Int, error) {
	return _CoreGoerli.Contract.GetApprovedPools(&_CoreGoerli.CallOpts)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_CoreGoerli *CoreGoerliCaller) GetAssociatedSystem(opts *bind.CallOpts, id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getAssociatedSystem", id)

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
func (_CoreGoerli *CoreGoerliSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _CoreGoerli.Contract.GetAssociatedSystem(&_CoreGoerli.CallOpts, id)
}

// GetAssociatedSystem is a free data retrieval call binding the contract method 0x60988e09.
//
// Solidity: function getAssociatedSystem(bytes32 id) view returns(address addr, bytes32 kind)
func (_CoreGoerli *CoreGoerliCallerSession) GetAssociatedSystem(id [32]byte) (struct {
	Addr common.Address
	Kind [32]byte
}, error) {
	return _CoreGoerli.Contract.GetAssociatedSystem(&_CoreGoerli.CallOpts, id)
}

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_CoreGoerli *CoreGoerliCaller) GetCollateralConfiguration(opts *bind.CallOpts, collateralType common.Address) (CollateralConfigurationData, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getCollateralConfiguration", collateralType)

	if err != nil {
		return *new(CollateralConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new(CollateralConfigurationData)).(*CollateralConfigurationData)

	return out0, err

}

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_CoreGoerli *CoreGoerliSession) GetCollateralConfiguration(collateralType common.Address) (CollateralConfigurationData, error) {
	return _CoreGoerli.Contract.GetCollateralConfiguration(&_CoreGoerli.CallOpts, collateralType)
}

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_CoreGoerli *CoreGoerliCallerSession) GetCollateralConfiguration(collateralType common.Address) (CollateralConfigurationData, error) {
	return _CoreGoerli.Contract.GetCollateralConfiguration(&_CoreGoerli.CallOpts, collateralType)
}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_CoreGoerli *CoreGoerliCaller) GetCollateralConfigurations(opts *bind.CallOpts, hideDisabled bool) ([]CollateralConfigurationData, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getCollateralConfigurations", hideDisabled)

	if err != nil {
		return *new([]CollateralConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CollateralConfigurationData)).(*[]CollateralConfigurationData)

	return out0, err

}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_CoreGoerli *CoreGoerliSession) GetCollateralConfigurations(hideDisabled bool) ([]CollateralConfigurationData, error) {
	return _CoreGoerli.Contract.GetCollateralConfigurations(&_CoreGoerli.CallOpts, hideDisabled)
}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_CoreGoerli *CoreGoerliCallerSession) GetCollateralConfigurations(hideDisabled bool) ([]CollateralConfigurationData, error) {
	return _CoreGoerli.Contract.GetCollateralConfigurations(&_CoreGoerli.CallOpts, hideDisabled)
}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_CoreGoerli *CoreGoerliCaller) GetCollateralPrice(opts *bind.CallOpts, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getCollateralPrice", collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetCollateralPrice(collateralType common.Address) (*big.Int, error) {
	return _CoreGoerli.Contract.GetCollateralPrice(&_CoreGoerli.CallOpts, collateralType)
}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_CoreGoerli *CoreGoerliCallerSession) GetCollateralPrice(collateralType common.Address) (*big.Int, error) {
	return _CoreGoerli.Contract.GetCollateralPrice(&_CoreGoerli.CallOpts, collateralType)
}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_CoreGoerli *CoreGoerliCaller) GetConfig(opts *bind.CallOpts, k [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getConfig", k)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_CoreGoerli *CoreGoerliSession) GetConfig(k [32]byte) ([32]byte, error) {
	return _CoreGoerli.Contract.GetConfig(&_CoreGoerli.CallOpts, k)
}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_CoreGoerli *CoreGoerliCallerSession) GetConfig(k [32]byte) ([32]byte, error) {
	return _CoreGoerli.Contract.GetConfig(&_CoreGoerli.CallOpts, k)
}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_CoreGoerli *CoreGoerliCaller) GetConfigAddress(opts *bind.CallOpts, k [32]byte) (common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getConfigAddress", k)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_CoreGoerli *CoreGoerliSession) GetConfigAddress(k [32]byte) (common.Address, error) {
	return _CoreGoerli.Contract.GetConfigAddress(&_CoreGoerli.CallOpts, k)
}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_CoreGoerli *CoreGoerliCallerSession) GetConfigAddress(k [32]byte) (common.Address, error) {
	return _CoreGoerli.Contract.GetConfigAddress(&_CoreGoerli.CallOpts, k)
}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_CoreGoerli *CoreGoerliCaller) GetConfigUint(opts *bind.CallOpts, k [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getConfigUint", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_CoreGoerli *CoreGoerliSession) GetConfigUint(k [32]byte) (*big.Int, error) {
	return _CoreGoerli.Contract.GetConfigUint(&_CoreGoerli.CallOpts, k)
}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_CoreGoerli *CoreGoerliCallerSession) GetConfigUint(k [32]byte) (*big.Int, error) {
	return _CoreGoerli.Contract.GetConfigUint(&_CoreGoerli.CallOpts, k)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_CoreGoerli *CoreGoerliCaller) GetDeniers(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getDeniers", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_CoreGoerli *CoreGoerliSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _CoreGoerli.Contract.GetDeniers(&_CoreGoerli.CallOpts, feature)
}

// GetDeniers is a free data retrieval call binding the contract method 0xed429cf7.
//
// Solidity: function getDeniers(bytes32 feature) view returns(address[])
func (_CoreGoerli *CoreGoerliCallerSession) GetDeniers(feature [32]byte) ([]common.Address, error) {
	return _CoreGoerli.Contract.GetDeniers(&_CoreGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_CoreGoerli *CoreGoerliCaller) GetFeatureFlagAllowAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getFeatureFlagAllowAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_CoreGoerli *CoreGoerliSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _CoreGoerli.Contract.GetFeatureFlagAllowAll(&_CoreGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowAll is a free data retrieval call binding the contract method 0x40a399ef.
//
// Solidity: function getFeatureFlagAllowAll(bytes32 feature) view returns(bool)
func (_CoreGoerli *CoreGoerliCallerSession) GetFeatureFlagAllowAll(feature [32]byte) (bool, error) {
	return _CoreGoerli.Contract.GetFeatureFlagAllowAll(&_CoreGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_CoreGoerli *CoreGoerliCaller) GetFeatureFlagAllowlist(opts *bind.CallOpts, feature [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getFeatureFlagAllowlist", feature)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_CoreGoerli *CoreGoerliSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _CoreGoerli.Contract.GetFeatureFlagAllowlist(&_CoreGoerli.CallOpts, feature)
}

// GetFeatureFlagAllowlist is a free data retrieval call binding the contract method 0xe12c8160.
//
// Solidity: function getFeatureFlagAllowlist(bytes32 feature) view returns(address[])
func (_CoreGoerli *CoreGoerliCallerSession) GetFeatureFlagAllowlist(feature [32]byte) ([]common.Address, error) {
	return _CoreGoerli.Contract.GetFeatureFlagAllowlist(&_CoreGoerli.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_CoreGoerli *CoreGoerliCaller) GetFeatureFlagDenyAll(opts *bind.CallOpts, feature [32]byte) (bool, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getFeatureFlagDenyAll", feature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_CoreGoerli *CoreGoerliSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _CoreGoerli.Contract.GetFeatureFlagDenyAll(&_CoreGoerli.CallOpts, feature)
}

// GetFeatureFlagDenyAll is a free data retrieval call binding the contract method 0xbcae3ea0.
//
// Solidity: function getFeatureFlagDenyAll(bytes32 feature) view returns(bool)
func (_CoreGoerli *CoreGoerliCallerSession) GetFeatureFlagDenyAll(feature [32]byte) (bool, error) {
	return _CoreGoerli.Contract.GetFeatureFlagDenyAll(&_CoreGoerli.CallOpts, feature)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_CoreGoerli *CoreGoerliCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_CoreGoerli *CoreGoerliSession) GetImplementation() (common.Address, error) {
	return _CoreGoerli.Contract.GetImplementation(&_CoreGoerli.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_CoreGoerli *CoreGoerliCallerSession) GetImplementation() (common.Address, error) {
	return _CoreGoerli.Contract.GetImplementation(&_CoreGoerli.CallOpts)
}

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_CoreGoerli *CoreGoerliCaller) GetLocks(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getLocks", accountId, collateralType, offset, count)

	if err != nil {
		return *new([]CollateralLockData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CollateralLockData)).(*[]CollateralLockData)

	return out0, err

}

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_CoreGoerli *CoreGoerliSession) GetLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	return _CoreGoerli.Contract.GetLocks(&_CoreGoerli.CallOpts, accountId, collateralType, offset, count)
}

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_CoreGoerli *CoreGoerliCallerSession) GetLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	return _CoreGoerli.Contract.GetLocks(&_CoreGoerli.CallOpts, accountId, collateralType, offset, count)
}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCaller) GetMarketCollateral(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getMarketCollateral", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetMarketCollateral(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketCollateral(&_CoreGoerli.CallOpts, marketId)
}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCallerSession) GetMarketCollateral(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketCollateral(&_CoreGoerli.CallOpts, marketId)
}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_CoreGoerli *CoreGoerliCaller) GetMarketCollateralAmount(opts *bind.CallOpts, marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getMarketCollateralAmount", marketId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_CoreGoerli *CoreGoerliSession) GetMarketCollateralAmount(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketCollateralAmount(&_CoreGoerli.CallOpts, marketId, collateralType)
}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_CoreGoerli *CoreGoerliCallerSession) GetMarketCollateralAmount(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketCollateralAmount(&_CoreGoerli.CallOpts, marketId, collateralType)
}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCaller) GetMarketCollateralValue(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getMarketCollateralValue", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetMarketCollateralValue(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketCollateralValue(&_CoreGoerli.CallOpts, marketId)
}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCallerSession) GetMarketCollateralValue(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketCollateralValue(&_CoreGoerli.CallOpts, marketId)
}

// GetMarketFees is a free data retrieval call binding the contract method 0xdfb83437.
//
// Solidity: function getMarketFees(uint128 , uint256 amount) view returns(uint256 depositFeeAmount, uint256 withdrawFeeAmount)
func (_CoreGoerli *CoreGoerliCaller) GetMarketFees(opts *bind.CallOpts, arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getMarketFees", arg0, amount)

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
func (_CoreGoerli *CoreGoerliSession) GetMarketFees(arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	return _CoreGoerli.Contract.GetMarketFees(&_CoreGoerli.CallOpts, arg0, amount)
}

// GetMarketFees is a free data retrieval call binding the contract method 0xdfb83437.
//
// Solidity: function getMarketFees(uint128 , uint256 amount) view returns(uint256 depositFeeAmount, uint256 withdrawFeeAmount)
func (_CoreGoerli *CoreGoerliCallerSession) GetMarketFees(arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	return _CoreGoerli.Contract.GetMarketFees(&_CoreGoerli.CallOpts, arg0, amount)
}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_CoreGoerli *CoreGoerliCaller) GetMarketMinDelegateTime(opts *bind.CallOpts, marketId *big.Int) (uint32, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getMarketMinDelegateTime", marketId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_CoreGoerli *CoreGoerliSession) GetMarketMinDelegateTime(marketId *big.Int) (uint32, error) {
	return _CoreGoerli.Contract.GetMarketMinDelegateTime(&_CoreGoerli.CallOpts, marketId)
}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_CoreGoerli *CoreGoerliCallerSession) GetMarketMinDelegateTime(marketId *big.Int) (uint32, error) {
	return _CoreGoerli.Contract.GetMarketMinDelegateTime(&_CoreGoerli.CallOpts, marketId)
}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_CoreGoerli *CoreGoerliCaller) GetMarketNetIssuance(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getMarketNetIssuance", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_CoreGoerli *CoreGoerliSession) GetMarketNetIssuance(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketNetIssuance(&_CoreGoerli.CallOpts, marketId)
}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_CoreGoerli *CoreGoerliCallerSession) GetMarketNetIssuance(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketNetIssuance(&_CoreGoerli.CallOpts, marketId)
}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCaller) GetMarketReportedDebt(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getMarketReportedDebt", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetMarketReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketReportedDebt(&_CoreGoerli.CallOpts, marketId)
}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCallerSession) GetMarketReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketReportedDebt(&_CoreGoerli.CallOpts, marketId)
}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_CoreGoerli *CoreGoerliCaller) GetMarketTotalDebt(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getMarketTotalDebt", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_CoreGoerli *CoreGoerliSession) GetMarketTotalDebt(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketTotalDebt(&_CoreGoerli.CallOpts, marketId)
}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_CoreGoerli *CoreGoerliCallerSession) GetMarketTotalDebt(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMarketTotalDebt(&_CoreGoerli.CallOpts, marketId)
}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_CoreGoerli *CoreGoerliCaller) GetMaximumMarketCollateral(opts *bind.CallOpts, marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getMaximumMarketCollateral", marketId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetMaximumMarketCollateral(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMaximumMarketCollateral(&_CoreGoerli.CallOpts, marketId, collateralType)
}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_CoreGoerli *CoreGoerliCallerSession) GetMaximumMarketCollateral(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMaximumMarketCollateral(&_CoreGoerli.CallOpts, marketId, collateralType)
}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCaller) GetMinLiquidityRatio(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getMinLiquidityRatio", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetMinLiquidityRatio(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMinLiquidityRatio(&_CoreGoerli.CallOpts, marketId)
}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCallerSession) GetMinLiquidityRatio(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetMinLiquidityRatio(&_CoreGoerli.CallOpts, marketId)
}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_CoreGoerli *CoreGoerliCaller) GetMinLiquidityRatio0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getMinLiquidityRatio0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetMinLiquidityRatio0() (*big.Int, error) {
	return _CoreGoerli.Contract.GetMinLiquidityRatio0(&_CoreGoerli.CallOpts)
}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_CoreGoerli *CoreGoerliCallerSession) GetMinLiquidityRatio0() (*big.Int, error) {
	return _CoreGoerli.Contract.GetMinLiquidityRatio0(&_CoreGoerli.CallOpts)
}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_CoreGoerli *CoreGoerliCaller) GetNominatedPoolOwner(opts *bind.CallOpts, poolId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getNominatedPoolOwner", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_CoreGoerli *CoreGoerliSession) GetNominatedPoolOwner(poolId *big.Int) (common.Address, error) {
	return _CoreGoerli.Contract.GetNominatedPoolOwner(&_CoreGoerli.CallOpts, poolId)
}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_CoreGoerli *CoreGoerliCallerSession) GetNominatedPoolOwner(poolId *big.Int) (common.Address, error) {
	return _CoreGoerli.Contract.GetNominatedPoolOwner(&_CoreGoerli.CallOpts, poolId)
}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_CoreGoerli *CoreGoerliCaller) GetOracleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getOracleManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_CoreGoerli *CoreGoerliSession) GetOracleManager() (common.Address, error) {
	return _CoreGoerli.Contract.GetOracleManager(&_CoreGoerli.CallOpts)
}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_CoreGoerli *CoreGoerliCallerSession) GetOracleManager() (common.Address, error) {
	return _CoreGoerli.Contract.GetOracleManager(&_CoreGoerli.CallOpts)
}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_CoreGoerli *CoreGoerliCaller) GetPoolConfiguration(opts *bind.CallOpts, poolId *big.Int) ([]MarketConfigurationData, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getPoolConfiguration", poolId)

	if err != nil {
		return *new([]MarketConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketConfigurationData)).(*[]MarketConfigurationData)

	return out0, err

}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_CoreGoerli *CoreGoerliSession) GetPoolConfiguration(poolId *big.Int) ([]MarketConfigurationData, error) {
	return _CoreGoerli.Contract.GetPoolConfiguration(&_CoreGoerli.CallOpts, poolId)
}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_CoreGoerli *CoreGoerliCallerSession) GetPoolConfiguration(poolId *big.Int) ([]MarketConfigurationData, error) {
	return _CoreGoerli.Contract.GetPoolConfiguration(&_CoreGoerli.CallOpts, poolId)
}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_CoreGoerli *CoreGoerliCaller) GetPoolName(opts *bind.CallOpts, poolId *big.Int) (string, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getPoolName", poolId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_CoreGoerli *CoreGoerliSession) GetPoolName(poolId *big.Int) (string, error) {
	return _CoreGoerli.Contract.GetPoolName(&_CoreGoerli.CallOpts, poolId)
}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_CoreGoerli *CoreGoerliCallerSession) GetPoolName(poolId *big.Int) (string, error) {
	return _CoreGoerli.Contract.GetPoolName(&_CoreGoerli.CallOpts, poolId)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_CoreGoerli *CoreGoerliCaller) GetPoolOwner(opts *bind.CallOpts, poolId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getPoolOwner", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_CoreGoerli *CoreGoerliSession) GetPoolOwner(poolId *big.Int) (common.Address, error) {
	return _CoreGoerli.Contract.GetPoolOwner(&_CoreGoerli.CallOpts, poolId)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_CoreGoerli *CoreGoerliCallerSession) GetPoolOwner(poolId *big.Int) (common.Address, error) {
	return _CoreGoerli.Contract.GetPoolOwner(&_CoreGoerli.CallOpts, poolId)
}

// GetPositionCollateral is a free data retrieval call binding the contract method 0x33cc422b.
//
// Solidity: function getPositionCollateral(uint128 accountId, uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_CoreGoerli *CoreGoerliCaller) GetPositionCollateral(opts *bind.CallOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getPositionCollateral", accountId, poolId, collateralType)

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
func (_CoreGoerli *CoreGoerliSession) GetPositionCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _CoreGoerli.Contract.GetPositionCollateral(&_CoreGoerli.CallOpts, accountId, poolId, collateralType)
}

// GetPositionCollateral is a free data retrieval call binding the contract method 0x33cc422b.
//
// Solidity: function getPositionCollateral(uint128 accountId, uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_CoreGoerli *CoreGoerliCallerSession) GetPositionCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _CoreGoerli.Contract.GetPositionCollateral(&_CoreGoerli.CallOpts, accountId, poolId, collateralType)
}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_CoreGoerli *CoreGoerliCaller) GetPreferredPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getPreferredPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_CoreGoerli *CoreGoerliSession) GetPreferredPool() (*big.Int, error) {
	return _CoreGoerli.Contract.GetPreferredPool(&_CoreGoerli.CallOpts)
}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_CoreGoerli *CoreGoerliCallerSession) GetPreferredPool() (*big.Int, error) {
	return _CoreGoerli.Contract.GetPreferredPool(&_CoreGoerli.CallOpts)
}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_CoreGoerli *CoreGoerliCaller) GetRewardRate(opts *bind.CallOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getRewardRate", poolId, collateralType, distributor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetRewardRate(poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	return _CoreGoerli.Contract.GetRewardRate(&_CoreGoerli.CallOpts, poolId, collateralType, distributor)
}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_CoreGoerli *CoreGoerliCallerSession) GetRewardRate(poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	return _CoreGoerli.Contract.GetRewardRate(&_CoreGoerli.CallOpts, poolId, collateralType, distributor)
}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_CoreGoerli *CoreGoerliCaller) GetUsdToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getUsdToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_CoreGoerli *CoreGoerliSession) GetUsdToken() (common.Address, error) {
	return _CoreGoerli.Contract.GetUsdToken(&_CoreGoerli.CallOpts)
}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_CoreGoerli *CoreGoerliCallerSession) GetUsdToken() (common.Address, error) {
	return _CoreGoerli.Contract.GetUsdToken(&_CoreGoerli.CallOpts)
}

// GetVaultCollateral is a free data retrieval call binding the contract method 0x078145a8.
//
// Solidity: function getVaultCollateral(uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_CoreGoerli *CoreGoerliCaller) GetVaultCollateral(opts *bind.CallOpts, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getVaultCollateral", poolId, collateralType)

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
func (_CoreGoerli *CoreGoerliSession) GetVaultCollateral(poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _CoreGoerli.Contract.GetVaultCollateral(&_CoreGoerli.CallOpts, poolId, collateralType)
}

// GetVaultCollateral is a free data retrieval call binding the contract method 0x078145a8.
//
// Solidity: function getVaultCollateral(uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_CoreGoerli *CoreGoerliCallerSession) GetVaultCollateral(poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _CoreGoerli.Contract.GetVaultCollateral(&_CoreGoerli.CallOpts, poolId, collateralType)
}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCaller) GetWithdrawableMarketUsd(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "getWithdrawableMarketUsd", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetWithdrawableMarketUsd(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetWithdrawableMarketUsd(&_CoreGoerli.CallOpts, marketId)
}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_CoreGoerli *CoreGoerliCallerSession) GetWithdrawableMarketUsd(marketId *big.Int) (*big.Int, error) {
	return _CoreGoerli.Contract.GetWithdrawableMarketUsd(&_CoreGoerli.CallOpts, marketId)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreGoerli *CoreGoerliCaller) HasPermission(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "hasPermission", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreGoerli *CoreGoerliSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _CoreGoerli.Contract.HasPermission(&_CoreGoerli.CallOpts, accountId, permission, user)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d34166b.
//
// Solidity: function hasPermission(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreGoerli *CoreGoerliCallerSession) HasPermission(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _CoreGoerli.Contract.HasPermission(&_CoreGoerli.CallOpts, accountId, permission, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreGoerli *CoreGoerliCaller) IsAuthorized(opts *bind.CallOpts, accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "isAuthorized", accountId, permission, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreGoerli *CoreGoerliSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _CoreGoerli.Contract.IsAuthorized(&_CoreGoerli.CallOpts, accountId, permission, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x1213d453.
//
// Solidity: function isAuthorized(uint128 accountId, bytes32 permission, address user) view returns(bool)
func (_CoreGoerli *CoreGoerliCallerSession) IsAuthorized(accountId *big.Int, permission [32]byte, user common.Address) (bool, error) {
	return _CoreGoerli.Contract.IsAuthorized(&_CoreGoerli.CallOpts, accountId, permission, user)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_CoreGoerli *CoreGoerliCaller) IsFeatureAllowed(opts *bind.CallOpts, feature [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "isFeatureAllowed", feature, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_CoreGoerli *CoreGoerliSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _CoreGoerli.Contract.IsFeatureAllowed(&_CoreGoerli.CallOpts, feature, account)
}

// IsFeatureAllowed is a free data retrieval call binding the contract method 0xcf635949.
//
// Solidity: function isFeatureAllowed(bytes32 feature, address account) view returns(bool)
func (_CoreGoerli *CoreGoerliCallerSession) IsFeatureAllowed(feature [32]byte, account common.Address) (bool, error) {
	return _CoreGoerli.Contract.IsFeatureAllowed(&_CoreGoerli.CallOpts, feature, account)
}

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_CoreGoerli *CoreGoerliCaller) IsMarketCapacityLocked(opts *bind.CallOpts, marketId *big.Int) (bool, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "isMarketCapacityLocked", marketId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_CoreGoerli *CoreGoerliSession) IsMarketCapacityLocked(marketId *big.Int) (bool, error) {
	return _CoreGoerli.Contract.IsMarketCapacityLocked(&_CoreGoerli.CallOpts, marketId)
}

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_CoreGoerli *CoreGoerliCallerSession) IsMarketCapacityLocked(marketId *big.Int) (bool, error) {
	return _CoreGoerli.Contract.IsMarketCapacityLocked(&_CoreGoerli.CallOpts, marketId)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_CoreGoerli *CoreGoerliCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_CoreGoerli *CoreGoerliSession) NominatedOwner() (common.Address, error) {
	return _CoreGoerli.Contract.NominatedOwner(&_CoreGoerli.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_CoreGoerli *CoreGoerliCallerSession) NominatedOwner() (common.Address, error) {
	return _CoreGoerli.Contract.NominatedOwner(&_CoreGoerli.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CoreGoerli *CoreGoerliCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CoreGoerli *CoreGoerliSession) Owner() (common.Address, error) {
	return _CoreGoerli.Contract.Owner(&_CoreGoerli.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CoreGoerli *CoreGoerliCallerSession) Owner() (common.Address, error) {
	return _CoreGoerli.Contract.Owner(&_CoreGoerli.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoreGoerli *CoreGoerliCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CoreGoerli.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoreGoerli *CoreGoerliSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CoreGoerli.Contract.SupportsInterface(&_CoreGoerli.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoreGoerli *CoreGoerliCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CoreGoerli.Contract.SupportsInterface(&_CoreGoerli.CallOpts, interfaceId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_CoreGoerli *CoreGoerliTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_CoreGoerli *CoreGoerliSession) AcceptOwnership() (*types.Transaction, error) {
	return _CoreGoerli.Contract.AcceptOwnership(&_CoreGoerli.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_CoreGoerli *CoreGoerliTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CoreGoerli.Contract.AcceptOwnership(&_CoreGoerli.TransactOpts)
}

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactor) AcceptPoolOwnership(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "acceptPoolOwnership", poolId)
}

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliSession) AcceptPoolOwnership(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.AcceptPoolOwnership(&_CoreGoerli.TransactOpts, poolId)
}

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) AcceptPoolOwnership(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.AcceptPoolOwnership(&_CoreGoerli.TransactOpts, poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactor) AddApprovedPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "addApprovedPool", poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliSession) AddApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.AddApprovedPool(&_CoreGoerli.TransactOpts, poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) AddApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.AddApprovedPool(&_CoreGoerli.TransactOpts, poolId)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreGoerli *CoreGoerliTransactor) AddToFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "addToFeatureFlagAllowlist", feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreGoerli *CoreGoerliSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.AddToFeatureFlagAllowlist(&_CoreGoerli.TransactOpts, feature, account)
}

// AddToFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xa0778144.
//
// Solidity: function addToFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) AddToFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.AddToFeatureFlagAllowlist(&_CoreGoerli.TransactOpts, feature, account)
}

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_CoreGoerli *CoreGoerliTransactor) AssociateDebt(opts *bind.TransactOpts, marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "associateDebt", marketId, poolId, collateralType, accountId, amount)
}

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_CoreGoerli *CoreGoerliSession) AssociateDebt(marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.AssociateDebt(&_CoreGoerli.TransactOpts, marketId, poolId, collateralType, accountId, amount)
}

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_CoreGoerli *CoreGoerliTransactorSession) AssociateDebt(marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.AssociateDebt(&_CoreGoerli.TransactOpts, marketId, poolId, collateralType, accountId, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreGoerli *CoreGoerliTransactor) BurnUsd(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "burnUsd", accountId, poolId, collateralType, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreGoerli *CoreGoerliSession) BurnUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.BurnUsd(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) BurnUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.BurnUsd(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType, amount)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_CoreGoerli *CoreGoerliTransactor) CcipReceive(opts *bind.TransactOpts, message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_CoreGoerli *CoreGoerliSession) CcipReceive(message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CcipReceive(&_CoreGoerli.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) CcipReceive(message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CcipReceive(&_CoreGoerli.TransactOpts, message)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_CoreGoerli *CoreGoerliTransactor) ClaimRewards(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "claimRewards", accountId, poolId, collateralType, distributor)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_CoreGoerli *CoreGoerliSession) ClaimRewards(accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.ClaimRewards(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType, distributor)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_CoreGoerli *CoreGoerliTransactorSession) ClaimRewards(accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.ClaimRewards(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType, distributor)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_CoreGoerli *CoreGoerliTransactor) CleanExpiredLocks(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "cleanExpiredLocks", accountId, collateralType, offset, count)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_CoreGoerli *CoreGoerliSession) CleanExpiredLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CleanExpiredLocks(&_CoreGoerli.TransactOpts, accountId, collateralType, offset, count)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_CoreGoerli *CoreGoerliTransactorSession) CleanExpiredLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CleanExpiredLocks(&_CoreGoerli.TransactOpts, accountId, collateralType, offset, count)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_CoreGoerli *CoreGoerliTransactor) ConfigureChainlinkCrossChain(opts *bind.TransactOpts, ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "configureChainlinkCrossChain", ccipRouter, ccipTokenPool)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_CoreGoerli *CoreGoerliSession) ConfigureChainlinkCrossChain(ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.ConfigureChainlinkCrossChain(&_CoreGoerli.TransactOpts, ccipRouter, ccipTokenPool)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) ConfigureChainlinkCrossChain(ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.ConfigureChainlinkCrossChain(&_CoreGoerli.TransactOpts, ccipRouter, ccipTokenPool)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_CoreGoerli *CoreGoerliTransactor) ConfigureCollateral(opts *bind.TransactOpts, config CollateralConfigurationData) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "configureCollateral", config)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_CoreGoerli *CoreGoerliSession) ConfigureCollateral(config CollateralConfigurationData) (*types.Transaction, error) {
	return _CoreGoerli.Contract.ConfigureCollateral(&_CoreGoerli.TransactOpts, config)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) ConfigureCollateral(config CollateralConfigurationData) (*types.Transaction, error) {
	return _CoreGoerli.Contract.ConfigureCollateral(&_CoreGoerli.TransactOpts, config)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_CoreGoerli *CoreGoerliTransactor) ConfigureMaximumMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "configureMaximumMarketCollateral", marketId, collateralType, amount)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_CoreGoerli *CoreGoerliSession) ConfigureMaximumMarketCollateral(marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.ConfigureMaximumMarketCollateral(&_CoreGoerli.TransactOpts, marketId, collateralType, amount)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) ConfigureMaximumMarketCollateral(marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.ConfigureMaximumMarketCollateral(&_CoreGoerli.TransactOpts, marketId, collateralType, amount)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_CoreGoerli *CoreGoerliTransactor) ConfigureOracleManager(opts *bind.TransactOpts, oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "configureOracleManager", oracleManagerAddress)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_CoreGoerli *CoreGoerliSession) ConfigureOracleManager(oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.ConfigureOracleManager(&_CoreGoerli.TransactOpts, oracleManagerAddress)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) ConfigureOracleManager(oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.ConfigureOracleManager(&_CoreGoerli.TransactOpts, oracleManagerAddress)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_CoreGoerli *CoreGoerliTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_CoreGoerli *CoreGoerliSession) CreateAccount() (*types.Transaction, error) {
	return _CoreGoerli.Contract.CreateAccount(&_CoreGoerli.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(uint128 accountId)
func (_CoreGoerli *CoreGoerliTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _CoreGoerli.Contract.CreateAccount(&_CoreGoerli.TransactOpts)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_CoreGoerli *CoreGoerliTransactor) CreateAccount0(opts *bind.TransactOpts, requestedAccountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "createAccount0", requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_CoreGoerli *CoreGoerliSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CreateAccount0(&_CoreGoerli.TransactOpts, requestedAccountId)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xcadb09a5.
//
// Solidity: function createAccount(uint128 requestedAccountId) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) CreateAccount0(requestedAccountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CreateAccount0(&_CoreGoerli.TransactOpts, requestedAccountId)
}

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_CoreGoerli *CoreGoerliTransactor) CreateLock(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "createLock", accountId, collateralType, amount, expireTimestamp)
}

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_CoreGoerli *CoreGoerliSession) CreateLock(accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CreateLock(&_CoreGoerli.TransactOpts, accountId, collateralType, amount, expireTimestamp)
}

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) CreateLock(accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CreateLock(&_CoreGoerli.TransactOpts, accountId, collateralType, amount, expireTimestamp)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_CoreGoerli *CoreGoerliTransactor) CreatePool(opts *bind.TransactOpts, requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "createPool", requestedPoolId, owner)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_CoreGoerli *CoreGoerliSession) CreatePool(requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CreatePool(&_CoreGoerli.TransactOpts, requestedPoolId, owner)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) CreatePool(requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.CreatePool(&_CoreGoerli.TransactOpts, requestedPoolId, owner)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_CoreGoerli *CoreGoerliTransactor) DelegateCollateral(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "delegateCollateral", accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_CoreGoerli *CoreGoerliSession) DelegateCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.DelegateCollateral(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) DelegateCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.DelegateCollateral(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliTransactor) Deposit(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "deposit", accountId, collateralType, tokenAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliSession) Deposit(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.Deposit(&_CoreGoerli.TransactOpts, accountId, collateralType, tokenAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) Deposit(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.Deposit(&_CoreGoerli.TransactOpts, accountId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliTransactor) DepositMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "depositMarketCollateral", marketId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliSession) DepositMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.DepositMarketCollateral(&_CoreGoerli.TransactOpts, marketId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) DepositMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.DepositMarketCollateral(&_CoreGoerli.TransactOpts, marketId, collateralType, tokenAmount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreGoerli *CoreGoerliTransactor) DepositMarketUsd(opts *bind.TransactOpts, marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "depositMarketUsd", marketId, target, amount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreGoerli *CoreGoerliSession) DepositMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.DepositMarketUsd(&_CoreGoerli.TransactOpts, marketId, target, amount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreGoerli *CoreGoerliTransactorSession) DepositMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.DepositMarketUsd(&_CoreGoerli.TransactOpts, marketId, target, amount)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_CoreGoerli *CoreGoerliTransactor) DistributeDebtToPools(opts *bind.TransactOpts, marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "distributeDebtToPools", marketId, maxIter)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_CoreGoerli *CoreGoerliSession) DistributeDebtToPools(marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.DistributeDebtToPools(&_CoreGoerli.TransactOpts, marketId, maxIter)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_CoreGoerli *CoreGoerliTransactorSession) DistributeDebtToPools(marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.DistributeDebtToPools(&_CoreGoerli.TransactOpts, marketId, maxIter)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_CoreGoerli *CoreGoerliTransactor) DistributeRewards(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "distributeRewards", poolId, collateralType, amount, start, duration)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_CoreGoerli *CoreGoerliSession) DistributeRewards(poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _CoreGoerli.Contract.DistributeRewards(&_CoreGoerli.TransactOpts, poolId, collateralType, amount, start, duration)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) DistributeRewards(poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _CoreGoerli.Contract.DistributeRewards(&_CoreGoerli.TransactOpts, poolId, collateralType, amount, start, duration)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_CoreGoerli *CoreGoerliTransactor) GetMarketDebtPerShare(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "getMarketDebtPerShare", marketId)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_CoreGoerli *CoreGoerliSession) GetMarketDebtPerShare(marketId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetMarketDebtPerShare(&_CoreGoerli.TransactOpts, marketId)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_CoreGoerli *CoreGoerliTransactorSession) GetMarketDebtPerShare(marketId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetMarketDebtPerShare(&_CoreGoerli.TransactOpts, marketId)
}

// GetMarketPoolDebtDistribution is a paid mutator transaction binding the contract method 0x25eeea4b.
//
// Solidity: function getMarketPoolDebtDistribution(uint128 marketId, uint128 poolId) returns(uint256 sharesD18, uint128 totalSharesD18, int128 valuePerShareD27)
func (_CoreGoerli *CoreGoerliTransactor) GetMarketPoolDebtDistribution(opts *bind.TransactOpts, marketId *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "getMarketPoolDebtDistribution", marketId, poolId)
}

// GetMarketPoolDebtDistribution is a paid mutator transaction binding the contract method 0x25eeea4b.
//
// Solidity: function getMarketPoolDebtDistribution(uint128 marketId, uint128 poolId) returns(uint256 sharesD18, uint128 totalSharesD18, int128 valuePerShareD27)
func (_CoreGoerli *CoreGoerliSession) GetMarketPoolDebtDistribution(marketId *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetMarketPoolDebtDistribution(&_CoreGoerli.TransactOpts, marketId, poolId)
}

// GetMarketPoolDebtDistribution is a paid mutator transaction binding the contract method 0x25eeea4b.
//
// Solidity: function getMarketPoolDebtDistribution(uint128 marketId, uint128 poolId) returns(uint256 sharesD18, uint128 totalSharesD18, int128 valuePerShareD27)
func (_CoreGoerli *CoreGoerliTransactorSession) GetMarketPoolDebtDistribution(marketId *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetMarketPoolDebtDistribution(&_CoreGoerli.TransactOpts, marketId, poolId)
}

// GetMarketPools is a paid mutator transaction binding the contract method 0xbe0b8e6f.
//
// Solidity: function getMarketPools(uint128 marketId) returns(uint128[] inRangePoolIds, uint128[] outRangePoolIds)
func (_CoreGoerli *CoreGoerliTransactor) GetMarketPools(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "getMarketPools", marketId)
}

// GetMarketPools is a paid mutator transaction binding the contract method 0xbe0b8e6f.
//
// Solidity: function getMarketPools(uint128 marketId) returns(uint128[] inRangePoolIds, uint128[] outRangePoolIds)
func (_CoreGoerli *CoreGoerliSession) GetMarketPools(marketId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetMarketPools(&_CoreGoerli.TransactOpts, marketId)
}

// GetMarketPools is a paid mutator transaction binding the contract method 0xbe0b8e6f.
//
// Solidity: function getMarketPools(uint128 marketId) returns(uint128[] inRangePoolIds, uint128[] outRangePoolIds)
func (_CoreGoerli *CoreGoerliTransactorSession) GetMarketPools(marketId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetMarketPools(&_CoreGoerli.TransactOpts, marketId)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_CoreGoerli *CoreGoerliTransactor) GetPosition(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "getPosition", accountId, poolId, collateralType)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_CoreGoerli *CoreGoerliSession) GetPosition(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetPosition(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_CoreGoerli *CoreGoerliTransactorSession) GetPosition(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetPosition(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_CoreGoerli *CoreGoerliTransactor) GetPositionCollateralRatio(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "getPositionCollateralRatio", accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetPositionCollateralRatio(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetPositionCollateralRatio(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_CoreGoerli *CoreGoerliTransactorSession) GetPositionCollateralRatio(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetPositionCollateralRatio(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_CoreGoerli *CoreGoerliTransactor) GetPositionDebt(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "getPositionDebt", accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_CoreGoerli *CoreGoerliSession) GetPositionDebt(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetPositionDebt(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_CoreGoerli *CoreGoerliTransactorSession) GetPositionDebt(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetPositionDebt(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_CoreGoerli *CoreGoerliTransactor) GetVaultCollateralRatio(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "getVaultCollateralRatio", poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_CoreGoerli *CoreGoerliSession) GetVaultCollateralRatio(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetVaultCollateralRatio(&_CoreGoerli.TransactOpts, poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_CoreGoerli *CoreGoerliTransactorSession) GetVaultCollateralRatio(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetVaultCollateralRatio(&_CoreGoerli.TransactOpts, poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_CoreGoerli *CoreGoerliTransactor) GetVaultDebt(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "getVaultDebt", poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_CoreGoerli *CoreGoerliSession) GetVaultDebt(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetVaultDebt(&_CoreGoerli.TransactOpts, poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_CoreGoerli *CoreGoerliTransactorSession) GetVaultDebt(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GetVaultDebt(&_CoreGoerli.TransactOpts, poolId, collateralType)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreGoerli *CoreGoerliTransactor) GrantPermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "grantPermission", accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreGoerli *CoreGoerliSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GrantPermission(&_CoreGoerli.TransactOpts, accountId, permission, user)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x00cd9ef3.
//
// Solidity: function grantPermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) GrantPermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.GrantPermission(&_CoreGoerli.TransactOpts, accountId, permission, user)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_CoreGoerli *CoreGoerliTransactor) InitOrUpgradeNft(opts *bind.TransactOpts, id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "initOrUpgradeNft", id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_CoreGoerli *CoreGoerliSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.InitOrUpgradeNft(&_CoreGoerli.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeNft is a paid mutator transaction binding the contract method 0x2d22bef9.
//
// Solidity: function initOrUpgradeNft(bytes32 id, string name, string symbol, string uri, address impl) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) InitOrUpgradeNft(id [32]byte, name string, symbol string, uri string, impl common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.InitOrUpgradeNft(&_CoreGoerli.TransactOpts, id, name, symbol, uri, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_CoreGoerli *CoreGoerliTransactor) InitOrUpgradeToken(opts *bind.TransactOpts, id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "initOrUpgradeToken", id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_CoreGoerli *CoreGoerliSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.InitOrUpgradeToken(&_CoreGoerli.TransactOpts, id, name, symbol, decimals, impl)
}

// InitOrUpgradeToken is a paid mutator transaction binding the contract method 0xc6f79537.
//
// Solidity: function initOrUpgradeToken(bytes32 id, string name, string symbol, uint8 decimals, address impl) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) InitOrUpgradeToken(id [32]byte, name string, symbol string, decimals uint8, impl common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.InitOrUpgradeToken(&_CoreGoerli.TransactOpts, id, name, symbol, decimals, impl)
}

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_CoreGoerli *CoreGoerliTransactor) IsPositionLiquidatable(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "isPositionLiquidatable", accountId, poolId, collateralType)
}

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_CoreGoerli *CoreGoerliSession) IsPositionLiquidatable(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.IsPositionLiquidatable(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType)
}

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_CoreGoerli *CoreGoerliTransactorSession) IsPositionLiquidatable(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.IsPositionLiquidatable(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_CoreGoerli *CoreGoerliTransactor) IsVaultLiquidatable(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "isVaultLiquidatable", poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_CoreGoerli *CoreGoerliSession) IsVaultLiquidatable(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.IsVaultLiquidatable(&_CoreGoerli.TransactOpts, poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_CoreGoerli *CoreGoerliTransactorSession) IsVaultLiquidatable(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.IsVaultLiquidatable(&_CoreGoerli.TransactOpts, poolId, collateralType)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_CoreGoerli *CoreGoerliTransactor) Liquidate(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "liquidate", accountId, poolId, collateralType, liquidateAsAccountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_CoreGoerli *CoreGoerliSession) Liquidate(accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.Liquidate(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType, liquidateAsAccountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_CoreGoerli *CoreGoerliTransactorSession) Liquidate(accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.Liquidate(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType, liquidateAsAccountId)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_CoreGoerli *CoreGoerliTransactor) LiquidateVault(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "liquidateVault", poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_CoreGoerli *CoreGoerliSession) LiquidateVault(poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.LiquidateVault(&_CoreGoerli.TransactOpts, poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_CoreGoerli *CoreGoerliTransactorSession) LiquidateVault(poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.LiquidateVault(&_CoreGoerli.TransactOpts, poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreGoerli *CoreGoerliTransactor) MintUsd(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "mintUsd", accountId, poolId, collateralType, amount)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreGoerli *CoreGoerliSession) MintUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.MintUsd(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType, amount)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) MintUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.MintUsd(&_CoreGoerli.TransactOpts, accountId, poolId, collateralType, amount)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_CoreGoerli *CoreGoerliTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_CoreGoerli *CoreGoerliSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _CoreGoerli.Contract.Multicall(&_CoreGoerli.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_CoreGoerli *CoreGoerliTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _CoreGoerli.Contract.Multicall(&_CoreGoerli.TransactOpts, data)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_CoreGoerli *CoreGoerliTransactor) NominateNewOwner(opts *bind.TransactOpts, newNominatedOwner common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "nominateNewOwner", newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_CoreGoerli *CoreGoerliSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.NominateNewOwner(&_CoreGoerli.TransactOpts, newNominatedOwner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address newNominatedOwner) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) NominateNewOwner(newNominatedOwner common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.NominateNewOwner(&_CoreGoerli.TransactOpts, newNominatedOwner)
}

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactor) NominatePoolOwner(opts *bind.TransactOpts, nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "nominatePoolOwner", nominatedOwner, poolId)
}

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliSession) NominatePoolOwner(nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.NominatePoolOwner(&_CoreGoerli.TransactOpts, nominatedOwner, poolId)
}

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) NominatePoolOwner(nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.NominatePoolOwner(&_CoreGoerli.TransactOpts, nominatedOwner, poolId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_CoreGoerli *CoreGoerliTransactor) NotifyAccountTransfer(opts *bind.TransactOpts, to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "notifyAccountTransfer", to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_CoreGoerli *CoreGoerliSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.NotifyAccountTransfer(&_CoreGoerli.TransactOpts, to, accountId)
}

// NotifyAccountTransfer is a paid mutator transaction binding the contract method 0x7dec8b55.
//
// Solidity: function notifyAccountTransfer(address to, uint128 accountId) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) NotifyAccountTransfer(to common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.NotifyAccountTransfer(&_CoreGoerli.TransactOpts, to, accountId)
}

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_CoreGoerli *CoreGoerliTransactor) RebalancePool(opts *bind.TransactOpts, poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "rebalancePool", poolId, optionalCollateralType)
}

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_CoreGoerli *CoreGoerliSession) RebalancePool(poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RebalancePool(&_CoreGoerli.TransactOpts, poolId, optionalCollateralType)
}

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) RebalancePool(poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RebalancePool(&_CoreGoerli.TransactOpts, poolId, optionalCollateralType)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_CoreGoerli *CoreGoerliTransactor) RegisterMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "registerMarket", market)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_CoreGoerli *CoreGoerliSession) RegisterMarket(market common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RegisterMarket(&_CoreGoerli.TransactOpts, market)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_CoreGoerli *CoreGoerliTransactorSession) RegisterMarket(market common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RegisterMarket(&_CoreGoerli.TransactOpts, market)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreGoerli *CoreGoerliTransactor) RegisterRewardsDistributor(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "registerRewardsDistributor", poolId, collateralType, distributor)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreGoerli *CoreGoerliSession) RegisterRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RegisterRewardsDistributor(&_CoreGoerli.TransactOpts, poolId, collateralType, distributor)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) RegisterRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RegisterRewardsDistributor(&_CoreGoerli.TransactOpts, poolId, collateralType, distributor)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_CoreGoerli *CoreGoerliTransactor) RegisterUnmanagedSystem(opts *bind.TransactOpts, id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "registerUnmanagedSystem", id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_CoreGoerli *CoreGoerliSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RegisterUnmanagedSystem(&_CoreGoerli.TransactOpts, id, endpoint)
}

// RegisterUnmanagedSystem is a paid mutator transaction binding the contract method 0xd245d983.
//
// Solidity: function registerUnmanagedSystem(bytes32 id, address endpoint) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) RegisterUnmanagedSystem(id [32]byte, endpoint common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RegisterUnmanagedSystem(&_CoreGoerli.TransactOpts, id, endpoint)
}

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactor) RemoveApprovedPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "removeApprovedPool", poolId)
}

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliSession) RemoveApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RemoveApprovedPool(&_CoreGoerli.TransactOpts, poolId)
}

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) RemoveApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RemoveApprovedPool(&_CoreGoerli.TransactOpts, poolId)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreGoerli *CoreGoerliTransactor) RemoveFromFeatureFlagAllowlist(opts *bind.TransactOpts, feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "removeFromFeatureFlagAllowlist", feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreGoerli *CoreGoerliSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RemoveFromFeatureFlagAllowlist(&_CoreGoerli.TransactOpts, feature, account)
}

// RemoveFromFeatureFlagAllowlist is a paid mutator transaction binding the contract method 0xb7746b59.
//
// Solidity: function removeFromFeatureFlagAllowlist(bytes32 feature, address account) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) RemoveFromFeatureFlagAllowlist(feature [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RemoveFromFeatureFlagAllowlist(&_CoreGoerli.TransactOpts, feature, account)
}

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreGoerli *CoreGoerliTransactor) RemoveRewardsDistributor(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "removeRewardsDistributor", poolId, collateralType, distributor)
}

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreGoerli *CoreGoerliSession) RemoveRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RemoveRewardsDistributor(&_CoreGoerli.TransactOpts, poolId, collateralType, distributor)
}

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) RemoveRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RemoveRewardsDistributor(&_CoreGoerli.TransactOpts, poolId, collateralType, distributor)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_CoreGoerli *CoreGoerliTransactor) RenounceNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "renounceNomination")
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_CoreGoerli *CoreGoerliSession) RenounceNomination() (*types.Transaction, error) {
	return _CoreGoerli.Contract.RenounceNomination(&_CoreGoerli.TransactOpts)
}

// RenounceNomination is a paid mutator transaction binding the contract method 0x718fe928.
//
// Solidity: function renounceNomination() returns()
func (_CoreGoerli *CoreGoerliTransactorSession) RenounceNomination() (*types.Transaction, error) {
	return _CoreGoerli.Contract.RenounceNomination(&_CoreGoerli.TransactOpts)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_CoreGoerli *CoreGoerliTransactor) RenouncePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "renouncePermission", accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_CoreGoerli *CoreGoerliSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RenouncePermission(&_CoreGoerli.TransactOpts, accountId, permission)
}

// RenouncePermission is a paid mutator transaction binding the contract method 0x47c1c561.
//
// Solidity: function renouncePermission(uint128 accountId, bytes32 permission) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) RenouncePermission(accountId *big.Int, permission [32]byte) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RenouncePermission(&_CoreGoerli.TransactOpts, accountId, permission)
}

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactor) RenouncePoolNomination(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "renouncePoolNomination", poolId)
}

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliSession) RenouncePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RenouncePoolNomination(&_CoreGoerli.TransactOpts, poolId)
}

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) RenouncePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RenouncePoolNomination(&_CoreGoerli.TransactOpts, poolId)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreGoerli *CoreGoerliTransactor) RevokePermission(opts *bind.TransactOpts, accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "revokePermission", accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreGoerli *CoreGoerliSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RevokePermission(&_CoreGoerli.TransactOpts, accountId, permission, user)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa7627288.
//
// Solidity: function revokePermission(uint128 accountId, bytes32 permission, address user) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) RevokePermission(accountId *big.Int, permission [32]byte, user common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RevokePermission(&_CoreGoerli.TransactOpts, accountId, permission, user)
}

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactor) RevokePoolNomination(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "revokePoolNomination", poolId)
}

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliSession) RevokePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RevokePoolNomination(&_CoreGoerli.TransactOpts, poolId)
}

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) RevokePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.RevokePoolNomination(&_CoreGoerli.TransactOpts, poolId)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_CoreGoerli *CoreGoerliTransactor) SetConfig(opts *bind.TransactOpts, k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "setConfig", k, v)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_CoreGoerli *CoreGoerliSession) SetConfig(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetConfig(&_CoreGoerli.TransactOpts, k, v)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) SetConfig(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetConfig(&_CoreGoerli.TransactOpts, k, v)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_CoreGoerli *CoreGoerliTransactor) SetDeniers(opts *bind.TransactOpts, feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "setDeniers", feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_CoreGoerli *CoreGoerliSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetDeniers(&_CoreGoerli.TransactOpts, feature, deniers)
}

// SetDeniers is a paid mutator transaction binding the contract method 0x715cb7d2.
//
// Solidity: function setDeniers(bytes32 feature, address[] deniers) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) SetDeniers(feature [32]byte, deniers []common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetDeniers(&_CoreGoerli.TransactOpts, feature, deniers)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_CoreGoerli *CoreGoerliTransactor) SetFeatureFlagAllowAll(opts *bind.TransactOpts, feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "setFeatureFlagAllowAll", feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_CoreGoerli *CoreGoerliSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetFeatureFlagAllowAll(&_CoreGoerli.TransactOpts, feature, allowAll)
}

// SetFeatureFlagAllowAll is a paid mutator transaction binding the contract method 0x7d632bd2.
//
// Solidity: function setFeatureFlagAllowAll(bytes32 feature, bool allowAll) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) SetFeatureFlagAllowAll(feature [32]byte, allowAll bool) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetFeatureFlagAllowAll(&_CoreGoerli.TransactOpts, feature, allowAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_CoreGoerli *CoreGoerliTransactor) SetFeatureFlagDenyAll(opts *bind.TransactOpts, feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "setFeatureFlagDenyAll", feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_CoreGoerli *CoreGoerliSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetFeatureFlagDenyAll(&_CoreGoerli.TransactOpts, feature, denyAll)
}

// SetFeatureFlagDenyAll is a paid mutator transaction binding the contract method 0x5e52ad6e.
//
// Solidity: function setFeatureFlagDenyAll(bytes32 feature, bool denyAll) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) SetFeatureFlagDenyAll(feature [32]byte, denyAll bool) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetFeatureFlagDenyAll(&_CoreGoerli.TransactOpts, feature, denyAll)
}

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_CoreGoerli *CoreGoerliTransactor) SetMarketMinDelegateTime(opts *bind.TransactOpts, marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "setMarketMinDelegateTime", marketId, minDelegateTime)
}

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_CoreGoerli *CoreGoerliSession) SetMarketMinDelegateTime(marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetMarketMinDelegateTime(&_CoreGoerli.TransactOpts, marketId, minDelegateTime)
}

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) SetMarketMinDelegateTime(marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetMarketMinDelegateTime(&_CoreGoerli.TransactOpts, marketId, minDelegateTime)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_CoreGoerli *CoreGoerliTransactor) SetMinLiquidityRatio(opts *bind.TransactOpts, marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "setMinLiquidityRatio", marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_CoreGoerli *CoreGoerliSession) SetMinLiquidityRatio(marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetMinLiquidityRatio(&_CoreGoerli.TransactOpts, marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) SetMinLiquidityRatio(marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetMinLiquidityRatio(&_CoreGoerli.TransactOpts, marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_CoreGoerli *CoreGoerliTransactor) SetMinLiquidityRatio0(opts *bind.TransactOpts, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "setMinLiquidityRatio0", minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_CoreGoerli *CoreGoerliSession) SetMinLiquidityRatio0(minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetMinLiquidityRatio0(&_CoreGoerli.TransactOpts, minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) SetMinLiquidityRatio0(minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetMinLiquidityRatio0(&_CoreGoerli.TransactOpts, minLiquidityRatio)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_CoreGoerli *CoreGoerliTransactor) SetPoolConfiguration(opts *bind.TransactOpts, poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "setPoolConfiguration", poolId, newMarketConfigurations)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_CoreGoerli *CoreGoerliSession) SetPoolConfiguration(poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetPoolConfiguration(&_CoreGoerli.TransactOpts, poolId, newMarketConfigurations)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) SetPoolConfiguration(poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetPoolConfiguration(&_CoreGoerli.TransactOpts, poolId, newMarketConfigurations)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_CoreGoerli *CoreGoerliTransactor) SetPoolName(opts *bind.TransactOpts, poolId *big.Int, name string) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "setPoolName", poolId, name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_CoreGoerli *CoreGoerliSession) SetPoolName(poolId *big.Int, name string) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetPoolName(&_CoreGoerli.TransactOpts, poolId, name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) SetPoolName(poolId *big.Int, name string) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetPoolName(&_CoreGoerli.TransactOpts, poolId, name)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactor) SetPreferredPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "setPreferredPool", poolId)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliSession) SetPreferredPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetPreferredPool(&_CoreGoerli.TransactOpts, poolId)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) SetPreferredPool(poolId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetPreferredPool(&_CoreGoerli.TransactOpts, poolId)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_CoreGoerli *CoreGoerliTransactor) SetSupportedCrossChainNetworks(opts *bind.TransactOpts, supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "setSupportedCrossChainNetworks", supportedNetworks, ccipSelectors)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_CoreGoerli *CoreGoerliSession) SetSupportedCrossChainNetworks(supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetSupportedCrossChainNetworks(&_CoreGoerli.TransactOpts, supportedNetworks, ccipSelectors)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_CoreGoerli *CoreGoerliTransactorSession) SetSupportedCrossChainNetworks(supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SetSupportedCrossChainNetworks(&_CoreGoerli.TransactOpts, supportedNetworks, ccipSelectors)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_CoreGoerli *CoreGoerliTransactor) SimulateUpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "simulateUpgradeTo", newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_CoreGoerli *CoreGoerliSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SimulateUpgradeTo(&_CoreGoerli.TransactOpts, newImplementation)
}

// SimulateUpgradeTo is a paid mutator transaction binding the contract method 0xc7f62cda.
//
// Solidity: function simulateUpgradeTo(address newImplementation) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) SimulateUpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.SimulateUpgradeTo(&_CoreGoerli.TransactOpts, newImplementation)
}

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_CoreGoerli *CoreGoerliTransactor) TransferCrossChain(opts *bind.TransactOpts, destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "transferCrossChain", destChainId, amount)
}

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_CoreGoerli *CoreGoerliSession) TransferCrossChain(destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.TransferCrossChain(&_CoreGoerli.TransactOpts, destChainId, amount)
}

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_CoreGoerli *CoreGoerliTransactorSession) TransferCrossChain(destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.TransferCrossChain(&_CoreGoerli.TransactOpts, destChainId, amount)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_CoreGoerli *CoreGoerliTransactor) UpdateRewards(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "updateRewards", poolId, collateralType, accountId)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_CoreGoerli *CoreGoerliSession) UpdateRewards(poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.UpdateRewards(&_CoreGoerli.TransactOpts, poolId, collateralType, accountId)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_CoreGoerli *CoreGoerliTransactorSession) UpdateRewards(poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.UpdateRewards(&_CoreGoerli.TransactOpts, poolId, collateralType, accountId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_CoreGoerli *CoreGoerliTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_CoreGoerli *CoreGoerliSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.UpgradeTo(&_CoreGoerli.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _CoreGoerli.Contract.UpgradeTo(&_CoreGoerli.TransactOpts, newImplementation)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliTransactor) Withdraw(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "withdraw", accountId, collateralType, tokenAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliSession) Withdraw(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.Withdraw(&_CoreGoerli.TransactOpts, accountId, collateralType, tokenAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) Withdraw(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.Withdraw(&_CoreGoerli.TransactOpts, accountId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliTransactor) WithdrawMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "withdrawMarketCollateral", marketId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliSession) WithdrawMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.WithdrawMarketCollateral(&_CoreGoerli.TransactOpts, marketId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_CoreGoerli *CoreGoerliTransactorSession) WithdrawMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.WithdrawMarketCollateral(&_CoreGoerli.TransactOpts, marketId, collateralType, tokenAmount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreGoerli *CoreGoerliTransactor) WithdrawMarketUsd(opts *bind.TransactOpts, marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.contract.Transact(opts, "withdrawMarketUsd", marketId, target, amount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreGoerli *CoreGoerliSession) WithdrawMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.WithdrawMarketUsd(&_CoreGoerli.TransactOpts, marketId, target, amount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_CoreGoerli *CoreGoerliTransactorSession) WithdrawMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoreGoerli.Contract.WithdrawMarketUsd(&_CoreGoerli.TransactOpts, marketId, target, amount)
}

// CoreGoerliAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the CoreGoerli contract.
type CoreGoerliAccountCreatedIterator struct {
	Event *CoreGoerliAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliAccountCreated represents a AccountCreated event raised by the CoreGoerli contract.
type CoreGoerliAccountCreated struct {
	AccountId *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) FilterAccountCreated(opts *bind.FilterOpts, accountId []*big.Int, owner []common.Address) (*CoreGoerliAccountCreatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliAccountCreatedIterator{contract: _CoreGoerli.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xa9e04d307e860938fa63307df8b8090e365276e59fcca12ed55656c25e538019.
//
// Solidity: event AccountCreated(uint128 indexed accountId, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *CoreGoerliAccountCreated, accountId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "AccountCreated", accountIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliAccountCreated)
				if err := _CoreGoerli.contract.UnpackLog(event, "AccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseAccountCreated(log types.Log) (*CoreGoerliAccountCreated, error) {
	event := new(CoreGoerliAccountCreated)
	if err := _CoreGoerli.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliAssociatedSystemSetIterator is returned from FilterAssociatedSystemSet and is used to iterate over the raw logs and unpacked data for AssociatedSystemSet events raised by the CoreGoerli contract.
type CoreGoerliAssociatedSystemSetIterator struct {
	Event *CoreGoerliAssociatedSystemSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliAssociatedSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliAssociatedSystemSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliAssociatedSystemSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliAssociatedSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliAssociatedSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliAssociatedSystemSet represents a AssociatedSystemSet event raised by the CoreGoerli contract.
type CoreGoerliAssociatedSystemSet struct {
	Kind  [32]byte
	Id    [32]byte
	Proxy common.Address
	Impl  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssociatedSystemSet is a free log retrieval operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_CoreGoerli *CoreGoerliFilterer) FilterAssociatedSystemSet(opts *bind.FilterOpts, kind [][32]byte, id [][32]byte) (*CoreGoerliAssociatedSystemSetIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliAssociatedSystemSetIterator{contract: _CoreGoerli.contract, event: "AssociatedSystemSet", logs: logs, sub: sub}, nil
}

// WatchAssociatedSystemSet is a free log subscription operation binding the contract event 0xc8551a5a03a7b06d5d20159b3b8839429a7aefab4bf3d020f1b65fa903ccb3d2.
//
// Solidity: event AssociatedSystemSet(bytes32 indexed kind, bytes32 indexed id, address proxy, address impl)
func (_CoreGoerli *CoreGoerliFilterer) WatchAssociatedSystemSet(opts *bind.WatchOpts, sink chan<- *CoreGoerliAssociatedSystemSet, kind [][32]byte, id [][32]byte) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "AssociatedSystemSet", kindRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliAssociatedSystemSet)
				if err := _CoreGoerli.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseAssociatedSystemSet(log types.Log) (*CoreGoerliAssociatedSystemSet, error) {
	event := new(CoreGoerliAssociatedSystemSet)
	if err := _CoreGoerli.contract.UnpackLog(event, "AssociatedSystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliCollateralConfiguredIterator is returned from FilterCollateralConfigured and is used to iterate over the raw logs and unpacked data for CollateralConfigured events raised by the CoreGoerli contract.
type CoreGoerliCollateralConfiguredIterator struct {
	Event *CoreGoerliCollateralConfigured // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliCollateralConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliCollateralConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliCollateralConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliCollateralConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliCollateralConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliCollateralConfigured represents a CollateralConfigured event raised by the CoreGoerli contract.
type CoreGoerliCollateralConfigured struct {
	CollateralType common.Address
	Config         CollateralConfigurationData
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCollateralConfigured is a free log retrieval operation binding the contract event 0xefc23317f58afd6b22480bd22174cc7da0913bce25c03d9859216dacddebe6fe.
//
// Solidity: event CollateralConfigured(address indexed collateralType, (bool,uint256,uint256,uint256,bytes32,address,uint256) config)
func (_CoreGoerli *CoreGoerliFilterer) FilterCollateralConfigured(opts *bind.FilterOpts, collateralType []common.Address) (*CoreGoerliCollateralConfiguredIterator, error) {

	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "CollateralConfigured", collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliCollateralConfiguredIterator{contract: _CoreGoerli.contract, event: "CollateralConfigured", logs: logs, sub: sub}, nil
}

// WatchCollateralConfigured is a free log subscription operation binding the contract event 0xefc23317f58afd6b22480bd22174cc7da0913bce25c03d9859216dacddebe6fe.
//
// Solidity: event CollateralConfigured(address indexed collateralType, (bool,uint256,uint256,uint256,bytes32,address,uint256) config)
func (_CoreGoerli *CoreGoerliFilterer) WatchCollateralConfigured(opts *bind.WatchOpts, sink chan<- *CoreGoerliCollateralConfigured, collateralType []common.Address) (event.Subscription, error) {

	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "CollateralConfigured", collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliCollateralConfigured)
				if err := _CoreGoerli.contract.UnpackLog(event, "CollateralConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseCollateralConfigured(log types.Log) (*CoreGoerliCollateralConfigured, error) {
	event := new(CoreGoerliCollateralConfigured)
	if err := _CoreGoerli.contract.UnpackLog(event, "CollateralConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliCollateralLockCreatedIterator is returned from FilterCollateralLockCreated and is used to iterate over the raw logs and unpacked data for CollateralLockCreated events raised by the CoreGoerli contract.
type CoreGoerliCollateralLockCreatedIterator struct {
	Event *CoreGoerliCollateralLockCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliCollateralLockCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliCollateralLockCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliCollateralLockCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliCollateralLockCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliCollateralLockCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliCollateralLockCreated represents a CollateralLockCreated event raised by the CoreGoerli contract.
type CoreGoerliCollateralLockCreated struct {
	AccountId       *big.Int
	CollateralType  common.Address
	TokenAmount     *big.Int
	ExpireTimestamp uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollateralLockCreated is a free log retrieval operation binding the contract event 0x8a78446a234220d3ee3f46aa6ea1ea5bc438bd153398ebcbd171843744b452a6.
//
// Solidity: event CollateralLockCreated(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_CoreGoerli *CoreGoerliFilterer) FilterCollateralLockCreated(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address) (*CoreGoerliCollateralLockCreatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "CollateralLockCreated", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliCollateralLockCreatedIterator{contract: _CoreGoerli.contract, event: "CollateralLockCreated", logs: logs, sub: sub}, nil
}

// WatchCollateralLockCreated is a free log subscription operation binding the contract event 0x8a78446a234220d3ee3f46aa6ea1ea5bc438bd153398ebcbd171843744b452a6.
//
// Solidity: event CollateralLockCreated(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_CoreGoerli *CoreGoerliFilterer) WatchCollateralLockCreated(opts *bind.WatchOpts, sink chan<- *CoreGoerliCollateralLockCreated, accountId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "CollateralLockCreated", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliCollateralLockCreated)
				if err := _CoreGoerli.contract.UnpackLog(event, "CollateralLockCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseCollateralLockCreated(log types.Log) (*CoreGoerliCollateralLockCreated, error) {
	event := new(CoreGoerliCollateralLockCreated)
	if err := _CoreGoerli.contract.UnpackLog(event, "CollateralLockCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliCollateralLockExpiredIterator is returned from FilterCollateralLockExpired and is used to iterate over the raw logs and unpacked data for CollateralLockExpired events raised by the CoreGoerli contract.
type CoreGoerliCollateralLockExpiredIterator struct {
	Event *CoreGoerliCollateralLockExpired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliCollateralLockExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliCollateralLockExpired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliCollateralLockExpired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliCollateralLockExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliCollateralLockExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliCollateralLockExpired represents a CollateralLockExpired event raised by the CoreGoerli contract.
type CoreGoerliCollateralLockExpired struct {
	AccountId       *big.Int
	CollateralType  common.Address
	TokenAmount     *big.Int
	ExpireTimestamp uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollateralLockExpired is a free log retrieval operation binding the contract event 0xc5e2c7afc4e54d998977e11260a0bfc0ad5678a3a8b6628162f9d4e642d7f160.
//
// Solidity: event CollateralLockExpired(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_CoreGoerli *CoreGoerliFilterer) FilterCollateralLockExpired(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address) (*CoreGoerliCollateralLockExpiredIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "CollateralLockExpired", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliCollateralLockExpiredIterator{contract: _CoreGoerli.contract, event: "CollateralLockExpired", logs: logs, sub: sub}, nil
}

// WatchCollateralLockExpired is a free log subscription operation binding the contract event 0xc5e2c7afc4e54d998977e11260a0bfc0ad5678a3a8b6628162f9d4e642d7f160.
//
// Solidity: event CollateralLockExpired(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_CoreGoerli *CoreGoerliFilterer) WatchCollateralLockExpired(opts *bind.WatchOpts, sink chan<- *CoreGoerliCollateralLockExpired, accountId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "CollateralLockExpired", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliCollateralLockExpired)
				if err := _CoreGoerli.contract.UnpackLog(event, "CollateralLockExpired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseCollateralLockExpired(log types.Log) (*CoreGoerliCollateralLockExpired, error) {
	event := new(CoreGoerliCollateralLockExpired)
	if err := _CoreGoerli.contract.UnpackLog(event, "CollateralLockExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliDebtAssociatedIterator is returned from FilterDebtAssociated and is used to iterate over the raw logs and unpacked data for DebtAssociated events raised by the CoreGoerli contract.
type CoreGoerliDebtAssociatedIterator struct {
	Event *CoreGoerliDebtAssociated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliDebtAssociatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliDebtAssociated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliDebtAssociated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliDebtAssociatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliDebtAssociatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliDebtAssociated represents a DebtAssociated event raised by the CoreGoerli contract.
type CoreGoerliDebtAssociated struct {
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
func (_CoreGoerli *CoreGoerliFilterer) FilterDebtAssociated(opts *bind.FilterOpts, marketId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*CoreGoerliDebtAssociatedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "DebtAssociated", marketIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliDebtAssociatedIterator{contract: _CoreGoerli.contract, event: "DebtAssociated", logs: logs, sub: sub}, nil
}

// WatchDebtAssociated is a free log subscription operation binding the contract event 0xb03bc7530b5650601d2c8bc86c81a45b6b50b0099defd3f17a6bdb48660180ad.
//
// Solidity: event DebtAssociated(uint128 indexed marketId, uint128 indexed poolId, address indexed collateralType, uint128 accountId, uint256 amount, int256 updatedDebt)
func (_CoreGoerli *CoreGoerliFilterer) WatchDebtAssociated(opts *bind.WatchOpts, sink chan<- *CoreGoerliDebtAssociated, marketId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "DebtAssociated", marketIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliDebtAssociated)
				if err := _CoreGoerli.contract.UnpackLog(event, "DebtAssociated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseDebtAssociated(log types.Log) (*CoreGoerliDebtAssociated, error) {
	event := new(CoreGoerliDebtAssociated)
	if err := _CoreGoerli.contract.UnpackLog(event, "DebtAssociated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliDelegationUpdatedIterator is returned from FilterDelegationUpdated and is used to iterate over the raw logs and unpacked data for DelegationUpdated events raised by the CoreGoerli contract.
type CoreGoerliDelegationUpdatedIterator struct {
	Event *CoreGoerliDelegationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliDelegationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliDelegationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliDelegationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliDelegationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliDelegationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliDelegationUpdated represents a DelegationUpdated event raised by the CoreGoerli contract.
type CoreGoerliDelegationUpdated struct {
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
func (_CoreGoerli *CoreGoerliFilterer) FilterDelegationUpdated(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*CoreGoerliDelegationUpdatedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "DelegationUpdated", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliDelegationUpdatedIterator{contract: _CoreGoerli.contract, event: "DelegationUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegationUpdated is a free log subscription operation binding the contract event 0x7b12dd38f18c0ff77ae702f6da13fbbcb28f53f807ecc7d39ee8d8b1ea8295ad.
//
// Solidity: event DelegationUpdated(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, uint256 leverage, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchDelegationUpdated(opts *bind.WatchOpts, sink chan<- *CoreGoerliDelegationUpdated, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "DelegationUpdated", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliDelegationUpdated)
				if err := _CoreGoerli.contract.UnpackLog(event, "DelegationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseDelegationUpdated(log types.Log) (*CoreGoerliDelegationUpdated, error) {
	event := new(CoreGoerliDelegationUpdated)
	if err := _CoreGoerli.contract.UnpackLog(event, "DelegationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the CoreGoerli contract.
type CoreGoerliDepositedIterator struct {
	Event *CoreGoerliDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliDeposited represents a Deposited event raised by the CoreGoerli contract.
type CoreGoerliDeposited struct {
	AccountId      *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xd92122e67326e9313bfae33ccb1fccf5194584c6bf93a8529a6b006d8c6e24a9.
//
// Solidity: event Deposited(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) FilterDeposited(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreGoerliDepositedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "Deposited", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliDepositedIterator{contract: _CoreGoerli.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xd92122e67326e9313bfae33ccb1fccf5194584c6bf93a8529a6b006d8c6e24a9.
//
// Solidity: event Deposited(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *CoreGoerliDeposited, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "Deposited", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliDeposited)
				if err := _CoreGoerli.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseDeposited(log types.Log) (*CoreGoerliDeposited, error) {
	event := new(CoreGoerliDeposited)
	if err := _CoreGoerli.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliFeatureFlagAllowAllSetIterator is returned from FilterFeatureFlagAllowAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowAllSet events raised by the CoreGoerli contract.
type CoreGoerliFeatureFlagAllowAllSetIterator struct {
	Event *CoreGoerliFeatureFlagAllowAllSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliFeatureFlagAllowAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliFeatureFlagAllowAllSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliFeatureFlagAllowAllSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliFeatureFlagAllowAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliFeatureFlagAllowAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliFeatureFlagAllowAllSet represents a FeatureFlagAllowAllSet event raised by the CoreGoerli contract.
type CoreGoerliFeatureFlagAllowAllSet struct {
	Feature  [32]byte
	AllowAll bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowAllSet is a free log retrieval operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_CoreGoerli *CoreGoerliFilterer) FilterFeatureFlagAllowAllSet(opts *bind.FilterOpts, feature [][32]byte) (*CoreGoerliFeatureFlagAllowAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliFeatureFlagAllowAllSetIterator{contract: _CoreGoerli.contract, event: "FeatureFlagAllowAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowAllSet is a free log subscription operation binding the contract event 0xa806035d8c8de7cd43725250d3fbf9ee7abe3b99ffb892897913d8a21721121d.
//
// Solidity: event FeatureFlagAllowAllSet(bytes32 indexed feature, bool allowAll)
func (_CoreGoerli *CoreGoerliFilterer) WatchFeatureFlagAllowAllSet(opts *bind.WatchOpts, sink chan<- *CoreGoerliFeatureFlagAllowAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "FeatureFlagAllowAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliFeatureFlagAllowAllSet)
				if err := _CoreGoerli.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseFeatureFlagAllowAllSet(log types.Log) (*CoreGoerliFeatureFlagAllowAllSet, error) {
	event := new(CoreGoerliFeatureFlagAllowAllSet)
	if err := _CoreGoerli.contract.UnpackLog(event, "FeatureFlagAllowAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliFeatureFlagAllowlistAddedIterator is returned from FilterFeatureFlagAllowlistAdded and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistAdded events raised by the CoreGoerli contract.
type CoreGoerliFeatureFlagAllowlistAddedIterator struct {
	Event *CoreGoerliFeatureFlagAllowlistAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliFeatureFlagAllowlistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliFeatureFlagAllowlistAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliFeatureFlagAllowlistAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliFeatureFlagAllowlistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliFeatureFlagAllowlistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliFeatureFlagAllowlistAdded represents a FeatureFlagAllowlistAdded event raised by the CoreGoerli contract.
type CoreGoerliFeatureFlagAllowlistAdded struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistAdded is a free log retrieval operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_CoreGoerli *CoreGoerliFilterer) FilterFeatureFlagAllowlistAdded(opts *bind.FilterOpts, feature [][32]byte) (*CoreGoerliFeatureFlagAllowlistAddedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliFeatureFlagAllowlistAddedIterator{contract: _CoreGoerli.contract, event: "FeatureFlagAllowlistAdded", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistAdded is a free log subscription operation binding the contract event 0x30b9b4104e2fb00b4f980e414dcd828e691c8fcb286f0c73d7267c3a2de49383.
//
// Solidity: event FeatureFlagAllowlistAdded(bytes32 indexed feature, address account)
func (_CoreGoerli *CoreGoerliFilterer) WatchFeatureFlagAllowlistAdded(opts *bind.WatchOpts, sink chan<- *CoreGoerliFeatureFlagAllowlistAdded, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "FeatureFlagAllowlistAdded", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliFeatureFlagAllowlistAdded)
				if err := _CoreGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseFeatureFlagAllowlistAdded(log types.Log) (*CoreGoerliFeatureFlagAllowlistAdded, error) {
	event := new(CoreGoerliFeatureFlagAllowlistAdded)
	if err := _CoreGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliFeatureFlagAllowlistRemovedIterator is returned from FilterFeatureFlagAllowlistRemoved and is used to iterate over the raw logs and unpacked data for FeatureFlagAllowlistRemoved events raised by the CoreGoerli contract.
type CoreGoerliFeatureFlagAllowlistRemovedIterator struct {
	Event *CoreGoerliFeatureFlagAllowlistRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliFeatureFlagAllowlistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliFeatureFlagAllowlistRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliFeatureFlagAllowlistRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliFeatureFlagAllowlistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliFeatureFlagAllowlistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliFeatureFlagAllowlistRemoved represents a FeatureFlagAllowlistRemoved event raised by the CoreGoerli contract.
type CoreGoerliFeatureFlagAllowlistRemoved struct {
	Feature [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagAllowlistRemoved is a free log retrieval operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_CoreGoerli *CoreGoerliFilterer) FilterFeatureFlagAllowlistRemoved(opts *bind.FilterOpts, feature [][32]byte) (*CoreGoerliFeatureFlagAllowlistRemovedIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliFeatureFlagAllowlistRemovedIterator{contract: _CoreGoerli.contract, event: "FeatureFlagAllowlistRemoved", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagAllowlistRemoved is a free log subscription operation binding the contract event 0xb44a47e11880cc865e8ea382561e406dea8c895366c58e3908f05708b2880890.
//
// Solidity: event FeatureFlagAllowlistRemoved(bytes32 indexed feature, address account)
func (_CoreGoerli *CoreGoerliFilterer) WatchFeatureFlagAllowlistRemoved(opts *bind.WatchOpts, sink chan<- *CoreGoerliFeatureFlagAllowlistRemoved, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "FeatureFlagAllowlistRemoved", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliFeatureFlagAllowlistRemoved)
				if err := _CoreGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseFeatureFlagAllowlistRemoved(log types.Log) (*CoreGoerliFeatureFlagAllowlistRemoved, error) {
	event := new(CoreGoerliFeatureFlagAllowlistRemoved)
	if err := _CoreGoerli.contract.UnpackLog(event, "FeatureFlagAllowlistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliFeatureFlagDeniersResetIterator is returned from FilterFeatureFlagDeniersReset and is used to iterate over the raw logs and unpacked data for FeatureFlagDeniersReset events raised by the CoreGoerli contract.
type CoreGoerliFeatureFlagDeniersResetIterator struct {
	Event *CoreGoerliFeatureFlagDeniersReset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliFeatureFlagDeniersResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliFeatureFlagDeniersReset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliFeatureFlagDeniersReset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliFeatureFlagDeniersResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliFeatureFlagDeniersResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliFeatureFlagDeniersReset represents a FeatureFlagDeniersReset event raised by the CoreGoerli contract.
type CoreGoerliFeatureFlagDeniersReset struct {
	Feature [32]byte
	Deniers []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDeniersReset is a free log retrieval operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_CoreGoerli *CoreGoerliFilterer) FilterFeatureFlagDeniersReset(opts *bind.FilterOpts, feature [][32]byte) (*CoreGoerliFeatureFlagDeniersResetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliFeatureFlagDeniersResetIterator{contract: _CoreGoerli.contract, event: "FeatureFlagDeniersReset", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDeniersReset is a free log subscription operation binding the contract event 0x74d48d0b51a70680130c00decd06b4d536fbb3cee16a3b0bdd2309c264dcbd13.
//
// Solidity: event FeatureFlagDeniersReset(bytes32 indexed feature, address[] deniers)
func (_CoreGoerli *CoreGoerliFilterer) WatchFeatureFlagDeniersReset(opts *bind.WatchOpts, sink chan<- *CoreGoerliFeatureFlagDeniersReset, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "FeatureFlagDeniersReset", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliFeatureFlagDeniersReset)
				if err := _CoreGoerli.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseFeatureFlagDeniersReset(log types.Log) (*CoreGoerliFeatureFlagDeniersReset, error) {
	event := new(CoreGoerliFeatureFlagDeniersReset)
	if err := _CoreGoerli.contract.UnpackLog(event, "FeatureFlagDeniersReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliFeatureFlagDenyAllSetIterator is returned from FilterFeatureFlagDenyAllSet and is used to iterate over the raw logs and unpacked data for FeatureFlagDenyAllSet events raised by the CoreGoerli contract.
type CoreGoerliFeatureFlagDenyAllSetIterator struct {
	Event *CoreGoerliFeatureFlagDenyAllSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliFeatureFlagDenyAllSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliFeatureFlagDenyAllSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliFeatureFlagDenyAllSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliFeatureFlagDenyAllSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliFeatureFlagDenyAllSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliFeatureFlagDenyAllSet represents a FeatureFlagDenyAllSet event raised by the CoreGoerli contract.
type CoreGoerliFeatureFlagDenyAllSet struct {
	Feature [32]byte
	DenyAll bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeatureFlagDenyAllSet is a free log retrieval operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_CoreGoerli *CoreGoerliFilterer) FilterFeatureFlagDenyAllSet(opts *bind.FilterOpts, feature [][32]byte) (*CoreGoerliFeatureFlagDenyAllSetIterator, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliFeatureFlagDenyAllSetIterator{contract: _CoreGoerli.contract, event: "FeatureFlagDenyAllSet", logs: logs, sub: sub}, nil
}

// WatchFeatureFlagDenyAllSet is a free log subscription operation binding the contract event 0x97f76d2e384948e28ddd4280a4e76e8600acc328a0c0910c93682a0fccc02018.
//
// Solidity: event FeatureFlagDenyAllSet(bytes32 indexed feature, bool denyAll)
func (_CoreGoerli *CoreGoerliFilterer) WatchFeatureFlagDenyAllSet(opts *bind.WatchOpts, sink chan<- *CoreGoerliFeatureFlagDenyAllSet, feature [][32]byte) (event.Subscription, error) {

	var featureRule []interface{}
	for _, featureItem := range feature {
		featureRule = append(featureRule, featureItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "FeatureFlagDenyAllSet", featureRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliFeatureFlagDenyAllSet)
				if err := _CoreGoerli.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseFeatureFlagDenyAllSet(log types.Log) (*CoreGoerliFeatureFlagDenyAllSet, error) {
	event := new(CoreGoerliFeatureFlagDenyAllSet)
	if err := _CoreGoerli.contract.UnpackLog(event, "FeatureFlagDenyAllSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliIssuanceFeePaidIterator is returned from FilterIssuanceFeePaid and is used to iterate over the raw logs and unpacked data for IssuanceFeePaid events raised by the CoreGoerli contract.
type CoreGoerliIssuanceFeePaidIterator struct {
	Event *CoreGoerliIssuanceFeePaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliIssuanceFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliIssuanceFeePaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliIssuanceFeePaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliIssuanceFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliIssuanceFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliIssuanceFeePaid represents a IssuanceFeePaid event raised by the CoreGoerli contract.
type CoreGoerliIssuanceFeePaid struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	FeeAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIssuanceFeePaid is a free log retrieval operation binding the contract event 0x28d0fb10e1c8ce51490a16fb3b40bf23f8064f7c624d3a3852ad66683a25995d.
//
// Solidity: event IssuanceFeePaid(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 feeAmount)
func (_CoreGoerli *CoreGoerliFilterer) FilterIssuanceFeePaid(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int) (*CoreGoerliIssuanceFeePaidIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "IssuanceFeePaid", accountIdRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliIssuanceFeePaidIterator{contract: _CoreGoerli.contract, event: "IssuanceFeePaid", logs: logs, sub: sub}, nil
}

// WatchIssuanceFeePaid is a free log subscription operation binding the contract event 0x28d0fb10e1c8ce51490a16fb3b40bf23f8064f7c624d3a3852ad66683a25995d.
//
// Solidity: event IssuanceFeePaid(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 feeAmount)
func (_CoreGoerli *CoreGoerliFilterer) WatchIssuanceFeePaid(opts *bind.WatchOpts, sink chan<- *CoreGoerliIssuanceFeePaid, accountId []*big.Int, poolId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "IssuanceFeePaid", accountIdRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliIssuanceFeePaid)
				if err := _CoreGoerli.contract.UnpackLog(event, "IssuanceFeePaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseIssuanceFeePaid(log types.Log) (*CoreGoerliIssuanceFeePaid, error) {
	event := new(CoreGoerliIssuanceFeePaid)
	if err := _CoreGoerli.contract.UnpackLog(event, "IssuanceFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliLiquidationIterator is returned from FilterLiquidation and is used to iterate over the raw logs and unpacked data for Liquidation events raised by the CoreGoerli contract.
type CoreGoerliLiquidationIterator struct {
	Event *CoreGoerliLiquidation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliLiquidation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliLiquidation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliLiquidation represents a Liquidation event raised by the CoreGoerli contract.
type CoreGoerliLiquidation struct {
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
func (_CoreGoerli *CoreGoerliFilterer) FilterLiquidation(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*CoreGoerliLiquidationIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "Liquidation", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliLiquidationIterator{contract: _CoreGoerli.contract, event: "Liquidation", logs: logs, sub: sub}, nil
}

// WatchLiquidation is a free log subscription operation binding the contract event 0xe6c1b26644f880854bf954d4186be9f0b2d06d50fa0484b596e79d409c07a5fd.
//
// Solidity: event Liquidation(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchLiquidation(opts *bind.WatchOpts, sink chan<- *CoreGoerliLiquidation, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "Liquidation", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliLiquidation)
				if err := _CoreGoerli.contract.UnpackLog(event, "Liquidation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseLiquidation(log types.Log) (*CoreGoerliLiquidation, error) {
	event := new(CoreGoerliLiquidation)
	if err := _CoreGoerli.contract.UnpackLog(event, "Liquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliMarketCollateralDepositedIterator is returned from FilterMarketCollateralDeposited and is used to iterate over the raw logs and unpacked data for MarketCollateralDeposited events raised by the CoreGoerli contract.
type CoreGoerliMarketCollateralDepositedIterator struct {
	Event *CoreGoerliMarketCollateralDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliMarketCollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliMarketCollateralDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliMarketCollateralDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliMarketCollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliMarketCollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliMarketCollateralDeposited represents a MarketCollateralDeposited event raised by the CoreGoerli contract.
type CoreGoerliMarketCollateralDeposited struct {
	MarketId       *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketCollateralDeposited is a free log retrieval operation binding the contract event 0x6affa35c1a0ba6198bd74578c93e3a0eb32a96af0af26ce8adf06b933212ca1b.
//
// Solidity: event MarketCollateralDeposited(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) FilterMarketCollateralDeposited(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreGoerliMarketCollateralDepositedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "MarketCollateralDeposited", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliMarketCollateralDepositedIterator{contract: _CoreGoerli.contract, event: "MarketCollateralDeposited", logs: logs, sub: sub}, nil
}

// WatchMarketCollateralDeposited is a free log subscription operation binding the contract event 0x6affa35c1a0ba6198bd74578c93e3a0eb32a96af0af26ce8adf06b933212ca1b.
//
// Solidity: event MarketCollateralDeposited(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchMarketCollateralDeposited(opts *bind.WatchOpts, sink chan<- *CoreGoerliMarketCollateralDeposited, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "MarketCollateralDeposited", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliMarketCollateralDeposited)
				if err := _CoreGoerli.contract.UnpackLog(event, "MarketCollateralDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseMarketCollateralDeposited(log types.Log) (*CoreGoerliMarketCollateralDeposited, error) {
	event := new(CoreGoerliMarketCollateralDeposited)
	if err := _CoreGoerli.contract.UnpackLog(event, "MarketCollateralDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliMarketCollateralWithdrawnIterator is returned from FilterMarketCollateralWithdrawn and is used to iterate over the raw logs and unpacked data for MarketCollateralWithdrawn events raised by the CoreGoerli contract.
type CoreGoerliMarketCollateralWithdrawnIterator struct {
	Event *CoreGoerliMarketCollateralWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliMarketCollateralWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliMarketCollateralWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliMarketCollateralWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliMarketCollateralWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliMarketCollateralWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliMarketCollateralWithdrawn represents a MarketCollateralWithdrawn event raised by the CoreGoerli contract.
type CoreGoerliMarketCollateralWithdrawn struct {
	MarketId       *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketCollateralWithdrawn is a free log retrieval operation binding the contract event 0x31fd88ca973dd7ee814313c50c2aaa7cd2d2142634582332b1e2671e6d97fd9a.
//
// Solidity: event MarketCollateralWithdrawn(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) FilterMarketCollateralWithdrawn(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreGoerliMarketCollateralWithdrawnIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "MarketCollateralWithdrawn", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliMarketCollateralWithdrawnIterator{contract: _CoreGoerli.contract, event: "MarketCollateralWithdrawn", logs: logs, sub: sub}, nil
}

// WatchMarketCollateralWithdrawn is a free log subscription operation binding the contract event 0x31fd88ca973dd7ee814313c50c2aaa7cd2d2142634582332b1e2671e6d97fd9a.
//
// Solidity: event MarketCollateralWithdrawn(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchMarketCollateralWithdrawn(opts *bind.WatchOpts, sink chan<- *CoreGoerliMarketCollateralWithdrawn, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "MarketCollateralWithdrawn", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliMarketCollateralWithdrawn)
				if err := _CoreGoerli.contract.UnpackLog(event, "MarketCollateralWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseMarketCollateralWithdrawn(log types.Log) (*CoreGoerliMarketCollateralWithdrawn, error) {
	event := new(CoreGoerliMarketCollateralWithdrawn)
	if err := _CoreGoerli.contract.UnpackLog(event, "MarketCollateralWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliMarketRegisteredIterator is returned from FilterMarketRegistered and is used to iterate over the raw logs and unpacked data for MarketRegistered events raised by the CoreGoerli contract.
type CoreGoerliMarketRegisteredIterator struct {
	Event *CoreGoerliMarketRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliMarketRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliMarketRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliMarketRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliMarketRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliMarketRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliMarketRegistered represents a MarketRegistered event raised by the CoreGoerli contract.
type CoreGoerliMarketRegistered struct {
	Market   common.Address
	MarketId *big.Int
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketRegistered is a free log retrieval operation binding the contract event 0xeb87361ace8c1947e121293eb214f68d889d9cf273c48246b38c3cbf185748d0.
//
// Solidity: event MarketRegistered(address indexed market, uint128 indexed marketId, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) FilterMarketRegistered(opts *bind.FilterOpts, market []common.Address, marketId []*big.Int, sender []common.Address) (*CoreGoerliMarketRegisteredIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "MarketRegistered", marketRule, marketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliMarketRegisteredIterator{contract: _CoreGoerli.contract, event: "MarketRegistered", logs: logs, sub: sub}, nil
}

// WatchMarketRegistered is a free log subscription operation binding the contract event 0xeb87361ace8c1947e121293eb214f68d889d9cf273c48246b38c3cbf185748d0.
//
// Solidity: event MarketRegistered(address indexed market, uint128 indexed marketId, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchMarketRegistered(opts *bind.WatchOpts, sink chan<- *CoreGoerliMarketRegistered, market []common.Address, marketId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "MarketRegistered", marketRule, marketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliMarketRegistered)
				if err := _CoreGoerli.contract.UnpackLog(event, "MarketRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseMarketRegistered(log types.Log) (*CoreGoerliMarketRegistered, error) {
	event := new(CoreGoerliMarketRegistered)
	if err := _CoreGoerli.contract.UnpackLog(event, "MarketRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliMarketSystemFeePaidIterator is returned from FilterMarketSystemFeePaid and is used to iterate over the raw logs and unpacked data for MarketSystemFeePaid events raised by the CoreGoerli contract.
type CoreGoerliMarketSystemFeePaidIterator struct {
	Event *CoreGoerliMarketSystemFeePaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliMarketSystemFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliMarketSystemFeePaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliMarketSystemFeePaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliMarketSystemFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliMarketSystemFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliMarketSystemFeePaid represents a MarketSystemFeePaid event raised by the CoreGoerli contract.
type CoreGoerliMarketSystemFeePaid struct {
	MarketId  *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMarketSystemFeePaid is a free log retrieval operation binding the contract event 0x8b69fed8aed97ef9572216662359ece45fa52f2b5ff44a78b7ec3c5ef05153f8.
//
// Solidity: event MarketSystemFeePaid(uint128 indexed marketId, uint256 feeAmount)
func (_CoreGoerli *CoreGoerliFilterer) FilterMarketSystemFeePaid(opts *bind.FilterOpts, marketId []*big.Int) (*CoreGoerliMarketSystemFeePaidIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "MarketSystemFeePaid", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliMarketSystemFeePaidIterator{contract: _CoreGoerli.contract, event: "MarketSystemFeePaid", logs: logs, sub: sub}, nil
}

// WatchMarketSystemFeePaid is a free log subscription operation binding the contract event 0x8b69fed8aed97ef9572216662359ece45fa52f2b5ff44a78b7ec3c5ef05153f8.
//
// Solidity: event MarketSystemFeePaid(uint128 indexed marketId, uint256 feeAmount)
func (_CoreGoerli *CoreGoerliFilterer) WatchMarketSystemFeePaid(opts *bind.WatchOpts, sink chan<- *CoreGoerliMarketSystemFeePaid, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "MarketSystemFeePaid", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliMarketSystemFeePaid)
				if err := _CoreGoerli.contract.UnpackLog(event, "MarketSystemFeePaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseMarketSystemFeePaid(log types.Log) (*CoreGoerliMarketSystemFeePaid, error) {
	event := new(CoreGoerliMarketSystemFeePaid)
	if err := _CoreGoerli.contract.UnpackLog(event, "MarketSystemFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliMarketUsdDepositedIterator is returned from FilterMarketUsdDeposited and is used to iterate over the raw logs and unpacked data for MarketUsdDeposited events raised by the CoreGoerli contract.
type CoreGoerliMarketUsdDepositedIterator struct {
	Event *CoreGoerliMarketUsdDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliMarketUsdDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliMarketUsdDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliMarketUsdDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliMarketUsdDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliMarketUsdDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliMarketUsdDeposited represents a MarketUsdDeposited event raised by the CoreGoerli contract.
type CoreGoerliMarketUsdDeposited struct {
	MarketId *big.Int
	Target   common.Address
	Amount   *big.Int
	Market   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketUsdDeposited is a free log retrieval operation binding the contract event 0xcbdbec7c74466ca0e764de5f6bca5bef859e31a7d6284be22ac6b4beabddf835.
//
// Solidity: event MarketUsdDeposited(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_CoreGoerli *CoreGoerliFilterer) FilterMarketUsdDeposited(opts *bind.FilterOpts, marketId []*big.Int, target []common.Address, market []common.Address) (*CoreGoerliMarketUsdDepositedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "MarketUsdDeposited", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliMarketUsdDepositedIterator{contract: _CoreGoerli.contract, event: "MarketUsdDeposited", logs: logs, sub: sub}, nil
}

// WatchMarketUsdDeposited is a free log subscription operation binding the contract event 0xcbdbec7c74466ca0e764de5f6bca5bef859e31a7d6284be22ac6b4beabddf835.
//
// Solidity: event MarketUsdDeposited(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_CoreGoerli *CoreGoerliFilterer) WatchMarketUsdDeposited(opts *bind.WatchOpts, sink chan<- *CoreGoerliMarketUsdDeposited, marketId []*big.Int, target []common.Address, market []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "MarketUsdDeposited", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliMarketUsdDeposited)
				if err := _CoreGoerli.contract.UnpackLog(event, "MarketUsdDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseMarketUsdDeposited(log types.Log) (*CoreGoerliMarketUsdDeposited, error) {
	event := new(CoreGoerliMarketUsdDeposited)
	if err := _CoreGoerli.contract.UnpackLog(event, "MarketUsdDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliMarketUsdWithdrawnIterator is returned from FilterMarketUsdWithdrawn and is used to iterate over the raw logs and unpacked data for MarketUsdWithdrawn events raised by the CoreGoerli contract.
type CoreGoerliMarketUsdWithdrawnIterator struct {
	Event *CoreGoerliMarketUsdWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliMarketUsdWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliMarketUsdWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliMarketUsdWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliMarketUsdWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliMarketUsdWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliMarketUsdWithdrawn represents a MarketUsdWithdrawn event raised by the CoreGoerli contract.
type CoreGoerliMarketUsdWithdrawn struct {
	MarketId *big.Int
	Target   common.Address
	Amount   *big.Int
	Market   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketUsdWithdrawn is a free log retrieval operation binding the contract event 0x8669e4526579b1f1da479be2f96d1c7953fff439a4f98362a1c9ca07fe42ba41.
//
// Solidity: event MarketUsdWithdrawn(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_CoreGoerli *CoreGoerliFilterer) FilterMarketUsdWithdrawn(opts *bind.FilterOpts, marketId []*big.Int, target []common.Address, market []common.Address) (*CoreGoerliMarketUsdWithdrawnIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "MarketUsdWithdrawn", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliMarketUsdWithdrawnIterator{contract: _CoreGoerli.contract, event: "MarketUsdWithdrawn", logs: logs, sub: sub}, nil
}

// WatchMarketUsdWithdrawn is a free log subscription operation binding the contract event 0x8669e4526579b1f1da479be2f96d1c7953fff439a4f98362a1c9ca07fe42ba41.
//
// Solidity: event MarketUsdWithdrawn(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_CoreGoerli *CoreGoerliFilterer) WatchMarketUsdWithdrawn(opts *bind.WatchOpts, sink chan<- *CoreGoerliMarketUsdWithdrawn, marketId []*big.Int, target []common.Address, market []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "MarketUsdWithdrawn", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliMarketUsdWithdrawn)
				if err := _CoreGoerli.contract.UnpackLog(event, "MarketUsdWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseMarketUsdWithdrawn(log types.Log) (*CoreGoerliMarketUsdWithdrawn, error) {
	event := new(CoreGoerliMarketUsdWithdrawn)
	if err := _CoreGoerli.contract.UnpackLog(event, "MarketUsdWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliMaximumMarketCollateralConfiguredIterator is returned from FilterMaximumMarketCollateralConfigured and is used to iterate over the raw logs and unpacked data for MaximumMarketCollateralConfigured events raised by the CoreGoerli contract.
type CoreGoerliMaximumMarketCollateralConfiguredIterator struct {
	Event *CoreGoerliMaximumMarketCollateralConfigured // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliMaximumMarketCollateralConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliMaximumMarketCollateralConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliMaximumMarketCollateralConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliMaximumMarketCollateralConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliMaximumMarketCollateralConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliMaximumMarketCollateralConfigured represents a MaximumMarketCollateralConfigured event raised by the CoreGoerli contract.
type CoreGoerliMaximumMarketCollateralConfigured struct {
	MarketId       *big.Int
	CollateralType common.Address
	SystemAmount   *big.Int
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMaximumMarketCollateralConfigured is a free log retrieval operation binding the contract event 0x499c8fcfbc4341c37dcf444c890d42ef888d46aa586f97ceb20577c2635c8075.
//
// Solidity: event MaximumMarketCollateralConfigured(uint128 indexed marketId, address indexed collateralType, uint256 systemAmount, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) FilterMaximumMarketCollateralConfigured(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, owner []common.Address) (*CoreGoerliMaximumMarketCollateralConfiguredIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "MaximumMarketCollateralConfigured", marketIdRule, collateralTypeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliMaximumMarketCollateralConfiguredIterator{contract: _CoreGoerli.contract, event: "MaximumMarketCollateralConfigured", logs: logs, sub: sub}, nil
}

// WatchMaximumMarketCollateralConfigured is a free log subscription operation binding the contract event 0x499c8fcfbc4341c37dcf444c890d42ef888d46aa586f97ceb20577c2635c8075.
//
// Solidity: event MaximumMarketCollateralConfigured(uint128 indexed marketId, address indexed collateralType, uint256 systemAmount, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) WatchMaximumMarketCollateralConfigured(opts *bind.WatchOpts, sink chan<- *CoreGoerliMaximumMarketCollateralConfigured, marketId []*big.Int, collateralType []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "MaximumMarketCollateralConfigured", marketIdRule, collateralTypeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliMaximumMarketCollateralConfigured)
				if err := _CoreGoerli.contract.UnpackLog(event, "MaximumMarketCollateralConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseMaximumMarketCollateralConfigured(log types.Log) (*CoreGoerliMaximumMarketCollateralConfigured, error) {
	event := new(CoreGoerliMaximumMarketCollateralConfigured)
	if err := _CoreGoerli.contract.UnpackLog(event, "MaximumMarketCollateralConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliNewSupportedCrossChainNetworkIterator is returned from FilterNewSupportedCrossChainNetwork and is used to iterate over the raw logs and unpacked data for NewSupportedCrossChainNetwork events raised by the CoreGoerli contract.
type CoreGoerliNewSupportedCrossChainNetworkIterator struct {
	Event *CoreGoerliNewSupportedCrossChainNetwork // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliNewSupportedCrossChainNetworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliNewSupportedCrossChainNetwork)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliNewSupportedCrossChainNetwork)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliNewSupportedCrossChainNetworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliNewSupportedCrossChainNetworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliNewSupportedCrossChainNetwork represents a NewSupportedCrossChainNetwork event raised by the CoreGoerli contract.
type CoreGoerliNewSupportedCrossChainNetwork struct {
	NewChainId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewSupportedCrossChainNetwork is a free log retrieval operation binding the contract event 0x1874eb2a5288e478dcedf1d33291bd7293eeef5946ec516d2ef54a364b3f63d8.
//
// Solidity: event NewSupportedCrossChainNetwork(uint64 newChainId)
func (_CoreGoerli *CoreGoerliFilterer) FilterNewSupportedCrossChainNetwork(opts *bind.FilterOpts) (*CoreGoerliNewSupportedCrossChainNetworkIterator, error) {

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "NewSupportedCrossChainNetwork")
	if err != nil {
		return nil, err
	}
	return &CoreGoerliNewSupportedCrossChainNetworkIterator{contract: _CoreGoerli.contract, event: "NewSupportedCrossChainNetwork", logs: logs, sub: sub}, nil
}

// WatchNewSupportedCrossChainNetwork is a free log subscription operation binding the contract event 0x1874eb2a5288e478dcedf1d33291bd7293eeef5946ec516d2ef54a364b3f63d8.
//
// Solidity: event NewSupportedCrossChainNetwork(uint64 newChainId)
func (_CoreGoerli *CoreGoerliFilterer) WatchNewSupportedCrossChainNetwork(opts *bind.WatchOpts, sink chan<- *CoreGoerliNewSupportedCrossChainNetwork) (event.Subscription, error) {

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "NewSupportedCrossChainNetwork")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliNewSupportedCrossChainNetwork)
				if err := _CoreGoerli.contract.UnpackLog(event, "NewSupportedCrossChainNetwork", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseNewSupportedCrossChainNetwork(log types.Log) (*CoreGoerliNewSupportedCrossChainNetwork, error) {
	event := new(CoreGoerliNewSupportedCrossChainNetwork)
	if err := _CoreGoerli.contract.UnpackLog(event, "NewSupportedCrossChainNetwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the CoreGoerli contract.
type CoreGoerliOwnerChangedIterator struct {
	Event *CoreGoerliOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliOwnerChanged represents a OwnerChanged event raised by the CoreGoerli contract.
type CoreGoerliOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_CoreGoerli *CoreGoerliFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*CoreGoerliOwnerChangedIterator, error) {

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &CoreGoerliOwnerChangedIterator{contract: _CoreGoerli.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_CoreGoerli *CoreGoerliFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *CoreGoerliOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliOwnerChanged)
				if err := _CoreGoerli.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseOwnerChanged(log types.Log) (*CoreGoerliOwnerChanged, error) {
	event := new(CoreGoerliOwnerChanged)
	if err := _CoreGoerli.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the CoreGoerli contract.
type CoreGoerliOwnerNominatedIterator struct {
	Event *CoreGoerliOwnerNominated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliOwnerNominated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliOwnerNominated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliOwnerNominated represents a OwnerNominated event raised by the CoreGoerli contract.
type CoreGoerliOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_CoreGoerli *CoreGoerliFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*CoreGoerliOwnerNominatedIterator, error) {

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &CoreGoerliOwnerNominatedIterator{contract: _CoreGoerli.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_CoreGoerli *CoreGoerliFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *CoreGoerliOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliOwnerNominated)
				if err := _CoreGoerli.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseOwnerNominated(log types.Log) (*CoreGoerliOwnerNominated, error) {
	event := new(CoreGoerliOwnerNominated)
	if err := _CoreGoerli.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPermissionGrantedIterator is returned from FilterPermissionGranted and is used to iterate over the raw logs and unpacked data for PermissionGranted events raised by the CoreGoerli contract.
type CoreGoerliPermissionGrantedIterator struct {
	Event *CoreGoerliPermissionGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPermissionGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPermissionGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPermissionGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPermissionGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPermissionGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPermissionGranted represents a PermissionGranted event raised by the CoreGoerli contract.
type CoreGoerliPermissionGranted struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionGranted is a free log retrieval operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_CoreGoerli *CoreGoerliFilterer) FilterPermissionGranted(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*CoreGoerliPermissionGrantedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPermissionGrantedIterator{contract: _CoreGoerli.contract, event: "PermissionGranted", logs: logs, sub: sub}, nil
}

// WatchPermissionGranted is a free log subscription operation binding the contract event 0x32ff7c3f84299a3543c1e89057e98ba962f4fbe7786c52289e184c57b9a36a50.
//
// Solidity: event PermissionGranted(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchPermissionGranted(opts *bind.WatchOpts, sink chan<- *CoreGoerliPermissionGranted, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PermissionGranted", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPermissionGranted)
				if err := _CoreGoerli.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePermissionGranted(log types.Log) (*CoreGoerliPermissionGranted, error) {
	event := new(CoreGoerliPermissionGranted)
	if err := _CoreGoerli.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPermissionRevokedIterator is returned from FilterPermissionRevoked and is used to iterate over the raw logs and unpacked data for PermissionRevoked events raised by the CoreGoerli contract.
type CoreGoerliPermissionRevokedIterator struct {
	Event *CoreGoerliPermissionRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPermissionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPermissionRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPermissionRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPermissionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPermissionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPermissionRevoked represents a PermissionRevoked event raised by the CoreGoerli contract.
type CoreGoerliPermissionRevoked struct {
	AccountId  *big.Int
	Permission [32]byte
	User       common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionRevoked is a free log retrieval operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_CoreGoerli *CoreGoerliFilterer) FilterPermissionRevoked(opts *bind.FilterOpts, accountId []*big.Int, permission [][32]byte, user []common.Address) (*CoreGoerliPermissionRevokedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPermissionRevokedIterator{contract: _CoreGoerli.contract, event: "PermissionRevoked", logs: logs, sub: sub}, nil
}

// WatchPermissionRevoked is a free log subscription operation binding the contract event 0x116c7e9cd2db316974fb473babcbcd625be1350842d0319e761d30aefb09a58a.
//
// Solidity: event PermissionRevoked(uint128 indexed accountId, bytes32 indexed permission, address indexed user, address sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchPermissionRevoked(opts *bind.WatchOpts, sink chan<- *CoreGoerliPermissionRevoked, accountId []*big.Int, permission [][32]byte, user []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PermissionRevoked", accountIdRule, permissionRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPermissionRevoked)
				if err := _CoreGoerli.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePermissionRevoked(log types.Log) (*CoreGoerliPermissionRevoked, error) {
	event := new(CoreGoerliPermissionRevoked)
	if err := _CoreGoerli.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPoolApprovedAddedIterator is returned from FilterPoolApprovedAdded and is used to iterate over the raw logs and unpacked data for PoolApprovedAdded events raised by the CoreGoerli contract.
type CoreGoerliPoolApprovedAddedIterator struct {
	Event *CoreGoerliPoolApprovedAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPoolApprovedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPoolApprovedAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPoolApprovedAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPoolApprovedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPoolApprovedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPoolApprovedAdded represents a PoolApprovedAdded event raised by the CoreGoerli contract.
type CoreGoerliPoolApprovedAdded struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolApprovedAdded is a free log retrieval operation binding the contract event 0x7d5bdf4e8c44e0b5a8249bf03c2a1febd848cc7f580efd7b1703301c5b1a9e4e.
//
// Solidity: event PoolApprovedAdded(uint256 poolId)
func (_CoreGoerli *CoreGoerliFilterer) FilterPoolApprovedAdded(opts *bind.FilterOpts) (*CoreGoerliPoolApprovedAddedIterator, error) {

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PoolApprovedAdded")
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPoolApprovedAddedIterator{contract: _CoreGoerli.contract, event: "PoolApprovedAdded", logs: logs, sub: sub}, nil
}

// WatchPoolApprovedAdded is a free log subscription operation binding the contract event 0x7d5bdf4e8c44e0b5a8249bf03c2a1febd848cc7f580efd7b1703301c5b1a9e4e.
//
// Solidity: event PoolApprovedAdded(uint256 poolId)
func (_CoreGoerli *CoreGoerliFilterer) WatchPoolApprovedAdded(opts *bind.WatchOpts, sink chan<- *CoreGoerliPoolApprovedAdded) (event.Subscription, error) {

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PoolApprovedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPoolApprovedAdded)
				if err := _CoreGoerli.contract.UnpackLog(event, "PoolApprovedAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePoolApprovedAdded(log types.Log) (*CoreGoerliPoolApprovedAdded, error) {
	event := new(CoreGoerliPoolApprovedAdded)
	if err := _CoreGoerli.contract.UnpackLog(event, "PoolApprovedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPoolApprovedRemovedIterator is returned from FilterPoolApprovedRemoved and is used to iterate over the raw logs and unpacked data for PoolApprovedRemoved events raised by the CoreGoerli contract.
type CoreGoerliPoolApprovedRemovedIterator struct {
	Event *CoreGoerliPoolApprovedRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPoolApprovedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPoolApprovedRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPoolApprovedRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPoolApprovedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPoolApprovedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPoolApprovedRemoved represents a PoolApprovedRemoved event raised by the CoreGoerli contract.
type CoreGoerliPoolApprovedRemoved struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolApprovedRemoved is a free log retrieval operation binding the contract event 0xc1567ee9983f306f073ea7d59a7fb5680ce07985f8b49cc50d00a3a9f748d3c2.
//
// Solidity: event PoolApprovedRemoved(uint256 poolId)
func (_CoreGoerli *CoreGoerliFilterer) FilterPoolApprovedRemoved(opts *bind.FilterOpts) (*CoreGoerliPoolApprovedRemovedIterator, error) {

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PoolApprovedRemoved")
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPoolApprovedRemovedIterator{contract: _CoreGoerli.contract, event: "PoolApprovedRemoved", logs: logs, sub: sub}, nil
}

// WatchPoolApprovedRemoved is a free log subscription operation binding the contract event 0xc1567ee9983f306f073ea7d59a7fb5680ce07985f8b49cc50d00a3a9f748d3c2.
//
// Solidity: event PoolApprovedRemoved(uint256 poolId)
func (_CoreGoerli *CoreGoerliFilterer) WatchPoolApprovedRemoved(opts *bind.WatchOpts, sink chan<- *CoreGoerliPoolApprovedRemoved) (event.Subscription, error) {

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PoolApprovedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPoolApprovedRemoved)
				if err := _CoreGoerli.contract.UnpackLog(event, "PoolApprovedRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePoolApprovedRemoved(log types.Log) (*CoreGoerliPoolApprovedRemoved, error) {
	event := new(CoreGoerliPoolApprovedRemoved)
	if err := _CoreGoerli.contract.UnpackLog(event, "PoolApprovedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPoolConfigurationSetIterator is returned from FilterPoolConfigurationSet and is used to iterate over the raw logs and unpacked data for PoolConfigurationSet events raised by the CoreGoerli contract.
type CoreGoerliPoolConfigurationSetIterator struct {
	Event *CoreGoerliPoolConfigurationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPoolConfigurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPoolConfigurationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPoolConfigurationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPoolConfigurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPoolConfigurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPoolConfigurationSet represents a PoolConfigurationSet event raised by the CoreGoerli contract.
type CoreGoerliPoolConfigurationSet struct {
	PoolId  *big.Int
	Markets []MarketConfigurationData
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolConfigurationSet is a free log retrieval operation binding the contract event 0xdd812c2e47943d98e6c66b2b9872d1f9270b8523c82eb60ad5c8d580a614081c.
//
// Solidity: event PoolConfigurationSet(uint128 indexed poolId, (uint128,uint128,int128)[] markets, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) FilterPoolConfigurationSet(opts *bind.FilterOpts, poolId []*big.Int, sender []common.Address) (*CoreGoerliPoolConfigurationSetIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PoolConfigurationSet", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPoolConfigurationSetIterator{contract: _CoreGoerli.contract, event: "PoolConfigurationSet", logs: logs, sub: sub}, nil
}

// WatchPoolConfigurationSet is a free log subscription operation binding the contract event 0xdd812c2e47943d98e6c66b2b9872d1f9270b8523c82eb60ad5c8d580a614081c.
//
// Solidity: event PoolConfigurationSet(uint128 indexed poolId, (uint128,uint128,int128)[] markets, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchPoolConfigurationSet(opts *bind.WatchOpts, sink chan<- *CoreGoerliPoolConfigurationSet, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PoolConfigurationSet", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPoolConfigurationSet)
				if err := _CoreGoerli.contract.UnpackLog(event, "PoolConfigurationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePoolConfigurationSet(log types.Log) (*CoreGoerliPoolConfigurationSet, error) {
	event := new(CoreGoerliPoolConfigurationSet)
	if err := _CoreGoerli.contract.UnpackLog(event, "PoolConfigurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the CoreGoerli contract.
type CoreGoerliPoolCreatedIterator struct {
	Event *CoreGoerliPoolCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPoolCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPoolCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPoolCreated represents a PoolCreated event raised by the CoreGoerli contract.
type CoreGoerliPoolCreated struct {
	PoolId *big.Int
	Owner  common.Address
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xb1517ad708e5f9a104c30d3f1ff749d55833b1d03bf472013c29888e741cf340.
//
// Solidity: event PoolCreated(uint128 indexed poolId, address indexed owner, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) FilterPoolCreated(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address, sender []common.Address) (*CoreGoerliPoolCreatedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PoolCreated", poolIdRule, ownerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPoolCreatedIterator{contract: _CoreGoerli.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xb1517ad708e5f9a104c30d3f1ff749d55833b1d03bf472013c29888e741cf340.
//
// Solidity: event PoolCreated(uint128 indexed poolId, address indexed owner, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *CoreGoerliPoolCreated, poolId []*big.Int, owner []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PoolCreated", poolIdRule, ownerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPoolCreated)
				if err := _CoreGoerli.contract.UnpackLog(event, "PoolCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePoolCreated(log types.Log) (*CoreGoerliPoolCreated, error) {
	event := new(CoreGoerliPoolCreated)
	if err := _CoreGoerli.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPoolNameUpdatedIterator is returned from FilterPoolNameUpdated and is used to iterate over the raw logs and unpacked data for PoolNameUpdated events raised by the CoreGoerli contract.
type CoreGoerliPoolNameUpdatedIterator struct {
	Event *CoreGoerliPoolNameUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPoolNameUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPoolNameUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPoolNameUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPoolNameUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPoolNameUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPoolNameUpdated represents a PoolNameUpdated event raised by the CoreGoerli contract.
type CoreGoerliPoolNameUpdated struct {
	PoolId *big.Int
	Name   string
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNameUpdated is a free log retrieval operation binding the contract event 0x63b42abaf7e145a993f20bc64259f45d09c43d18838ab0bca078b15093ac55f4.
//
// Solidity: event PoolNameUpdated(uint128 indexed poolId, string name, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) FilterPoolNameUpdated(opts *bind.FilterOpts, poolId []*big.Int, sender []common.Address) (*CoreGoerliPoolNameUpdatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PoolNameUpdated", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPoolNameUpdatedIterator{contract: _CoreGoerli.contract, event: "PoolNameUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolNameUpdated is a free log subscription operation binding the contract event 0x63b42abaf7e145a993f20bc64259f45d09c43d18838ab0bca078b15093ac55f4.
//
// Solidity: event PoolNameUpdated(uint128 indexed poolId, string name, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchPoolNameUpdated(opts *bind.WatchOpts, sink chan<- *CoreGoerliPoolNameUpdated, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PoolNameUpdated", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPoolNameUpdated)
				if err := _CoreGoerli.contract.UnpackLog(event, "PoolNameUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePoolNameUpdated(log types.Log) (*CoreGoerliPoolNameUpdated, error) {
	event := new(CoreGoerliPoolNameUpdated)
	if err := _CoreGoerli.contract.UnpackLog(event, "PoolNameUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPoolNominationRenouncedIterator is returned from FilterPoolNominationRenounced and is used to iterate over the raw logs and unpacked data for PoolNominationRenounced events raised by the CoreGoerli contract.
type CoreGoerliPoolNominationRenouncedIterator struct {
	Event *CoreGoerliPoolNominationRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPoolNominationRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPoolNominationRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPoolNominationRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPoolNominationRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPoolNominationRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPoolNominationRenounced represents a PoolNominationRenounced event raised by the CoreGoerli contract.
type CoreGoerliPoolNominationRenounced struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNominationRenounced is a free log retrieval operation binding the contract event 0x28301da5fb0feefb138efa6310af4547a74f415d62616f90519436dc169c3ae0.
//
// Solidity: event PoolNominationRenounced(uint128 indexed poolId, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) FilterPoolNominationRenounced(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*CoreGoerliPoolNominationRenouncedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PoolNominationRenounced", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPoolNominationRenouncedIterator{contract: _CoreGoerli.contract, event: "PoolNominationRenounced", logs: logs, sub: sub}, nil
}

// WatchPoolNominationRenounced is a free log subscription operation binding the contract event 0x28301da5fb0feefb138efa6310af4547a74f415d62616f90519436dc169c3ae0.
//
// Solidity: event PoolNominationRenounced(uint128 indexed poolId, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) WatchPoolNominationRenounced(opts *bind.WatchOpts, sink chan<- *CoreGoerliPoolNominationRenounced, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PoolNominationRenounced", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPoolNominationRenounced)
				if err := _CoreGoerli.contract.UnpackLog(event, "PoolNominationRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePoolNominationRenounced(log types.Log) (*CoreGoerliPoolNominationRenounced, error) {
	event := new(CoreGoerliPoolNominationRenounced)
	if err := _CoreGoerli.contract.UnpackLog(event, "PoolNominationRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPoolNominationRevokedIterator is returned from FilterPoolNominationRevoked and is used to iterate over the raw logs and unpacked data for PoolNominationRevoked events raised by the CoreGoerli contract.
type CoreGoerliPoolNominationRevokedIterator struct {
	Event *CoreGoerliPoolNominationRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPoolNominationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPoolNominationRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPoolNominationRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPoolNominationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPoolNominationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPoolNominationRevoked represents a PoolNominationRevoked event raised by the CoreGoerli contract.
type CoreGoerliPoolNominationRevoked struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNominationRevoked is a free log retrieval operation binding the contract event 0xa20a605599b6da4a06e0662f1284c442a576bc452b77a38c8c55805cb82a1865.
//
// Solidity: event PoolNominationRevoked(uint128 indexed poolId, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) FilterPoolNominationRevoked(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*CoreGoerliPoolNominationRevokedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PoolNominationRevoked", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPoolNominationRevokedIterator{contract: _CoreGoerli.contract, event: "PoolNominationRevoked", logs: logs, sub: sub}, nil
}

// WatchPoolNominationRevoked is a free log subscription operation binding the contract event 0xa20a605599b6da4a06e0662f1284c442a576bc452b77a38c8c55805cb82a1865.
//
// Solidity: event PoolNominationRevoked(uint128 indexed poolId, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) WatchPoolNominationRevoked(opts *bind.WatchOpts, sink chan<- *CoreGoerliPoolNominationRevoked, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PoolNominationRevoked", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPoolNominationRevoked)
				if err := _CoreGoerli.contract.UnpackLog(event, "PoolNominationRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePoolNominationRevoked(log types.Log) (*CoreGoerliPoolNominationRevoked, error) {
	event := new(CoreGoerliPoolNominationRevoked)
	if err := _CoreGoerli.contract.UnpackLog(event, "PoolNominationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPoolOwnerNominatedIterator is returned from FilterPoolOwnerNominated and is used to iterate over the raw logs and unpacked data for PoolOwnerNominated events raised by the CoreGoerli contract.
type CoreGoerliPoolOwnerNominatedIterator struct {
	Event *CoreGoerliPoolOwnerNominated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPoolOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPoolOwnerNominated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPoolOwnerNominated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPoolOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPoolOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPoolOwnerNominated represents a PoolOwnerNominated event raised by the CoreGoerli contract.
type CoreGoerliPoolOwnerNominated struct {
	PoolId         *big.Int
	NominatedOwner common.Address
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolOwnerNominated is a free log retrieval operation binding the contract event 0x55d98f82a53fb5776e9ea48d624ab9cb015b51a45249b1ed8425fc857c82f4f8.
//
// Solidity: event PoolOwnerNominated(uint128 indexed poolId, address indexed nominatedOwner, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) FilterPoolOwnerNominated(opts *bind.FilterOpts, poolId []*big.Int, nominatedOwner []common.Address, owner []common.Address) (*CoreGoerliPoolOwnerNominatedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PoolOwnerNominated", poolIdRule, nominatedOwnerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPoolOwnerNominatedIterator{contract: _CoreGoerli.contract, event: "PoolOwnerNominated", logs: logs, sub: sub}, nil
}

// WatchPoolOwnerNominated is a free log subscription operation binding the contract event 0x55d98f82a53fb5776e9ea48d624ab9cb015b51a45249b1ed8425fc857c82f4f8.
//
// Solidity: event PoolOwnerNominated(uint128 indexed poolId, address indexed nominatedOwner, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) WatchPoolOwnerNominated(opts *bind.WatchOpts, sink chan<- *CoreGoerliPoolOwnerNominated, poolId []*big.Int, nominatedOwner []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PoolOwnerNominated", poolIdRule, nominatedOwnerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPoolOwnerNominated)
				if err := _CoreGoerli.contract.UnpackLog(event, "PoolOwnerNominated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePoolOwnerNominated(log types.Log) (*CoreGoerliPoolOwnerNominated, error) {
	event := new(CoreGoerliPoolOwnerNominated)
	if err := _CoreGoerli.contract.UnpackLog(event, "PoolOwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPoolOwnershipAcceptedIterator is returned from FilterPoolOwnershipAccepted and is used to iterate over the raw logs and unpacked data for PoolOwnershipAccepted events raised by the CoreGoerli contract.
type CoreGoerliPoolOwnershipAcceptedIterator struct {
	Event *CoreGoerliPoolOwnershipAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPoolOwnershipAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPoolOwnershipAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPoolOwnershipAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPoolOwnershipAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPoolOwnershipAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPoolOwnershipAccepted represents a PoolOwnershipAccepted event raised by the CoreGoerli contract.
type CoreGoerliPoolOwnershipAccepted struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolOwnershipAccepted is a free log retrieval operation binding the contract event 0x4f86f2ce8b08e27d0e470f4269b71c3bbc68407d51a2e692f6573236074ebc5a.
//
// Solidity: event PoolOwnershipAccepted(uint128 indexed poolId, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) FilterPoolOwnershipAccepted(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*CoreGoerliPoolOwnershipAcceptedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PoolOwnershipAccepted", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPoolOwnershipAcceptedIterator{contract: _CoreGoerli.contract, event: "PoolOwnershipAccepted", logs: logs, sub: sub}, nil
}

// WatchPoolOwnershipAccepted is a free log subscription operation binding the contract event 0x4f86f2ce8b08e27d0e470f4269b71c3bbc68407d51a2e692f6573236074ebc5a.
//
// Solidity: event PoolOwnershipAccepted(uint128 indexed poolId, address indexed owner)
func (_CoreGoerli *CoreGoerliFilterer) WatchPoolOwnershipAccepted(opts *bind.WatchOpts, sink chan<- *CoreGoerliPoolOwnershipAccepted, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PoolOwnershipAccepted", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPoolOwnershipAccepted)
				if err := _CoreGoerli.contract.UnpackLog(event, "PoolOwnershipAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePoolOwnershipAccepted(log types.Log) (*CoreGoerliPoolOwnershipAccepted, error) {
	event := new(CoreGoerliPoolOwnershipAccepted)
	if err := _CoreGoerli.contract.UnpackLog(event, "PoolOwnershipAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliPreferredPoolSetIterator is returned from FilterPreferredPoolSet and is used to iterate over the raw logs and unpacked data for PreferredPoolSet events raised by the CoreGoerli contract.
type CoreGoerliPreferredPoolSetIterator struct {
	Event *CoreGoerliPreferredPoolSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliPreferredPoolSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliPreferredPoolSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliPreferredPoolSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliPreferredPoolSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliPreferredPoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliPreferredPoolSet represents a PreferredPoolSet event raised by the CoreGoerli contract.
type CoreGoerliPreferredPoolSet struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPreferredPoolSet is a free log retrieval operation binding the contract event 0x7e7cb4726e710dc12fad41f158c37a4a71af3a6f053b8b13670d35c710139a56.
//
// Solidity: event PreferredPoolSet(uint256 poolId)
func (_CoreGoerli *CoreGoerliFilterer) FilterPreferredPoolSet(opts *bind.FilterOpts) (*CoreGoerliPreferredPoolSetIterator, error) {

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "PreferredPoolSet")
	if err != nil {
		return nil, err
	}
	return &CoreGoerliPreferredPoolSetIterator{contract: _CoreGoerli.contract, event: "PreferredPoolSet", logs: logs, sub: sub}, nil
}

// WatchPreferredPoolSet is a free log subscription operation binding the contract event 0x7e7cb4726e710dc12fad41f158c37a4a71af3a6f053b8b13670d35c710139a56.
//
// Solidity: event PreferredPoolSet(uint256 poolId)
func (_CoreGoerli *CoreGoerliFilterer) WatchPreferredPoolSet(opts *bind.WatchOpts, sink chan<- *CoreGoerliPreferredPoolSet) (event.Subscription, error) {

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "PreferredPoolSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliPreferredPoolSet)
				if err := _CoreGoerli.contract.UnpackLog(event, "PreferredPoolSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParsePreferredPoolSet(log types.Log) (*CoreGoerliPreferredPoolSet, error) {
	event := new(CoreGoerliPreferredPoolSet)
	if err := _CoreGoerli.contract.UnpackLog(event, "PreferredPoolSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the CoreGoerli contract.
type CoreGoerliRewardsClaimedIterator struct {
	Event *CoreGoerliRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliRewardsClaimed represents a RewardsClaimed event raised by the CoreGoerli contract.
type CoreGoerliRewardsClaimed struct {
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
func (_CoreGoerli *CoreGoerliFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*CoreGoerliRewardsClaimedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "RewardsClaimed", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliRewardsClaimedIterator{contract: _CoreGoerli.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xa4a60be4203e7975e54ab5314c7e9e18aba9ad71e8da714d8de987f4f05410f2.
//
// Solidity: event RewardsClaimed(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount)
func (_CoreGoerli *CoreGoerliFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *CoreGoerliRewardsClaimed, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "RewardsClaimed", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliRewardsClaimed)
				if err := _CoreGoerli.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseRewardsClaimed(log types.Log) (*CoreGoerliRewardsClaimed, error) {
	event := new(CoreGoerliRewardsClaimed)
	if err := _CoreGoerli.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliRewardsDistributedIterator is returned from FilterRewardsDistributed and is used to iterate over the raw logs and unpacked data for RewardsDistributed events raised by the CoreGoerli contract.
type CoreGoerliRewardsDistributedIterator struct {
	Event *CoreGoerliRewardsDistributed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliRewardsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliRewardsDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliRewardsDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliRewardsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliRewardsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliRewardsDistributed represents a RewardsDistributed event raised by the CoreGoerli contract.
type CoreGoerliRewardsDistributed struct {
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
func (_CoreGoerli *CoreGoerliFilterer) FilterRewardsDistributed(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address) (*CoreGoerliRewardsDistributedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "RewardsDistributed", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliRewardsDistributedIterator{contract: _CoreGoerli.contract, event: "RewardsDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributed is a free log subscription operation binding the contract event 0x19ced31d71d1db45f99d5a8d3a7616fe9d78828df58f2a28feb68c9f9ab876ca.
//
// Solidity: event RewardsDistributed(uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount, uint256 start, uint256 duration)
func (_CoreGoerli *CoreGoerliFilterer) WatchRewardsDistributed(opts *bind.WatchOpts, sink chan<- *CoreGoerliRewardsDistributed, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "RewardsDistributed", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliRewardsDistributed)
				if err := _CoreGoerli.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseRewardsDistributed(log types.Log) (*CoreGoerliRewardsDistributed, error) {
	event := new(CoreGoerliRewardsDistributed)
	if err := _CoreGoerli.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliRewardsDistributorRegisteredIterator is returned from FilterRewardsDistributorRegistered and is used to iterate over the raw logs and unpacked data for RewardsDistributorRegistered events raised by the CoreGoerli contract.
type CoreGoerliRewardsDistributorRegisteredIterator struct {
	Event *CoreGoerliRewardsDistributorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliRewardsDistributorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliRewardsDistributorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliRewardsDistributorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliRewardsDistributorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliRewardsDistributorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliRewardsDistributorRegistered represents a RewardsDistributorRegistered event raised by the CoreGoerli contract.
type CoreGoerliRewardsDistributorRegistered struct {
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributorRegistered is a free log retrieval operation binding the contract event 0x9d3609c05a83dc93a5b355d62c2b37dfde8f0833b1184d4d05c6f51cd46b6e5b.
//
// Solidity: event RewardsDistributorRegistered(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_CoreGoerli *CoreGoerliFilterer) FilterRewardsDistributorRegistered(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (*CoreGoerliRewardsDistributorRegisteredIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "RewardsDistributorRegistered", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliRewardsDistributorRegisteredIterator{contract: _CoreGoerli.contract, event: "RewardsDistributorRegistered", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributorRegistered is a free log subscription operation binding the contract event 0x9d3609c05a83dc93a5b355d62c2b37dfde8f0833b1184d4d05c6f51cd46b6e5b.
//
// Solidity: event RewardsDistributorRegistered(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_CoreGoerli *CoreGoerliFilterer) WatchRewardsDistributorRegistered(opts *bind.WatchOpts, sink chan<- *CoreGoerliRewardsDistributorRegistered, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "RewardsDistributorRegistered", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliRewardsDistributorRegistered)
				if err := _CoreGoerli.contract.UnpackLog(event, "RewardsDistributorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseRewardsDistributorRegistered(log types.Log) (*CoreGoerliRewardsDistributorRegistered, error) {
	event := new(CoreGoerliRewardsDistributorRegistered)
	if err := _CoreGoerli.contract.UnpackLog(event, "RewardsDistributorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliRewardsDistributorRemovedIterator is returned from FilterRewardsDistributorRemoved and is used to iterate over the raw logs and unpacked data for RewardsDistributorRemoved events raised by the CoreGoerli contract.
type CoreGoerliRewardsDistributorRemovedIterator struct {
	Event *CoreGoerliRewardsDistributorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliRewardsDistributorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliRewardsDistributorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliRewardsDistributorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliRewardsDistributorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliRewardsDistributorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliRewardsDistributorRemoved represents a RewardsDistributorRemoved event raised by the CoreGoerli contract.
type CoreGoerliRewardsDistributorRemoved struct {
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributorRemoved is a free log retrieval operation binding the contract event 0x375c4507f463c55a506be95e2cfd3cfdc0610be055087eac6049588a1bcfacba.
//
// Solidity: event RewardsDistributorRemoved(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_CoreGoerli *CoreGoerliFilterer) FilterRewardsDistributorRemoved(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (*CoreGoerliRewardsDistributorRemovedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "RewardsDistributorRemoved", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliRewardsDistributorRemovedIterator{contract: _CoreGoerli.contract, event: "RewardsDistributorRemoved", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributorRemoved is a free log subscription operation binding the contract event 0x375c4507f463c55a506be95e2cfd3cfdc0610be055087eac6049588a1bcfacba.
//
// Solidity: event RewardsDistributorRemoved(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_CoreGoerli *CoreGoerliFilterer) WatchRewardsDistributorRemoved(opts *bind.WatchOpts, sink chan<- *CoreGoerliRewardsDistributorRemoved, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "RewardsDistributorRemoved", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliRewardsDistributorRemoved)
				if err := _CoreGoerli.contract.UnpackLog(event, "RewardsDistributorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseRewardsDistributorRemoved(log types.Log) (*CoreGoerliRewardsDistributorRemoved, error) {
	event := new(CoreGoerliRewardsDistributorRemoved)
	if err := _CoreGoerli.contract.UnpackLog(event, "RewardsDistributorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliSetMarketMinLiquidityRatioIterator is returned from FilterSetMarketMinLiquidityRatio and is used to iterate over the raw logs and unpacked data for SetMarketMinLiquidityRatio events raised by the CoreGoerli contract.
type CoreGoerliSetMarketMinLiquidityRatioIterator struct {
	Event *CoreGoerliSetMarketMinLiquidityRatio // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliSetMarketMinLiquidityRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliSetMarketMinLiquidityRatio)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliSetMarketMinLiquidityRatio)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliSetMarketMinLiquidityRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliSetMarketMinLiquidityRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliSetMarketMinLiquidityRatio represents a SetMarketMinLiquidityRatio event raised by the CoreGoerli contract.
type CoreGoerliSetMarketMinLiquidityRatio struct {
	MarketId          *big.Int
	MinLiquidityRatio *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMarketMinLiquidityRatio is a free log retrieval operation binding the contract event 0x563eb723f21b3e87ec8932cfb4ffa64d1b68c42053c28d6b4db019a40f6daf47.
//
// Solidity: event SetMarketMinLiquidityRatio(uint128 indexed marketId, uint256 minLiquidityRatio)
func (_CoreGoerli *CoreGoerliFilterer) FilterSetMarketMinLiquidityRatio(opts *bind.FilterOpts, marketId []*big.Int) (*CoreGoerliSetMarketMinLiquidityRatioIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "SetMarketMinLiquidityRatio", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliSetMarketMinLiquidityRatioIterator{contract: _CoreGoerli.contract, event: "SetMarketMinLiquidityRatio", logs: logs, sub: sub}, nil
}

// WatchSetMarketMinLiquidityRatio is a free log subscription operation binding the contract event 0x563eb723f21b3e87ec8932cfb4ffa64d1b68c42053c28d6b4db019a40f6daf47.
//
// Solidity: event SetMarketMinLiquidityRatio(uint128 indexed marketId, uint256 minLiquidityRatio)
func (_CoreGoerli *CoreGoerliFilterer) WatchSetMarketMinLiquidityRatio(opts *bind.WatchOpts, sink chan<- *CoreGoerliSetMarketMinLiquidityRatio, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "SetMarketMinLiquidityRatio", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliSetMarketMinLiquidityRatio)
				if err := _CoreGoerli.contract.UnpackLog(event, "SetMarketMinLiquidityRatio", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseSetMarketMinLiquidityRatio(log types.Log) (*CoreGoerliSetMarketMinLiquidityRatio, error) {
	event := new(CoreGoerliSetMarketMinLiquidityRatio)
	if err := _CoreGoerli.contract.UnpackLog(event, "SetMarketMinLiquidityRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliSetMinDelegateTimeIterator is returned from FilterSetMinDelegateTime and is used to iterate over the raw logs and unpacked data for SetMinDelegateTime events raised by the CoreGoerli contract.
type CoreGoerliSetMinDelegateTimeIterator struct {
	Event *CoreGoerliSetMinDelegateTime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliSetMinDelegateTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliSetMinDelegateTime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliSetMinDelegateTime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliSetMinDelegateTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliSetMinDelegateTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliSetMinDelegateTime represents a SetMinDelegateTime event raised by the CoreGoerli contract.
type CoreGoerliSetMinDelegateTime struct {
	MarketId        *big.Int
	MinDelegateTime uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetMinDelegateTime is a free log retrieval operation binding the contract event 0x6942a68d151863c1fed3c0c4c5f3258af738218527147ac69290ab23ca7d26c6.
//
// Solidity: event SetMinDelegateTime(uint128 indexed marketId, uint32 minDelegateTime)
func (_CoreGoerli *CoreGoerliFilterer) FilterSetMinDelegateTime(opts *bind.FilterOpts, marketId []*big.Int) (*CoreGoerliSetMinDelegateTimeIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "SetMinDelegateTime", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliSetMinDelegateTimeIterator{contract: _CoreGoerli.contract, event: "SetMinDelegateTime", logs: logs, sub: sub}, nil
}

// WatchSetMinDelegateTime is a free log subscription operation binding the contract event 0x6942a68d151863c1fed3c0c4c5f3258af738218527147ac69290ab23ca7d26c6.
//
// Solidity: event SetMinDelegateTime(uint128 indexed marketId, uint32 minDelegateTime)
func (_CoreGoerli *CoreGoerliFilterer) WatchSetMinDelegateTime(opts *bind.WatchOpts, sink chan<- *CoreGoerliSetMinDelegateTime, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "SetMinDelegateTime", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliSetMinDelegateTime)
				if err := _CoreGoerli.contract.UnpackLog(event, "SetMinDelegateTime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseSetMinDelegateTime(log types.Log) (*CoreGoerliSetMinDelegateTime, error) {
	event := new(CoreGoerliSetMinDelegateTime)
	if err := _CoreGoerli.contract.UnpackLog(event, "SetMinDelegateTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliSetMinLiquidityRatioIterator is returned from FilterSetMinLiquidityRatio and is used to iterate over the raw logs and unpacked data for SetMinLiquidityRatio events raised by the CoreGoerli contract.
type CoreGoerliSetMinLiquidityRatioIterator struct {
	Event *CoreGoerliSetMinLiquidityRatio // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliSetMinLiquidityRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliSetMinLiquidityRatio)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliSetMinLiquidityRatio)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliSetMinLiquidityRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliSetMinLiquidityRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliSetMinLiquidityRatio represents a SetMinLiquidityRatio event raised by the CoreGoerli contract.
type CoreGoerliSetMinLiquidityRatio struct {
	MinLiquidityRatio *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMinLiquidityRatio is a free log retrieval operation binding the contract event 0x66fd484d9868d1faddc8fef1f3faed0ed25eb4e6acde49dd1f2cbf0fba903635.
//
// Solidity: event SetMinLiquidityRatio(uint256 minLiquidityRatio)
func (_CoreGoerli *CoreGoerliFilterer) FilterSetMinLiquidityRatio(opts *bind.FilterOpts) (*CoreGoerliSetMinLiquidityRatioIterator, error) {

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "SetMinLiquidityRatio")
	if err != nil {
		return nil, err
	}
	return &CoreGoerliSetMinLiquidityRatioIterator{contract: _CoreGoerli.contract, event: "SetMinLiquidityRatio", logs: logs, sub: sub}, nil
}

// WatchSetMinLiquidityRatio is a free log subscription operation binding the contract event 0x66fd484d9868d1faddc8fef1f3faed0ed25eb4e6acde49dd1f2cbf0fba903635.
//
// Solidity: event SetMinLiquidityRatio(uint256 minLiquidityRatio)
func (_CoreGoerli *CoreGoerliFilterer) WatchSetMinLiquidityRatio(opts *bind.WatchOpts, sink chan<- *CoreGoerliSetMinLiquidityRatio) (event.Subscription, error) {

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "SetMinLiquidityRatio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliSetMinLiquidityRatio)
				if err := _CoreGoerli.contract.UnpackLog(event, "SetMinLiquidityRatio", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseSetMinLiquidityRatio(log types.Log) (*CoreGoerliSetMinLiquidityRatio, error) {
	event := new(CoreGoerliSetMinLiquidityRatio)
	if err := _CoreGoerli.contract.UnpackLog(event, "SetMinLiquidityRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliTransferCrossChainInitiatedIterator is returned from FilterTransferCrossChainInitiated and is used to iterate over the raw logs and unpacked data for TransferCrossChainInitiated events raised by the CoreGoerli contract.
type CoreGoerliTransferCrossChainInitiatedIterator struct {
	Event *CoreGoerliTransferCrossChainInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliTransferCrossChainInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliTransferCrossChainInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliTransferCrossChainInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliTransferCrossChainInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliTransferCrossChainInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliTransferCrossChainInitiated represents a TransferCrossChainInitiated event raised by the CoreGoerli contract.
type CoreGoerliTransferCrossChainInitiated struct {
	DestChainId uint64
	Amount      *big.Int
	Sender      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferCrossChainInitiated is a free log retrieval operation binding the contract event 0xb87c3097d7f9145a4915e8e434f04a1b7b91646d8a6e66a5cdab25caccb644c4.
//
// Solidity: event TransferCrossChainInitiated(uint64 indexed destChainId, uint256 indexed amount, address sender)
func (_CoreGoerli *CoreGoerliFilterer) FilterTransferCrossChainInitiated(opts *bind.FilterOpts, destChainId []uint64, amount []*big.Int) (*CoreGoerliTransferCrossChainInitiatedIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "TransferCrossChainInitiated", destChainIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliTransferCrossChainInitiatedIterator{contract: _CoreGoerli.contract, event: "TransferCrossChainInitiated", logs: logs, sub: sub}, nil
}

// WatchTransferCrossChainInitiated is a free log subscription operation binding the contract event 0xb87c3097d7f9145a4915e8e434f04a1b7b91646d8a6e66a5cdab25caccb644c4.
//
// Solidity: event TransferCrossChainInitiated(uint64 indexed destChainId, uint256 indexed amount, address sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchTransferCrossChainInitiated(opts *bind.WatchOpts, sink chan<- *CoreGoerliTransferCrossChainInitiated, destChainId []uint64, amount []*big.Int) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "TransferCrossChainInitiated", destChainIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliTransferCrossChainInitiated)
				if err := _CoreGoerli.contract.UnpackLog(event, "TransferCrossChainInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseTransferCrossChainInitiated(log types.Log) (*CoreGoerliTransferCrossChainInitiated, error) {
	event := new(CoreGoerliTransferCrossChainInitiated)
	if err := _CoreGoerli.contract.UnpackLog(event, "TransferCrossChainInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CoreGoerli contract.
type CoreGoerliUpgradedIterator struct {
	Event *CoreGoerliUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliUpgraded represents a Upgraded event raised by the CoreGoerli contract.
type CoreGoerliUpgraded struct {
	Self           common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_CoreGoerli *CoreGoerliFilterer) FilterUpgraded(opts *bind.FilterOpts, self []common.Address) (*CoreGoerliUpgradedIterator, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliUpgradedIterator{contract: _CoreGoerli.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed self, address implementation)
func (_CoreGoerli *CoreGoerliFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CoreGoerliUpgraded, self []common.Address) (event.Subscription, error) {

	var selfRule []interface{}
	for _, selfItem := range self {
		selfRule = append(selfRule, selfItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "Upgraded", selfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliUpgraded)
				if err := _CoreGoerli.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseUpgraded(log types.Log) (*CoreGoerliUpgraded, error) {
	event := new(CoreGoerliUpgraded)
	if err := _CoreGoerli.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliUsdBurnedIterator is returned from FilterUsdBurned and is used to iterate over the raw logs and unpacked data for UsdBurned events raised by the CoreGoerli contract.
type CoreGoerliUsdBurnedIterator struct {
	Event *CoreGoerliUsdBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliUsdBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliUsdBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliUsdBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliUsdBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliUsdBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliUsdBurned represents a UsdBurned event raised by the CoreGoerli contract.
type CoreGoerliUsdBurned struct {
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
func (_CoreGoerli *CoreGoerliFilterer) FilterUsdBurned(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*CoreGoerliUsdBurnedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "UsdBurned", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliUsdBurnedIterator{contract: _CoreGoerli.contract, event: "UsdBurned", logs: logs, sub: sub}, nil
}

// WatchUsdBurned is a free log subscription operation binding the contract event 0x6b0230f0abe9188cbdbb1c816a4c5e213758b5b743d8a94af056280cc1e7aeb1.
//
// Solidity: event UsdBurned(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchUsdBurned(opts *bind.WatchOpts, sink chan<- *CoreGoerliUsdBurned, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "UsdBurned", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliUsdBurned)
				if err := _CoreGoerli.contract.UnpackLog(event, "UsdBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseUsdBurned(log types.Log) (*CoreGoerliUsdBurned, error) {
	event := new(CoreGoerliUsdBurned)
	if err := _CoreGoerli.contract.UnpackLog(event, "UsdBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliUsdMintedIterator is returned from FilterUsdMinted and is used to iterate over the raw logs and unpacked data for UsdMinted events raised by the CoreGoerli contract.
type CoreGoerliUsdMintedIterator struct {
	Event *CoreGoerliUsdMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliUsdMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliUsdMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliUsdMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliUsdMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliUsdMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliUsdMinted represents a UsdMinted event raised by the CoreGoerli contract.
type CoreGoerliUsdMinted struct {
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
func (_CoreGoerli *CoreGoerliFilterer) FilterUsdMinted(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*CoreGoerliUsdMintedIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "UsdMinted", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliUsdMintedIterator{contract: _CoreGoerli.contract, event: "UsdMinted", logs: logs, sub: sub}, nil
}

// WatchUsdMinted is a free log subscription operation binding the contract event 0x2100f67dc9a5917400f799bb13194553e3f74c19a207c56350d2c223ac9c36c9.
//
// Solidity: event UsdMinted(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchUsdMinted(opts *bind.WatchOpts, sink chan<- *CoreGoerliUsdMinted, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "UsdMinted", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliUsdMinted)
				if err := _CoreGoerli.contract.UnpackLog(event, "UsdMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseUsdMinted(log types.Log) (*CoreGoerliUsdMinted, error) {
	event := new(CoreGoerliUsdMinted)
	if err := _CoreGoerli.contract.UnpackLog(event, "UsdMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliVaultLiquidationIterator is returned from FilterVaultLiquidation and is used to iterate over the raw logs and unpacked data for VaultLiquidation events raised by the CoreGoerli contract.
type CoreGoerliVaultLiquidationIterator struct {
	Event *CoreGoerliVaultLiquidation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliVaultLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliVaultLiquidation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliVaultLiquidation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliVaultLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliVaultLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliVaultLiquidation represents a VaultLiquidation event raised by the CoreGoerli contract.
type CoreGoerliVaultLiquidation struct {
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
func (_CoreGoerli *CoreGoerliFilterer) FilterVaultLiquidation(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address) (*CoreGoerliVaultLiquidationIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "VaultLiquidation", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliVaultLiquidationIterator{contract: _CoreGoerli.contract, event: "VaultLiquidation", logs: logs, sub: sub}, nil
}

// WatchVaultLiquidation is a free log subscription operation binding the contract event 0x1834a7cc9d14f9bfa482df5c0404dadd1b8ec123b41f082e76ae28a3b2ea68d5.
//
// Solidity: event VaultLiquidation(uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchVaultLiquidation(opts *bind.WatchOpts, sink chan<- *CoreGoerliVaultLiquidation, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "VaultLiquidation", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliVaultLiquidation)
				if err := _CoreGoerli.contract.UnpackLog(event, "VaultLiquidation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseVaultLiquidation(log types.Log) (*CoreGoerliVaultLiquidation, error) {
	event := new(CoreGoerliVaultLiquidation)
	if err := _CoreGoerli.contract.UnpackLog(event, "VaultLiquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreGoerliWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the CoreGoerli contract.
type CoreGoerliWithdrawnIterator struct {
	Event *CoreGoerliWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreGoerliWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreGoerliWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreGoerliWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreGoerliWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreGoerliWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreGoerliWithdrawn represents a Withdrawn event raised by the CoreGoerli contract.
type CoreGoerliWithdrawn struct {
	AccountId      *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x8b5f9d7ce522936589c630db08c0fa2405b21c4a5ff8ef19899900172736ba38.
//
// Solidity: event Withdrawn(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) FilterWithdrawn(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (*CoreGoerliWithdrawnIterator, error) {

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

	logs, sub, err := _CoreGoerli.contract.FilterLogs(opts, "Withdrawn", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreGoerliWithdrawnIterator{contract: _CoreGoerli.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x8b5f9d7ce522936589c630db08c0fa2405b21c4a5ff8ef19899900172736ba38.
//
// Solidity: event Withdrawn(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_CoreGoerli *CoreGoerliFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *CoreGoerliWithdrawn, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreGoerli.contract.WatchLogs(opts, "Withdrawn", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreGoerliWithdrawn)
				if err := _CoreGoerli.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_CoreGoerli *CoreGoerliFilterer) ParseWithdrawn(log types.Log) (*CoreGoerliWithdrawn, error) {
	event := new(CoreGoerliWithdrawn)
	if err := _CoreGoerli.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
