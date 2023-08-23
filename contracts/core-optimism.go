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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationIsSterile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotNominated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeSimulationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"self\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNominatedOwner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"simulateUpgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValueAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotInSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagAllowAllSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeatureFlagAllowlistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"FeatureFlagDeniersReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"FeatureFlagDenyAllSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getDeniers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"}],\"name\":\"getFeatureFlagDenyAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeatureAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromFeatureFlagAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"deniers\",\"type\":\"address[]\"}],\"name\":\"setDeniers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagAllowAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"denyAll\",\"type\":\"bool\"}],\"name\":\"setFeatureFlagDenyAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"FeatureUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"InvalidAccountId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"InvalidPermission\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"}],\"name\":\"OnlyAccountTokenProxy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"PermissionNotGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionOutOfBounds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PermissionRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedAccountId\",\"type\":\"uint128\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountLastInteraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"getAccountPermissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"permissions\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIAccountModule.AccountPermissions[]\",\"name\":\"accountPerms\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"notifyAccountTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"}],\"name\":\"renouncePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"permission\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revokePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"AccountNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyDistribution\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateralRatio\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"MarketNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"NotFundedByPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt256ToUint256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint128ToInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToInt256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint128\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"updatedDebt\",\"type\":\"int256\"}],\"name\":\"DebtAssociated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"associateDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MismatchAssociatedSystemKind\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MissingAssociatedSystem\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"AssociatedSystemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssociatedSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kind\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initOrUpgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"endpoint\",\"type\":\"address\"}],\"name\":\"registerUnmanagedSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InvalidMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotCcipRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNetwork\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCcipClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCcipClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredTime\",\"type\":\"uint256\"}],\"name\":\"AccountActivityTimeoutPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"CollateralDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InsufficientAccountCollateral\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existing\",\"type\":\"uint256\"}],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parameter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"PrecisionLost\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"CollateralLockCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"CollateralLockExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"cleanExpiredLocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cleared\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expireTimestamp\",\"type\":\"uint64\"}],\"name\":\"createLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getAccountAvailableCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getAccountCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDeposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAssigned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalLocked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getLocks\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amountD18\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"lockExpirationTime\",\"type\":\"uint64\"}],\"internalType\":\"structCollateralLock.Data[]\",\"name\":\"locks\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"CollateralConfigured\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"configureCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getCollateralConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"hideDisabled\",\"type\":\"bool\"}],\"name\":\"getCollateralConfigurations\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositingEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuanceRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatioD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRewardD18\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleNodeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDelegationD18\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralConfiguration.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getCollateralPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientCcipFee\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TransferCrossChainInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCrossChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gasTokenUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentDebt\",\"type\":\"int256\"}],\"name\":\"InsufficientDebt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"PoolNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"IssuanceFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UsdBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UsdMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnUsd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintUsd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CannotScaleEmptyMapping\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"currentCRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cratio\",\"type\":\"uint256\"}],\"name\":\"IneligibleForLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientMappedAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeVaultLiquidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowInt128ToUint128\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Liquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"VaultLiquidation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"isPositionLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"isVaultLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"}],\"name\":\"liquidate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"liquidateAsAccountId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxUsd\",\"type\":\"uint256\"}],\"name\":\"liquidateVault\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debtLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRewarded\",\"type\":\"uint256\"}],\"internalType\":\"structILiquidationModule.LiquidationData\",\"name\":\"liquidationData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountToDeposit\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarketCollateralDepositable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"InsufficientMarketCollateralWithdrawable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MarketCollateralDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MarketCollateralWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"systemAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MaximumMarketCollateralConfigured\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"configureMaximumMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"depositMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getMarketCollateralAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralAmountD18\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getMaximumMarketCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawMarketCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"IncorrectMarketInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NotEnoughLiquidity\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MarketRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"MarketSystemFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketUsdDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketUsdWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"SetMarketMinLiquidityRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"minDelegateTime\",\"type\":\"uint32\"}],\"name\":\"SetMinDelegateTime\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxIter\",\"type\":\"uint256\"}],\"name\":\"distributeDebtToPools\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketDebtPerShare\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getMarketFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFeeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketMinDelegateTime\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketNetIssuance\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketReportedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMarketTotalDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getMinLiquidityRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleManager\",\"outputs\":[{\"internalType\":\"contractIOracleManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsdToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"getWithdrawableMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"name\":\"isMarketCapacityLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"registerMarket\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"minDelegateTime\",\"type\":\"uint32\"}],\"name\":\"setMarketMinDelegateTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"setMinLiquidityRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawMarketUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RecursiveMulticall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PoolApprovedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PoolApprovedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PreferredPoolSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"addApprovedPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApprovedPools\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreferredPool\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"removeApprovedPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"setPreferredPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"CapacityLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"timeRemaining\",\"type\":\"uint32\"}],\"name\":\"MinDelegationTimeoutPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"PoolAlreadyExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"indexed\":false,\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"markets\",\"type\":\"tuple[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolConfigurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PoolNameUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolNominationRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolNominationRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolOwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolOwnershipAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"SetMinLiquidityRatio\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"acceptPoolOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"requestedPoolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinLiquidityRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getNominatedPoolOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"poolName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"getPoolOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"nominatePoolOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"optionalCollateralType\",\"type\":\"address\"}],\"name\":\"rebalancePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"renouncePoolNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"name\":\"revokePoolNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minLiquidityRatio\",\"type\":\"uint256\"}],\"name\":\"setMinLiquidityRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"marketId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"weightD18\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"maxDebtShareValueD18\",\"type\":\"int128\"}],\"internalType\":\"structMarketConfiguration.Data[]\",\"name\":\"newMarketConfigurations\",\"type\":\"tuple[]\"}],\"name\":\"setPoolConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setPoolName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OverflowUint256ToUint32\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint32ToInt32\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverflowUint64ToInt64\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardDistributorNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardUnavailable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"RewardsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardsDistributorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"RewardsDistributorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"}],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"getRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"registerRewardsDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"removeRewardsDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"}],\"name\":\"updateRewards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newChainId\",\"type\":\"uint64\"}],\"name\":\"NewSupportedCrossChainNetwork\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ccipRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ccipTokenPool\",\"type\":\"address\"}],\"name\":\"configureChainlinkCrossChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleManagerAddress\",\"type\":\"address\"}],\"name\":\"configureOracleManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"v\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfigAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getConfigUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"v\",\"type\":\"bytes32\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"supportedNetworks\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"ccipSelectors\",\"type\":\"uint64[]\"}],\"name\":\"setSupportedCrossChainNetworks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numRegistered\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minDelegation\",\"type\":\"uint256\"}],\"name\":\"InsufficientDelegation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCollateralAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"}],\"name\":\"InvalidLeverage\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"DelegationUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralAmountD18\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"}],\"name\":\"delegateCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"collateralizationRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionCollateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"accountId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getPositionDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"debt\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultCollateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"}],\"name\":\"getVaultDebt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_Contracts *ContractsCaller) GetAccountAvailableCollateral(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAccountAvailableCollateral", accountId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_Contracts *ContractsSession) GetAccountAvailableCollateral(accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetAccountAvailableCollateral(&_Contracts.CallOpts, accountId, collateralType)
}

// GetAccountAvailableCollateral is a free data retrieval call binding the contract method 0x927482ff.
//
// Solidity: function getAccountAvailableCollateral(uint128 accountId, address collateralType) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetAccountAvailableCollateral(accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetAccountAvailableCollateral(&_Contracts.CallOpts, accountId, collateralType)
}

// GetAccountCollateral is a free data retrieval call binding the contract method 0xef45148e.
//
// Solidity: function getAccountCollateral(uint128 accountId, address collateralType) view returns(uint256 totalDeposited, uint256 totalAssigned, uint256 totalLocked)
func (_Contracts *ContractsCaller) GetAccountCollateral(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAccountCollateral", accountId, collateralType)

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
func (_Contracts *ContractsSession) GetAccountCollateral(accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	return _Contracts.Contract.GetAccountCollateral(&_Contracts.CallOpts, accountId, collateralType)
}

// GetAccountCollateral is a free data retrieval call binding the contract method 0xef45148e.
//
// Solidity: function getAccountCollateral(uint128 accountId, address collateralType) view returns(uint256 totalDeposited, uint256 totalAssigned, uint256 totalLocked)
func (_Contracts *ContractsCallerSession) GetAccountCollateral(accountId *big.Int, collateralType common.Address) (struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}, error) {
	return _Contracts.Contract.GetAccountCollateral(&_Contracts.CallOpts, accountId, collateralType)
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

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_Contracts *ContractsCaller) GetApprovedPools(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getApprovedPools")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_Contracts *ContractsSession) GetApprovedPools() ([]*big.Int, error) {
	return _Contracts.Contract.GetApprovedPools(&_Contracts.CallOpts)
}

// GetApprovedPools is a free data retrieval call binding the contract method 0x48741626.
//
// Solidity: function getApprovedPools() view returns(uint256[])
func (_Contracts *ContractsCallerSession) GetApprovedPools() ([]*big.Int, error) {
	return _Contracts.Contract.GetApprovedPools(&_Contracts.CallOpts)
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

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_Contracts *ContractsCaller) GetCollateralConfiguration(opts *bind.CallOpts, collateralType common.Address) (CollateralConfigurationData, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCollateralConfiguration", collateralType)

	if err != nil {
		return *new(CollateralConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new(CollateralConfigurationData)).(*CollateralConfigurationData)

	return out0, err

}

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_Contracts *ContractsSession) GetCollateralConfiguration(collateralType common.Address) (CollateralConfigurationData, error) {
	return _Contracts.Contract.GetCollateralConfiguration(&_Contracts.CallOpts, collateralType)
}

// GetCollateralConfiguration is a free data retrieval call binding the contract method 0xdc0b3f52.
//
// Solidity: function getCollateralConfiguration(address collateralType) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256))
func (_Contracts *ContractsCallerSession) GetCollateralConfiguration(collateralType common.Address) (CollateralConfigurationData, error) {
	return _Contracts.Contract.GetCollateralConfiguration(&_Contracts.CallOpts, collateralType)
}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_Contracts *ContractsCaller) GetCollateralConfigurations(opts *bind.CallOpts, hideDisabled bool) ([]CollateralConfigurationData, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCollateralConfigurations", hideDisabled)

	if err != nil {
		return *new([]CollateralConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CollateralConfigurationData)).(*[]CollateralConfigurationData)

	return out0, err

}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_Contracts *ContractsSession) GetCollateralConfigurations(hideDisabled bool) ([]CollateralConfigurationData, error) {
	return _Contracts.Contract.GetCollateralConfigurations(&_Contracts.CallOpts, hideDisabled)
}

// GetCollateralConfigurations is a free data retrieval call binding the contract method 0x75bf2444.
//
// Solidity: function getCollateralConfigurations(bool hideDisabled) view returns((bool,uint256,uint256,uint256,bytes32,address,uint256)[])
func (_Contracts *ContractsCallerSession) GetCollateralConfigurations(hideDisabled bool) ([]CollateralConfigurationData, error) {
	return _Contracts.Contract.GetCollateralConfigurations(&_Contracts.CallOpts, hideDisabled)
}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_Contracts *ContractsCaller) GetCollateralPrice(opts *bind.CallOpts, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCollateralPrice", collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_Contracts *ContractsSession) GetCollateralPrice(collateralType common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetCollateralPrice(&_Contracts.CallOpts, collateralType)
}

// GetCollateralPrice is a free data retrieval call binding the contract method 0x51a40994.
//
// Solidity: function getCollateralPrice(address collateralType) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetCollateralPrice(collateralType common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetCollateralPrice(&_Contracts.CallOpts, collateralType)
}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_Contracts *ContractsCaller) GetConfig(opts *bind.CallOpts, k [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getConfig", k)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_Contracts *ContractsSession) GetConfig(k [32]byte) ([32]byte, error) {
	return _Contracts.Contract.GetConfig(&_Contracts.CallOpts, k)
}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 k) view returns(bytes32 v)
func (_Contracts *ContractsCallerSession) GetConfig(k [32]byte) ([32]byte, error) {
	return _Contracts.Contract.GetConfig(&_Contracts.CallOpts, k)
}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_Contracts *ContractsCaller) GetConfigAddress(opts *bind.CallOpts, k [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getConfigAddress", k)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_Contracts *ContractsSession) GetConfigAddress(k [32]byte) (common.Address, error) {
	return _Contracts.Contract.GetConfigAddress(&_Contracts.CallOpts, k)
}

// GetConfigAddress is a free data retrieval call binding the contract method 0xf896503a.
//
// Solidity: function getConfigAddress(bytes32 k) view returns(address v)
func (_Contracts *ContractsCallerSession) GetConfigAddress(k [32]byte) (common.Address, error) {
	return _Contracts.Contract.GetConfigAddress(&_Contracts.CallOpts, k)
}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_Contracts *ContractsCaller) GetConfigUint(opts *bind.CallOpts, k [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getConfigUint", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_Contracts *ContractsSession) GetConfigUint(k [32]byte) (*big.Int, error) {
	return _Contracts.Contract.GetConfigUint(&_Contracts.CallOpts, k)
}

// GetConfigUint is a free data retrieval call binding the contract method 0xf92bb8c9.
//
// Solidity: function getConfigUint(bytes32 k) view returns(uint256 v)
func (_Contracts *ContractsCallerSession) GetConfigUint(k [32]byte) (*big.Int, error) {
	return _Contracts.Contract.GetConfigUint(&_Contracts.CallOpts, k)
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

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_Contracts *ContractsCaller) GetLocks(opts *bind.CallOpts, accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLocks", accountId, collateralType, offset, count)

	if err != nil {
		return *new([]CollateralLockData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CollateralLockData)).(*[]CollateralLockData)

	return out0, err

}

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_Contracts *ContractsSession) GetLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	return _Contracts.Contract.GetLocks(&_Contracts.CallOpts, accountId, collateralType, offset, count)
}

// GetLocks is a free data retrieval call binding the contract method 0xaa8c6369.
//
// Solidity: function getLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) view returns((uint128,uint64)[] locks)
func (_Contracts *ContractsCallerSession) GetLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) ([]CollateralLockData, error) {
	return _Contracts.Contract.GetLocks(&_Contracts.CallOpts, accountId, collateralType, offset, count)
}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCaller) GetMarketCollateral(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketCollateral", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsSession) GetMarketCollateral(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketCollateral(&_Contracts.CallOpts, marketId)
}

// GetMarketCollateral is a free data retrieval call binding the contract method 0x150834a3.
//
// Solidity: function getMarketCollateral(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetMarketCollateral(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketCollateral(&_Contracts.CallOpts, marketId)
}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_Contracts *ContractsCaller) GetMarketCollateralAmount(opts *bind.CallOpts, marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketCollateralAmount", marketId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_Contracts *ContractsSession) GetMarketCollateralAmount(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetMarketCollateralAmount(&_Contracts.CallOpts, marketId, collateralType)
}

// GetMarketCollateralAmount is a free data retrieval call binding the contract method 0xc2b0cf41.
//
// Solidity: function getMarketCollateralAmount(uint128 marketId, address collateralType) view returns(uint256 collateralAmountD18)
func (_Contracts *ContractsCallerSession) GetMarketCollateralAmount(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetMarketCollateralAmount(&_Contracts.CallOpts, marketId, collateralType)
}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCaller) GetMarketCollateralValue(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketCollateralValue", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsSession) GetMarketCollateralValue(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketCollateralValue(&_Contracts.CallOpts, marketId)
}

// GetMarketCollateralValue is a free data retrieval call binding the contract method 0xd4f88381.
//
// Solidity: function getMarketCollateralValue(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetMarketCollateralValue(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketCollateralValue(&_Contracts.CallOpts, marketId)
}

// GetMarketFees is a free data retrieval call binding the contract method 0xdfb83437.
//
// Solidity: function getMarketFees(uint128 , uint256 amount) view returns(uint256 depositFeeAmount, uint256 withdrawFeeAmount)
func (_Contracts *ContractsCaller) GetMarketFees(opts *bind.CallOpts, arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketFees", arg0, amount)

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
func (_Contracts *ContractsSession) GetMarketFees(arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	return _Contracts.Contract.GetMarketFees(&_Contracts.CallOpts, arg0, amount)
}

// GetMarketFees is a free data retrieval call binding the contract method 0xdfb83437.
//
// Solidity: function getMarketFees(uint128 , uint256 amount) view returns(uint256 depositFeeAmount, uint256 withdrawFeeAmount)
func (_Contracts *ContractsCallerSession) GetMarketFees(arg0 *big.Int, amount *big.Int) (struct {
	DepositFeeAmount  *big.Int
	WithdrawFeeAmount *big.Int
}, error) {
	return _Contracts.Contract.GetMarketFees(&_Contracts.CallOpts, arg0, amount)
}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_Contracts *ContractsCaller) GetMarketMinDelegateTime(opts *bind.CallOpts, marketId *big.Int) (uint32, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketMinDelegateTime", marketId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_Contracts *ContractsSession) GetMarketMinDelegateTime(marketId *big.Int) (uint32, error) {
	return _Contracts.Contract.GetMarketMinDelegateTime(&_Contracts.CallOpts, marketId)
}

// GetMarketMinDelegateTime is a free data retrieval call binding the contract method 0x5424901b.
//
// Solidity: function getMarketMinDelegateTime(uint128 marketId) view returns(uint32)
func (_Contracts *ContractsCallerSession) GetMarketMinDelegateTime(marketId *big.Int) (uint32, error) {
	return _Contracts.Contract.GetMarketMinDelegateTime(&_Contracts.CallOpts, marketId)
}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_Contracts *ContractsCaller) GetMarketNetIssuance(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketNetIssuance", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_Contracts *ContractsSession) GetMarketNetIssuance(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketNetIssuance(&_Contracts.CallOpts, marketId)
}

// GetMarketNetIssuance is a free data retrieval call binding the contract method 0x85d99ebc.
//
// Solidity: function getMarketNetIssuance(uint128 marketId) view returns(int128)
func (_Contracts *ContractsCallerSession) GetMarketNetIssuance(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketNetIssuance(&_Contracts.CallOpts, marketId)
}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCaller) GetMarketReportedDebt(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketReportedDebt", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsSession) GetMarketReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketReportedDebt(&_Contracts.CallOpts, marketId)
}

// GetMarketReportedDebt is a free data retrieval call binding the contract method 0x86e3b1cf.
//
// Solidity: function getMarketReportedDebt(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetMarketReportedDebt(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketReportedDebt(&_Contracts.CallOpts, marketId)
}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_Contracts *ContractsCaller) GetMarketTotalDebt(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarketTotalDebt", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_Contracts *ContractsSession) GetMarketTotalDebt(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketTotalDebt(&_Contracts.CallOpts, marketId)
}

// GetMarketTotalDebt is a free data retrieval call binding the contract method 0xbaa2a264.
//
// Solidity: function getMarketTotalDebt(uint128 marketId) view returns(int256)
func (_Contracts *ContractsCallerSession) GetMarketTotalDebt(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMarketTotalDebt(&_Contracts.CallOpts, marketId)
}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_Contracts *ContractsCaller) GetMaximumMarketCollateral(opts *bind.CallOpts, marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMaximumMarketCollateral", marketId, collateralType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_Contracts *ContractsSession) GetMaximumMarketCollateral(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetMaximumMarketCollateral(&_Contracts.CallOpts, marketId, collateralType)
}

// GetMaximumMarketCollateral is a free data retrieval call binding the contract method 0x12e1c673.
//
// Solidity: function getMaximumMarketCollateral(uint128 marketId, address collateralType) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetMaximumMarketCollateral(marketId *big.Int, collateralType common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetMaximumMarketCollateral(&_Contracts.CallOpts, marketId, collateralType)
}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCaller) GetMinLiquidityRatio(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMinLiquidityRatio", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsSession) GetMinLiquidityRatio(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMinLiquidityRatio(&_Contracts.CallOpts, marketId)
}

// GetMinLiquidityRatio is a free data retrieval call binding the contract method 0x84f29b6d.
//
// Solidity: function getMinLiquidityRatio(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetMinLiquidityRatio(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetMinLiquidityRatio(&_Contracts.CallOpts, marketId)
}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_Contracts *ContractsCaller) GetMinLiquidityRatio0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMinLiquidityRatio0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_Contracts *ContractsSession) GetMinLiquidityRatio0() (*big.Int, error) {
	return _Contracts.Contract.GetMinLiquidityRatio0(&_Contracts.CallOpts)
}

// GetMinLiquidityRatio0 is a free data retrieval call binding the contract method 0xfd85c1f8.
//
// Solidity: function getMinLiquidityRatio() view returns(uint256)
func (_Contracts *ContractsCallerSession) GetMinLiquidityRatio0() (*big.Int, error) {
	return _Contracts.Contract.GetMinLiquidityRatio0(&_Contracts.CallOpts)
}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_Contracts *ContractsCaller) GetNominatedPoolOwner(opts *bind.CallOpts, poolId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getNominatedPoolOwner", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_Contracts *ContractsSession) GetNominatedPoolOwner(poolId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetNominatedPoolOwner(&_Contracts.CallOpts, poolId)
}

// GetNominatedPoolOwner is a free data retrieval call binding the contract method 0x9851af01.
//
// Solidity: function getNominatedPoolOwner(uint128 poolId) view returns(address)
func (_Contracts *ContractsCallerSession) GetNominatedPoolOwner(poolId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetNominatedPoolOwner(&_Contracts.CallOpts, poolId)
}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_Contracts *ContractsCaller) GetOracleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getOracleManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_Contracts *ContractsSession) GetOracleManager() (common.Address, error) {
	return _Contracts.Contract.GetOracleManager(&_Contracts.CallOpts)
}

// GetOracleManager is a free data retrieval call binding the contract method 0xb01ceccd.
//
// Solidity: function getOracleManager() view returns(address)
func (_Contracts *ContractsCallerSession) GetOracleManager() (common.Address, error) {
	return _Contracts.Contract.GetOracleManager(&_Contracts.CallOpts)
}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_Contracts *ContractsCaller) GetPoolConfiguration(opts *bind.CallOpts, poolId *big.Int) ([]MarketConfigurationData, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPoolConfiguration", poolId)

	if err != nil {
		return *new([]MarketConfigurationData), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketConfigurationData)).(*[]MarketConfigurationData)

	return out0, err

}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_Contracts *ContractsSession) GetPoolConfiguration(poolId *big.Int) ([]MarketConfigurationData, error) {
	return _Contracts.Contract.GetPoolConfiguration(&_Contracts.CallOpts, poolId)
}

// GetPoolConfiguration is a free data retrieval call binding the contract method 0xefecf137.
//
// Solidity: function getPoolConfiguration(uint128 poolId) view returns((uint128,uint128,int128)[])
func (_Contracts *ContractsCallerSession) GetPoolConfiguration(poolId *big.Int) ([]MarketConfigurationData, error) {
	return _Contracts.Contract.GetPoolConfiguration(&_Contracts.CallOpts, poolId)
}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_Contracts *ContractsCaller) GetPoolName(opts *bind.CallOpts, poolId *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPoolName", poolId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_Contracts *ContractsSession) GetPoolName(poolId *big.Int) (string, error) {
	return _Contracts.Contract.GetPoolName(&_Contracts.CallOpts, poolId)
}

// GetPoolName is a free data retrieval call binding the contract method 0xf86e6f91.
//
// Solidity: function getPoolName(uint128 poolId) view returns(string poolName)
func (_Contracts *ContractsCallerSession) GetPoolName(poolId *big.Int) (string, error) {
	return _Contracts.Contract.GetPoolName(&_Contracts.CallOpts, poolId)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_Contracts *ContractsCaller) GetPoolOwner(opts *bind.CallOpts, poolId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPoolOwner", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_Contracts *ContractsSession) GetPoolOwner(poolId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetPoolOwner(&_Contracts.CallOpts, poolId)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0xbbdd7c5a.
//
// Solidity: function getPoolOwner(uint128 poolId) view returns(address)
func (_Contracts *ContractsCallerSession) GetPoolOwner(poolId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetPoolOwner(&_Contracts.CallOpts, poolId)
}

// GetPositionCollateral is a free data retrieval call binding the contract method 0x33cc422b.
//
// Solidity: function getPositionCollateral(uint128 accountId, uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_Contracts *ContractsCaller) GetPositionCollateral(opts *bind.CallOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPositionCollateral", accountId, poolId, collateralType)

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
func (_Contracts *ContractsSession) GetPositionCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _Contracts.Contract.GetPositionCollateral(&_Contracts.CallOpts, accountId, poolId, collateralType)
}

// GetPositionCollateral is a free data retrieval call binding the contract method 0x33cc422b.
//
// Solidity: function getPositionCollateral(uint128 accountId, uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_Contracts *ContractsCallerSession) GetPositionCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _Contracts.Contract.GetPositionCollateral(&_Contracts.CallOpts, accountId, poolId, collateralType)
}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_Contracts *ContractsCaller) GetPreferredPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPreferredPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_Contracts *ContractsSession) GetPreferredPool() (*big.Int, error) {
	return _Contracts.Contract.GetPreferredPool(&_Contracts.CallOpts)
}

// GetPreferredPool is a free data retrieval call binding the contract method 0x3b390b57.
//
// Solidity: function getPreferredPool() view returns(uint128)
func (_Contracts *ContractsCallerSession) GetPreferredPool() (*big.Int, error) {
	return _Contracts.Contract.GetPreferredPool(&_Contracts.CallOpts)
}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_Contracts *ContractsCaller) GetRewardRate(opts *bind.CallOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getRewardRate", poolId, collateralType, distributor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_Contracts *ContractsSession) GetRewardRate(poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetRewardRate(&_Contracts.CallOpts, poolId, collateralType, distributor)
}

// GetRewardRate is a free data retrieval call binding the contract method 0x0dd2395a.
//
// Solidity: function getRewardRate(uint128 poolId, address collateralType, address distributor) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetRewardRate(poolId *big.Int, collateralType common.Address, distributor common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetRewardRate(&_Contracts.CallOpts, poolId, collateralType, distributor)
}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_Contracts *ContractsCaller) GetUsdToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getUsdToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_Contracts *ContractsSession) GetUsdToken() (common.Address, error) {
	return _Contracts.Contract.GetUsdToken(&_Contracts.CallOpts)
}

// GetUsdToken is a free data retrieval call binding the contract method 0x21f1d9e5.
//
// Solidity: function getUsdToken() view returns(address)
func (_Contracts *ContractsCallerSession) GetUsdToken() (common.Address, error) {
	return _Contracts.Contract.GetUsdToken(&_Contracts.CallOpts)
}

// GetVaultCollateral is a free data retrieval call binding the contract method 0x078145a8.
//
// Solidity: function getVaultCollateral(uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_Contracts *ContractsCaller) GetVaultCollateral(opts *bind.CallOpts, poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getVaultCollateral", poolId, collateralType)

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
func (_Contracts *ContractsSession) GetVaultCollateral(poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _Contracts.Contract.GetVaultCollateral(&_Contracts.CallOpts, poolId, collateralType)
}

// GetVaultCollateral is a free data retrieval call binding the contract method 0x078145a8.
//
// Solidity: function getVaultCollateral(uint128 poolId, address collateralType) view returns(uint256 amount, uint256 value)
func (_Contracts *ContractsCallerSession) GetVaultCollateral(poolId *big.Int, collateralType common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
}, error) {
	return _Contracts.Contract.GetVaultCollateral(&_Contracts.CallOpts, poolId, collateralType)
}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCaller) GetWithdrawableMarketUsd(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getWithdrawableMarketUsd", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsSession) GetWithdrawableMarketUsd(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetWithdrawableMarketUsd(&_Contracts.CallOpts, marketId)
}

// GetWithdrawableMarketUsd is a free data retrieval call binding the contract method 0x1eb60770.
//
// Solidity: function getWithdrawableMarketUsd(uint128 marketId) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetWithdrawableMarketUsd(marketId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetWithdrawableMarketUsd(&_Contracts.CallOpts, marketId)
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

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_Contracts *ContractsCaller) IsMarketCapacityLocked(opts *bind.CallOpts, marketId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isMarketCapacityLocked", marketId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_Contracts *ContractsSession) IsMarketCapacityLocked(marketId *big.Int) (bool, error) {
	return _Contracts.Contract.IsMarketCapacityLocked(&_Contracts.CallOpts, marketId)
}

// IsMarketCapacityLocked is a free data retrieval call binding the contract method 0x07003f0a.
//
// Solidity: function isMarketCapacityLocked(uint128 marketId) view returns(bool)
func (_Contracts *ContractsCallerSession) IsMarketCapacityLocked(marketId *big.Int) (bool, error) {
	return _Contracts.Contract.IsMarketCapacityLocked(&_Contracts.CallOpts, marketId)
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

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_Contracts *ContractsTransactor) AcceptPoolOwnership(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "acceptPoolOwnership", poolId)
}

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_Contracts *ContractsSession) AcceptPoolOwnership(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AcceptPoolOwnership(&_Contracts.TransactOpts, poolId)
}

// AcceptPoolOwnership is a paid mutator transaction binding the contract method 0xc707a39f.
//
// Solidity: function acceptPoolOwnership(uint128 poolId) returns()
func (_Contracts *ContractsTransactorSession) AcceptPoolOwnership(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AcceptPoolOwnership(&_Contracts.TransactOpts, poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_Contracts *ContractsTransactor) AddApprovedPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addApprovedPool", poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_Contracts *ContractsSession) AddApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddApprovedPool(&_Contracts.TransactOpts, poolId)
}

// AddApprovedPool is a paid mutator transaction binding the contract method 0xb790a1ae.
//
// Solidity: function addApprovedPool(uint128 poolId) returns()
func (_Contracts *ContractsTransactorSession) AddApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddApprovedPool(&_Contracts.TransactOpts, poolId)
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

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_Contracts *ContractsTransactor) AssociateDebt(opts *bind.TransactOpts, marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "associateDebt", marketId, poolId, collateralType, accountId, amount)
}

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_Contracts *ContractsSession) AssociateDebt(marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AssociateDebt(&_Contracts.TransactOpts, marketId, poolId, collateralType, accountId, amount)
}

// AssociateDebt is a paid mutator transaction binding the contract method 0x11aa282d.
//
// Solidity: function associateDebt(uint128 marketId, uint128 poolId, address collateralType, uint128 accountId, uint256 amount) returns(int256)
func (_Contracts *ContractsTransactorSession) AssociateDebt(marketId *big.Int, poolId *big.Int, collateralType common.Address, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AssociateDebt(&_Contracts.TransactOpts, marketId, poolId, collateralType, accountId, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Contracts *ContractsTransactor) BurnUsd(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "burnUsd", accountId, poolId, collateralType, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Contracts *ContractsSession) BurnUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.BurnUsd(&_Contracts.TransactOpts, accountId, poolId, collateralType, amount)
}

// BurnUsd is a paid mutator transaction binding the contract method 0xd3264e43.
//
// Solidity: function burnUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) BurnUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.BurnUsd(&_Contracts.TransactOpts, accountId, poolId, collateralType, amount)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Contracts *ContractsTransactor) CcipReceive(opts *bind.TransactOpts, message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Contracts *ContractsSession) CcipReceive(message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _Contracts.Contract.CcipReceive(&_Contracts.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Contracts *ContractsTransactorSession) CcipReceive(message CcipClientAny2EVMMessage) (*types.Transaction, error) {
	return _Contracts.Contract.CcipReceive(&_Contracts.TransactOpts, message)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_Contracts *ContractsTransactor) ClaimRewards(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "claimRewards", accountId, poolId, collateralType, distributor)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_Contracts *ContractsSession) ClaimRewards(accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ClaimRewards(&_Contracts.TransactOpts, accountId, poolId, collateralType, distributor)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x460d2049.
//
// Solidity: function claimRewards(uint128 accountId, uint128 poolId, address collateralType, address distributor) returns(uint256)
func (_Contracts *ContractsTransactorSession) ClaimRewards(accountId *big.Int, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ClaimRewards(&_Contracts.TransactOpts, accountId, poolId, collateralType, distributor)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_Contracts *ContractsTransactor) CleanExpiredLocks(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "cleanExpiredLocks", accountId, collateralType, offset, count)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_Contracts *ContractsSession) CleanExpiredLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CleanExpiredLocks(&_Contracts.TransactOpts, accountId, collateralType, offset, count)
}

// CleanExpiredLocks is a paid mutator transaction binding the contract method 0x198f0aa1.
//
// Solidity: function cleanExpiredLocks(uint128 accountId, address collateralType, uint256 offset, uint256 count) returns(uint256 cleared)
func (_Contracts *ContractsTransactorSession) CleanExpiredLocks(accountId *big.Int, collateralType common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CleanExpiredLocks(&_Contracts.TransactOpts, accountId, collateralType, offset, count)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_Contracts *ContractsTransactor) ConfigureChainlinkCrossChain(opts *bind.TransactOpts, ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "configureChainlinkCrossChain", ccipRouter, ccipTokenPool)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_Contracts *ContractsSession) ConfigureChainlinkCrossChain(ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ConfigureChainlinkCrossChain(&_Contracts.TransactOpts, ccipRouter, ccipTokenPool)
}

// ConfigureChainlinkCrossChain is a paid mutator transaction binding the contract method 0x10d52805.
//
// Solidity: function configureChainlinkCrossChain(address ccipRouter, address ccipTokenPool) returns()
func (_Contracts *ContractsTransactorSession) ConfigureChainlinkCrossChain(ccipRouter common.Address, ccipTokenPool common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ConfigureChainlinkCrossChain(&_Contracts.TransactOpts, ccipRouter, ccipTokenPool)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_Contracts *ContractsTransactor) ConfigureCollateral(opts *bind.TransactOpts, config CollateralConfigurationData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "configureCollateral", config)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_Contracts *ContractsSession) ConfigureCollateral(config CollateralConfigurationData) (*types.Transaction, error) {
	return _Contracts.Contract.ConfigureCollateral(&_Contracts.TransactOpts, config)
}

// ConfigureCollateral is a paid mutator transaction binding the contract method 0x644cb0f3.
//
// Solidity: function configureCollateral((bool,uint256,uint256,uint256,bytes32,address,uint256) config) returns()
func (_Contracts *ContractsTransactorSession) ConfigureCollateral(config CollateralConfigurationData) (*types.Transaction, error) {
	return _Contracts.Contract.ConfigureCollateral(&_Contracts.TransactOpts, config)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_Contracts *ContractsTransactor) ConfigureMaximumMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "configureMaximumMarketCollateral", marketId, collateralType, amount)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_Contracts *ContractsSession) ConfigureMaximumMarketCollateral(marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ConfigureMaximumMarketCollateral(&_Contracts.TransactOpts, marketId, collateralType, amount)
}

// ConfigureMaximumMarketCollateral is a paid mutator transaction binding the contract method 0xdbdea94c.
//
// Solidity: function configureMaximumMarketCollateral(uint128 marketId, address collateralType, uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) ConfigureMaximumMarketCollateral(marketId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ConfigureMaximumMarketCollateral(&_Contracts.TransactOpts, marketId, collateralType, amount)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_Contracts *ContractsTransactor) ConfigureOracleManager(opts *bind.TransactOpts, oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "configureOracleManager", oracleManagerAddress)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_Contracts *ContractsSession) ConfigureOracleManager(oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ConfigureOracleManager(&_Contracts.TransactOpts, oracleManagerAddress)
}

// ConfigureOracleManager is a paid mutator transaction binding the contract method 0xa5d49393.
//
// Solidity: function configureOracleManager(address oracleManagerAddress) returns()
func (_Contracts *ContractsTransactorSession) ConfigureOracleManager(oracleManagerAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ConfigureOracleManager(&_Contracts.TransactOpts, oracleManagerAddress)
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

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_Contracts *ContractsTransactor) CreateLock(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createLock", accountId, collateralType, amount, expireTimestamp)
}

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_Contracts *ContractsSession) CreateLock(accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _Contracts.Contract.CreateLock(&_Contracts.TransactOpts, accountId, collateralType, amount, expireTimestamp)
}

// CreateLock is a paid mutator transaction binding the contract method 0x0bae9893.
//
// Solidity: function createLock(uint128 accountId, address collateralType, uint256 amount, uint64 expireTimestamp) returns()
func (_Contracts *ContractsTransactorSession) CreateLock(accountId *big.Int, collateralType common.Address, amount *big.Int, expireTimestamp uint64) (*types.Transaction, error) {
	return _Contracts.Contract.CreateLock(&_Contracts.TransactOpts, accountId, collateralType, amount, expireTimestamp)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_Contracts *ContractsTransactor) CreatePool(opts *bind.TransactOpts, requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createPool", requestedPoolId, owner)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_Contracts *ContractsSession) CreatePool(requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.CreatePool(&_Contracts.TransactOpts, requestedPoolId, owner)
}

// CreatePool is a paid mutator transaction binding the contract method 0xcaab529b.
//
// Solidity: function createPool(uint128 requestedPoolId, address owner) returns()
func (_Contracts *ContractsTransactorSession) CreatePool(requestedPoolId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.CreatePool(&_Contracts.TransactOpts, requestedPoolId, owner)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_Contracts *ContractsTransactor) DelegateCollateral(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "delegateCollateral", accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_Contracts *ContractsSession) DelegateCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DelegateCollateral(&_Contracts.TransactOpts, accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// DelegateCollateral is a paid mutator transaction binding the contract method 0x7b0532a4.
//
// Solidity: function delegateCollateral(uint128 accountId, uint128 poolId, address collateralType, uint256 newCollateralAmountD18, uint256 leverage) returns()
func (_Contracts *ContractsTransactorSession) DelegateCollateral(accountId *big.Int, poolId *big.Int, collateralType common.Address, newCollateralAmountD18 *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DelegateCollateral(&_Contracts.TransactOpts, accountId, poolId, collateralType, newCollateralAmountD18, leverage)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsTransactor) Deposit(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deposit", accountId, collateralType, tokenAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsSession) Deposit(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Deposit(&_Contracts.TransactOpts, accountId, collateralType, tokenAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x83802968.
//
// Solidity: function deposit(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsTransactorSession) Deposit(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Deposit(&_Contracts.TransactOpts, accountId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsTransactor) DepositMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "depositMarketCollateral", marketId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsSession) DepositMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DepositMarketCollateral(&_Contracts.TransactOpts, marketId, collateralType, tokenAmount)
}

// DepositMarketCollateral is a paid mutator transaction binding the contract method 0xa4e6306b.
//
// Solidity: function depositMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsTransactorSession) DepositMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DepositMarketCollateral(&_Contracts.TransactOpts, marketId, collateralType, tokenAmount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Contracts *ContractsTransactor) DepositMarketUsd(opts *bind.TransactOpts, marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "depositMarketUsd", marketId, target, amount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Contracts *ContractsSession) DepositMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DepositMarketUsd(&_Contracts.TransactOpts, marketId, target, amount)
}

// DepositMarketUsd is a paid mutator transaction binding the contract method 0x10b0cf76.
//
// Solidity: function depositMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Contracts *ContractsTransactorSession) DepositMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DepositMarketUsd(&_Contracts.TransactOpts, marketId, target, amount)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_Contracts *ContractsTransactor) DistributeDebtToPools(opts *bind.TransactOpts, marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "distributeDebtToPools", marketId, maxIter)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_Contracts *ContractsSession) DistributeDebtToPools(marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DistributeDebtToPools(&_Contracts.TransactOpts, marketId, maxIter)
}

// DistributeDebtToPools is a paid mutator transaction binding the contract method 0xa0c12269.
//
// Solidity: function distributeDebtToPools(uint128 marketId, uint256 maxIter) returns(bool)
func (_Contracts *ContractsTransactorSession) DistributeDebtToPools(marketId *big.Int, maxIter *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DistributeDebtToPools(&_Contracts.TransactOpts, marketId, maxIter)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_Contracts *ContractsTransactor) DistributeRewards(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "distributeRewards", poolId, collateralType, amount, start, duration)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_Contracts *ContractsSession) DistributeRewards(poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _Contracts.Contract.DistributeRewards(&_Contracts.TransactOpts, poolId, collateralType, amount, start, duration)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5a7ff7c5.
//
// Solidity: function distributeRewards(uint128 poolId, address collateralType, uint256 amount, uint64 start, uint32 duration) returns()
func (_Contracts *ContractsTransactorSession) DistributeRewards(poolId *big.Int, collateralType common.Address, amount *big.Int, start uint64, duration uint32) (*types.Transaction, error) {
	return _Contracts.Contract.DistributeRewards(&_Contracts.TransactOpts, poolId, collateralType, amount, start, duration)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_Contracts *ContractsTransactor) GetMarketDebtPerShare(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "getMarketDebtPerShare", marketId)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_Contracts *ContractsSession) GetMarketDebtPerShare(marketId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.GetMarketDebtPerShare(&_Contracts.TransactOpts, marketId)
}

// GetMarketDebtPerShare is a paid mutator transaction binding the contract method 0x95909ba3.
//
// Solidity: function getMarketDebtPerShare(uint128 marketId) returns(int256)
func (_Contracts *ContractsTransactorSession) GetMarketDebtPerShare(marketId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.GetMarketDebtPerShare(&_Contracts.TransactOpts, marketId)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_Contracts *ContractsTransactor) GetPosition(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "getPosition", accountId, poolId, collateralType)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_Contracts *ContractsSession) GetPosition(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GetPosition(&_Contracts.TransactOpts, accountId, poolId, collateralType)
}

// GetPosition is a paid mutator transaction binding the contract method 0xf544d66e.
//
// Solidity: function getPosition(uint128 accountId, uint128 poolId, address collateralType) returns(uint256 collateralAmount, uint256 collateralValue, int256 debt, uint256 collateralizationRatio)
func (_Contracts *ContractsTransactorSession) GetPosition(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GetPosition(&_Contracts.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_Contracts *ContractsTransactor) GetPositionCollateralRatio(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "getPositionCollateralRatio", accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_Contracts *ContractsSession) GetPositionCollateralRatio(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GetPositionCollateralRatio(&_Contracts.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionCollateralRatio is a paid mutator transaction binding the contract method 0xdc0a5384.
//
// Solidity: function getPositionCollateralRatio(uint128 accountId, uint128 poolId, address collateralType) returns(uint256)
func (_Contracts *ContractsTransactorSession) GetPositionCollateralRatio(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GetPositionCollateralRatio(&_Contracts.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_Contracts *ContractsTransactor) GetPositionDebt(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "getPositionDebt", accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_Contracts *ContractsSession) GetPositionDebt(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GetPositionDebt(&_Contracts.TransactOpts, accountId, poolId, collateralType)
}

// GetPositionDebt is a paid mutator transaction binding the contract method 0x3593bbd2.
//
// Solidity: function getPositionDebt(uint128 accountId, uint128 poolId, address collateralType) returns(int256 debt)
func (_Contracts *ContractsTransactorSession) GetPositionDebt(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GetPositionDebt(&_Contracts.TransactOpts, accountId, poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_Contracts *ContractsTransactor) GetVaultCollateralRatio(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "getVaultCollateralRatio", poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_Contracts *ContractsSession) GetVaultCollateralRatio(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GetVaultCollateralRatio(&_Contracts.TransactOpts, poolId, collateralType)
}

// GetVaultCollateralRatio is a paid mutator transaction binding the contract method 0x60248c55.
//
// Solidity: function getVaultCollateralRatio(uint128 poolId, address collateralType) returns(uint256)
func (_Contracts *ContractsTransactorSession) GetVaultCollateralRatio(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GetVaultCollateralRatio(&_Contracts.TransactOpts, poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_Contracts *ContractsTransactor) GetVaultDebt(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "getVaultDebt", poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_Contracts *ContractsSession) GetVaultDebt(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GetVaultDebt(&_Contracts.TransactOpts, poolId, collateralType)
}

// GetVaultDebt is a paid mutator transaction binding the contract method 0x2fb8ff24.
//
// Solidity: function getVaultDebt(uint128 poolId, address collateralType) returns(int256)
func (_Contracts *ContractsTransactorSession) GetVaultDebt(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GetVaultDebt(&_Contracts.TransactOpts, poolId, collateralType)
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

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_Contracts *ContractsTransactor) IsPositionLiquidatable(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "isPositionLiquidatable", accountId, poolId, collateralType)
}

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_Contracts *ContractsSession) IsPositionLiquidatable(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IsPositionLiquidatable(&_Contracts.TransactOpts, accountId, poolId, collateralType)
}

// IsPositionLiquidatable is a paid mutator transaction binding the contract method 0x2fa7bb65.
//
// Solidity: function isPositionLiquidatable(uint128 accountId, uint128 poolId, address collateralType) returns(bool)
func (_Contracts *ContractsTransactorSession) IsPositionLiquidatable(accountId *big.Int, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IsPositionLiquidatable(&_Contracts.TransactOpts, accountId, poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_Contracts *ContractsTransactor) IsVaultLiquidatable(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "isVaultLiquidatable", poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_Contracts *ContractsSession) IsVaultLiquidatable(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IsVaultLiquidatable(&_Contracts.TransactOpts, poolId, collateralType)
}

// IsVaultLiquidatable is a paid mutator transaction binding the contract method 0x2a5354d2.
//
// Solidity: function isVaultLiquidatable(uint128 poolId, address collateralType) returns(bool)
func (_Contracts *ContractsTransactorSession) IsVaultLiquidatable(poolId *big.Int, collateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IsVaultLiquidatable(&_Contracts.TransactOpts, poolId, collateralType)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_Contracts *ContractsTransactor) Liquidate(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "liquidate", accountId, poolId, collateralType, liquidateAsAccountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_Contracts *ContractsSession) Liquidate(accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Liquidate(&_Contracts.TransactOpts, accountId, poolId, collateralType, liquidateAsAccountId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x3e033a06.
//
// Solidity: function liquidate(uint128 accountId, uint128 poolId, address collateralType, uint128 liquidateAsAccountId) returns((uint256,uint256,uint256) liquidationData)
func (_Contracts *ContractsTransactorSession) Liquidate(accountId *big.Int, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Liquidate(&_Contracts.TransactOpts, accountId, poolId, collateralType, liquidateAsAccountId)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_Contracts *ContractsTransactor) LiquidateVault(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "liquidateVault", poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_Contracts *ContractsSession) LiquidateVault(poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.LiquidateVault(&_Contracts.TransactOpts, poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// LiquidateVault is a paid mutator transaction binding the contract method 0x7d8a4140.
//
// Solidity: function liquidateVault(uint128 poolId, address collateralType, uint128 liquidateAsAccountId, uint256 maxUsd) returns((uint256,uint256,uint256) liquidationData)
func (_Contracts *ContractsTransactorSession) LiquidateVault(poolId *big.Int, collateralType common.Address, liquidateAsAccountId *big.Int, maxUsd *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.LiquidateVault(&_Contracts.TransactOpts, poolId, collateralType, liquidateAsAccountId, maxUsd)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Contracts *ContractsTransactor) MintUsd(opts *bind.TransactOpts, accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "mintUsd", accountId, poolId, collateralType, amount)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Contracts *ContractsSession) MintUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.MintUsd(&_Contracts.TransactOpts, accountId, poolId, collateralType, amount)
}

// MintUsd is a paid mutator transaction binding the contract method 0xdf16a074.
//
// Solidity: function mintUsd(uint128 accountId, uint128 poolId, address collateralType, uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) MintUsd(accountId *big.Int, poolId *big.Int, collateralType common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.MintUsd(&_Contracts.TransactOpts, accountId, poolId, collateralType, amount)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Contracts *ContractsTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Contracts *ContractsSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Multicall(&_Contracts.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Contracts *ContractsTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Multicall(&_Contracts.TransactOpts, data)
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

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_Contracts *ContractsTransactor) NominatePoolOwner(opts *bind.TransactOpts, nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "nominatePoolOwner", nominatedOwner, poolId)
}

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_Contracts *ContractsSession) NominatePoolOwner(nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NominatePoolOwner(&_Contracts.TransactOpts, nominatedOwner, poolId)
}

// NominatePoolOwner is a paid mutator transaction binding the contract method 0x6141f7a2.
//
// Solidity: function nominatePoolOwner(address nominatedOwner, uint128 poolId) returns()
func (_Contracts *ContractsTransactorSession) NominatePoolOwner(nominatedOwner common.Address, poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NominatePoolOwner(&_Contracts.TransactOpts, nominatedOwner, poolId)
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

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_Contracts *ContractsTransactor) RebalancePool(opts *bind.TransactOpts, poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "rebalancePool", poolId, optionalCollateralType)
}

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_Contracts *ContractsSession) RebalancePool(poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RebalancePool(&_Contracts.TransactOpts, poolId, optionalCollateralType)
}

// RebalancePool is a paid mutator transaction binding the contract method 0x183231d7.
//
// Solidity: function rebalancePool(uint128 poolId, address optionalCollateralType) returns()
func (_Contracts *ContractsTransactorSession) RebalancePool(poolId *big.Int, optionalCollateralType common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RebalancePool(&_Contracts.TransactOpts, poolId, optionalCollateralType)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_Contracts *ContractsTransactor) RegisterMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerMarket", market)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_Contracts *ContractsSession) RegisterMarket(market common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterMarket(&_Contracts.TransactOpts, market)
}

// RegisterMarket is a paid mutator transaction binding the contract method 0xa79b9ec9.
//
// Solidity: function registerMarket(address market) returns(uint128 marketId)
func (_Contracts *ContractsTransactorSession) RegisterMarket(market common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterMarket(&_Contracts.TransactOpts, market)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Contracts *ContractsTransactor) RegisterRewardsDistributor(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerRewardsDistributor", poolId, collateralType, distributor)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Contracts *ContractsSession) RegisterRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterRewardsDistributor(&_Contracts.TransactOpts, poolId, collateralType, distributor)
}

// RegisterRewardsDistributor is a paid mutator transaction binding the contract method 0x170c1351.
//
// Solidity: function registerRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Contracts *ContractsTransactorSession) RegisterRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterRewardsDistributor(&_Contracts.TransactOpts, poolId, collateralType, distributor)
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

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_Contracts *ContractsTransactor) RemoveApprovedPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeApprovedPool", poolId)
}

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_Contracts *ContractsSession) RemoveApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveApprovedPool(&_Contracts.TransactOpts, poolId)
}

// RemoveApprovedPool is a paid mutator transaction binding the contract method 0xe1b440d0.
//
// Solidity: function removeApprovedPool(uint128 poolId) returns()
func (_Contracts *ContractsTransactorSession) RemoveApprovedPool(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveApprovedPool(&_Contracts.TransactOpts, poolId)
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

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Contracts *ContractsTransactor) RemoveRewardsDistributor(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeRewardsDistributor", poolId, collateralType, distributor)
}

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Contracts *ContractsSession) RemoveRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveRewardsDistributor(&_Contracts.TransactOpts, poolId, collateralType, distributor)
}

// RemoveRewardsDistributor is a paid mutator transaction binding the contract method 0x2685f42b.
//
// Solidity: function removeRewardsDistributor(uint128 poolId, address collateralType, address distributor) returns()
func (_Contracts *ContractsTransactorSession) RemoveRewardsDistributor(poolId *big.Int, collateralType common.Address, distributor common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveRewardsDistributor(&_Contracts.TransactOpts, poolId, collateralType, distributor)
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

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_Contracts *ContractsTransactor) RenouncePoolNomination(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renouncePoolNomination", poolId)
}

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_Contracts *ContractsSession) RenouncePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RenouncePoolNomination(&_Contracts.TransactOpts, poolId)
}

// RenouncePoolNomination is a paid mutator transaction binding the contract method 0xca5bed77.
//
// Solidity: function renouncePoolNomination(uint128 poolId) returns()
func (_Contracts *ContractsTransactorSession) RenouncePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RenouncePoolNomination(&_Contracts.TransactOpts, poolId)
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

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_Contracts *ContractsTransactor) RevokePoolNomination(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "revokePoolNomination", poolId)
}

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_Contracts *ContractsSession) RevokePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RevokePoolNomination(&_Contracts.TransactOpts, poolId)
}

// RevokePoolNomination is a paid mutator transaction binding the contract method 0x1f1b33b9.
//
// Solidity: function revokePoolNomination(uint128 poolId) returns()
func (_Contracts *ContractsTransactorSession) RevokePoolNomination(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RevokePoolNomination(&_Contracts.TransactOpts, poolId)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_Contracts *ContractsTransactor) SetConfig(opts *bind.TransactOpts, k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setConfig", k, v)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_Contracts *ContractsSession) SetConfig(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetConfig(&_Contracts.TransactOpts, k, v)
}

// SetConfig is a paid mutator transaction binding the contract method 0xd1fd27b3.
//
// Solidity: function setConfig(bytes32 k, bytes32 v) returns()
func (_Contracts *ContractsTransactorSession) SetConfig(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetConfig(&_Contracts.TransactOpts, k, v)
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

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_Contracts *ContractsTransactor) SetMarketMinDelegateTime(opts *bind.TransactOpts, marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMarketMinDelegateTime", marketId, minDelegateTime)
}

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_Contracts *ContractsSession) SetMarketMinDelegateTime(marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _Contracts.Contract.SetMarketMinDelegateTime(&_Contracts.TransactOpts, marketId, minDelegateTime)
}

// SetMarketMinDelegateTime is a paid mutator transaction binding the contract method 0x1d90e392.
//
// Solidity: function setMarketMinDelegateTime(uint128 marketId, uint32 minDelegateTime) returns()
func (_Contracts *ContractsTransactorSession) SetMarketMinDelegateTime(marketId *big.Int, minDelegateTime uint32) (*types.Transaction, error) {
	return _Contracts.Contract.SetMarketMinDelegateTime(&_Contracts.TransactOpts, marketId, minDelegateTime)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_Contracts *ContractsTransactor) SetMinLiquidityRatio(opts *bind.TransactOpts, marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMinLiquidityRatio", marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_Contracts *ContractsSession) SetMinLiquidityRatio(marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinLiquidityRatio(&_Contracts.TransactOpts, marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio is a paid mutator transaction binding the contract method 0x6fd5bdce.
//
// Solidity: function setMinLiquidityRatio(uint128 marketId, uint256 minLiquidityRatio) returns()
func (_Contracts *ContractsTransactorSession) SetMinLiquidityRatio(marketId *big.Int, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinLiquidityRatio(&_Contracts.TransactOpts, marketId, minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_Contracts *ContractsTransactor) SetMinLiquidityRatio0(opts *bind.TransactOpts, minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMinLiquidityRatio0", minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_Contracts *ContractsSession) SetMinLiquidityRatio0(minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinLiquidityRatio0(&_Contracts.TransactOpts, minLiquidityRatio)
}

// SetMinLiquidityRatio0 is a paid mutator transaction binding the contract method 0x34078a01.
//
// Solidity: function setMinLiquidityRatio(uint256 minLiquidityRatio) returns()
func (_Contracts *ContractsTransactorSession) SetMinLiquidityRatio0(minLiquidityRatio *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinLiquidityRatio0(&_Contracts.TransactOpts, minLiquidityRatio)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_Contracts *ContractsTransactor) SetPoolConfiguration(opts *bind.TransactOpts, poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setPoolConfiguration", poolId, newMarketConfigurations)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_Contracts *ContractsSession) SetPoolConfiguration(poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _Contracts.Contract.SetPoolConfiguration(&_Contracts.TransactOpts, poolId, newMarketConfigurations)
}

// SetPoolConfiguration is a paid mutator transaction binding the contract method 0x5d8c8844.
//
// Solidity: function setPoolConfiguration(uint128 poolId, (uint128,uint128,int128)[] newMarketConfigurations) returns()
func (_Contracts *ContractsTransactorSession) SetPoolConfiguration(poolId *big.Int, newMarketConfigurations []MarketConfigurationData) (*types.Transaction, error) {
	return _Contracts.Contract.SetPoolConfiguration(&_Contracts.TransactOpts, poolId, newMarketConfigurations)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_Contracts *ContractsTransactor) SetPoolName(opts *bind.TransactOpts, poolId *big.Int, name string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setPoolName", poolId, name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_Contracts *ContractsSession) SetPoolName(poolId *big.Int, name string) (*types.Transaction, error) {
	return _Contracts.Contract.SetPoolName(&_Contracts.TransactOpts, poolId, name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x11e72a43.
//
// Solidity: function setPoolName(uint128 poolId, string name) returns()
func (_Contracts *ContractsTransactorSession) SetPoolName(poolId *big.Int, name string) (*types.Transaction, error) {
	return _Contracts.Contract.SetPoolName(&_Contracts.TransactOpts, poolId, name)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_Contracts *ContractsTransactor) SetPreferredPool(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setPreferredPool", poolId)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_Contracts *ContractsSession) SetPreferredPool(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetPreferredPool(&_Contracts.TransactOpts, poolId)
}

// SetPreferredPool is a paid mutator transaction binding the contract method 0xe7098c0c.
//
// Solidity: function setPreferredPool(uint128 poolId) returns()
func (_Contracts *ContractsTransactorSession) SetPreferredPool(poolId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetPreferredPool(&_Contracts.TransactOpts, poolId)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_Contracts *ContractsTransactor) SetSupportedCrossChainNetworks(opts *bind.TransactOpts, supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setSupportedCrossChainNetworks", supportedNetworks, ccipSelectors)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_Contracts *ContractsSession) SetSupportedCrossChainNetworks(supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _Contracts.Contract.SetSupportedCrossChainNetworks(&_Contracts.TransactOpts, supportedNetworks, ccipSelectors)
}

// SetSupportedCrossChainNetworks is a paid mutator transaction binding the contract method 0x830e23b5.
//
// Solidity: function setSupportedCrossChainNetworks(uint64[] supportedNetworks, uint64[] ccipSelectors) returns(uint256 numRegistered)
func (_Contracts *ContractsTransactorSession) SetSupportedCrossChainNetworks(supportedNetworks []uint64, ccipSelectors []uint64) (*types.Transaction, error) {
	return _Contracts.Contract.SetSupportedCrossChainNetworks(&_Contracts.TransactOpts, supportedNetworks, ccipSelectors)
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

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_Contracts *ContractsTransactor) TransferCrossChain(opts *bind.TransactOpts, destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferCrossChain", destChainId, amount)
}

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_Contracts *ContractsSession) TransferCrossChain(destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferCrossChain(&_Contracts.TransactOpts, destChainId, amount)
}

// TransferCrossChain is a paid mutator transaction binding the contract method 0x340824d7.
//
// Solidity: function transferCrossChain(uint64 destChainId, uint256 amount) payable returns(uint256 gasTokenUsed)
func (_Contracts *ContractsTransactorSession) TransferCrossChain(destChainId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferCrossChain(&_Contracts.TransactOpts, destChainId, amount)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_Contracts *ContractsTransactor) UpdateRewards(opts *bind.TransactOpts, poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateRewards", poolId, collateralType, accountId)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_Contracts *ContractsSession) UpdateRewards(poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateRewards(&_Contracts.TransactOpts, poolId, collateralType, accountId)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x645657d8.
//
// Solidity: function updateRewards(uint128 poolId, address collateralType, uint128 accountId) returns(uint256[], address[])
func (_Contracts *ContractsTransactorSession) UpdateRewards(poolId *big.Int, collateralType common.Address, accountId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateRewards(&_Contracts.TransactOpts, poolId, collateralType, accountId)
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

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsTransactor) Withdraw(opts *bind.TransactOpts, accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdraw", accountId, collateralType, tokenAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsSession) Withdraw(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, accountId, collateralType, tokenAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x95997c51.
//
// Solidity: function withdraw(uint128 accountId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsTransactorSession) Withdraw(accountId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, accountId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsTransactor) WithdrawMarketCollateral(opts *bind.TransactOpts, marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawMarketCollateral", marketId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsSession) WithdrawMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawMarketCollateral(&_Contracts.TransactOpts, marketId, collateralType, tokenAmount)
}

// WithdrawMarketCollateral is a paid mutator transaction binding the contract method 0xa3aa8b51.
//
// Solidity: function withdrawMarketCollateral(uint128 marketId, address collateralType, uint256 tokenAmount) returns()
func (_Contracts *ContractsTransactorSession) WithdrawMarketCollateral(marketId *big.Int, collateralType common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawMarketCollateral(&_Contracts.TransactOpts, marketId, collateralType, tokenAmount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Contracts *ContractsTransactor) WithdrawMarketUsd(opts *bind.TransactOpts, marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawMarketUsd", marketId, target, amount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Contracts *ContractsSession) WithdrawMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawMarketUsd(&_Contracts.TransactOpts, marketId, target, amount)
}

// WithdrawMarketUsd is a paid mutator transaction binding the contract method 0x140a7cfe.
//
// Solidity: function withdrawMarketUsd(uint128 marketId, address target, uint256 amount) returns(uint256 feeAmount)
func (_Contracts *ContractsTransactorSession) WithdrawMarketUsd(marketId *big.Int, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawMarketUsd(&_Contracts.TransactOpts, marketId, target, amount)
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

// ContractsCollateralConfiguredIterator is returned from FilterCollateralConfigured and is used to iterate over the raw logs and unpacked data for CollateralConfigured events raised by the Contracts contract.
type ContractsCollateralConfiguredIterator struct {
	Event *ContractsCollateralConfigured // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsCollateralConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCollateralConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsCollateralConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsCollateralConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCollateralConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCollateralConfigured represents a CollateralConfigured event raised by the Contracts contract.
type ContractsCollateralConfigured struct {
	CollateralType common.Address
	Config         CollateralConfigurationData
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCollateralConfigured is a free log retrieval operation binding the contract event 0xefc23317f58afd6b22480bd22174cc7da0913bce25c03d9859216dacddebe6fe.
//
// Solidity: event CollateralConfigured(address indexed collateralType, (bool,uint256,uint256,uint256,bytes32,address,uint256) config)
func (_Contracts *ContractsFilterer) FilterCollateralConfigured(opts *bind.FilterOpts, collateralType []common.Address) (*ContractsCollateralConfiguredIterator, error) {

	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CollateralConfigured", collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsCollateralConfiguredIterator{contract: _Contracts.contract, event: "CollateralConfigured", logs: logs, sub: sub}, nil
}

// WatchCollateralConfigured is a free log subscription operation binding the contract event 0xefc23317f58afd6b22480bd22174cc7da0913bce25c03d9859216dacddebe6fe.
//
// Solidity: event CollateralConfigured(address indexed collateralType, (bool,uint256,uint256,uint256,bytes32,address,uint256) config)
func (_Contracts *ContractsFilterer) WatchCollateralConfigured(opts *bind.WatchOpts, sink chan<- *ContractsCollateralConfigured, collateralType []common.Address) (event.Subscription, error) {

	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CollateralConfigured", collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCollateralConfigured)
				if err := _Contracts.contract.UnpackLog(event, "CollateralConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseCollateralConfigured(log types.Log) (*ContractsCollateralConfigured, error) {
	event := new(ContractsCollateralConfigured)
	if err := _Contracts.contract.UnpackLog(event, "CollateralConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCollateralLockCreatedIterator is returned from FilterCollateralLockCreated and is used to iterate over the raw logs and unpacked data for CollateralLockCreated events raised by the Contracts contract.
type ContractsCollateralLockCreatedIterator struct {
	Event *ContractsCollateralLockCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsCollateralLockCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCollateralLockCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsCollateralLockCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsCollateralLockCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCollateralLockCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCollateralLockCreated represents a CollateralLockCreated event raised by the Contracts contract.
type ContractsCollateralLockCreated struct {
	AccountId       *big.Int
	CollateralType  common.Address
	TokenAmount     *big.Int
	ExpireTimestamp uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollateralLockCreated is a free log retrieval operation binding the contract event 0x8a78446a234220d3ee3f46aa6ea1ea5bc438bd153398ebcbd171843744b452a6.
//
// Solidity: event CollateralLockCreated(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_Contracts *ContractsFilterer) FilterCollateralLockCreated(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address) (*ContractsCollateralLockCreatedIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CollateralLockCreated", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsCollateralLockCreatedIterator{contract: _Contracts.contract, event: "CollateralLockCreated", logs: logs, sub: sub}, nil
}

// WatchCollateralLockCreated is a free log subscription operation binding the contract event 0x8a78446a234220d3ee3f46aa6ea1ea5bc438bd153398ebcbd171843744b452a6.
//
// Solidity: event CollateralLockCreated(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_Contracts *ContractsFilterer) WatchCollateralLockCreated(opts *bind.WatchOpts, sink chan<- *ContractsCollateralLockCreated, accountId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CollateralLockCreated", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCollateralLockCreated)
				if err := _Contracts.contract.UnpackLog(event, "CollateralLockCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseCollateralLockCreated(log types.Log) (*ContractsCollateralLockCreated, error) {
	event := new(ContractsCollateralLockCreated)
	if err := _Contracts.contract.UnpackLog(event, "CollateralLockCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCollateralLockExpiredIterator is returned from FilterCollateralLockExpired and is used to iterate over the raw logs and unpacked data for CollateralLockExpired events raised by the Contracts contract.
type ContractsCollateralLockExpiredIterator struct {
	Event *ContractsCollateralLockExpired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsCollateralLockExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCollateralLockExpired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsCollateralLockExpired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsCollateralLockExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCollateralLockExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCollateralLockExpired represents a CollateralLockExpired event raised by the Contracts contract.
type ContractsCollateralLockExpired struct {
	AccountId       *big.Int
	CollateralType  common.Address
	TokenAmount     *big.Int
	ExpireTimestamp uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollateralLockExpired is a free log retrieval operation binding the contract event 0xc5e2c7afc4e54d998977e11260a0bfc0ad5678a3a8b6628162f9d4e642d7f160.
//
// Solidity: event CollateralLockExpired(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_Contracts *ContractsFilterer) FilterCollateralLockExpired(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address) (*ContractsCollateralLockExpiredIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CollateralLockExpired", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsCollateralLockExpiredIterator{contract: _Contracts.contract, event: "CollateralLockExpired", logs: logs, sub: sub}, nil
}

// WatchCollateralLockExpired is a free log subscription operation binding the contract event 0xc5e2c7afc4e54d998977e11260a0bfc0ad5678a3a8b6628162f9d4e642d7f160.
//
// Solidity: event CollateralLockExpired(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, uint64 expireTimestamp)
func (_Contracts *ContractsFilterer) WatchCollateralLockExpired(opts *bind.WatchOpts, sink chan<- *ContractsCollateralLockExpired, accountId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CollateralLockExpired", accountIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCollateralLockExpired)
				if err := _Contracts.contract.UnpackLog(event, "CollateralLockExpired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseCollateralLockExpired(log types.Log) (*ContractsCollateralLockExpired, error) {
	event := new(ContractsCollateralLockExpired)
	if err := _Contracts.contract.UnpackLog(event, "CollateralLockExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDebtAssociatedIterator is returned from FilterDebtAssociated and is used to iterate over the raw logs and unpacked data for DebtAssociated events raised by the Contracts contract.
type ContractsDebtAssociatedIterator struct {
	Event *ContractsDebtAssociated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsDebtAssociatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDebtAssociated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsDebtAssociated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsDebtAssociatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDebtAssociatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDebtAssociated represents a DebtAssociated event raised by the Contracts contract.
type ContractsDebtAssociated struct {
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
func (_Contracts *ContractsFilterer) FilterDebtAssociated(opts *bind.FilterOpts, marketId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*ContractsDebtAssociatedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "DebtAssociated", marketIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsDebtAssociatedIterator{contract: _Contracts.contract, event: "DebtAssociated", logs: logs, sub: sub}, nil
}

// WatchDebtAssociated is a free log subscription operation binding the contract event 0xb03bc7530b5650601d2c8bc86c81a45b6b50b0099defd3f17a6bdb48660180ad.
//
// Solidity: event DebtAssociated(uint128 indexed marketId, uint128 indexed poolId, address indexed collateralType, uint128 accountId, uint256 amount, int256 updatedDebt)
func (_Contracts *ContractsFilterer) WatchDebtAssociated(opts *bind.WatchOpts, sink chan<- *ContractsDebtAssociated, marketId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "DebtAssociated", marketIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDebtAssociated)
				if err := _Contracts.contract.UnpackLog(event, "DebtAssociated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseDebtAssociated(log types.Log) (*ContractsDebtAssociated, error) {
	event := new(ContractsDebtAssociated)
	if err := _Contracts.contract.UnpackLog(event, "DebtAssociated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDelegationUpdatedIterator is returned from FilterDelegationUpdated and is used to iterate over the raw logs and unpacked data for DelegationUpdated events raised by the Contracts contract.
type ContractsDelegationUpdatedIterator struct {
	Event *ContractsDelegationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsDelegationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDelegationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsDelegationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsDelegationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDelegationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDelegationUpdated represents a DelegationUpdated event raised by the Contracts contract.
type ContractsDelegationUpdated struct {
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
func (_Contracts *ContractsFilterer) FilterDelegationUpdated(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*ContractsDelegationUpdatedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "DelegationUpdated", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsDelegationUpdatedIterator{contract: _Contracts.contract, event: "DelegationUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegationUpdated is a free log subscription operation binding the contract event 0x7b12dd38f18c0ff77ae702f6da13fbbcb28f53f807ecc7d39ee8d8b1ea8295ad.
//
// Solidity: event DelegationUpdated(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, uint256 leverage, address indexed sender)
func (_Contracts *ContractsFilterer) WatchDelegationUpdated(opts *bind.WatchOpts, sink chan<- *ContractsDelegationUpdated, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "DelegationUpdated", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDelegationUpdated)
				if err := _Contracts.contract.UnpackLog(event, "DelegationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseDelegationUpdated(log types.Log) (*ContractsDelegationUpdated, error) {
	event := new(ContractsDelegationUpdated)
	if err := _Contracts.contract.UnpackLog(event, "DelegationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Contracts contract.
type ContractsDepositedIterator struct {
	Event *ContractsDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDeposited represents a Deposited event raised by the Contracts contract.
type ContractsDeposited struct {
	AccountId      *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xd92122e67326e9313bfae33ccb1fccf5194584c6bf93a8529a6b006d8c6e24a9.
//
// Solidity: event Deposited(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Contracts *ContractsFilterer) FilterDeposited(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (*ContractsDepositedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Deposited", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsDepositedIterator{contract: _Contracts.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xd92122e67326e9313bfae33ccb1fccf5194584c6bf93a8529a6b006d8c6e24a9.
//
// Solidity: event Deposited(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Contracts *ContractsFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ContractsDeposited, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Deposited", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDeposited)
				if err := _Contracts.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseDeposited(log types.Log) (*ContractsDeposited, error) {
	event := new(ContractsDeposited)
	if err := _Contracts.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ContractsIssuanceFeePaidIterator is returned from FilterIssuanceFeePaid and is used to iterate over the raw logs and unpacked data for IssuanceFeePaid events raised by the Contracts contract.
type ContractsIssuanceFeePaidIterator struct {
	Event *ContractsIssuanceFeePaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsIssuanceFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsIssuanceFeePaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsIssuanceFeePaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsIssuanceFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsIssuanceFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsIssuanceFeePaid represents a IssuanceFeePaid event raised by the Contracts contract.
type ContractsIssuanceFeePaid struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	FeeAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIssuanceFeePaid is a free log retrieval operation binding the contract event 0x28d0fb10e1c8ce51490a16fb3b40bf23f8064f7c624d3a3852ad66683a25995d.
//
// Solidity: event IssuanceFeePaid(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 feeAmount)
func (_Contracts *ContractsFilterer) FilterIssuanceFeePaid(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int) (*ContractsIssuanceFeePaidIterator, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "IssuanceFeePaid", accountIdRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsIssuanceFeePaidIterator{contract: _Contracts.contract, event: "IssuanceFeePaid", logs: logs, sub: sub}, nil
}

// WatchIssuanceFeePaid is a free log subscription operation binding the contract event 0x28d0fb10e1c8ce51490a16fb3b40bf23f8064f7c624d3a3852ad66683a25995d.
//
// Solidity: event IssuanceFeePaid(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 feeAmount)
func (_Contracts *ContractsFilterer) WatchIssuanceFeePaid(opts *bind.WatchOpts, sink chan<- *ContractsIssuanceFeePaid, accountId []*big.Int, poolId []*big.Int) (event.Subscription, error) {

	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "IssuanceFeePaid", accountIdRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsIssuanceFeePaid)
				if err := _Contracts.contract.UnpackLog(event, "IssuanceFeePaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseIssuanceFeePaid(log types.Log) (*ContractsIssuanceFeePaid, error) {
	event := new(ContractsIssuanceFeePaid)
	if err := _Contracts.contract.UnpackLog(event, "IssuanceFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLiquidationIterator is returned from FilterLiquidation and is used to iterate over the raw logs and unpacked data for Liquidation events raised by the Contracts contract.
type ContractsLiquidationIterator struct {
	Event *ContractsLiquidation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLiquidation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsLiquidation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLiquidation represents a Liquidation event raised by the Contracts contract.
type ContractsLiquidation struct {
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
func (_Contracts *ContractsFilterer) FilterLiquidation(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*ContractsLiquidationIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Liquidation", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLiquidationIterator{contract: _Contracts.contract, event: "Liquidation", logs: logs, sub: sub}, nil
}

// WatchLiquidation is a free log subscription operation binding the contract event 0xe6c1b26644f880854bf954d4186be9f0b2d06d50fa0484b596e79d409c07a5fd.
//
// Solidity: event Liquidation(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_Contracts *ContractsFilterer) WatchLiquidation(opts *bind.WatchOpts, sink chan<- *ContractsLiquidation, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Liquidation", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLiquidation)
				if err := _Contracts.contract.UnpackLog(event, "Liquidation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseLiquidation(log types.Log) (*ContractsLiquidation, error) {
	event := new(ContractsLiquidation)
	if err := _Contracts.contract.UnpackLog(event, "Liquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketCollateralDepositedIterator is returned from FilterMarketCollateralDeposited and is used to iterate over the raw logs and unpacked data for MarketCollateralDeposited events raised by the Contracts contract.
type ContractsMarketCollateralDepositedIterator struct {
	Event *ContractsMarketCollateralDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsMarketCollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketCollateralDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsMarketCollateralDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsMarketCollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketCollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketCollateralDeposited represents a MarketCollateralDeposited event raised by the Contracts contract.
type ContractsMarketCollateralDeposited struct {
	MarketId       *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketCollateralDeposited is a free log retrieval operation binding the contract event 0x6affa35c1a0ba6198bd74578c93e3a0eb32a96af0af26ce8adf06b933212ca1b.
//
// Solidity: event MarketCollateralDeposited(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Contracts *ContractsFilterer) FilterMarketCollateralDeposited(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (*ContractsMarketCollateralDepositedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketCollateralDeposited", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketCollateralDepositedIterator{contract: _Contracts.contract, event: "MarketCollateralDeposited", logs: logs, sub: sub}, nil
}

// WatchMarketCollateralDeposited is a free log subscription operation binding the contract event 0x6affa35c1a0ba6198bd74578c93e3a0eb32a96af0af26ce8adf06b933212ca1b.
//
// Solidity: event MarketCollateralDeposited(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Contracts *ContractsFilterer) WatchMarketCollateralDeposited(opts *bind.WatchOpts, sink chan<- *ContractsMarketCollateralDeposited, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketCollateralDeposited", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketCollateralDeposited)
				if err := _Contracts.contract.UnpackLog(event, "MarketCollateralDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseMarketCollateralDeposited(log types.Log) (*ContractsMarketCollateralDeposited, error) {
	event := new(ContractsMarketCollateralDeposited)
	if err := _Contracts.contract.UnpackLog(event, "MarketCollateralDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketCollateralWithdrawnIterator is returned from FilterMarketCollateralWithdrawn and is used to iterate over the raw logs and unpacked data for MarketCollateralWithdrawn events raised by the Contracts contract.
type ContractsMarketCollateralWithdrawnIterator struct {
	Event *ContractsMarketCollateralWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsMarketCollateralWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketCollateralWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsMarketCollateralWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsMarketCollateralWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketCollateralWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketCollateralWithdrawn represents a MarketCollateralWithdrawn event raised by the Contracts contract.
type ContractsMarketCollateralWithdrawn struct {
	MarketId       *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketCollateralWithdrawn is a free log retrieval operation binding the contract event 0x31fd88ca973dd7ee814313c50c2aaa7cd2d2142634582332b1e2671e6d97fd9a.
//
// Solidity: event MarketCollateralWithdrawn(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Contracts *ContractsFilterer) FilterMarketCollateralWithdrawn(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (*ContractsMarketCollateralWithdrawnIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketCollateralWithdrawn", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketCollateralWithdrawnIterator{contract: _Contracts.contract, event: "MarketCollateralWithdrawn", logs: logs, sub: sub}, nil
}

// WatchMarketCollateralWithdrawn is a free log subscription operation binding the contract event 0x31fd88ca973dd7ee814313c50c2aaa7cd2d2142634582332b1e2671e6d97fd9a.
//
// Solidity: event MarketCollateralWithdrawn(uint128 indexed marketId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Contracts *ContractsFilterer) WatchMarketCollateralWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractsMarketCollateralWithdrawn, marketId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketCollateralWithdrawn", marketIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketCollateralWithdrawn)
				if err := _Contracts.contract.UnpackLog(event, "MarketCollateralWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseMarketCollateralWithdrawn(log types.Log) (*ContractsMarketCollateralWithdrawn, error) {
	event := new(ContractsMarketCollateralWithdrawn)
	if err := _Contracts.contract.UnpackLog(event, "MarketCollateralWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketRegisteredIterator is returned from FilterMarketRegistered and is used to iterate over the raw logs and unpacked data for MarketRegistered events raised by the Contracts contract.
type ContractsMarketRegisteredIterator struct {
	Event *ContractsMarketRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsMarketRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsMarketRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsMarketRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketRegistered represents a MarketRegistered event raised by the Contracts contract.
type ContractsMarketRegistered struct {
	Market   common.Address
	MarketId *big.Int
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketRegistered is a free log retrieval operation binding the contract event 0xeb87361ace8c1947e121293eb214f68d889d9cf273c48246b38c3cbf185748d0.
//
// Solidity: event MarketRegistered(address indexed market, uint128 indexed marketId, address indexed sender)
func (_Contracts *ContractsFilterer) FilterMarketRegistered(opts *bind.FilterOpts, market []common.Address, marketId []*big.Int, sender []common.Address) (*ContractsMarketRegisteredIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketRegistered", marketRule, marketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketRegisteredIterator{contract: _Contracts.contract, event: "MarketRegistered", logs: logs, sub: sub}, nil
}

// WatchMarketRegistered is a free log subscription operation binding the contract event 0xeb87361ace8c1947e121293eb214f68d889d9cf273c48246b38c3cbf185748d0.
//
// Solidity: event MarketRegistered(address indexed market, uint128 indexed marketId, address indexed sender)
func (_Contracts *ContractsFilterer) WatchMarketRegistered(opts *bind.WatchOpts, sink chan<- *ContractsMarketRegistered, market []common.Address, marketId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketRegistered", marketRule, marketIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketRegistered)
				if err := _Contracts.contract.UnpackLog(event, "MarketRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseMarketRegistered(log types.Log) (*ContractsMarketRegistered, error) {
	event := new(ContractsMarketRegistered)
	if err := _Contracts.contract.UnpackLog(event, "MarketRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketSystemFeePaidIterator is returned from FilterMarketSystemFeePaid and is used to iterate over the raw logs and unpacked data for MarketSystemFeePaid events raised by the Contracts contract.
type ContractsMarketSystemFeePaidIterator struct {
	Event *ContractsMarketSystemFeePaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsMarketSystemFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketSystemFeePaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsMarketSystemFeePaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsMarketSystemFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketSystemFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketSystemFeePaid represents a MarketSystemFeePaid event raised by the Contracts contract.
type ContractsMarketSystemFeePaid struct {
	MarketId  *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMarketSystemFeePaid is a free log retrieval operation binding the contract event 0x8b69fed8aed97ef9572216662359ece45fa52f2b5ff44a78b7ec3c5ef05153f8.
//
// Solidity: event MarketSystemFeePaid(uint128 indexed marketId, uint256 feeAmount)
func (_Contracts *ContractsFilterer) FilterMarketSystemFeePaid(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsMarketSystemFeePaidIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketSystemFeePaid", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketSystemFeePaidIterator{contract: _Contracts.contract, event: "MarketSystemFeePaid", logs: logs, sub: sub}, nil
}

// WatchMarketSystemFeePaid is a free log subscription operation binding the contract event 0x8b69fed8aed97ef9572216662359ece45fa52f2b5ff44a78b7ec3c5ef05153f8.
//
// Solidity: event MarketSystemFeePaid(uint128 indexed marketId, uint256 feeAmount)
func (_Contracts *ContractsFilterer) WatchMarketSystemFeePaid(opts *bind.WatchOpts, sink chan<- *ContractsMarketSystemFeePaid, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketSystemFeePaid", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketSystemFeePaid)
				if err := _Contracts.contract.UnpackLog(event, "MarketSystemFeePaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseMarketSystemFeePaid(log types.Log) (*ContractsMarketSystemFeePaid, error) {
	event := new(ContractsMarketSystemFeePaid)
	if err := _Contracts.contract.UnpackLog(event, "MarketSystemFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketUsdDepositedIterator is returned from FilterMarketUsdDeposited and is used to iterate over the raw logs and unpacked data for MarketUsdDeposited events raised by the Contracts contract.
type ContractsMarketUsdDepositedIterator struct {
	Event *ContractsMarketUsdDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsMarketUsdDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketUsdDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsMarketUsdDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsMarketUsdDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketUsdDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketUsdDeposited represents a MarketUsdDeposited event raised by the Contracts contract.
type ContractsMarketUsdDeposited struct {
	MarketId *big.Int
	Target   common.Address
	Amount   *big.Int
	Market   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketUsdDeposited is a free log retrieval operation binding the contract event 0xcbdbec7c74466ca0e764de5f6bca5bef859e31a7d6284be22ac6b4beabddf835.
//
// Solidity: event MarketUsdDeposited(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_Contracts *ContractsFilterer) FilterMarketUsdDeposited(opts *bind.FilterOpts, marketId []*big.Int, target []common.Address, market []common.Address) (*ContractsMarketUsdDepositedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketUsdDeposited", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketUsdDepositedIterator{contract: _Contracts.contract, event: "MarketUsdDeposited", logs: logs, sub: sub}, nil
}

// WatchMarketUsdDeposited is a free log subscription operation binding the contract event 0xcbdbec7c74466ca0e764de5f6bca5bef859e31a7d6284be22ac6b4beabddf835.
//
// Solidity: event MarketUsdDeposited(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_Contracts *ContractsFilterer) WatchMarketUsdDeposited(opts *bind.WatchOpts, sink chan<- *ContractsMarketUsdDeposited, marketId []*big.Int, target []common.Address, market []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketUsdDeposited", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketUsdDeposited)
				if err := _Contracts.contract.UnpackLog(event, "MarketUsdDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseMarketUsdDeposited(log types.Log) (*ContractsMarketUsdDeposited, error) {
	event := new(ContractsMarketUsdDeposited)
	if err := _Contracts.contract.UnpackLog(event, "MarketUsdDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMarketUsdWithdrawnIterator is returned from FilterMarketUsdWithdrawn and is used to iterate over the raw logs and unpacked data for MarketUsdWithdrawn events raised by the Contracts contract.
type ContractsMarketUsdWithdrawnIterator struct {
	Event *ContractsMarketUsdWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsMarketUsdWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketUsdWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsMarketUsdWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsMarketUsdWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketUsdWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketUsdWithdrawn represents a MarketUsdWithdrawn event raised by the Contracts contract.
type ContractsMarketUsdWithdrawn struct {
	MarketId *big.Int
	Target   common.Address
	Amount   *big.Int
	Market   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketUsdWithdrawn is a free log retrieval operation binding the contract event 0x8669e4526579b1f1da479be2f96d1c7953fff439a4f98362a1c9ca07fe42ba41.
//
// Solidity: event MarketUsdWithdrawn(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_Contracts *ContractsFilterer) FilterMarketUsdWithdrawn(opts *bind.FilterOpts, marketId []*big.Int, target []common.Address, market []common.Address) (*ContractsMarketUsdWithdrawnIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketUsdWithdrawn", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketUsdWithdrawnIterator{contract: _Contracts.contract, event: "MarketUsdWithdrawn", logs: logs, sub: sub}, nil
}

// WatchMarketUsdWithdrawn is a free log subscription operation binding the contract event 0x8669e4526579b1f1da479be2f96d1c7953fff439a4f98362a1c9ca07fe42ba41.
//
// Solidity: event MarketUsdWithdrawn(uint128 indexed marketId, address indexed target, uint256 amount, address indexed market)
func (_Contracts *ContractsFilterer) WatchMarketUsdWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractsMarketUsdWithdrawn, marketId []*big.Int, target []common.Address, market []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketUsdWithdrawn", marketIdRule, targetRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketUsdWithdrawn)
				if err := _Contracts.contract.UnpackLog(event, "MarketUsdWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseMarketUsdWithdrawn(log types.Log) (*ContractsMarketUsdWithdrawn, error) {
	event := new(ContractsMarketUsdWithdrawn)
	if err := _Contracts.contract.UnpackLog(event, "MarketUsdWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMaximumMarketCollateralConfiguredIterator is returned from FilterMaximumMarketCollateralConfigured and is used to iterate over the raw logs and unpacked data for MaximumMarketCollateralConfigured events raised by the Contracts contract.
type ContractsMaximumMarketCollateralConfiguredIterator struct {
	Event *ContractsMaximumMarketCollateralConfigured // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsMaximumMarketCollateralConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMaximumMarketCollateralConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsMaximumMarketCollateralConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsMaximumMarketCollateralConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMaximumMarketCollateralConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMaximumMarketCollateralConfigured represents a MaximumMarketCollateralConfigured event raised by the Contracts contract.
type ContractsMaximumMarketCollateralConfigured struct {
	MarketId       *big.Int
	CollateralType common.Address
	SystemAmount   *big.Int
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMaximumMarketCollateralConfigured is a free log retrieval operation binding the contract event 0x499c8fcfbc4341c37dcf444c890d42ef888d46aa586f97ceb20577c2635c8075.
//
// Solidity: event MaximumMarketCollateralConfigured(uint128 indexed marketId, address indexed collateralType, uint256 systemAmount, address indexed owner)
func (_Contracts *ContractsFilterer) FilterMaximumMarketCollateralConfigured(opts *bind.FilterOpts, marketId []*big.Int, collateralType []common.Address, owner []common.Address) (*ContractsMaximumMarketCollateralConfiguredIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MaximumMarketCollateralConfigured", marketIdRule, collateralTypeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMaximumMarketCollateralConfiguredIterator{contract: _Contracts.contract, event: "MaximumMarketCollateralConfigured", logs: logs, sub: sub}, nil
}

// WatchMaximumMarketCollateralConfigured is a free log subscription operation binding the contract event 0x499c8fcfbc4341c37dcf444c890d42ef888d46aa586f97ceb20577c2635c8075.
//
// Solidity: event MaximumMarketCollateralConfigured(uint128 indexed marketId, address indexed collateralType, uint256 systemAmount, address indexed owner)
func (_Contracts *ContractsFilterer) WatchMaximumMarketCollateralConfigured(opts *bind.WatchOpts, sink chan<- *ContractsMaximumMarketCollateralConfigured, marketId []*big.Int, collateralType []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MaximumMarketCollateralConfigured", marketIdRule, collateralTypeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMaximumMarketCollateralConfigured)
				if err := _Contracts.contract.UnpackLog(event, "MaximumMarketCollateralConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseMaximumMarketCollateralConfigured(log types.Log) (*ContractsMaximumMarketCollateralConfigured, error) {
	event := new(ContractsMaximumMarketCollateralConfigured)
	if err := _Contracts.contract.UnpackLog(event, "MaximumMarketCollateralConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsNewSupportedCrossChainNetworkIterator is returned from FilterNewSupportedCrossChainNetwork and is used to iterate over the raw logs and unpacked data for NewSupportedCrossChainNetwork events raised by the Contracts contract.
type ContractsNewSupportedCrossChainNetworkIterator struct {
	Event *ContractsNewSupportedCrossChainNetwork // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsNewSupportedCrossChainNetworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsNewSupportedCrossChainNetwork)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsNewSupportedCrossChainNetwork)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsNewSupportedCrossChainNetworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsNewSupportedCrossChainNetworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsNewSupportedCrossChainNetwork represents a NewSupportedCrossChainNetwork event raised by the Contracts contract.
type ContractsNewSupportedCrossChainNetwork struct {
	NewChainId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewSupportedCrossChainNetwork is a free log retrieval operation binding the contract event 0x1874eb2a5288e478dcedf1d33291bd7293eeef5946ec516d2ef54a364b3f63d8.
//
// Solidity: event NewSupportedCrossChainNetwork(uint64 newChainId)
func (_Contracts *ContractsFilterer) FilterNewSupportedCrossChainNetwork(opts *bind.FilterOpts) (*ContractsNewSupportedCrossChainNetworkIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "NewSupportedCrossChainNetwork")
	if err != nil {
		return nil, err
	}
	return &ContractsNewSupportedCrossChainNetworkIterator{contract: _Contracts.contract, event: "NewSupportedCrossChainNetwork", logs: logs, sub: sub}, nil
}

// WatchNewSupportedCrossChainNetwork is a free log subscription operation binding the contract event 0x1874eb2a5288e478dcedf1d33291bd7293eeef5946ec516d2ef54a364b3f63d8.
//
// Solidity: event NewSupportedCrossChainNetwork(uint64 newChainId)
func (_Contracts *ContractsFilterer) WatchNewSupportedCrossChainNetwork(opts *bind.WatchOpts, sink chan<- *ContractsNewSupportedCrossChainNetwork) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "NewSupportedCrossChainNetwork")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsNewSupportedCrossChainNetwork)
				if err := _Contracts.contract.UnpackLog(event, "NewSupportedCrossChainNetwork", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseNewSupportedCrossChainNetwork(log types.Log) (*ContractsNewSupportedCrossChainNetwork, error) {
	event := new(ContractsNewSupportedCrossChainNetwork)
	if err := _Contracts.contract.UnpackLog(event, "NewSupportedCrossChainNetwork", log); err != nil {
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

// ContractsPoolApprovedAddedIterator is returned from FilterPoolApprovedAdded and is used to iterate over the raw logs and unpacked data for PoolApprovedAdded events raised by the Contracts contract.
type ContractsPoolApprovedAddedIterator struct {
	Event *ContractsPoolApprovedAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPoolApprovedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPoolApprovedAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPoolApprovedAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPoolApprovedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPoolApprovedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPoolApprovedAdded represents a PoolApprovedAdded event raised by the Contracts contract.
type ContractsPoolApprovedAdded struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolApprovedAdded is a free log retrieval operation binding the contract event 0x7d5bdf4e8c44e0b5a8249bf03c2a1febd848cc7f580efd7b1703301c5b1a9e4e.
//
// Solidity: event PoolApprovedAdded(uint256 poolId)
func (_Contracts *ContractsFilterer) FilterPoolApprovedAdded(opts *bind.FilterOpts) (*ContractsPoolApprovedAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PoolApprovedAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsPoolApprovedAddedIterator{contract: _Contracts.contract, event: "PoolApprovedAdded", logs: logs, sub: sub}, nil
}

// WatchPoolApprovedAdded is a free log subscription operation binding the contract event 0x7d5bdf4e8c44e0b5a8249bf03c2a1febd848cc7f580efd7b1703301c5b1a9e4e.
//
// Solidity: event PoolApprovedAdded(uint256 poolId)
func (_Contracts *ContractsFilterer) WatchPoolApprovedAdded(opts *bind.WatchOpts, sink chan<- *ContractsPoolApprovedAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PoolApprovedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPoolApprovedAdded)
				if err := _Contracts.contract.UnpackLog(event, "PoolApprovedAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParsePoolApprovedAdded(log types.Log) (*ContractsPoolApprovedAdded, error) {
	event := new(ContractsPoolApprovedAdded)
	if err := _Contracts.contract.UnpackLog(event, "PoolApprovedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPoolApprovedRemovedIterator is returned from FilterPoolApprovedRemoved and is used to iterate over the raw logs and unpacked data for PoolApprovedRemoved events raised by the Contracts contract.
type ContractsPoolApprovedRemovedIterator struct {
	Event *ContractsPoolApprovedRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPoolApprovedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPoolApprovedRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPoolApprovedRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPoolApprovedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPoolApprovedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPoolApprovedRemoved represents a PoolApprovedRemoved event raised by the Contracts contract.
type ContractsPoolApprovedRemoved struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolApprovedRemoved is a free log retrieval operation binding the contract event 0xc1567ee9983f306f073ea7d59a7fb5680ce07985f8b49cc50d00a3a9f748d3c2.
//
// Solidity: event PoolApprovedRemoved(uint256 poolId)
func (_Contracts *ContractsFilterer) FilterPoolApprovedRemoved(opts *bind.FilterOpts) (*ContractsPoolApprovedRemovedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PoolApprovedRemoved")
	if err != nil {
		return nil, err
	}
	return &ContractsPoolApprovedRemovedIterator{contract: _Contracts.contract, event: "PoolApprovedRemoved", logs: logs, sub: sub}, nil
}

// WatchPoolApprovedRemoved is a free log subscription operation binding the contract event 0xc1567ee9983f306f073ea7d59a7fb5680ce07985f8b49cc50d00a3a9f748d3c2.
//
// Solidity: event PoolApprovedRemoved(uint256 poolId)
func (_Contracts *ContractsFilterer) WatchPoolApprovedRemoved(opts *bind.WatchOpts, sink chan<- *ContractsPoolApprovedRemoved) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PoolApprovedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPoolApprovedRemoved)
				if err := _Contracts.contract.UnpackLog(event, "PoolApprovedRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParsePoolApprovedRemoved(log types.Log) (*ContractsPoolApprovedRemoved, error) {
	event := new(ContractsPoolApprovedRemoved)
	if err := _Contracts.contract.UnpackLog(event, "PoolApprovedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPoolConfigurationSetIterator is returned from FilterPoolConfigurationSet and is used to iterate over the raw logs and unpacked data for PoolConfigurationSet events raised by the Contracts contract.
type ContractsPoolConfigurationSetIterator struct {
	Event *ContractsPoolConfigurationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPoolConfigurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPoolConfigurationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPoolConfigurationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPoolConfigurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPoolConfigurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPoolConfigurationSet represents a PoolConfigurationSet event raised by the Contracts contract.
type ContractsPoolConfigurationSet struct {
	PoolId  *big.Int
	Markets []MarketConfigurationData
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolConfigurationSet is a free log retrieval operation binding the contract event 0xdd812c2e47943d98e6c66b2b9872d1f9270b8523c82eb60ad5c8d580a614081c.
//
// Solidity: event PoolConfigurationSet(uint128 indexed poolId, (uint128,uint128,int128)[] markets, address indexed sender)
func (_Contracts *ContractsFilterer) FilterPoolConfigurationSet(opts *bind.FilterOpts, poolId []*big.Int, sender []common.Address) (*ContractsPoolConfigurationSetIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PoolConfigurationSet", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPoolConfigurationSetIterator{contract: _Contracts.contract, event: "PoolConfigurationSet", logs: logs, sub: sub}, nil
}

// WatchPoolConfigurationSet is a free log subscription operation binding the contract event 0xdd812c2e47943d98e6c66b2b9872d1f9270b8523c82eb60ad5c8d580a614081c.
//
// Solidity: event PoolConfigurationSet(uint128 indexed poolId, (uint128,uint128,int128)[] markets, address indexed sender)
func (_Contracts *ContractsFilterer) WatchPoolConfigurationSet(opts *bind.WatchOpts, sink chan<- *ContractsPoolConfigurationSet, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PoolConfigurationSet", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPoolConfigurationSet)
				if err := _Contracts.contract.UnpackLog(event, "PoolConfigurationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParsePoolConfigurationSet(log types.Log) (*ContractsPoolConfigurationSet, error) {
	event := new(ContractsPoolConfigurationSet)
	if err := _Contracts.contract.UnpackLog(event, "PoolConfigurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the Contracts contract.
type ContractsPoolCreatedIterator struct {
	Event *ContractsPoolCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPoolCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPoolCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPoolCreated represents a PoolCreated event raised by the Contracts contract.
type ContractsPoolCreated struct {
	PoolId *big.Int
	Owner  common.Address
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xb1517ad708e5f9a104c30d3f1ff749d55833b1d03bf472013c29888e741cf340.
//
// Solidity: event PoolCreated(uint128 indexed poolId, address indexed owner, address indexed sender)
func (_Contracts *ContractsFilterer) FilterPoolCreated(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address, sender []common.Address) (*ContractsPoolCreatedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PoolCreated", poolIdRule, ownerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPoolCreatedIterator{contract: _Contracts.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xb1517ad708e5f9a104c30d3f1ff749d55833b1d03bf472013c29888e741cf340.
//
// Solidity: event PoolCreated(uint128 indexed poolId, address indexed owner, address indexed sender)
func (_Contracts *ContractsFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *ContractsPoolCreated, poolId []*big.Int, owner []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PoolCreated", poolIdRule, ownerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPoolCreated)
				if err := _Contracts.contract.UnpackLog(event, "PoolCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParsePoolCreated(log types.Log) (*ContractsPoolCreated, error) {
	event := new(ContractsPoolCreated)
	if err := _Contracts.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPoolNameUpdatedIterator is returned from FilterPoolNameUpdated and is used to iterate over the raw logs and unpacked data for PoolNameUpdated events raised by the Contracts contract.
type ContractsPoolNameUpdatedIterator struct {
	Event *ContractsPoolNameUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPoolNameUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPoolNameUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPoolNameUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPoolNameUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPoolNameUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPoolNameUpdated represents a PoolNameUpdated event raised by the Contracts contract.
type ContractsPoolNameUpdated struct {
	PoolId *big.Int
	Name   string
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNameUpdated is a free log retrieval operation binding the contract event 0x63b42abaf7e145a993f20bc64259f45d09c43d18838ab0bca078b15093ac55f4.
//
// Solidity: event PoolNameUpdated(uint128 indexed poolId, string name, address indexed sender)
func (_Contracts *ContractsFilterer) FilterPoolNameUpdated(opts *bind.FilterOpts, poolId []*big.Int, sender []common.Address) (*ContractsPoolNameUpdatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PoolNameUpdated", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPoolNameUpdatedIterator{contract: _Contracts.contract, event: "PoolNameUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolNameUpdated is a free log subscription operation binding the contract event 0x63b42abaf7e145a993f20bc64259f45d09c43d18838ab0bca078b15093ac55f4.
//
// Solidity: event PoolNameUpdated(uint128 indexed poolId, string name, address indexed sender)
func (_Contracts *ContractsFilterer) WatchPoolNameUpdated(opts *bind.WatchOpts, sink chan<- *ContractsPoolNameUpdated, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PoolNameUpdated", poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPoolNameUpdated)
				if err := _Contracts.contract.UnpackLog(event, "PoolNameUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParsePoolNameUpdated(log types.Log) (*ContractsPoolNameUpdated, error) {
	event := new(ContractsPoolNameUpdated)
	if err := _Contracts.contract.UnpackLog(event, "PoolNameUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPoolNominationRenouncedIterator is returned from FilterPoolNominationRenounced and is used to iterate over the raw logs and unpacked data for PoolNominationRenounced events raised by the Contracts contract.
type ContractsPoolNominationRenouncedIterator struct {
	Event *ContractsPoolNominationRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPoolNominationRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPoolNominationRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPoolNominationRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPoolNominationRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPoolNominationRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPoolNominationRenounced represents a PoolNominationRenounced event raised by the Contracts contract.
type ContractsPoolNominationRenounced struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNominationRenounced is a free log retrieval operation binding the contract event 0x28301da5fb0feefb138efa6310af4547a74f415d62616f90519436dc169c3ae0.
//
// Solidity: event PoolNominationRenounced(uint128 indexed poolId, address indexed owner)
func (_Contracts *ContractsFilterer) FilterPoolNominationRenounced(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*ContractsPoolNominationRenouncedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PoolNominationRenounced", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPoolNominationRenouncedIterator{contract: _Contracts.contract, event: "PoolNominationRenounced", logs: logs, sub: sub}, nil
}

// WatchPoolNominationRenounced is a free log subscription operation binding the contract event 0x28301da5fb0feefb138efa6310af4547a74f415d62616f90519436dc169c3ae0.
//
// Solidity: event PoolNominationRenounced(uint128 indexed poolId, address indexed owner)
func (_Contracts *ContractsFilterer) WatchPoolNominationRenounced(opts *bind.WatchOpts, sink chan<- *ContractsPoolNominationRenounced, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PoolNominationRenounced", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPoolNominationRenounced)
				if err := _Contracts.contract.UnpackLog(event, "PoolNominationRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParsePoolNominationRenounced(log types.Log) (*ContractsPoolNominationRenounced, error) {
	event := new(ContractsPoolNominationRenounced)
	if err := _Contracts.contract.UnpackLog(event, "PoolNominationRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPoolNominationRevokedIterator is returned from FilterPoolNominationRevoked and is used to iterate over the raw logs and unpacked data for PoolNominationRevoked events raised by the Contracts contract.
type ContractsPoolNominationRevokedIterator struct {
	Event *ContractsPoolNominationRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPoolNominationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPoolNominationRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPoolNominationRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPoolNominationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPoolNominationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPoolNominationRevoked represents a PoolNominationRevoked event raised by the Contracts contract.
type ContractsPoolNominationRevoked struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolNominationRevoked is a free log retrieval operation binding the contract event 0xa20a605599b6da4a06e0662f1284c442a576bc452b77a38c8c55805cb82a1865.
//
// Solidity: event PoolNominationRevoked(uint128 indexed poolId, address indexed owner)
func (_Contracts *ContractsFilterer) FilterPoolNominationRevoked(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*ContractsPoolNominationRevokedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PoolNominationRevoked", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPoolNominationRevokedIterator{contract: _Contracts.contract, event: "PoolNominationRevoked", logs: logs, sub: sub}, nil
}

// WatchPoolNominationRevoked is a free log subscription operation binding the contract event 0xa20a605599b6da4a06e0662f1284c442a576bc452b77a38c8c55805cb82a1865.
//
// Solidity: event PoolNominationRevoked(uint128 indexed poolId, address indexed owner)
func (_Contracts *ContractsFilterer) WatchPoolNominationRevoked(opts *bind.WatchOpts, sink chan<- *ContractsPoolNominationRevoked, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PoolNominationRevoked", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPoolNominationRevoked)
				if err := _Contracts.contract.UnpackLog(event, "PoolNominationRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParsePoolNominationRevoked(log types.Log) (*ContractsPoolNominationRevoked, error) {
	event := new(ContractsPoolNominationRevoked)
	if err := _Contracts.contract.UnpackLog(event, "PoolNominationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPoolOwnerNominatedIterator is returned from FilterPoolOwnerNominated and is used to iterate over the raw logs and unpacked data for PoolOwnerNominated events raised by the Contracts contract.
type ContractsPoolOwnerNominatedIterator struct {
	Event *ContractsPoolOwnerNominated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPoolOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPoolOwnerNominated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPoolOwnerNominated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPoolOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPoolOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPoolOwnerNominated represents a PoolOwnerNominated event raised by the Contracts contract.
type ContractsPoolOwnerNominated struct {
	PoolId         *big.Int
	NominatedOwner common.Address
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolOwnerNominated is a free log retrieval operation binding the contract event 0x55d98f82a53fb5776e9ea48d624ab9cb015b51a45249b1ed8425fc857c82f4f8.
//
// Solidity: event PoolOwnerNominated(uint128 indexed poolId, address indexed nominatedOwner, address indexed owner)
func (_Contracts *ContractsFilterer) FilterPoolOwnerNominated(opts *bind.FilterOpts, poolId []*big.Int, nominatedOwner []common.Address, owner []common.Address) (*ContractsPoolOwnerNominatedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PoolOwnerNominated", poolIdRule, nominatedOwnerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPoolOwnerNominatedIterator{contract: _Contracts.contract, event: "PoolOwnerNominated", logs: logs, sub: sub}, nil
}

// WatchPoolOwnerNominated is a free log subscription operation binding the contract event 0x55d98f82a53fb5776e9ea48d624ab9cb015b51a45249b1ed8425fc857c82f4f8.
//
// Solidity: event PoolOwnerNominated(uint128 indexed poolId, address indexed nominatedOwner, address indexed owner)
func (_Contracts *ContractsFilterer) WatchPoolOwnerNominated(opts *bind.WatchOpts, sink chan<- *ContractsPoolOwnerNominated, poolId []*big.Int, nominatedOwner []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PoolOwnerNominated", poolIdRule, nominatedOwnerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPoolOwnerNominated)
				if err := _Contracts.contract.UnpackLog(event, "PoolOwnerNominated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParsePoolOwnerNominated(log types.Log) (*ContractsPoolOwnerNominated, error) {
	event := new(ContractsPoolOwnerNominated)
	if err := _Contracts.contract.UnpackLog(event, "PoolOwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPoolOwnershipAcceptedIterator is returned from FilterPoolOwnershipAccepted and is used to iterate over the raw logs and unpacked data for PoolOwnershipAccepted events raised by the Contracts contract.
type ContractsPoolOwnershipAcceptedIterator struct {
	Event *ContractsPoolOwnershipAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPoolOwnershipAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPoolOwnershipAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPoolOwnershipAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPoolOwnershipAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPoolOwnershipAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPoolOwnershipAccepted represents a PoolOwnershipAccepted event raised by the Contracts contract.
type ContractsPoolOwnershipAccepted struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolOwnershipAccepted is a free log retrieval operation binding the contract event 0x4f86f2ce8b08e27d0e470f4269b71c3bbc68407d51a2e692f6573236074ebc5a.
//
// Solidity: event PoolOwnershipAccepted(uint128 indexed poolId, address indexed owner)
func (_Contracts *ContractsFilterer) FilterPoolOwnershipAccepted(opts *bind.FilterOpts, poolId []*big.Int, owner []common.Address) (*ContractsPoolOwnershipAcceptedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PoolOwnershipAccepted", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPoolOwnershipAcceptedIterator{contract: _Contracts.contract, event: "PoolOwnershipAccepted", logs: logs, sub: sub}, nil
}

// WatchPoolOwnershipAccepted is a free log subscription operation binding the contract event 0x4f86f2ce8b08e27d0e470f4269b71c3bbc68407d51a2e692f6573236074ebc5a.
//
// Solidity: event PoolOwnershipAccepted(uint128 indexed poolId, address indexed owner)
func (_Contracts *ContractsFilterer) WatchPoolOwnershipAccepted(opts *bind.WatchOpts, sink chan<- *ContractsPoolOwnershipAccepted, poolId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PoolOwnershipAccepted", poolIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPoolOwnershipAccepted)
				if err := _Contracts.contract.UnpackLog(event, "PoolOwnershipAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParsePoolOwnershipAccepted(log types.Log) (*ContractsPoolOwnershipAccepted, error) {
	event := new(ContractsPoolOwnershipAccepted)
	if err := _Contracts.contract.UnpackLog(event, "PoolOwnershipAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPreferredPoolSetIterator is returned from FilterPreferredPoolSet and is used to iterate over the raw logs and unpacked data for PreferredPoolSet events raised by the Contracts contract.
type ContractsPreferredPoolSetIterator struct {
	Event *ContractsPreferredPoolSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPreferredPoolSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPreferredPoolSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPreferredPoolSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPreferredPoolSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPreferredPoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPreferredPoolSet represents a PreferredPoolSet event raised by the Contracts contract.
type ContractsPreferredPoolSet struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPreferredPoolSet is a free log retrieval operation binding the contract event 0x7e7cb4726e710dc12fad41f158c37a4a71af3a6f053b8b13670d35c710139a56.
//
// Solidity: event PreferredPoolSet(uint256 poolId)
func (_Contracts *ContractsFilterer) FilterPreferredPoolSet(opts *bind.FilterOpts) (*ContractsPreferredPoolSetIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PreferredPoolSet")
	if err != nil {
		return nil, err
	}
	return &ContractsPreferredPoolSetIterator{contract: _Contracts.contract, event: "PreferredPoolSet", logs: logs, sub: sub}, nil
}

// WatchPreferredPoolSet is a free log subscription operation binding the contract event 0x7e7cb4726e710dc12fad41f158c37a4a71af3a6f053b8b13670d35c710139a56.
//
// Solidity: event PreferredPoolSet(uint256 poolId)
func (_Contracts *ContractsFilterer) WatchPreferredPoolSet(opts *bind.WatchOpts, sink chan<- *ContractsPreferredPoolSet) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PreferredPoolSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPreferredPoolSet)
				if err := _Contracts.contract.UnpackLog(event, "PreferredPoolSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParsePreferredPoolSet(log types.Log) (*ContractsPreferredPoolSet, error) {
	event := new(ContractsPreferredPoolSet)
	if err := _Contracts.contract.UnpackLog(event, "PreferredPoolSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the Contracts contract.
type ContractsRewardsClaimedIterator struct {
	Event *ContractsRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRewardsClaimed represents a RewardsClaimed event raised by the Contracts contract.
type ContractsRewardsClaimed struct {
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
func (_Contracts *ContractsFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (*ContractsRewardsClaimedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RewardsClaimed", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRewardsClaimedIterator{contract: _Contracts.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xa4a60be4203e7975e54ab5314c7e9e18aba9ad71e8da714d8de987f4f05410f2.
//
// Solidity: event RewardsClaimed(uint128 indexed accountId, uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount)
func (_Contracts *ContractsFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *ContractsRewardsClaimed, accountId []*big.Int, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RewardsClaimed", accountIdRule, poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRewardsClaimed)
				if err := _Contracts.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseRewardsClaimed(log types.Log) (*ContractsRewardsClaimed, error) {
	event := new(ContractsRewardsClaimed)
	if err := _Contracts.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRewardsDistributedIterator is returned from FilterRewardsDistributed and is used to iterate over the raw logs and unpacked data for RewardsDistributed events raised by the Contracts contract.
type ContractsRewardsDistributedIterator struct {
	Event *ContractsRewardsDistributed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsRewardsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRewardsDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsRewardsDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsRewardsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRewardsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRewardsDistributed represents a RewardsDistributed event raised by the Contracts contract.
type ContractsRewardsDistributed struct {
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
func (_Contracts *ContractsFilterer) FilterRewardsDistributed(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address) (*ContractsRewardsDistributedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RewardsDistributed", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRewardsDistributedIterator{contract: _Contracts.contract, event: "RewardsDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributed is a free log subscription operation binding the contract event 0x19ced31d71d1db45f99d5a8d3a7616fe9d78828df58f2a28feb68c9f9ab876ca.
//
// Solidity: event RewardsDistributed(uint128 indexed poolId, address indexed collateralType, address distributor, uint256 amount, uint256 start, uint256 duration)
func (_Contracts *ContractsFilterer) WatchRewardsDistributed(opts *bind.WatchOpts, sink chan<- *ContractsRewardsDistributed, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RewardsDistributed", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRewardsDistributed)
				if err := _Contracts.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseRewardsDistributed(log types.Log) (*ContractsRewardsDistributed, error) {
	event := new(ContractsRewardsDistributed)
	if err := _Contracts.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRewardsDistributorRegisteredIterator is returned from FilterRewardsDistributorRegistered and is used to iterate over the raw logs and unpacked data for RewardsDistributorRegistered events raised by the Contracts contract.
type ContractsRewardsDistributorRegisteredIterator struct {
	Event *ContractsRewardsDistributorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsRewardsDistributorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRewardsDistributorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsRewardsDistributorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsRewardsDistributorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRewardsDistributorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRewardsDistributorRegistered represents a RewardsDistributorRegistered event raised by the Contracts contract.
type ContractsRewardsDistributorRegistered struct {
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributorRegistered is a free log retrieval operation binding the contract event 0x9d3609c05a83dc93a5b355d62c2b37dfde8f0833b1184d4d05c6f51cd46b6e5b.
//
// Solidity: event RewardsDistributorRegistered(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_Contracts *ContractsFilterer) FilterRewardsDistributorRegistered(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (*ContractsRewardsDistributorRegisteredIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RewardsDistributorRegistered", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRewardsDistributorRegisteredIterator{contract: _Contracts.contract, event: "RewardsDistributorRegistered", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributorRegistered is a free log subscription operation binding the contract event 0x9d3609c05a83dc93a5b355d62c2b37dfde8f0833b1184d4d05c6f51cd46b6e5b.
//
// Solidity: event RewardsDistributorRegistered(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_Contracts *ContractsFilterer) WatchRewardsDistributorRegistered(opts *bind.WatchOpts, sink chan<- *ContractsRewardsDistributorRegistered, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RewardsDistributorRegistered", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRewardsDistributorRegistered)
				if err := _Contracts.contract.UnpackLog(event, "RewardsDistributorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseRewardsDistributorRegistered(log types.Log) (*ContractsRewardsDistributorRegistered, error) {
	event := new(ContractsRewardsDistributorRegistered)
	if err := _Contracts.contract.UnpackLog(event, "RewardsDistributorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRewardsDistributorRemovedIterator is returned from FilterRewardsDistributorRemoved and is used to iterate over the raw logs and unpacked data for RewardsDistributorRemoved events raised by the Contracts contract.
type ContractsRewardsDistributorRemovedIterator struct {
	Event *ContractsRewardsDistributorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsRewardsDistributorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRewardsDistributorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsRewardsDistributorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsRewardsDistributorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRewardsDistributorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRewardsDistributorRemoved represents a RewardsDistributorRemoved event raised by the Contracts contract.
type ContractsRewardsDistributorRemoved struct {
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributorRemoved is a free log retrieval operation binding the contract event 0x375c4507f463c55a506be95e2cfd3cfdc0610be055087eac6049588a1bcfacba.
//
// Solidity: event RewardsDistributorRemoved(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_Contracts *ContractsFilterer) FilterRewardsDistributorRemoved(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (*ContractsRewardsDistributorRemovedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RewardsDistributorRemoved", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRewardsDistributorRemovedIterator{contract: _Contracts.contract, event: "RewardsDistributorRemoved", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributorRemoved is a free log subscription operation binding the contract event 0x375c4507f463c55a506be95e2cfd3cfdc0610be055087eac6049588a1bcfacba.
//
// Solidity: event RewardsDistributorRemoved(uint128 indexed poolId, address indexed collateralType, address indexed distributor)
func (_Contracts *ContractsFilterer) WatchRewardsDistributorRemoved(opts *bind.WatchOpts, sink chan<- *ContractsRewardsDistributorRemoved, poolId []*big.Int, collateralType []common.Address, distributor []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RewardsDistributorRemoved", poolIdRule, collateralTypeRule, distributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRewardsDistributorRemoved)
				if err := _Contracts.contract.UnpackLog(event, "RewardsDistributorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseRewardsDistributorRemoved(log types.Log) (*ContractsRewardsDistributorRemoved, error) {
	event := new(ContractsRewardsDistributorRemoved)
	if err := _Contracts.contract.UnpackLog(event, "RewardsDistributorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSetMarketMinLiquidityRatioIterator is returned from FilterSetMarketMinLiquidityRatio and is used to iterate over the raw logs and unpacked data for SetMarketMinLiquidityRatio events raised by the Contracts contract.
type ContractsSetMarketMinLiquidityRatioIterator struct {
	Event *ContractsSetMarketMinLiquidityRatio // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsSetMarketMinLiquidityRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSetMarketMinLiquidityRatio)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsSetMarketMinLiquidityRatio)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsSetMarketMinLiquidityRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSetMarketMinLiquidityRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSetMarketMinLiquidityRatio represents a SetMarketMinLiquidityRatio event raised by the Contracts contract.
type ContractsSetMarketMinLiquidityRatio struct {
	MarketId          *big.Int
	MinLiquidityRatio *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMarketMinLiquidityRatio is a free log retrieval operation binding the contract event 0x563eb723f21b3e87ec8932cfb4ffa64d1b68c42053c28d6b4db019a40f6daf47.
//
// Solidity: event SetMarketMinLiquidityRatio(uint128 indexed marketId, uint256 minLiquidityRatio)
func (_Contracts *ContractsFilterer) FilterSetMarketMinLiquidityRatio(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsSetMarketMinLiquidityRatioIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SetMarketMinLiquidityRatio", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSetMarketMinLiquidityRatioIterator{contract: _Contracts.contract, event: "SetMarketMinLiquidityRatio", logs: logs, sub: sub}, nil
}

// WatchSetMarketMinLiquidityRatio is a free log subscription operation binding the contract event 0x563eb723f21b3e87ec8932cfb4ffa64d1b68c42053c28d6b4db019a40f6daf47.
//
// Solidity: event SetMarketMinLiquidityRatio(uint128 indexed marketId, uint256 minLiquidityRatio)
func (_Contracts *ContractsFilterer) WatchSetMarketMinLiquidityRatio(opts *bind.WatchOpts, sink chan<- *ContractsSetMarketMinLiquidityRatio, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SetMarketMinLiquidityRatio", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSetMarketMinLiquidityRatio)
				if err := _Contracts.contract.UnpackLog(event, "SetMarketMinLiquidityRatio", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseSetMarketMinLiquidityRatio(log types.Log) (*ContractsSetMarketMinLiquidityRatio, error) {
	event := new(ContractsSetMarketMinLiquidityRatio)
	if err := _Contracts.contract.UnpackLog(event, "SetMarketMinLiquidityRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSetMinDelegateTimeIterator is returned from FilterSetMinDelegateTime and is used to iterate over the raw logs and unpacked data for SetMinDelegateTime events raised by the Contracts contract.
type ContractsSetMinDelegateTimeIterator struct {
	Event *ContractsSetMinDelegateTime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsSetMinDelegateTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSetMinDelegateTime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsSetMinDelegateTime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsSetMinDelegateTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSetMinDelegateTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSetMinDelegateTime represents a SetMinDelegateTime event raised by the Contracts contract.
type ContractsSetMinDelegateTime struct {
	MarketId        *big.Int
	MinDelegateTime uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetMinDelegateTime is a free log retrieval operation binding the contract event 0x6942a68d151863c1fed3c0c4c5f3258af738218527147ac69290ab23ca7d26c6.
//
// Solidity: event SetMinDelegateTime(uint128 indexed marketId, uint32 minDelegateTime)
func (_Contracts *ContractsFilterer) FilterSetMinDelegateTime(opts *bind.FilterOpts, marketId []*big.Int) (*ContractsSetMinDelegateTimeIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SetMinDelegateTime", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSetMinDelegateTimeIterator{contract: _Contracts.contract, event: "SetMinDelegateTime", logs: logs, sub: sub}, nil
}

// WatchSetMinDelegateTime is a free log subscription operation binding the contract event 0x6942a68d151863c1fed3c0c4c5f3258af738218527147ac69290ab23ca7d26c6.
//
// Solidity: event SetMinDelegateTime(uint128 indexed marketId, uint32 minDelegateTime)
func (_Contracts *ContractsFilterer) WatchSetMinDelegateTime(opts *bind.WatchOpts, sink chan<- *ContractsSetMinDelegateTime, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SetMinDelegateTime", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSetMinDelegateTime)
				if err := _Contracts.contract.UnpackLog(event, "SetMinDelegateTime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseSetMinDelegateTime(log types.Log) (*ContractsSetMinDelegateTime, error) {
	event := new(ContractsSetMinDelegateTime)
	if err := _Contracts.contract.UnpackLog(event, "SetMinDelegateTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSetMinLiquidityRatioIterator is returned from FilterSetMinLiquidityRatio and is used to iterate over the raw logs and unpacked data for SetMinLiquidityRatio events raised by the Contracts contract.
type ContractsSetMinLiquidityRatioIterator struct {
	Event *ContractsSetMinLiquidityRatio // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsSetMinLiquidityRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSetMinLiquidityRatio)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsSetMinLiquidityRatio)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsSetMinLiquidityRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSetMinLiquidityRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSetMinLiquidityRatio represents a SetMinLiquidityRatio event raised by the Contracts contract.
type ContractsSetMinLiquidityRatio struct {
	MinLiquidityRatio *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMinLiquidityRatio is a free log retrieval operation binding the contract event 0x66fd484d9868d1faddc8fef1f3faed0ed25eb4e6acde49dd1f2cbf0fba903635.
//
// Solidity: event SetMinLiquidityRatio(uint256 minLiquidityRatio)
func (_Contracts *ContractsFilterer) FilterSetMinLiquidityRatio(opts *bind.FilterOpts) (*ContractsSetMinLiquidityRatioIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SetMinLiquidityRatio")
	if err != nil {
		return nil, err
	}
	return &ContractsSetMinLiquidityRatioIterator{contract: _Contracts.contract, event: "SetMinLiquidityRatio", logs: logs, sub: sub}, nil
}

// WatchSetMinLiquidityRatio is a free log subscription operation binding the contract event 0x66fd484d9868d1faddc8fef1f3faed0ed25eb4e6acde49dd1f2cbf0fba903635.
//
// Solidity: event SetMinLiquidityRatio(uint256 minLiquidityRatio)
func (_Contracts *ContractsFilterer) WatchSetMinLiquidityRatio(opts *bind.WatchOpts, sink chan<- *ContractsSetMinLiquidityRatio) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SetMinLiquidityRatio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSetMinLiquidityRatio)
				if err := _Contracts.contract.UnpackLog(event, "SetMinLiquidityRatio", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseSetMinLiquidityRatio(log types.Log) (*ContractsSetMinLiquidityRatio, error) {
	event := new(ContractsSetMinLiquidityRatio)
	if err := _Contracts.contract.UnpackLog(event, "SetMinLiquidityRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTransferCrossChainInitiatedIterator is returned from FilterTransferCrossChainInitiated and is used to iterate over the raw logs and unpacked data for TransferCrossChainInitiated events raised by the Contracts contract.
type ContractsTransferCrossChainInitiatedIterator struct {
	Event *ContractsTransferCrossChainInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsTransferCrossChainInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransferCrossChainInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsTransferCrossChainInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsTransferCrossChainInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransferCrossChainInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransferCrossChainInitiated represents a TransferCrossChainInitiated event raised by the Contracts contract.
type ContractsTransferCrossChainInitiated struct {
	DestChainId uint64
	Amount      *big.Int
	Sender      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferCrossChainInitiated is a free log retrieval operation binding the contract event 0xb87c3097d7f9145a4915e8e434f04a1b7b91646d8a6e66a5cdab25caccb644c4.
//
// Solidity: event TransferCrossChainInitiated(uint64 indexed destChainId, uint256 indexed amount, address sender)
func (_Contracts *ContractsFilterer) FilterTransferCrossChainInitiated(opts *bind.FilterOpts, destChainId []uint64, amount []*big.Int) (*ContractsTransferCrossChainInitiatedIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "TransferCrossChainInitiated", destChainIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTransferCrossChainInitiatedIterator{contract: _Contracts.contract, event: "TransferCrossChainInitiated", logs: logs, sub: sub}, nil
}

// WatchTransferCrossChainInitiated is a free log subscription operation binding the contract event 0xb87c3097d7f9145a4915e8e434f04a1b7b91646d8a6e66a5cdab25caccb644c4.
//
// Solidity: event TransferCrossChainInitiated(uint64 indexed destChainId, uint256 indexed amount, address sender)
func (_Contracts *ContractsFilterer) WatchTransferCrossChainInitiated(opts *bind.WatchOpts, sink chan<- *ContractsTransferCrossChainInitiated, destChainId []uint64, amount []*big.Int) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "TransferCrossChainInitiated", destChainIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransferCrossChainInitiated)
				if err := _Contracts.contract.UnpackLog(event, "TransferCrossChainInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseTransferCrossChainInitiated(log types.Log) (*ContractsTransferCrossChainInitiated, error) {
	event := new(ContractsTransferCrossChainInitiated)
	if err := _Contracts.contract.UnpackLog(event, "TransferCrossChainInitiated", log); err != nil {
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

// ContractsUsdBurnedIterator is returned from FilterUsdBurned and is used to iterate over the raw logs and unpacked data for UsdBurned events raised by the Contracts contract.
type ContractsUsdBurnedIterator struct {
	Event *ContractsUsdBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsUsdBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUsdBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsUsdBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsUsdBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUsdBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUsdBurned represents a UsdBurned event raised by the Contracts contract.
type ContractsUsdBurned struct {
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
func (_Contracts *ContractsFilterer) FilterUsdBurned(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*ContractsUsdBurnedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "UsdBurned", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsUsdBurnedIterator{contract: _Contracts.contract, event: "UsdBurned", logs: logs, sub: sub}, nil
}

// WatchUsdBurned is a free log subscription operation binding the contract event 0x6b0230f0abe9188cbdbb1c816a4c5e213758b5b743d8a94af056280cc1e7aeb1.
//
// Solidity: event UsdBurned(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_Contracts *ContractsFilterer) WatchUsdBurned(opts *bind.WatchOpts, sink chan<- *ContractsUsdBurned, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "UsdBurned", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUsdBurned)
				if err := _Contracts.contract.UnpackLog(event, "UsdBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseUsdBurned(log types.Log) (*ContractsUsdBurned, error) {
	event := new(ContractsUsdBurned)
	if err := _Contracts.contract.UnpackLog(event, "UsdBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsUsdMintedIterator is returned from FilterUsdMinted and is used to iterate over the raw logs and unpacked data for UsdMinted events raised by the Contracts contract.
type ContractsUsdMintedIterator struct {
	Event *ContractsUsdMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsUsdMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUsdMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsUsdMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsUsdMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUsdMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUsdMinted represents a UsdMinted event raised by the Contracts contract.
type ContractsUsdMinted struct {
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
func (_Contracts *ContractsFilterer) FilterUsdMinted(opts *bind.FilterOpts, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (*ContractsUsdMintedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "UsdMinted", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsUsdMintedIterator{contract: _Contracts.contract, event: "UsdMinted", logs: logs, sub: sub}, nil
}

// WatchUsdMinted is a free log subscription operation binding the contract event 0x2100f67dc9a5917400f799bb13194553e3f74c19a207c56350d2c223ac9c36c9.
//
// Solidity: event UsdMinted(uint128 indexed accountId, uint128 indexed poolId, address collateralType, uint256 amount, address indexed sender)
func (_Contracts *ContractsFilterer) WatchUsdMinted(opts *bind.WatchOpts, sink chan<- *ContractsUsdMinted, accountId []*big.Int, poolId []*big.Int, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "UsdMinted", accountIdRule, poolIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUsdMinted)
				if err := _Contracts.contract.UnpackLog(event, "UsdMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseUsdMinted(log types.Log) (*ContractsUsdMinted, error) {
	event := new(ContractsUsdMinted)
	if err := _Contracts.contract.UnpackLog(event, "UsdMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVaultLiquidationIterator is returned from FilterVaultLiquidation and is used to iterate over the raw logs and unpacked data for VaultLiquidation events raised by the Contracts contract.
type ContractsVaultLiquidationIterator struct {
	Event *ContractsVaultLiquidation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsVaultLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVaultLiquidation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsVaultLiquidation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsVaultLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVaultLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVaultLiquidation represents a VaultLiquidation event raised by the Contracts contract.
type ContractsVaultLiquidation struct {
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
func (_Contracts *ContractsFilterer) FilterVaultLiquidation(opts *bind.FilterOpts, poolId []*big.Int, collateralType []common.Address) (*ContractsVaultLiquidationIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VaultLiquidation", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsVaultLiquidationIterator{contract: _Contracts.contract, event: "VaultLiquidation", logs: logs, sub: sub}, nil
}

// WatchVaultLiquidation is a free log subscription operation binding the contract event 0x1834a7cc9d14f9bfa482df5c0404dadd1b8ec123b41f082e76ae28a3b2ea68d5.
//
// Solidity: event VaultLiquidation(uint128 indexed poolId, address indexed collateralType, (uint256,uint256,uint256) liquidationData, uint128 liquidateAsAccountId, address sender)
func (_Contracts *ContractsFilterer) WatchVaultLiquidation(opts *bind.WatchOpts, sink chan<- *ContractsVaultLiquidation, poolId []*big.Int, collateralType []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var collateralTypeRule []interface{}
	for _, collateralTypeItem := range collateralType {
		collateralTypeRule = append(collateralTypeRule, collateralTypeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VaultLiquidation", poolIdRule, collateralTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVaultLiquidation)
				if err := _Contracts.contract.UnpackLog(event, "VaultLiquidation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseVaultLiquidation(log types.Log) (*ContractsVaultLiquidation, error) {
	event := new(ContractsVaultLiquidation)
	if err := _Contracts.contract.UnpackLog(event, "VaultLiquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Contracts contract.
type ContractsWithdrawnIterator struct {
	Event *ContractsWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsWithdrawn represents a Withdrawn event raised by the Contracts contract.
type ContractsWithdrawn struct {
	AccountId      *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x8b5f9d7ce522936589c630db08c0fa2405b21c4a5ff8ef19899900172736ba38.
//
// Solidity: event Withdrawn(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Contracts *ContractsFilterer) FilterWithdrawn(opts *bind.FilterOpts, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (*ContractsWithdrawnIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Withdrawn", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsWithdrawnIterator{contract: _Contracts.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x8b5f9d7ce522936589c630db08c0fa2405b21c4a5ff8ef19899900172736ba38.
//
// Solidity: event Withdrawn(uint128 indexed accountId, address indexed collateralType, uint256 tokenAmount, address indexed sender)
func (_Contracts *ContractsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractsWithdrawn, accountId []*big.Int, collateralType []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Withdrawn", accountIdRule, collateralTypeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsWithdrawn)
				if err := _Contracts.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Contracts *ContractsFilterer) ParseWithdrawn(log types.Log) (*ContractsWithdrawn, error) {
	event := new(ContractsWithdrawn)
	if err := _Contracts.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
