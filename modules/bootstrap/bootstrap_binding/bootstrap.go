// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bootstrap_binding

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BeaconChainProofsValidatorContainerProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsValidatorContainerProof struct {
	BeaconBlockTimestamp        *big.Int
	ValidatorIndex              *big.Int
	StateRoot                   [32]byte
	StateRootProof              [][32]byte
	ValidatorContainerRootProof [][32]byte
}

// BootstrapStorageTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type BootstrapStorageTokenInfo struct {
	Name          string
	Symbol        string
	TokenAddress  common.Address
	Decimals      uint8
	DepositAmount *big.Int
}

// IValidatorRegistryCommission is an auto generated low-level Go binding around an user-defined struct.
type IValidatorRegistryCommission struct {
	Rate          *big.Int
	MaxRate       *big.Int
	MaxChangeRate *big.Int
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// BootstrapABI is the input ABI used to generate the binding from.
const BootstrapABI = "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"endpoint_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structBootstrapStorage.ImmutableConfig\",\"components\":[{\"name\":\"imuachainChainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"beaconOracleAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeacon\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"imuaCapsuleBeacon\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"beaconProxyBytecode\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"networkConfig\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AFTER_PECTRA_MAX_EFFECTIVE_BALANCE_ETH_PER_VALIDATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"AFTER_PECTRA_MIN_ACTIVATION_BALANCE_ETH_PER_VALIDATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BEACON_ORACLE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BEACON_PROXY_BYTECODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBeaconProxyBytecode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IMUACHAIN_CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IMUA_ADDRESS_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IMUA_CAPSULE_BEACON\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VAULT_BEACON\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addWhitelistTokens\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tvlLimits\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowInitializePath\",\"inputs\":[{\"name\":\"origin\",\"type\":\"tuple\",\"internalType\":\"structOrigin\",\"components\":[{\"name\":\"srcEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bootstrapped\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimNSTFromImuachain\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimPrincipalFromImuachain\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimRewardFromImuachain\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"clientChainGatewayLogic\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clientChainInitializationData\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"commissionEdited\",\"inputs\":[{\"name\":\"imAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"hasEdited\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"composeMsgSender\",\"inputs\":[],\"outputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"consensusPublicKeyInUse\",\"inputs\":[{\"name\":\"consensusKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"used\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createImuaCapsule\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"customProxyAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegateTo\",\"inputs\":[{\"name\":\"validator\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegations\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"imAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationsByValidator\",\"inputs\":[{\"name\":\"imAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositThenDelegateTo\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validator\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositors\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositsByToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"endpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILayerZeroEndpointV2\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethToImAddress\",\"inputs\":[{\"name\":\"ethAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"imAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositorsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPubkeysCount\",\"inputs\":[{\"name\":\"stakerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorsCountForStakerToken\",\"inputs\":[{\"name\":\"stakerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWhitelistedTokenAtIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBootstrapStorage.TokenInfo\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWhitelistedTokensCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inboundNonce\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spawnTime_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"offsetDuration_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"whitelistTokens_\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tvlLimits_\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"customProxyAdmin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"clientChainGatewayLogic_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"clientChainInitializationData_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isCommissionValid\",\"inputs\":[{\"name\":\"commission\",\"type\":\"tuple\",\"internalType\":\"structIValidatorRegistry.Commission\",\"components\":[{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxChangeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isDepositor\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"hasDeposited\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLocked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidImAddress\",\"inputs\":[{\"name\":\"addressToValidate\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isWhitelistedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"whitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lzReceive\",\"inputs\":[{\"name\":\"_origin\",\"type\":\"tuple\",\"internalType\":\"structOrigin\",\"components\":[{\"name\":\"srcEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"_guid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_executor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"markBootstrapped\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextNonce\",\"inputs\":[{\"name\":\"srcEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oAppVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"senderVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiverVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"offsetDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerToCapsule\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"capsule\",\"type\":\"address\",\"internalType\":\"contractIImuaCapsule\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"peers\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"peer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerValidator\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"commission\",\"type\":\"tuple\",\"internalType\":\"structIValidatorRegistry.Commission\",\"components\":[{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxChangeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"consensusPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredValidators\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"replaceKey\",\"inputs\":[{\"name\":\"newKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestBeaconFullWithdrawal\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestBeaconPartialWithdrawal\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setClientChainGatewayLogic\",\"inputs\":[{\"name\":\"_clientChainGatewayLogic\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_clientChainInitializationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDelegate\",\"inputs\":[{\"name\":\"_delegate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOffsetDuration\",\"inputs\":[{\"name\":\"offsetDuration_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPeer\",\"inputs\":[{\"name\":\"_eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_peer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSpawnTime\",\"inputs\":[{\"name\":\"spawnTime_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spawnTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"stakerToPubkeyIDs\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerToTokenToValidators\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitReward\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"tokenToVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"contractIVault\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalDepositAmounts\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegateFrom\",\"inputs\":[{\"name\":\"validator\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"instantUnbond\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateRate\",\"inputs\":[{\"name\":\"newRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTvlLimit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tvlLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorNameInUse\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"used\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validators\",\"inputs\":[{\"name\":\"imAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"commission\",\"type\":\"tuple\",\"internalType\":\"structIValidatorRegistry.Commission\",\"components\":[{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxChangeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"consensusPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyAndDepositNativeStake\",\"inputs\":[{\"name\":\"validatorContainer\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.ValidatorContainerProof\",\"components\":[{\"name\":\"beaconBlockTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRootProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"validatorContainerRootProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"whitelistTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawPrincipal\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawReward\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawableAmounts\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BootstrapNotTimeYet\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BootstrapUpgradeFailed\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Bootstrapped\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BootstrappedAlready\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CapsuleCreated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"capsule\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimPrincipalResult\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClientChainGatewayLogicUpdated\",\"inputs\":[{\"name\":\"newLogic\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"initializationData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateResult\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegatee\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositResult\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositThenDelegateResult\",\"inputs\":[{\"name\":\"delegateSuccess\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegatee\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delegatedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageExecuted\",\"inputs\":[{\"name\":\"act\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumAction\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSent\",\"inputs\":[{\"name\":\"act\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumAction\"},{\"name\":\"packetId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"nativeFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OffsetDurationUpdated\",\"inputs\":[{\"name\":\"newOffsetDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PeerSet\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"peer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SpawnTimeUpdated\",\"inputs\":[{\"name\":\"newSpawnTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakedWithCapsule\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"capsule\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UndelegateResult\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"undelegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"undelegatee\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorCommissionUpdated\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorKeyReplaced\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newConsensusPublicKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegistered\",\"inputs\":[{\"name\":\"ethAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"validatorAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"commission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIValidatorRegistry.Commission\",\"components\":[{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxChangeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"consensusPublicKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultCreated\",\"inputs\":[{\"name\":\"underlyingToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WhitelistTokenAdded\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]"

// Bootstrap is an auto generated Go binding around an Ethereum contract.
type Bootstrap struct {
	BootstrapCaller     // Read-only binding to the contract
	BootstrapTransactor // Write-only binding to the contract
	BootstrapFilterer   // Log filterer for contract events
}

// BootstrapCaller is an auto generated read-only Go binding around an Ethereum contract.
type BootstrapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootstrapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BootstrapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootstrapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BootstrapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootstrapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BootstrapSession struct {
	Contract     *Bootstrap        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BootstrapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BootstrapCallerSession struct {
	Contract *BootstrapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BootstrapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BootstrapTransactorSession struct {
	Contract     *BootstrapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BootstrapRaw is an auto generated low-level Go binding around an Ethereum contract.
type BootstrapRaw struct {
	Contract *Bootstrap // Generic contract binding to access the raw methods on
}

// BootstrapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BootstrapCallerRaw struct {
	Contract *BootstrapCaller // Generic read-only contract binding to access the raw methods on
}

// BootstrapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BootstrapTransactorRaw struct {
	Contract *BootstrapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBootstrap creates a new instance of Bootstrap, bound to a specific deployed contract.
func NewBootstrap(address common.Address, backend bind.ContractBackend) (*Bootstrap, error) {
	contract, err := bindBootstrap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bootstrap{BootstrapCaller: BootstrapCaller{contract: contract}, BootstrapTransactor: BootstrapTransactor{contract: contract}, BootstrapFilterer: BootstrapFilterer{contract: contract}}, nil
}

// NewBootstrapCaller creates a new read-only instance of Bootstrap, bound to a specific deployed contract.
func NewBootstrapCaller(address common.Address, caller bind.ContractCaller) (*BootstrapCaller, error) {
	contract, err := bindBootstrap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BootstrapCaller{contract: contract}, nil
}

// NewBootstrapTransactor creates a new write-only instance of Bootstrap, bound to a specific deployed contract.
func NewBootstrapTransactor(address common.Address, transactor bind.ContractTransactor) (*BootstrapTransactor, error) {
	contract, err := bindBootstrap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BootstrapTransactor{contract: contract}, nil
}

// NewBootstrapFilterer creates a new log filterer instance of Bootstrap, bound to a specific deployed contract.
func NewBootstrapFilterer(address common.Address, filterer bind.ContractFilterer) (*BootstrapFilterer, error) {
	contract, err := bindBootstrap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BootstrapFilterer{contract: contract}, nil
}

// bindBootstrap binds a generic wrapper to an already deployed contract.
func bindBootstrap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BootstrapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bootstrap *BootstrapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bootstrap.Contract.BootstrapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bootstrap *BootstrapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bootstrap.Contract.BootstrapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bootstrap *BootstrapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bootstrap.Contract.BootstrapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bootstrap *BootstrapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bootstrap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bootstrap *BootstrapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bootstrap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bootstrap *BootstrapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bootstrap.Contract.contract.Transact(opts, method, params...)
}

// AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0x5fe60b0d.
//
// Solidity: function AFTER_PECTRA_MAX_EFFECTIVE_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_Bootstrap *BootstrapCaller) AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "AFTER_PECTRA_MAX_EFFECTIVE_BALANCE_ETH_PER_VALIDATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0x5fe60b0d.
//
// Solidity: function AFTER_PECTRA_MAX_EFFECTIVE_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_Bootstrap *BootstrapSession) AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR() (*big.Int, error) {
	return _Bootstrap.Contract.AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR(&_Bootstrap.CallOpts)
}

// AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0x5fe60b0d.
//
// Solidity: function AFTER_PECTRA_MAX_EFFECTIVE_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_Bootstrap *BootstrapCallerSession) AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR() (*big.Int, error) {
	return _Bootstrap.Contract.AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR(&_Bootstrap.CallOpts)
}

// AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0xf442193e.
//
// Solidity: function AFTER_PECTRA_MIN_ACTIVATION_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_Bootstrap *BootstrapCaller) AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "AFTER_PECTRA_MIN_ACTIVATION_BALANCE_ETH_PER_VALIDATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0xf442193e.
//
// Solidity: function AFTER_PECTRA_MIN_ACTIVATION_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_Bootstrap *BootstrapSession) AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR() (*big.Int, error) {
	return _Bootstrap.Contract.AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR(&_Bootstrap.CallOpts)
}

// AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0xf442193e.
//
// Solidity: function AFTER_PECTRA_MIN_ACTIVATION_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_Bootstrap *BootstrapCallerSession) AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR() (*big.Int, error) {
	return _Bootstrap.Contract.AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR(&_Bootstrap.CallOpts)
}

// BEACONORACLEADDRESS is a free data retrieval call binding the contract method 0x90a9b42e.
//
// Solidity: function BEACON_ORACLE_ADDRESS() view returns(address)
func (_Bootstrap *BootstrapCaller) BEACONORACLEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "BEACON_ORACLE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BEACONORACLEADDRESS is a free data retrieval call binding the contract method 0x90a9b42e.
//
// Solidity: function BEACON_ORACLE_ADDRESS() view returns(address)
func (_Bootstrap *BootstrapSession) BEACONORACLEADDRESS() (common.Address, error) {
	return _Bootstrap.Contract.BEACONORACLEADDRESS(&_Bootstrap.CallOpts)
}

// BEACONORACLEADDRESS is a free data retrieval call binding the contract method 0x90a9b42e.
//
// Solidity: function BEACON_ORACLE_ADDRESS() view returns(address)
func (_Bootstrap *BootstrapCallerSession) BEACONORACLEADDRESS() (common.Address, error) {
	return _Bootstrap.Contract.BEACONORACLEADDRESS(&_Bootstrap.CallOpts)
}

// BEACONPROXYBYTECODE is a free data retrieval call binding the contract method 0xcda3ef36.
//
// Solidity: function BEACON_PROXY_BYTECODE() view returns(address)
func (_Bootstrap *BootstrapCaller) BEACONPROXYBYTECODE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "BEACON_PROXY_BYTECODE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BEACONPROXYBYTECODE is a free data retrieval call binding the contract method 0xcda3ef36.
//
// Solidity: function BEACON_PROXY_BYTECODE() view returns(address)
func (_Bootstrap *BootstrapSession) BEACONPROXYBYTECODE() (common.Address, error) {
	return _Bootstrap.Contract.BEACONPROXYBYTECODE(&_Bootstrap.CallOpts)
}

// BEACONPROXYBYTECODE is a free data retrieval call binding the contract method 0xcda3ef36.
//
// Solidity: function BEACON_PROXY_BYTECODE() view returns(address)
func (_Bootstrap *BootstrapCallerSession) BEACONPROXYBYTECODE() (common.Address, error) {
	return _Bootstrap.Contract.BEACONPROXYBYTECODE(&_Bootstrap.CallOpts)
}

// IMUACHAINCHAINID is a free data retrieval call binding the contract method 0x1ef3a2df.
//
// Solidity: function IMUACHAIN_CHAIN_ID() view returns(uint32)
func (_Bootstrap *BootstrapCaller) IMUACHAINCHAINID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "IMUACHAIN_CHAIN_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// IMUACHAINCHAINID is a free data retrieval call binding the contract method 0x1ef3a2df.
//
// Solidity: function IMUACHAIN_CHAIN_ID() view returns(uint32)
func (_Bootstrap *BootstrapSession) IMUACHAINCHAINID() (uint32, error) {
	return _Bootstrap.Contract.IMUACHAINCHAINID(&_Bootstrap.CallOpts)
}

// IMUACHAINCHAINID is a free data retrieval call binding the contract method 0x1ef3a2df.
//
// Solidity: function IMUACHAIN_CHAIN_ID() view returns(uint32)
func (_Bootstrap *BootstrapCallerSession) IMUACHAINCHAINID() (uint32, error) {
	return _Bootstrap.Contract.IMUACHAINCHAINID(&_Bootstrap.CallOpts)
}

// IMUAADDRESSPREFIX is a free data retrieval call binding the contract method 0xe09b8274.
//
// Solidity: function IMUA_ADDRESS_PREFIX() view returns(bytes)
func (_Bootstrap *BootstrapCaller) IMUAADDRESSPREFIX(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "IMUA_ADDRESS_PREFIX")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// IMUAADDRESSPREFIX is a free data retrieval call binding the contract method 0xe09b8274.
//
// Solidity: function IMUA_ADDRESS_PREFIX() view returns(bytes)
func (_Bootstrap *BootstrapSession) IMUAADDRESSPREFIX() ([]byte, error) {
	return _Bootstrap.Contract.IMUAADDRESSPREFIX(&_Bootstrap.CallOpts)
}

// IMUAADDRESSPREFIX is a free data retrieval call binding the contract method 0xe09b8274.
//
// Solidity: function IMUA_ADDRESS_PREFIX() view returns(bytes)
func (_Bootstrap *BootstrapCallerSession) IMUAADDRESSPREFIX() ([]byte, error) {
	return _Bootstrap.Contract.IMUAADDRESSPREFIX(&_Bootstrap.CallOpts)
}

// IMUACAPSULEBEACON is a free data retrieval call binding the contract method 0xf67fadda.
//
// Solidity: function IMUA_CAPSULE_BEACON() view returns(address)
func (_Bootstrap *BootstrapCaller) IMUACAPSULEBEACON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "IMUA_CAPSULE_BEACON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IMUACAPSULEBEACON is a free data retrieval call binding the contract method 0xf67fadda.
//
// Solidity: function IMUA_CAPSULE_BEACON() view returns(address)
func (_Bootstrap *BootstrapSession) IMUACAPSULEBEACON() (common.Address, error) {
	return _Bootstrap.Contract.IMUACAPSULEBEACON(&_Bootstrap.CallOpts)
}

// IMUACAPSULEBEACON is a free data retrieval call binding the contract method 0xf67fadda.
//
// Solidity: function IMUA_CAPSULE_BEACON() view returns(address)
func (_Bootstrap *BootstrapCallerSession) IMUACAPSULEBEACON() (common.Address, error) {
	return _Bootstrap.Contract.IMUACAPSULEBEACON(&_Bootstrap.CallOpts)
}

// VAULTBEACON is a free data retrieval call binding the contract method 0x979c75d9.
//
// Solidity: function VAULT_BEACON() view returns(address)
func (_Bootstrap *BootstrapCaller) VAULTBEACON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "VAULT_BEACON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VAULTBEACON is a free data retrieval call binding the contract method 0x979c75d9.
//
// Solidity: function VAULT_BEACON() view returns(address)
func (_Bootstrap *BootstrapSession) VAULTBEACON() (common.Address, error) {
	return _Bootstrap.Contract.VAULTBEACON(&_Bootstrap.CallOpts)
}

// VAULTBEACON is a free data retrieval call binding the contract method 0x979c75d9.
//
// Solidity: function VAULT_BEACON() view returns(address)
func (_Bootstrap *BootstrapCallerSession) VAULTBEACON() (common.Address, error) {
	return _Bootstrap.Contract.VAULTBEACON(&_Bootstrap.CallOpts)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Bootstrap *BootstrapCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Bootstrap *BootstrapSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Bootstrap.Contract.AllowInitializePath(&_Bootstrap.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Bootstrap *BootstrapCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Bootstrap.Contract.AllowInitializePath(&_Bootstrap.CallOpts, origin)
}

// Bootstrapped is a free data retrieval call binding the contract method 0x35142c8c.
//
// Solidity: function bootstrapped() view returns(bool)
func (_Bootstrap *BootstrapCaller) Bootstrapped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "bootstrapped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Bootstrapped is a free data retrieval call binding the contract method 0x35142c8c.
//
// Solidity: function bootstrapped() view returns(bool)
func (_Bootstrap *BootstrapSession) Bootstrapped() (bool, error) {
	return _Bootstrap.Contract.Bootstrapped(&_Bootstrap.CallOpts)
}

// Bootstrapped is a free data retrieval call binding the contract method 0x35142c8c.
//
// Solidity: function bootstrapped() view returns(bool)
func (_Bootstrap *BootstrapCallerSession) Bootstrapped() (bool, error) {
	return _Bootstrap.Contract.Bootstrapped(&_Bootstrap.CallOpts)
}

// ClientChainGatewayLogic is a free data retrieval call binding the contract method 0x669aac44.
//
// Solidity: function clientChainGatewayLogic() view returns(address)
func (_Bootstrap *BootstrapCaller) ClientChainGatewayLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "clientChainGatewayLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClientChainGatewayLogic is a free data retrieval call binding the contract method 0x669aac44.
//
// Solidity: function clientChainGatewayLogic() view returns(address)
func (_Bootstrap *BootstrapSession) ClientChainGatewayLogic() (common.Address, error) {
	return _Bootstrap.Contract.ClientChainGatewayLogic(&_Bootstrap.CallOpts)
}

// ClientChainGatewayLogic is a free data retrieval call binding the contract method 0x669aac44.
//
// Solidity: function clientChainGatewayLogic() view returns(address)
func (_Bootstrap *BootstrapCallerSession) ClientChainGatewayLogic() (common.Address, error) {
	return _Bootstrap.Contract.ClientChainGatewayLogic(&_Bootstrap.CallOpts)
}

// ClientChainInitializationData is a free data retrieval call binding the contract method 0xf613fbec.
//
// Solidity: function clientChainInitializationData() view returns(bytes)
func (_Bootstrap *BootstrapCaller) ClientChainInitializationData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "clientChainInitializationData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ClientChainInitializationData is a free data retrieval call binding the contract method 0xf613fbec.
//
// Solidity: function clientChainInitializationData() view returns(bytes)
func (_Bootstrap *BootstrapSession) ClientChainInitializationData() ([]byte, error) {
	return _Bootstrap.Contract.ClientChainInitializationData(&_Bootstrap.CallOpts)
}

// ClientChainInitializationData is a free data retrieval call binding the contract method 0xf613fbec.
//
// Solidity: function clientChainInitializationData() view returns(bytes)
func (_Bootstrap *BootstrapCallerSession) ClientChainInitializationData() ([]byte, error) {
	return _Bootstrap.Contract.ClientChainInitializationData(&_Bootstrap.CallOpts)
}

// CommissionEdited is a free data retrieval call binding the contract method 0x31f13d30.
//
// Solidity: function commissionEdited(string imAddress) view returns(bool hasEdited)
func (_Bootstrap *BootstrapCaller) CommissionEdited(opts *bind.CallOpts, imAddress string) (bool, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "commissionEdited", imAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CommissionEdited is a free data retrieval call binding the contract method 0x31f13d30.
//
// Solidity: function commissionEdited(string imAddress) view returns(bool hasEdited)
func (_Bootstrap *BootstrapSession) CommissionEdited(imAddress string) (bool, error) {
	return _Bootstrap.Contract.CommissionEdited(&_Bootstrap.CallOpts, imAddress)
}

// CommissionEdited is a free data retrieval call binding the contract method 0x31f13d30.
//
// Solidity: function commissionEdited(string imAddress) view returns(bool hasEdited)
func (_Bootstrap *BootstrapCallerSession) CommissionEdited(imAddress string) (bool, error) {
	return _Bootstrap.Contract.CommissionEdited(&_Bootstrap.CallOpts, imAddress)
}

// ComposeMsgSender is a free data retrieval call binding the contract method 0xb92d0eff.
//
// Solidity: function composeMsgSender() view returns(address sender)
func (_Bootstrap *BootstrapCaller) ComposeMsgSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "composeMsgSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComposeMsgSender is a free data retrieval call binding the contract method 0xb92d0eff.
//
// Solidity: function composeMsgSender() view returns(address sender)
func (_Bootstrap *BootstrapSession) ComposeMsgSender() (common.Address, error) {
	return _Bootstrap.Contract.ComposeMsgSender(&_Bootstrap.CallOpts)
}

// ComposeMsgSender is a free data retrieval call binding the contract method 0xb92d0eff.
//
// Solidity: function composeMsgSender() view returns(address sender)
func (_Bootstrap *BootstrapCallerSession) ComposeMsgSender() (common.Address, error) {
	return _Bootstrap.Contract.ComposeMsgSender(&_Bootstrap.CallOpts)
}

// ConsensusPublicKeyInUse is a free data retrieval call binding the contract method 0xc489f42f.
//
// Solidity: function consensusPublicKeyInUse(bytes32 consensusKey) view returns(bool used)
func (_Bootstrap *BootstrapCaller) ConsensusPublicKeyInUse(opts *bind.CallOpts, consensusKey [32]byte) (bool, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "consensusPublicKeyInUse", consensusKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConsensusPublicKeyInUse is a free data retrieval call binding the contract method 0xc489f42f.
//
// Solidity: function consensusPublicKeyInUse(bytes32 consensusKey) view returns(bool used)
func (_Bootstrap *BootstrapSession) ConsensusPublicKeyInUse(consensusKey [32]byte) (bool, error) {
	return _Bootstrap.Contract.ConsensusPublicKeyInUse(&_Bootstrap.CallOpts, consensusKey)
}

// ConsensusPublicKeyInUse is a free data retrieval call binding the contract method 0xc489f42f.
//
// Solidity: function consensusPublicKeyInUse(bytes32 consensusKey) view returns(bool used)
func (_Bootstrap *BootstrapCallerSession) ConsensusPublicKeyInUse(consensusKey [32]byte) (bool, error) {
	return _Bootstrap.Contract.ConsensusPublicKeyInUse(&_Bootstrap.CallOpts, consensusKey)
}

// CustomProxyAdmin is a free data retrieval call binding the contract method 0x4caa408d.
//
// Solidity: function customProxyAdmin() view returns(address)
func (_Bootstrap *BootstrapCaller) CustomProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "customProxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CustomProxyAdmin is a free data retrieval call binding the contract method 0x4caa408d.
//
// Solidity: function customProxyAdmin() view returns(address)
func (_Bootstrap *BootstrapSession) CustomProxyAdmin() (common.Address, error) {
	return _Bootstrap.Contract.CustomProxyAdmin(&_Bootstrap.CallOpts)
}

// CustomProxyAdmin is a free data retrieval call binding the contract method 0x4caa408d.
//
// Solidity: function customProxyAdmin() view returns(address)
func (_Bootstrap *BootstrapCallerSession) CustomProxyAdmin() (common.Address, error) {
	return _Bootstrap.Contract.CustomProxyAdmin(&_Bootstrap.CallOpts)
}

// Delegations is a free data retrieval call binding the contract method 0xd428dc46.
//
// Solidity: function delegations(address delegator, string imAddress, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapCaller) Delegations(opts *bind.CallOpts, delegator common.Address, imAddress string, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "delegations", delegator, imAddress, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delegations is a free data retrieval call binding the contract method 0xd428dc46.
//
// Solidity: function delegations(address delegator, string imAddress, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapSession) Delegations(delegator common.Address, imAddress string, tokenAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.Delegations(&_Bootstrap.CallOpts, delegator, imAddress, tokenAddress)
}

// Delegations is a free data retrieval call binding the contract method 0xd428dc46.
//
// Solidity: function delegations(address delegator, string imAddress, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapCallerSession) Delegations(delegator common.Address, imAddress string, tokenAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.Delegations(&_Bootstrap.CallOpts, delegator, imAddress, tokenAddress)
}

// DelegationsByValidator is a free data retrieval call binding the contract method 0x2795ae23.
//
// Solidity: function delegationsByValidator(string imAddress, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapCaller) DelegationsByValidator(opts *bind.CallOpts, imAddress string, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "delegationsByValidator", imAddress, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationsByValidator is a free data retrieval call binding the contract method 0x2795ae23.
//
// Solidity: function delegationsByValidator(string imAddress, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapSession) DelegationsByValidator(imAddress string, tokenAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.DelegationsByValidator(&_Bootstrap.CallOpts, imAddress, tokenAddress)
}

// DelegationsByValidator is a free data retrieval call binding the contract method 0x2795ae23.
//
// Solidity: function delegationsByValidator(string imAddress, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapCallerSession) DelegationsByValidator(imAddress string, tokenAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.DelegationsByValidator(&_Bootstrap.CallOpts, imAddress, tokenAddress)
}

// Depositors is a free data retrieval call binding the contract method 0xe4b2fb79.
//
// Solidity: function depositors(uint256 ) view returns(address)
func (_Bootstrap *BootstrapCaller) Depositors(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "depositors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Depositors is a free data retrieval call binding the contract method 0xe4b2fb79.
//
// Solidity: function depositors(uint256 ) view returns(address)
func (_Bootstrap *BootstrapSession) Depositors(arg0 *big.Int) (common.Address, error) {
	return _Bootstrap.Contract.Depositors(&_Bootstrap.CallOpts, arg0)
}

// Depositors is a free data retrieval call binding the contract method 0xe4b2fb79.
//
// Solidity: function depositors(uint256 ) view returns(address)
func (_Bootstrap *BootstrapCallerSession) Depositors(arg0 *big.Int) (common.Address, error) {
	return _Bootstrap.Contract.Depositors(&_Bootstrap.CallOpts, arg0)
}

// DepositsByToken is a free data retrieval call binding the contract method 0x6edad9a6.
//
// Solidity: function depositsByToken(address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapCaller) DepositsByToken(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "depositsByToken", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositsByToken is a free data retrieval call binding the contract method 0x6edad9a6.
//
// Solidity: function depositsByToken(address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapSession) DepositsByToken(tokenAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.DepositsByToken(&_Bootstrap.CallOpts, tokenAddress)
}

// DepositsByToken is a free data retrieval call binding the contract method 0x6edad9a6.
//
// Solidity: function depositsByToken(address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapCallerSession) DepositsByToken(tokenAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.DepositsByToken(&_Bootstrap.CallOpts, tokenAddress)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Bootstrap *BootstrapCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Bootstrap *BootstrapSession) Endpoint() (common.Address, error) {
	return _Bootstrap.Contract.Endpoint(&_Bootstrap.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Bootstrap *BootstrapCallerSession) Endpoint() (common.Address, error) {
	return _Bootstrap.Contract.Endpoint(&_Bootstrap.CallOpts)
}

// EthToImAddress is a free data retrieval call binding the contract method 0x92f026f9.
//
// Solidity: function ethToImAddress(address ethAddress) view returns(string imAddress)
func (_Bootstrap *BootstrapCaller) EthToImAddress(opts *bind.CallOpts, ethAddress common.Address) (string, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "ethToImAddress", ethAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EthToImAddress is a free data retrieval call binding the contract method 0x92f026f9.
//
// Solidity: function ethToImAddress(address ethAddress) view returns(string imAddress)
func (_Bootstrap *BootstrapSession) EthToImAddress(ethAddress common.Address) (string, error) {
	return _Bootstrap.Contract.EthToImAddress(&_Bootstrap.CallOpts, ethAddress)
}

// EthToImAddress is a free data retrieval call binding the contract method 0x92f026f9.
//
// Solidity: function ethToImAddress(address ethAddress) view returns(string imAddress)
func (_Bootstrap *BootstrapCallerSession) EthToImAddress(ethAddress common.Address) (string, error) {
	return _Bootstrap.Contract.EthToImAddress(&_Bootstrap.CallOpts, ethAddress)
}

// GetDepositorsCount is a free data retrieval call binding the contract method 0x933ee626.
//
// Solidity: function getDepositorsCount() view returns(uint256)
func (_Bootstrap *BootstrapCaller) GetDepositorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "getDepositorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositorsCount is a free data retrieval call binding the contract method 0x933ee626.
//
// Solidity: function getDepositorsCount() view returns(uint256)
func (_Bootstrap *BootstrapSession) GetDepositorsCount() (*big.Int, error) {
	return _Bootstrap.Contract.GetDepositorsCount(&_Bootstrap.CallOpts)
}

// GetDepositorsCount is a free data retrieval call binding the contract method 0x933ee626.
//
// Solidity: function getDepositorsCount() view returns(uint256)
func (_Bootstrap *BootstrapCallerSession) GetDepositorsCount() (*big.Int, error) {
	return _Bootstrap.Contract.GetDepositorsCount(&_Bootstrap.CallOpts)
}

// GetPubkeysCount is a free data retrieval call binding the contract method 0x54f79dd2.
//
// Solidity: function getPubkeysCount(address stakerAddress) view returns(uint256)
func (_Bootstrap *BootstrapCaller) GetPubkeysCount(opts *bind.CallOpts, stakerAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "getPubkeysCount", stakerAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPubkeysCount is a free data retrieval call binding the contract method 0x54f79dd2.
//
// Solidity: function getPubkeysCount(address stakerAddress) view returns(uint256)
func (_Bootstrap *BootstrapSession) GetPubkeysCount(stakerAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.GetPubkeysCount(&_Bootstrap.CallOpts, stakerAddress)
}

// GetPubkeysCount is a free data retrieval call binding the contract method 0x54f79dd2.
//
// Solidity: function getPubkeysCount(address stakerAddress) view returns(uint256)
func (_Bootstrap *BootstrapCallerSession) GetPubkeysCount(stakerAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.GetPubkeysCount(&_Bootstrap.CallOpts, stakerAddress)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Bootstrap *BootstrapCaller) GetValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "getValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Bootstrap *BootstrapSession) GetValidatorsCount() (*big.Int, error) {
	return _Bootstrap.Contract.GetValidatorsCount(&_Bootstrap.CallOpts)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Bootstrap *BootstrapCallerSession) GetValidatorsCount() (*big.Int, error) {
	return _Bootstrap.Contract.GetValidatorsCount(&_Bootstrap.CallOpts)
}

// GetValidatorsCountForStakerToken is a free data retrieval call binding the contract method 0xcc43aa91.
//
// Solidity: function getValidatorsCountForStakerToken(address stakerAddress, address token) view returns(uint256)
func (_Bootstrap *BootstrapCaller) GetValidatorsCountForStakerToken(opts *bind.CallOpts, stakerAddress common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "getValidatorsCountForStakerToken", stakerAddress, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorsCountForStakerToken is a free data retrieval call binding the contract method 0xcc43aa91.
//
// Solidity: function getValidatorsCountForStakerToken(address stakerAddress, address token) view returns(uint256)
func (_Bootstrap *BootstrapSession) GetValidatorsCountForStakerToken(stakerAddress common.Address, token common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.GetValidatorsCountForStakerToken(&_Bootstrap.CallOpts, stakerAddress, token)
}

// GetValidatorsCountForStakerToken is a free data retrieval call binding the contract method 0xcc43aa91.
//
// Solidity: function getValidatorsCountForStakerToken(address stakerAddress, address token) view returns(uint256)
func (_Bootstrap *BootstrapCallerSession) GetValidatorsCountForStakerToken(stakerAddress common.Address, token common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.GetValidatorsCountForStakerToken(&_Bootstrap.CallOpts, stakerAddress, token)
}

// GetWhitelistedTokenAtIndex is a free data retrieval call binding the contract method 0x7c70c0a3.
//
// Solidity: function getWhitelistedTokenAtIndex(uint256 index) view returns((string,string,address,uint8,uint256))
func (_Bootstrap *BootstrapCaller) GetWhitelistedTokenAtIndex(opts *bind.CallOpts, index *big.Int) (BootstrapStorageTokenInfo, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "getWhitelistedTokenAtIndex", index)

	if err != nil {
		return *new(BootstrapStorageTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(BootstrapStorageTokenInfo)).(*BootstrapStorageTokenInfo)

	return out0, err

}

// GetWhitelistedTokenAtIndex is a free data retrieval call binding the contract method 0x7c70c0a3.
//
// Solidity: function getWhitelistedTokenAtIndex(uint256 index) view returns((string,string,address,uint8,uint256))
func (_Bootstrap *BootstrapSession) GetWhitelistedTokenAtIndex(index *big.Int) (BootstrapStorageTokenInfo, error) {
	return _Bootstrap.Contract.GetWhitelistedTokenAtIndex(&_Bootstrap.CallOpts, index)
}

// GetWhitelistedTokenAtIndex is a free data retrieval call binding the contract method 0x7c70c0a3.
//
// Solidity: function getWhitelistedTokenAtIndex(uint256 index) view returns((string,string,address,uint8,uint256))
func (_Bootstrap *BootstrapCallerSession) GetWhitelistedTokenAtIndex(index *big.Int) (BootstrapStorageTokenInfo, error) {
	return _Bootstrap.Contract.GetWhitelistedTokenAtIndex(&_Bootstrap.CallOpts, index)
}

// GetWhitelistedTokensCount is a free data retrieval call binding the contract method 0x9cd3f384.
//
// Solidity: function getWhitelistedTokensCount() view returns(uint256)
func (_Bootstrap *BootstrapCaller) GetWhitelistedTokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "getWhitelistedTokensCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWhitelistedTokensCount is a free data retrieval call binding the contract method 0x9cd3f384.
//
// Solidity: function getWhitelistedTokensCount() view returns(uint256)
func (_Bootstrap *BootstrapSession) GetWhitelistedTokensCount() (*big.Int, error) {
	return _Bootstrap.Contract.GetWhitelistedTokensCount(&_Bootstrap.CallOpts)
}

// GetWhitelistedTokensCount is a free data retrieval call binding the contract method 0x9cd3f384.
//
// Solidity: function getWhitelistedTokensCount() view returns(uint256)
func (_Bootstrap *BootstrapCallerSession) GetWhitelistedTokensCount() (*big.Int, error) {
	return _Bootstrap.Contract.GetWhitelistedTokensCount(&_Bootstrap.CallOpts)
}

// InboundNonce is a free data retrieval call binding the contract method 0x632284fd.
//
// Solidity: function inboundNonce(uint32 eid, bytes32 sender) view returns(uint64 nonce)
func (_Bootstrap *BootstrapCaller) InboundNonce(opts *bind.CallOpts, eid uint32, sender [32]byte) (uint64, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "inboundNonce", eid, sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InboundNonce is a free data retrieval call binding the contract method 0x632284fd.
//
// Solidity: function inboundNonce(uint32 eid, bytes32 sender) view returns(uint64 nonce)
func (_Bootstrap *BootstrapSession) InboundNonce(eid uint32, sender [32]byte) (uint64, error) {
	return _Bootstrap.Contract.InboundNonce(&_Bootstrap.CallOpts, eid, sender)
}

// InboundNonce is a free data retrieval call binding the contract method 0x632284fd.
//
// Solidity: function inboundNonce(uint32 eid, bytes32 sender) view returns(uint64 nonce)
func (_Bootstrap *BootstrapCallerSession) InboundNonce(eid uint32, sender [32]byte) (uint64, error) {
	return _Bootstrap.Contract.InboundNonce(&_Bootstrap.CallOpts, eid, sender)
}

// IsCommissionValid is a free data retrieval call binding the contract method 0xe6331ddc.
//
// Solidity: function isCommissionValid((uint256,uint256,uint256) commission) pure returns(bool)
func (_Bootstrap *BootstrapCaller) IsCommissionValid(opts *bind.CallOpts, commission IValidatorRegistryCommission) (bool, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "isCommissionValid", commission)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCommissionValid is a free data retrieval call binding the contract method 0xe6331ddc.
//
// Solidity: function isCommissionValid((uint256,uint256,uint256) commission) pure returns(bool)
func (_Bootstrap *BootstrapSession) IsCommissionValid(commission IValidatorRegistryCommission) (bool, error) {
	return _Bootstrap.Contract.IsCommissionValid(&_Bootstrap.CallOpts, commission)
}

// IsCommissionValid is a free data retrieval call binding the contract method 0xe6331ddc.
//
// Solidity: function isCommissionValid((uint256,uint256,uint256) commission) pure returns(bool)
func (_Bootstrap *BootstrapCallerSession) IsCommissionValid(commission IValidatorRegistryCommission) (bool, error) {
	return _Bootstrap.Contract.IsCommissionValid(&_Bootstrap.CallOpts, commission)
}

// IsDepositor is a free data retrieval call binding the contract method 0x2f70d1ba.
//
// Solidity: function isDepositor(address depositor) view returns(bool hasDeposited)
func (_Bootstrap *BootstrapCaller) IsDepositor(opts *bind.CallOpts, depositor common.Address) (bool, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "isDepositor", depositor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositor is a free data retrieval call binding the contract method 0x2f70d1ba.
//
// Solidity: function isDepositor(address depositor) view returns(bool hasDeposited)
func (_Bootstrap *BootstrapSession) IsDepositor(depositor common.Address) (bool, error) {
	return _Bootstrap.Contract.IsDepositor(&_Bootstrap.CallOpts, depositor)
}

// IsDepositor is a free data retrieval call binding the contract method 0x2f70d1ba.
//
// Solidity: function isDepositor(address depositor) view returns(bool hasDeposited)
func (_Bootstrap *BootstrapCallerSession) IsDepositor(depositor common.Address) (bool, error) {
	return _Bootstrap.Contract.IsDepositor(&_Bootstrap.CallOpts, depositor)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_Bootstrap *BootstrapCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_Bootstrap *BootstrapSession) IsLocked() (bool, error) {
	return _Bootstrap.Contract.IsLocked(&_Bootstrap.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_Bootstrap *BootstrapCallerSession) IsLocked() (bool, error) {
	return _Bootstrap.Contract.IsLocked(&_Bootstrap.CallOpts)
}

// IsValidImAddress is a free data retrieval call binding the contract method 0x72ac3ab6.
//
// Solidity: function isValidImAddress(string addressToValidate) pure returns(bool)
func (_Bootstrap *BootstrapCaller) IsValidImAddress(opts *bind.CallOpts, addressToValidate string) (bool, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "isValidImAddress", addressToValidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidImAddress is a free data retrieval call binding the contract method 0x72ac3ab6.
//
// Solidity: function isValidImAddress(string addressToValidate) pure returns(bool)
func (_Bootstrap *BootstrapSession) IsValidImAddress(addressToValidate string) (bool, error) {
	return _Bootstrap.Contract.IsValidImAddress(&_Bootstrap.CallOpts, addressToValidate)
}

// IsValidImAddress is a free data retrieval call binding the contract method 0x72ac3ab6.
//
// Solidity: function isValidImAddress(string addressToValidate) pure returns(bool)
func (_Bootstrap *BootstrapCallerSession) IsValidImAddress(addressToValidate string) (bool, error) {
	return _Bootstrap.Contract.IsValidImAddress(&_Bootstrap.CallOpts, addressToValidate)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token) view returns(bool whitelisted)
func (_Bootstrap *BootstrapCaller) IsWhitelistedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "isWhitelistedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token) view returns(bool whitelisted)
func (_Bootstrap *BootstrapSession) IsWhitelistedToken(token common.Address) (bool, error) {
	return _Bootstrap.Contract.IsWhitelistedToken(&_Bootstrap.CallOpts, token)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token) view returns(bool whitelisted)
func (_Bootstrap *BootstrapCallerSession) IsWhitelistedToken(token common.Address) (bool, error) {
	return _Bootstrap.Contract.IsWhitelistedToken(&_Bootstrap.CallOpts, token)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 srcEid, bytes32 sender) view returns(uint64)
func (_Bootstrap *BootstrapCaller) NextNonce(opts *bind.CallOpts, srcEid uint32, sender [32]byte) (uint64, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "nextNonce", srcEid, sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 srcEid, bytes32 sender) view returns(uint64)
func (_Bootstrap *BootstrapSession) NextNonce(srcEid uint32, sender [32]byte) (uint64, error) {
	return _Bootstrap.Contract.NextNonce(&_Bootstrap.CallOpts, srcEid, sender)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 srcEid, bytes32 sender) view returns(uint64)
func (_Bootstrap *BootstrapCallerSession) NextNonce(srcEid uint32, sender [32]byte) (uint64, error) {
	return _Bootstrap.Contract.NextNonce(&_Bootstrap.CallOpts, srcEid, sender)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() view returns(uint64 senderVersion, uint64 receiverVersion)
func (_Bootstrap *BootstrapCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "oAppVersion")

	outstruct := new(struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SenderVersion = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReceiverVersion = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() view returns(uint64 senderVersion, uint64 receiverVersion)
func (_Bootstrap *BootstrapSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Bootstrap.Contract.OAppVersion(&_Bootstrap.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() view returns(uint64 senderVersion, uint64 receiverVersion)
func (_Bootstrap *BootstrapCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Bootstrap.Contract.OAppVersion(&_Bootstrap.CallOpts)
}

// OffsetDuration is a free data retrieval call binding the contract method 0xec246031.
//
// Solidity: function offsetDuration() view returns(uint256)
func (_Bootstrap *BootstrapCaller) OffsetDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "offsetDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetDuration is a free data retrieval call binding the contract method 0xec246031.
//
// Solidity: function offsetDuration() view returns(uint256)
func (_Bootstrap *BootstrapSession) OffsetDuration() (*big.Int, error) {
	return _Bootstrap.Contract.OffsetDuration(&_Bootstrap.CallOpts)
}

// OffsetDuration is a free data retrieval call binding the contract method 0xec246031.
//
// Solidity: function offsetDuration() view returns(uint256)
func (_Bootstrap *BootstrapCallerSession) OffsetDuration() (*big.Int, error) {
	return _Bootstrap.Contract.OffsetDuration(&_Bootstrap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bootstrap *BootstrapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bootstrap *BootstrapSession) Owner() (common.Address, error) {
	return _Bootstrap.Contract.Owner(&_Bootstrap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bootstrap *BootstrapCallerSession) Owner() (common.Address, error) {
	return _Bootstrap.Contract.Owner(&_Bootstrap.CallOpts)
}

// OwnerToCapsule is a free data retrieval call binding the contract method 0x098ad41a.
//
// Solidity: function ownerToCapsule(address owner) view returns(address capsule)
func (_Bootstrap *BootstrapCaller) OwnerToCapsule(opts *bind.CallOpts, owner common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "ownerToCapsule", owner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerToCapsule is a free data retrieval call binding the contract method 0x098ad41a.
//
// Solidity: function ownerToCapsule(address owner) view returns(address capsule)
func (_Bootstrap *BootstrapSession) OwnerToCapsule(owner common.Address) (common.Address, error) {
	return _Bootstrap.Contract.OwnerToCapsule(&_Bootstrap.CallOpts, owner)
}

// OwnerToCapsule is a free data retrieval call binding the contract method 0x098ad41a.
//
// Solidity: function ownerToCapsule(address owner) view returns(address capsule)
func (_Bootstrap *BootstrapCallerSession) OwnerToCapsule(owner common.Address) (common.Address, error) {
	return _Bootstrap.Contract.OwnerToCapsule(&_Bootstrap.CallOpts, owner)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bootstrap *BootstrapCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bootstrap *BootstrapSession) Paused() (bool, error) {
	return _Bootstrap.Contract.Paused(&_Bootstrap.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bootstrap *BootstrapCallerSession) Paused() (bool, error) {
	return _Bootstrap.Contract.Paused(&_Bootstrap.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Bootstrap *BootstrapCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Bootstrap *BootstrapSession) Peers(eid uint32) ([32]byte, error) {
	return _Bootstrap.Contract.Peers(&_Bootstrap.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Bootstrap *BootstrapCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _Bootstrap.Contract.Peers(&_Bootstrap.CallOpts, eid)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xc306d47d.
//
// Solidity: function registeredValidators(uint256 ) view returns(address)
func (_Bootstrap *BootstrapCaller) RegisteredValidators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "registeredValidators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegisteredValidators is a free data retrieval call binding the contract method 0xc306d47d.
//
// Solidity: function registeredValidators(uint256 ) view returns(address)
func (_Bootstrap *BootstrapSession) RegisteredValidators(arg0 *big.Int) (common.Address, error) {
	return _Bootstrap.Contract.RegisteredValidators(&_Bootstrap.CallOpts, arg0)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xc306d47d.
//
// Solidity: function registeredValidators(uint256 ) view returns(address)
func (_Bootstrap *BootstrapCallerSession) RegisteredValidators(arg0 *big.Int) (common.Address, error) {
	return _Bootstrap.Contract.RegisteredValidators(&_Bootstrap.CallOpts, arg0)
}

// SpawnTime is a free data retrieval call binding the contract method 0x76fd464f.
//
// Solidity: function spawnTime() view returns(uint256)
func (_Bootstrap *BootstrapCaller) SpawnTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "spawnTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpawnTime is a free data retrieval call binding the contract method 0x76fd464f.
//
// Solidity: function spawnTime() view returns(uint256)
func (_Bootstrap *BootstrapSession) SpawnTime() (*big.Int, error) {
	return _Bootstrap.Contract.SpawnTime(&_Bootstrap.CallOpts)
}

// SpawnTime is a free data retrieval call binding the contract method 0x76fd464f.
//
// Solidity: function spawnTime() view returns(uint256)
func (_Bootstrap *BootstrapCallerSession) SpawnTime() (*big.Int, error) {
	return _Bootstrap.Contract.SpawnTime(&_Bootstrap.CallOpts)
}

// StakerToPubkeyIDs is a free data retrieval call binding the contract method 0x7e4af2d2.
//
// Solidity: function stakerToPubkeyIDs(address staker, uint256 ) view returns(bytes32)
func (_Bootstrap *BootstrapCaller) StakerToPubkeyIDs(opts *bind.CallOpts, staker common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "stakerToPubkeyIDs", staker, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StakerToPubkeyIDs is a free data retrieval call binding the contract method 0x7e4af2d2.
//
// Solidity: function stakerToPubkeyIDs(address staker, uint256 ) view returns(bytes32)
func (_Bootstrap *BootstrapSession) StakerToPubkeyIDs(staker common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Bootstrap.Contract.StakerToPubkeyIDs(&_Bootstrap.CallOpts, staker, arg1)
}

// StakerToPubkeyIDs is a free data retrieval call binding the contract method 0x7e4af2d2.
//
// Solidity: function stakerToPubkeyIDs(address staker, uint256 ) view returns(bytes32)
func (_Bootstrap *BootstrapCallerSession) StakerToPubkeyIDs(staker common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Bootstrap.Contract.StakerToPubkeyIDs(&_Bootstrap.CallOpts, staker, arg1)
}

// StakerToTokenToValidators is a free data retrieval call binding the contract method 0x304884d7.
//
// Solidity: function stakerToTokenToValidators(address staker, address token, uint256 ) view returns(string)
func (_Bootstrap *BootstrapCaller) StakerToTokenToValidators(opts *bind.CallOpts, staker common.Address, token common.Address, arg2 *big.Int) (string, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "stakerToTokenToValidators", staker, token, arg2)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StakerToTokenToValidators is a free data retrieval call binding the contract method 0x304884d7.
//
// Solidity: function stakerToTokenToValidators(address staker, address token, uint256 ) view returns(string)
func (_Bootstrap *BootstrapSession) StakerToTokenToValidators(staker common.Address, token common.Address, arg2 *big.Int) (string, error) {
	return _Bootstrap.Contract.StakerToTokenToValidators(&_Bootstrap.CallOpts, staker, token, arg2)
}

// StakerToTokenToValidators is a free data retrieval call binding the contract method 0x304884d7.
//
// Solidity: function stakerToTokenToValidators(address staker, address token, uint256 ) view returns(string)
func (_Bootstrap *BootstrapCallerSession) StakerToTokenToValidators(staker common.Address, token common.Address, arg2 *big.Int) (string, error) {
	return _Bootstrap.Contract.StakerToTokenToValidators(&_Bootstrap.CallOpts, staker, token, arg2)
}

// TokenToVault is a free data retrieval call binding the contract method 0x0c7e1725.
//
// Solidity: function tokenToVault(address token) view returns(address vault)
func (_Bootstrap *BootstrapCaller) TokenToVault(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "tokenToVault", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenToVault is a free data retrieval call binding the contract method 0x0c7e1725.
//
// Solidity: function tokenToVault(address token) view returns(address vault)
func (_Bootstrap *BootstrapSession) TokenToVault(token common.Address) (common.Address, error) {
	return _Bootstrap.Contract.TokenToVault(&_Bootstrap.CallOpts, token)
}

// TokenToVault is a free data retrieval call binding the contract method 0x0c7e1725.
//
// Solidity: function tokenToVault(address token) view returns(address vault)
func (_Bootstrap *BootstrapCallerSession) TokenToVault(token common.Address) (common.Address, error) {
	return _Bootstrap.Contract.TokenToVault(&_Bootstrap.CallOpts, token)
}

// TotalDepositAmounts is a free data retrieval call binding the contract method 0x1af20d25.
//
// Solidity: function totalDepositAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapCaller) TotalDepositAmounts(opts *bind.CallOpts, depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "totalDepositAmounts", depositor, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositAmounts is a free data retrieval call binding the contract method 0x1af20d25.
//
// Solidity: function totalDepositAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapSession) TotalDepositAmounts(depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.TotalDepositAmounts(&_Bootstrap.CallOpts, depositor, tokenAddress)
}

// TotalDepositAmounts is a free data retrieval call binding the contract method 0x1af20d25.
//
// Solidity: function totalDepositAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapCallerSession) TotalDepositAmounts(depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.TotalDepositAmounts(&_Bootstrap.CallOpts, depositor, tokenAddress)
}

// ValidatorNameInUse is a free data retrieval call binding the contract method 0x598e7fbd.
//
// Solidity: function validatorNameInUse(string name) view returns(bool used)
func (_Bootstrap *BootstrapCaller) ValidatorNameInUse(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "validatorNameInUse", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidatorNameInUse is a free data retrieval call binding the contract method 0x598e7fbd.
//
// Solidity: function validatorNameInUse(string name) view returns(bool used)
func (_Bootstrap *BootstrapSession) ValidatorNameInUse(name string) (bool, error) {
	return _Bootstrap.Contract.ValidatorNameInUse(&_Bootstrap.CallOpts, name)
}

// ValidatorNameInUse is a free data retrieval call binding the contract method 0x598e7fbd.
//
// Solidity: function validatorNameInUse(string name) view returns(bool used)
func (_Bootstrap *BootstrapCallerSession) ValidatorNameInUse(name string) (bool, error) {
	return _Bootstrap.Contract.ValidatorNameInUse(&_Bootstrap.CallOpts, name)
}

// Validators is a free data retrieval call binding the contract method 0xfacbd0e3.
//
// Solidity: function validators(string imAddress) view returns(string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey)
func (_Bootstrap *BootstrapCaller) Validators(opts *bind.CallOpts, imAddress string) (struct {
	Name               string
	Commission         IValidatorRegistryCommission
	ConsensusPublicKey [32]byte
}, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "validators", imAddress)

	outstruct := new(struct {
		Name               string
		Commission         IValidatorRegistryCommission
		ConsensusPublicKey [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Commission = *abi.ConvertType(out[1], new(IValidatorRegistryCommission)).(*IValidatorRegistryCommission)
	outstruct.ConsensusPublicKey = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfacbd0e3.
//
// Solidity: function validators(string imAddress) view returns(string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey)
func (_Bootstrap *BootstrapSession) Validators(imAddress string) (struct {
	Name               string
	Commission         IValidatorRegistryCommission
	ConsensusPublicKey [32]byte
}, error) {
	return _Bootstrap.Contract.Validators(&_Bootstrap.CallOpts, imAddress)
}

// Validators is a free data retrieval call binding the contract method 0xfacbd0e3.
//
// Solidity: function validators(string imAddress) view returns(string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey)
func (_Bootstrap *BootstrapCallerSession) Validators(imAddress string) (struct {
	Name               string
	Commission         IValidatorRegistryCommission
	ConsensusPublicKey [32]byte
}, error) {
	return _Bootstrap.Contract.Validators(&_Bootstrap.CallOpts, imAddress)
}

// WhitelistTokens is a free data retrieval call binding the contract method 0x602a70f1.
//
// Solidity: function whitelistTokens(uint256 ) view returns(address)
func (_Bootstrap *BootstrapCaller) WhitelistTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "whitelistTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistTokens is a free data retrieval call binding the contract method 0x602a70f1.
//
// Solidity: function whitelistTokens(uint256 ) view returns(address)
func (_Bootstrap *BootstrapSession) WhitelistTokens(arg0 *big.Int) (common.Address, error) {
	return _Bootstrap.Contract.WhitelistTokens(&_Bootstrap.CallOpts, arg0)
}

// WhitelistTokens is a free data retrieval call binding the contract method 0x602a70f1.
//
// Solidity: function whitelistTokens(uint256 ) view returns(address)
func (_Bootstrap *BootstrapCallerSession) WhitelistTokens(arg0 *big.Int) (common.Address, error) {
	return _Bootstrap.Contract.WhitelistTokens(&_Bootstrap.CallOpts, arg0)
}

// WithdrawReward is a free data retrieval call binding the contract method 0xb9c3da0f.
//
// Solidity: function withdrawReward(address , address , uint256 ) view returns()
func (_Bootstrap *BootstrapCaller) WithdrawReward(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "withdrawReward", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// WithdrawReward is a free data retrieval call binding the contract method 0xb9c3da0f.
//
// Solidity: function withdrawReward(address , address , uint256 ) view returns()
func (_Bootstrap *BootstrapSession) WithdrawReward(arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	return _Bootstrap.Contract.WithdrawReward(&_Bootstrap.CallOpts, arg0, arg1, arg2)
}

// WithdrawReward is a free data retrieval call binding the contract method 0xb9c3da0f.
//
// Solidity: function withdrawReward(address , address , uint256 ) view returns()
func (_Bootstrap *BootstrapCallerSession) WithdrawReward(arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	return _Bootstrap.Contract.WithdrawReward(&_Bootstrap.CallOpts, arg0, arg1, arg2)
}

// WithdrawableAmounts is a free data retrieval call binding the contract method 0x3a5a6389.
//
// Solidity: function withdrawableAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapCaller) WithdrawableAmounts(opts *bind.CallOpts, depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootstrap.contract.Call(opts, &out, "withdrawableAmounts", depositor, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableAmounts is a free data retrieval call binding the contract method 0x3a5a6389.
//
// Solidity: function withdrawableAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapSession) WithdrawableAmounts(depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.WithdrawableAmounts(&_Bootstrap.CallOpts, depositor, tokenAddress)
}

// WithdrawableAmounts is a free data retrieval call binding the contract method 0x3a5a6389.
//
// Solidity: function withdrawableAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_Bootstrap *BootstrapCallerSession) WithdrawableAmounts(depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _Bootstrap.Contract.WithdrawableAmounts(&_Bootstrap.CallOpts, depositor, tokenAddress)
}

// AddWhitelistTokens is a paid mutator transaction binding the contract method 0x8f733dbf.
//
// Solidity: function addWhitelistTokens(address[] tokens, uint256[] tvlLimits) returns()
func (_Bootstrap *BootstrapTransactor) AddWhitelistTokens(opts *bind.TransactOpts, tokens []common.Address, tvlLimits []*big.Int) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "addWhitelistTokens", tokens, tvlLimits)
}

// AddWhitelistTokens is a paid mutator transaction binding the contract method 0x8f733dbf.
//
// Solidity: function addWhitelistTokens(address[] tokens, uint256[] tvlLimits) returns()
func (_Bootstrap *BootstrapSession) AddWhitelistTokens(tokens []common.Address, tvlLimits []*big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.AddWhitelistTokens(&_Bootstrap.TransactOpts, tokens, tvlLimits)
}

// AddWhitelistTokens is a paid mutator transaction binding the contract method 0x8f733dbf.
//
// Solidity: function addWhitelistTokens(address[] tokens, uint256[] tvlLimits) returns()
func (_Bootstrap *BootstrapTransactorSession) AddWhitelistTokens(tokens []common.Address, tvlLimits []*big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.AddWhitelistTokens(&_Bootstrap.TransactOpts, tokens, tvlLimits)
}

// ClaimNSTFromImuachain is a paid mutator transaction binding the contract method 0xb7cb1f4b.
//
// Solidity: function claimNSTFromImuachain(uint256 ) payable returns()
func (_Bootstrap *BootstrapTransactor) ClaimNSTFromImuachain(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "claimNSTFromImuachain", arg0)
}

// ClaimNSTFromImuachain is a paid mutator transaction binding the contract method 0xb7cb1f4b.
//
// Solidity: function claimNSTFromImuachain(uint256 ) payable returns()
func (_Bootstrap *BootstrapSession) ClaimNSTFromImuachain(arg0 *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.ClaimNSTFromImuachain(&_Bootstrap.TransactOpts, arg0)
}

// ClaimNSTFromImuachain is a paid mutator transaction binding the contract method 0xb7cb1f4b.
//
// Solidity: function claimNSTFromImuachain(uint256 ) payable returns()
func (_Bootstrap *BootstrapTransactorSession) ClaimNSTFromImuachain(arg0 *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.ClaimNSTFromImuachain(&_Bootstrap.TransactOpts, arg0)
}

// ClaimPrincipalFromImuachain is a paid mutator transaction binding the contract method 0xc697d7e2.
//
// Solidity: function claimPrincipalFromImuachain(address token, uint256 amount) payable returns()
func (_Bootstrap *BootstrapTransactor) ClaimPrincipalFromImuachain(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "claimPrincipalFromImuachain", token, amount)
}

// ClaimPrincipalFromImuachain is a paid mutator transaction binding the contract method 0xc697d7e2.
//
// Solidity: function claimPrincipalFromImuachain(address token, uint256 amount) payable returns()
func (_Bootstrap *BootstrapSession) ClaimPrincipalFromImuachain(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.ClaimPrincipalFromImuachain(&_Bootstrap.TransactOpts, token, amount)
}

// ClaimPrincipalFromImuachain is a paid mutator transaction binding the contract method 0xc697d7e2.
//
// Solidity: function claimPrincipalFromImuachain(address token, uint256 amount) payable returns()
func (_Bootstrap *BootstrapTransactorSession) ClaimPrincipalFromImuachain(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.ClaimPrincipalFromImuachain(&_Bootstrap.TransactOpts, token, amount)
}

// ClaimRewardFromImuachain is a paid mutator transaction binding the contract method 0x8f70494d.
//
// Solidity: function claimRewardFromImuachain(address , uint256 ) payable returns()
func (_Bootstrap *BootstrapTransactor) ClaimRewardFromImuachain(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "claimRewardFromImuachain", arg0, arg1)
}

// ClaimRewardFromImuachain is a paid mutator transaction binding the contract method 0x8f70494d.
//
// Solidity: function claimRewardFromImuachain(address , uint256 ) payable returns()
func (_Bootstrap *BootstrapSession) ClaimRewardFromImuachain(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.ClaimRewardFromImuachain(&_Bootstrap.TransactOpts, arg0, arg1)
}

// ClaimRewardFromImuachain is a paid mutator transaction binding the contract method 0x8f70494d.
//
// Solidity: function claimRewardFromImuachain(address , uint256 ) payable returns()
func (_Bootstrap *BootstrapTransactorSession) ClaimRewardFromImuachain(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.ClaimRewardFromImuachain(&_Bootstrap.TransactOpts, arg0, arg1)
}

// CreateImuaCapsule is a paid mutator transaction binding the contract method 0x35a83698.
//
// Solidity: function createImuaCapsule() returns(address)
func (_Bootstrap *BootstrapTransactor) CreateImuaCapsule(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "createImuaCapsule")
}

// CreateImuaCapsule is a paid mutator transaction binding the contract method 0x35a83698.
//
// Solidity: function createImuaCapsule() returns(address)
func (_Bootstrap *BootstrapSession) CreateImuaCapsule() (*types.Transaction, error) {
	return _Bootstrap.Contract.CreateImuaCapsule(&_Bootstrap.TransactOpts)
}

// CreateImuaCapsule is a paid mutator transaction binding the contract method 0x35a83698.
//
// Solidity: function createImuaCapsule() returns(address)
func (_Bootstrap *BootstrapTransactorSession) CreateImuaCapsule() (*types.Transaction, error) {
	return _Bootstrap.Contract.CreateImuaCapsule(&_Bootstrap.TransactOpts)
}

// DelegateTo is a paid mutator transaction binding the contract method 0x7b357fcd.
//
// Solidity: function delegateTo(string validator, address token, uint256 amount) payable returns()
func (_Bootstrap *BootstrapTransactor) DelegateTo(opts *bind.TransactOpts, validator string, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "delegateTo", validator, token, amount)
}

// DelegateTo is a paid mutator transaction binding the contract method 0x7b357fcd.
//
// Solidity: function delegateTo(string validator, address token, uint256 amount) payable returns()
func (_Bootstrap *BootstrapSession) DelegateTo(validator string, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.DelegateTo(&_Bootstrap.TransactOpts, validator, token, amount)
}

// DelegateTo is a paid mutator transaction binding the contract method 0x7b357fcd.
//
// Solidity: function delegateTo(string validator, address token, uint256 amount) payable returns()
func (_Bootstrap *BootstrapTransactorSession) DelegateTo(validator string, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.DelegateTo(&_Bootstrap.TransactOpts, validator, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_Bootstrap *BootstrapTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_Bootstrap *BootstrapSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.Deposit(&_Bootstrap.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_Bootstrap *BootstrapTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.Deposit(&_Bootstrap.TransactOpts, token, amount)
}

// DepositThenDelegateTo is a paid mutator transaction binding the contract method 0x2e9fe843.
//
// Solidity: function depositThenDelegateTo(address token, uint256 amount, string validator) payable returns()
func (_Bootstrap *BootstrapTransactor) DepositThenDelegateTo(opts *bind.TransactOpts, token common.Address, amount *big.Int, validator string) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "depositThenDelegateTo", token, amount, validator)
}

// DepositThenDelegateTo is a paid mutator transaction binding the contract method 0x2e9fe843.
//
// Solidity: function depositThenDelegateTo(address token, uint256 amount, string validator) payable returns()
func (_Bootstrap *BootstrapSession) DepositThenDelegateTo(token common.Address, amount *big.Int, validator string) (*types.Transaction, error) {
	return _Bootstrap.Contract.DepositThenDelegateTo(&_Bootstrap.TransactOpts, token, amount, validator)
}

// DepositThenDelegateTo is a paid mutator transaction binding the contract method 0x2e9fe843.
//
// Solidity: function depositThenDelegateTo(address token, uint256 amount, string validator) payable returns()
func (_Bootstrap *BootstrapTransactorSession) DepositThenDelegateTo(token common.Address, amount *big.Int, validator string) (*types.Transaction, error) {
	return _Bootstrap.Contract.DepositThenDelegateTo(&_Bootstrap.TransactOpts, token, amount, validator)
}

// Initialize is a paid mutator transaction binding the contract method 0xba9abed1.
//
// Solidity: function initialize(address owner, uint256 spawnTime_, uint256 offsetDuration_, address[] whitelistTokens_, uint256[] tvlLimits_, address customProxyAdmin_, address clientChainGatewayLogic_, bytes clientChainInitializationData_) returns()
func (_Bootstrap *BootstrapTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, spawnTime_ *big.Int, offsetDuration_ *big.Int, whitelistTokens_ []common.Address, tvlLimits_ []*big.Int, customProxyAdmin_ common.Address, clientChainGatewayLogic_ common.Address, clientChainInitializationData_ []byte) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "initialize", owner, spawnTime_, offsetDuration_, whitelistTokens_, tvlLimits_, customProxyAdmin_, clientChainGatewayLogic_, clientChainInitializationData_)
}

// Initialize is a paid mutator transaction binding the contract method 0xba9abed1.
//
// Solidity: function initialize(address owner, uint256 spawnTime_, uint256 offsetDuration_, address[] whitelistTokens_, uint256[] tvlLimits_, address customProxyAdmin_, address clientChainGatewayLogic_, bytes clientChainInitializationData_) returns()
func (_Bootstrap *BootstrapSession) Initialize(owner common.Address, spawnTime_ *big.Int, offsetDuration_ *big.Int, whitelistTokens_ []common.Address, tvlLimits_ []*big.Int, customProxyAdmin_ common.Address, clientChainGatewayLogic_ common.Address, clientChainInitializationData_ []byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.Initialize(&_Bootstrap.TransactOpts, owner, spawnTime_, offsetDuration_, whitelistTokens_, tvlLimits_, customProxyAdmin_, clientChainGatewayLogic_, clientChainInitializationData_)
}

// Initialize is a paid mutator transaction binding the contract method 0xba9abed1.
//
// Solidity: function initialize(address owner, uint256 spawnTime_, uint256 offsetDuration_, address[] whitelistTokens_, uint256[] tvlLimits_, address customProxyAdmin_, address clientChainGatewayLogic_, bytes clientChainInitializationData_) returns()
func (_Bootstrap *BootstrapTransactorSession) Initialize(owner common.Address, spawnTime_ *big.Int, offsetDuration_ *big.Int, whitelistTokens_ []common.Address, tvlLimits_ []*big.Int, customProxyAdmin_ common.Address, clientChainGatewayLogic_ common.Address, clientChainInitializationData_ []byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.Initialize(&_Bootstrap.TransactOpts, owner, spawnTime_, offsetDuration_, whitelistTokens_, tvlLimits_, customProxyAdmin_, clientChainGatewayLogic_, clientChainInitializationData_)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Bootstrap *BootstrapTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Bootstrap *BootstrapSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.LzReceive(&_Bootstrap.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Bootstrap *BootstrapTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.LzReceive(&_Bootstrap.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// MarkBootstrapped is a paid mutator transaction binding the contract method 0xc605af52.
//
// Solidity: function markBootstrapped() returns()
func (_Bootstrap *BootstrapTransactor) MarkBootstrapped(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "markBootstrapped")
}

// MarkBootstrapped is a paid mutator transaction binding the contract method 0xc605af52.
//
// Solidity: function markBootstrapped() returns()
func (_Bootstrap *BootstrapSession) MarkBootstrapped() (*types.Transaction, error) {
	return _Bootstrap.Contract.MarkBootstrapped(&_Bootstrap.TransactOpts)
}

// MarkBootstrapped is a paid mutator transaction binding the contract method 0xc605af52.
//
// Solidity: function markBootstrapped() returns()
func (_Bootstrap *BootstrapTransactorSession) MarkBootstrapped() (*types.Transaction, error) {
	return _Bootstrap.Contract.MarkBootstrapped(&_Bootstrap.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bootstrap *BootstrapTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bootstrap *BootstrapSession) Pause() (*types.Transaction, error) {
	return _Bootstrap.Contract.Pause(&_Bootstrap.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bootstrap *BootstrapTransactorSession) Pause() (*types.Transaction, error) {
	return _Bootstrap.Contract.Pause(&_Bootstrap.TransactOpts)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x67b2353d.
//
// Solidity: function registerValidator(string validatorAddress, string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey) returns()
func (_Bootstrap *BootstrapTransactor) RegisterValidator(opts *bind.TransactOpts, validatorAddress string, name string, commission IValidatorRegistryCommission, consensusPublicKey [32]byte) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "registerValidator", validatorAddress, name, commission, consensusPublicKey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x67b2353d.
//
// Solidity: function registerValidator(string validatorAddress, string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey) returns()
func (_Bootstrap *BootstrapSession) RegisterValidator(validatorAddress string, name string, commission IValidatorRegistryCommission, consensusPublicKey [32]byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.RegisterValidator(&_Bootstrap.TransactOpts, validatorAddress, name, commission, consensusPublicKey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x67b2353d.
//
// Solidity: function registerValidator(string validatorAddress, string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey) returns()
func (_Bootstrap *BootstrapTransactorSession) RegisterValidator(validatorAddress string, name string, commission IValidatorRegistryCommission, consensusPublicKey [32]byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.RegisterValidator(&_Bootstrap.TransactOpts, validatorAddress, name, commission, consensusPublicKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bootstrap *BootstrapTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bootstrap *BootstrapSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bootstrap.Contract.RenounceOwnership(&_Bootstrap.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bootstrap *BootstrapTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bootstrap.Contract.RenounceOwnership(&_Bootstrap.TransactOpts)
}

// ReplaceKey is a paid mutator transaction binding the contract method 0x64fa45ca.
//
// Solidity: function replaceKey(bytes32 newKey) returns()
func (_Bootstrap *BootstrapTransactor) ReplaceKey(opts *bind.TransactOpts, newKey [32]byte) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "replaceKey", newKey)
}

// ReplaceKey is a paid mutator transaction binding the contract method 0x64fa45ca.
//
// Solidity: function replaceKey(bytes32 newKey) returns()
func (_Bootstrap *BootstrapSession) ReplaceKey(newKey [32]byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.ReplaceKey(&_Bootstrap.TransactOpts, newKey)
}

// ReplaceKey is a paid mutator transaction binding the contract method 0x64fa45ca.
//
// Solidity: function replaceKey(bytes32 newKey) returns()
func (_Bootstrap *BootstrapTransactorSession) ReplaceKey(newKey [32]byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.ReplaceKey(&_Bootstrap.TransactOpts, newKey)
}

// RequestBeaconFullWithdrawal is a paid mutator transaction binding the contract method 0x08004d30.
//
// Solidity: function requestBeaconFullWithdrawal(bytes ) payable returns()
func (_Bootstrap *BootstrapTransactor) RequestBeaconFullWithdrawal(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "requestBeaconFullWithdrawal", arg0)
}

// RequestBeaconFullWithdrawal is a paid mutator transaction binding the contract method 0x08004d30.
//
// Solidity: function requestBeaconFullWithdrawal(bytes ) payable returns()
func (_Bootstrap *BootstrapSession) RequestBeaconFullWithdrawal(arg0 []byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.RequestBeaconFullWithdrawal(&_Bootstrap.TransactOpts, arg0)
}

// RequestBeaconFullWithdrawal is a paid mutator transaction binding the contract method 0x08004d30.
//
// Solidity: function requestBeaconFullWithdrawal(bytes ) payable returns()
func (_Bootstrap *BootstrapTransactorSession) RequestBeaconFullWithdrawal(arg0 []byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.RequestBeaconFullWithdrawal(&_Bootstrap.TransactOpts, arg0)
}

// RequestBeaconPartialWithdrawal is a paid mutator transaction binding the contract method 0x7d0304e8.
//
// Solidity: function requestBeaconPartialWithdrawal(bytes , uint64 ) payable returns()
func (_Bootstrap *BootstrapTransactor) RequestBeaconPartialWithdrawal(opts *bind.TransactOpts, arg0 []byte, arg1 uint64) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "requestBeaconPartialWithdrawal", arg0, arg1)
}

// RequestBeaconPartialWithdrawal is a paid mutator transaction binding the contract method 0x7d0304e8.
//
// Solidity: function requestBeaconPartialWithdrawal(bytes , uint64 ) payable returns()
func (_Bootstrap *BootstrapSession) RequestBeaconPartialWithdrawal(arg0 []byte, arg1 uint64) (*types.Transaction, error) {
	return _Bootstrap.Contract.RequestBeaconPartialWithdrawal(&_Bootstrap.TransactOpts, arg0, arg1)
}

// RequestBeaconPartialWithdrawal is a paid mutator transaction binding the contract method 0x7d0304e8.
//
// Solidity: function requestBeaconPartialWithdrawal(bytes , uint64 ) payable returns()
func (_Bootstrap *BootstrapTransactorSession) RequestBeaconPartialWithdrawal(arg0 []byte, arg1 uint64) (*types.Transaction, error) {
	return _Bootstrap.Contract.RequestBeaconPartialWithdrawal(&_Bootstrap.TransactOpts, arg0, arg1)
}

// SetClientChainGatewayLogic is a paid mutator transaction binding the contract method 0xb7ec7448.
//
// Solidity: function setClientChainGatewayLogic(address _clientChainGatewayLogic, bytes _clientChainInitializationData) returns()
func (_Bootstrap *BootstrapTransactor) SetClientChainGatewayLogic(opts *bind.TransactOpts, _clientChainGatewayLogic common.Address, _clientChainInitializationData []byte) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "setClientChainGatewayLogic", _clientChainGatewayLogic, _clientChainInitializationData)
}

// SetClientChainGatewayLogic is a paid mutator transaction binding the contract method 0xb7ec7448.
//
// Solidity: function setClientChainGatewayLogic(address _clientChainGatewayLogic, bytes _clientChainInitializationData) returns()
func (_Bootstrap *BootstrapSession) SetClientChainGatewayLogic(_clientChainGatewayLogic common.Address, _clientChainInitializationData []byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.SetClientChainGatewayLogic(&_Bootstrap.TransactOpts, _clientChainGatewayLogic, _clientChainInitializationData)
}

// SetClientChainGatewayLogic is a paid mutator transaction binding the contract method 0xb7ec7448.
//
// Solidity: function setClientChainGatewayLogic(address _clientChainGatewayLogic, bytes _clientChainInitializationData) returns()
func (_Bootstrap *BootstrapTransactorSession) SetClientChainGatewayLogic(_clientChainGatewayLogic common.Address, _clientChainInitializationData []byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.SetClientChainGatewayLogic(&_Bootstrap.TransactOpts, _clientChainGatewayLogic, _clientChainInitializationData)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Bootstrap *BootstrapTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Bootstrap *BootstrapSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Bootstrap.Contract.SetDelegate(&_Bootstrap.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Bootstrap *BootstrapTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Bootstrap.Contract.SetDelegate(&_Bootstrap.TransactOpts, _delegate)
}

// SetOffsetDuration is a paid mutator transaction binding the contract method 0x4cc32457.
//
// Solidity: function setOffsetDuration(uint256 offsetDuration_) returns()
func (_Bootstrap *BootstrapTransactor) SetOffsetDuration(opts *bind.TransactOpts, offsetDuration_ *big.Int) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "setOffsetDuration", offsetDuration_)
}

// SetOffsetDuration is a paid mutator transaction binding the contract method 0x4cc32457.
//
// Solidity: function setOffsetDuration(uint256 offsetDuration_) returns()
func (_Bootstrap *BootstrapSession) SetOffsetDuration(offsetDuration_ *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.SetOffsetDuration(&_Bootstrap.TransactOpts, offsetDuration_)
}

// SetOffsetDuration is a paid mutator transaction binding the contract method 0x4cc32457.
//
// Solidity: function setOffsetDuration(uint256 offsetDuration_) returns()
func (_Bootstrap *BootstrapTransactorSession) SetOffsetDuration(offsetDuration_ *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.SetOffsetDuration(&_Bootstrap.TransactOpts, offsetDuration_)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Bootstrap *BootstrapTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Bootstrap *BootstrapSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.SetPeer(&_Bootstrap.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Bootstrap *BootstrapTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.SetPeer(&_Bootstrap.TransactOpts, _eid, _peer)
}

// SetSpawnTime is a paid mutator transaction binding the contract method 0x61163dfb.
//
// Solidity: function setSpawnTime(uint256 spawnTime_) returns()
func (_Bootstrap *BootstrapTransactor) SetSpawnTime(opts *bind.TransactOpts, spawnTime_ *big.Int) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "setSpawnTime", spawnTime_)
}

// SetSpawnTime is a paid mutator transaction binding the contract method 0x61163dfb.
//
// Solidity: function setSpawnTime(uint256 spawnTime_) returns()
func (_Bootstrap *BootstrapSession) SetSpawnTime(spawnTime_ *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.SetSpawnTime(&_Bootstrap.TransactOpts, spawnTime_)
}

// SetSpawnTime is a paid mutator transaction binding the contract method 0x61163dfb.
//
// Solidity: function setSpawnTime(uint256 spawnTime_) returns()
func (_Bootstrap *BootstrapTransactorSession) SetSpawnTime(spawnTime_ *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.SetSpawnTime(&_Bootstrap.TransactOpts, spawnTime_)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_Bootstrap *BootstrapTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_Bootstrap *BootstrapSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.Stake(&_Bootstrap.TransactOpts, pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_Bootstrap *BootstrapTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _Bootstrap.Contract.Stake(&_Bootstrap.TransactOpts, pubkey, signature, depositDataRoot)
}

// SubmitReward is a paid mutator transaction binding the contract method 0x9c48b825.
//
// Solidity: function submitReward(address , address , uint256 ) payable returns()
func (_Bootstrap *BootstrapTransactor) SubmitReward(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "submitReward", arg0, arg1, arg2)
}

// SubmitReward is a paid mutator transaction binding the contract method 0x9c48b825.
//
// Solidity: function submitReward(address , address , uint256 ) payable returns()
func (_Bootstrap *BootstrapSession) SubmitReward(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.SubmitReward(&_Bootstrap.TransactOpts, arg0, arg1, arg2)
}

// SubmitReward is a paid mutator transaction binding the contract method 0x9c48b825.
//
// Solidity: function submitReward(address , address , uint256 ) payable returns()
func (_Bootstrap *BootstrapTransactorSession) SubmitReward(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.SubmitReward(&_Bootstrap.TransactOpts, arg0, arg1, arg2)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bootstrap *BootstrapTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bootstrap *BootstrapSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bootstrap.Contract.TransferOwnership(&_Bootstrap.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bootstrap *BootstrapTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bootstrap.Contract.TransferOwnership(&_Bootstrap.TransactOpts, newOwner)
}

// UndelegateFrom is a paid mutator transaction binding the contract method 0x234c352c.
//
// Solidity: function undelegateFrom(string validator, address token, uint256 amount, bool instantUnbond) payable returns()
func (_Bootstrap *BootstrapTransactor) UndelegateFrom(opts *bind.TransactOpts, validator string, token common.Address, amount *big.Int, instantUnbond bool) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "undelegateFrom", validator, token, amount, instantUnbond)
}

// UndelegateFrom is a paid mutator transaction binding the contract method 0x234c352c.
//
// Solidity: function undelegateFrom(string validator, address token, uint256 amount, bool instantUnbond) payable returns()
func (_Bootstrap *BootstrapSession) UndelegateFrom(validator string, token common.Address, amount *big.Int, instantUnbond bool) (*types.Transaction, error) {
	return _Bootstrap.Contract.UndelegateFrom(&_Bootstrap.TransactOpts, validator, token, amount, instantUnbond)
}

// UndelegateFrom is a paid mutator transaction binding the contract method 0x234c352c.
//
// Solidity: function undelegateFrom(string validator, address token, uint256 amount, bool instantUnbond) payable returns()
func (_Bootstrap *BootstrapTransactorSession) UndelegateFrom(validator string, token common.Address, amount *big.Int, instantUnbond bool) (*types.Transaction, error) {
	return _Bootstrap.Contract.UndelegateFrom(&_Bootstrap.TransactOpts, validator, token, amount, instantUnbond)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bootstrap *BootstrapTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bootstrap *BootstrapSession) Unpause() (*types.Transaction, error) {
	return _Bootstrap.Contract.Unpause(&_Bootstrap.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bootstrap *BootstrapTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bootstrap.Contract.Unpause(&_Bootstrap.TransactOpts)
}

// UpdateRate is a paid mutator transaction binding the contract method 0x69ea1771.
//
// Solidity: function updateRate(uint256 newRate) returns()
func (_Bootstrap *BootstrapTransactor) UpdateRate(opts *bind.TransactOpts, newRate *big.Int) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "updateRate", newRate)
}

// UpdateRate is a paid mutator transaction binding the contract method 0x69ea1771.
//
// Solidity: function updateRate(uint256 newRate) returns()
func (_Bootstrap *BootstrapSession) UpdateRate(newRate *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.UpdateRate(&_Bootstrap.TransactOpts, newRate)
}

// UpdateRate is a paid mutator transaction binding the contract method 0x69ea1771.
//
// Solidity: function updateRate(uint256 newRate) returns()
func (_Bootstrap *BootstrapTransactorSession) UpdateRate(newRate *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.UpdateRate(&_Bootstrap.TransactOpts, newRate)
}

// UpdateTvlLimit is a paid mutator transaction binding the contract method 0x803e1500.
//
// Solidity: function updateTvlLimit(address token, uint256 tvlLimit) returns()
func (_Bootstrap *BootstrapTransactor) UpdateTvlLimit(opts *bind.TransactOpts, token common.Address, tvlLimit *big.Int) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "updateTvlLimit", token, tvlLimit)
}

// UpdateTvlLimit is a paid mutator transaction binding the contract method 0x803e1500.
//
// Solidity: function updateTvlLimit(address token, uint256 tvlLimit) returns()
func (_Bootstrap *BootstrapSession) UpdateTvlLimit(token common.Address, tvlLimit *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.UpdateTvlLimit(&_Bootstrap.TransactOpts, token, tvlLimit)
}

// UpdateTvlLimit is a paid mutator transaction binding the contract method 0x803e1500.
//
// Solidity: function updateTvlLimit(address token, uint256 tvlLimit) returns()
func (_Bootstrap *BootstrapTransactorSession) UpdateTvlLimit(token common.Address, tvlLimit *big.Int) (*types.Transaction, error) {
	return _Bootstrap.Contract.UpdateTvlLimit(&_Bootstrap.TransactOpts, token, tvlLimit)
}

// VerifyAndDepositNativeStake is a paid mutator transaction binding the contract method 0x8bedd239.
//
// Solidity: function verifyAndDepositNativeStake(bytes32[] validatorContainer, (uint256,uint256,bytes32,bytes32[],bytes32[]) proof) payable returns()
func (_Bootstrap *BootstrapTransactor) VerifyAndDepositNativeStake(opts *bind.TransactOpts, validatorContainer [][32]byte, proof BeaconChainProofsValidatorContainerProof) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "verifyAndDepositNativeStake", validatorContainer, proof)
}

// VerifyAndDepositNativeStake is a paid mutator transaction binding the contract method 0x8bedd239.
//
// Solidity: function verifyAndDepositNativeStake(bytes32[] validatorContainer, (uint256,uint256,bytes32,bytes32[],bytes32[]) proof) payable returns()
func (_Bootstrap *BootstrapSession) VerifyAndDepositNativeStake(validatorContainer [][32]byte, proof BeaconChainProofsValidatorContainerProof) (*types.Transaction, error) {
	return _Bootstrap.Contract.VerifyAndDepositNativeStake(&_Bootstrap.TransactOpts, validatorContainer, proof)
}

// VerifyAndDepositNativeStake is a paid mutator transaction binding the contract method 0x8bedd239.
//
// Solidity: function verifyAndDepositNativeStake(bytes32[] validatorContainer, (uint256,uint256,bytes32,bytes32[],bytes32[]) proof) payable returns()
func (_Bootstrap *BootstrapTransactorSession) VerifyAndDepositNativeStake(validatorContainer [][32]byte, proof BeaconChainProofsValidatorContainerProof) (*types.Transaction, error) {
	return _Bootstrap.Contract.VerifyAndDepositNativeStake(&_Bootstrap.TransactOpts, validatorContainer, proof)
}

// WithdrawPrincipal is a paid mutator transaction binding the contract method 0x20350767.
//
// Solidity: function withdrawPrincipal(address token, uint256 amount, address recipient) returns()
func (_Bootstrap *BootstrapTransactor) WithdrawPrincipal(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Bootstrap.contract.Transact(opts, "withdrawPrincipal", token, amount, recipient)
}

// WithdrawPrincipal is a paid mutator transaction binding the contract method 0x20350767.
//
// Solidity: function withdrawPrincipal(address token, uint256 amount, address recipient) returns()
func (_Bootstrap *BootstrapSession) WithdrawPrincipal(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Bootstrap.Contract.WithdrawPrincipal(&_Bootstrap.TransactOpts, token, amount, recipient)
}

// WithdrawPrincipal is a paid mutator transaction binding the contract method 0x20350767.
//
// Solidity: function withdrawPrincipal(address token, uint256 amount, address recipient) returns()
func (_Bootstrap *BootstrapTransactorSession) WithdrawPrincipal(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Bootstrap.Contract.WithdrawPrincipal(&_Bootstrap.TransactOpts, token, amount, recipient)
}

// BootstrapBootstrapNotTimeYetIterator is returned from FilterBootstrapNotTimeYet and is used to iterate over the raw logs and unpacked data for BootstrapNotTimeYet events raised by the Bootstrap contract.
type BootstrapBootstrapNotTimeYetIterator struct {
	Event *BootstrapBootstrapNotTimeYet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapBootstrapNotTimeYetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapBootstrapNotTimeYet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapBootstrapNotTimeYet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapBootstrapNotTimeYetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapBootstrapNotTimeYetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapBootstrapNotTimeYet represents a BootstrapNotTimeYet event raised by the Bootstrap contract.
type BootstrapBootstrapNotTimeYet struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBootstrapNotTimeYet is a free log retrieval operation binding the contract event 0xd8bb15fbd9470bdd3978999d7cefc45ea8566f88468b9ca55852fc9cdee48845.
//
// Solidity: event BootstrapNotTimeYet()
func (_Bootstrap *BootstrapFilterer) FilterBootstrapNotTimeYet(opts *bind.FilterOpts) (*BootstrapBootstrapNotTimeYetIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "BootstrapNotTimeYet")
	if err != nil {
		return nil, err
	}
	return &BootstrapBootstrapNotTimeYetIterator{contract: _Bootstrap.contract, event: "BootstrapNotTimeYet", logs: logs, sub: sub}, nil
}

// WatchBootstrapNotTimeYet is a free log subscription operation binding the contract event 0xd8bb15fbd9470bdd3978999d7cefc45ea8566f88468b9ca55852fc9cdee48845.
//
// Solidity: event BootstrapNotTimeYet()
func (_Bootstrap *BootstrapFilterer) WatchBootstrapNotTimeYet(opts *bind.WatchOpts, sink chan<- *BootstrapBootstrapNotTimeYet) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "BootstrapNotTimeYet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapBootstrapNotTimeYet)
				if err := _Bootstrap.contract.UnpackLog(event, "BootstrapNotTimeYet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBootstrapNotTimeYet is a log parse operation binding the contract event 0xd8bb15fbd9470bdd3978999d7cefc45ea8566f88468b9ca55852fc9cdee48845.
//
// Solidity: event BootstrapNotTimeYet()
func (_Bootstrap *BootstrapFilterer) ParseBootstrapNotTimeYet(log types.Log) (*BootstrapBootstrapNotTimeYet, error) {
	event := new(BootstrapBootstrapNotTimeYet)
	if err := _Bootstrap.contract.UnpackLog(event, "BootstrapNotTimeYet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapBootstrapUpgradeFailedIterator is returned from FilterBootstrapUpgradeFailed and is used to iterate over the raw logs and unpacked data for BootstrapUpgradeFailed events raised by the Bootstrap contract.
type BootstrapBootstrapUpgradeFailedIterator struct {
	Event *BootstrapBootstrapUpgradeFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapBootstrapUpgradeFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapBootstrapUpgradeFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapBootstrapUpgradeFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapBootstrapUpgradeFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapBootstrapUpgradeFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapBootstrapUpgradeFailed represents a BootstrapUpgradeFailed event raised by the Bootstrap contract.
type BootstrapBootstrapUpgradeFailed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBootstrapUpgradeFailed is a free log retrieval operation binding the contract event 0xdcb7e6fbb96e73875566de62d28669c83aff5eb7dfae2271b7f78ab7727d4b84.
//
// Solidity: event BootstrapUpgradeFailed()
func (_Bootstrap *BootstrapFilterer) FilterBootstrapUpgradeFailed(opts *bind.FilterOpts) (*BootstrapBootstrapUpgradeFailedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "BootstrapUpgradeFailed")
	if err != nil {
		return nil, err
	}
	return &BootstrapBootstrapUpgradeFailedIterator{contract: _Bootstrap.contract, event: "BootstrapUpgradeFailed", logs: logs, sub: sub}, nil
}

// WatchBootstrapUpgradeFailed is a free log subscription operation binding the contract event 0xdcb7e6fbb96e73875566de62d28669c83aff5eb7dfae2271b7f78ab7727d4b84.
//
// Solidity: event BootstrapUpgradeFailed()
func (_Bootstrap *BootstrapFilterer) WatchBootstrapUpgradeFailed(opts *bind.WatchOpts, sink chan<- *BootstrapBootstrapUpgradeFailed) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "BootstrapUpgradeFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapBootstrapUpgradeFailed)
				if err := _Bootstrap.contract.UnpackLog(event, "BootstrapUpgradeFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBootstrapUpgradeFailed is a log parse operation binding the contract event 0xdcb7e6fbb96e73875566de62d28669c83aff5eb7dfae2271b7f78ab7727d4b84.
//
// Solidity: event BootstrapUpgradeFailed()
func (_Bootstrap *BootstrapFilterer) ParseBootstrapUpgradeFailed(log types.Log) (*BootstrapBootstrapUpgradeFailed, error) {
	event := new(BootstrapBootstrapUpgradeFailed)
	if err := _Bootstrap.contract.UnpackLog(event, "BootstrapUpgradeFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapBootstrappedIterator is returned from FilterBootstrapped and is used to iterate over the raw logs and unpacked data for Bootstrapped events raised by the Bootstrap contract.
type BootstrapBootstrappedIterator struct {
	Event *BootstrapBootstrapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapBootstrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapBootstrapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapBootstrapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapBootstrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapBootstrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapBootstrapped represents a Bootstrapped event raised by the Bootstrap contract.
type BootstrapBootstrapped struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBootstrapped is a free log retrieval operation binding the contract event 0x90af5282d84f2aaaa645bb396ffbb1a1929fda2af47defb09244b31d458a719d.
//
// Solidity: event Bootstrapped()
func (_Bootstrap *BootstrapFilterer) FilterBootstrapped(opts *bind.FilterOpts) (*BootstrapBootstrappedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "Bootstrapped")
	if err != nil {
		return nil, err
	}
	return &BootstrapBootstrappedIterator{contract: _Bootstrap.contract, event: "Bootstrapped", logs: logs, sub: sub}, nil
}

// WatchBootstrapped is a free log subscription operation binding the contract event 0x90af5282d84f2aaaa645bb396ffbb1a1929fda2af47defb09244b31d458a719d.
//
// Solidity: event Bootstrapped()
func (_Bootstrap *BootstrapFilterer) WatchBootstrapped(opts *bind.WatchOpts, sink chan<- *BootstrapBootstrapped) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "Bootstrapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapBootstrapped)
				if err := _Bootstrap.contract.UnpackLog(event, "Bootstrapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBootstrapped is a log parse operation binding the contract event 0x90af5282d84f2aaaa645bb396ffbb1a1929fda2af47defb09244b31d458a719d.
//
// Solidity: event Bootstrapped()
func (_Bootstrap *BootstrapFilterer) ParseBootstrapped(log types.Log) (*BootstrapBootstrapped, error) {
	event := new(BootstrapBootstrapped)
	if err := _Bootstrap.contract.UnpackLog(event, "Bootstrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapBootstrappedAlreadyIterator is returned from FilterBootstrappedAlready and is used to iterate over the raw logs and unpacked data for BootstrappedAlready events raised by the Bootstrap contract.
type BootstrapBootstrappedAlreadyIterator struct {
	Event *BootstrapBootstrappedAlready // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapBootstrappedAlreadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapBootstrappedAlready)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapBootstrappedAlready)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapBootstrappedAlreadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapBootstrappedAlreadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapBootstrappedAlready represents a BootstrappedAlready event raised by the Bootstrap contract.
type BootstrapBootstrappedAlready struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBootstrappedAlready is a free log retrieval operation binding the contract event 0x6e5e65216df72fd325095b47d04105d3ca27805ca44babb89b131cb8640892ff.
//
// Solidity: event BootstrappedAlready()
func (_Bootstrap *BootstrapFilterer) FilterBootstrappedAlready(opts *bind.FilterOpts) (*BootstrapBootstrappedAlreadyIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "BootstrappedAlready")
	if err != nil {
		return nil, err
	}
	return &BootstrapBootstrappedAlreadyIterator{contract: _Bootstrap.contract, event: "BootstrappedAlready", logs: logs, sub: sub}, nil
}

// WatchBootstrappedAlready is a free log subscription operation binding the contract event 0x6e5e65216df72fd325095b47d04105d3ca27805ca44babb89b131cb8640892ff.
//
// Solidity: event BootstrappedAlready()
func (_Bootstrap *BootstrapFilterer) WatchBootstrappedAlready(opts *bind.WatchOpts, sink chan<- *BootstrapBootstrappedAlready) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "BootstrappedAlready")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapBootstrappedAlready)
				if err := _Bootstrap.contract.UnpackLog(event, "BootstrappedAlready", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBootstrappedAlready is a log parse operation binding the contract event 0x6e5e65216df72fd325095b47d04105d3ca27805ca44babb89b131cb8640892ff.
//
// Solidity: event BootstrappedAlready()
func (_Bootstrap *BootstrapFilterer) ParseBootstrappedAlready(log types.Log) (*BootstrapBootstrappedAlready, error) {
	event := new(BootstrapBootstrappedAlready)
	if err := _Bootstrap.contract.UnpackLog(event, "BootstrappedAlready", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapCapsuleCreatedIterator is returned from FilterCapsuleCreated and is used to iterate over the raw logs and unpacked data for CapsuleCreated events raised by the Bootstrap contract.
type BootstrapCapsuleCreatedIterator struct {
	Event *BootstrapCapsuleCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapCapsuleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapCapsuleCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapCapsuleCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapCapsuleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapCapsuleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapCapsuleCreated represents a CapsuleCreated event raised by the Bootstrap contract.
type BootstrapCapsuleCreated struct {
	Owner   common.Address
	Capsule common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCapsuleCreated is a free log retrieval operation binding the contract event 0x9b7b96ddfa41a5d76e277b661fe9d90e452022c554009c095d97269a615a11a3.
//
// Solidity: event CapsuleCreated(address indexed owner, address indexed capsule)
func (_Bootstrap *BootstrapFilterer) FilterCapsuleCreated(opts *bind.FilterOpts, owner []common.Address, capsule []common.Address) (*BootstrapCapsuleCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var capsuleRule []interface{}
	for _, capsuleItem := range capsule {
		capsuleRule = append(capsuleRule, capsuleItem)
	}

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "CapsuleCreated", ownerRule, capsuleRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapCapsuleCreatedIterator{contract: _Bootstrap.contract, event: "CapsuleCreated", logs: logs, sub: sub}, nil
}

// WatchCapsuleCreated is a free log subscription operation binding the contract event 0x9b7b96ddfa41a5d76e277b661fe9d90e452022c554009c095d97269a615a11a3.
//
// Solidity: event CapsuleCreated(address indexed owner, address indexed capsule)
func (_Bootstrap *BootstrapFilterer) WatchCapsuleCreated(opts *bind.WatchOpts, sink chan<- *BootstrapCapsuleCreated, owner []common.Address, capsule []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var capsuleRule []interface{}
	for _, capsuleItem := range capsule {
		capsuleRule = append(capsuleRule, capsuleItem)
	}

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "CapsuleCreated", ownerRule, capsuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapCapsuleCreated)
				if err := _Bootstrap.contract.UnpackLog(event, "CapsuleCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCapsuleCreated is a log parse operation binding the contract event 0x9b7b96ddfa41a5d76e277b661fe9d90e452022c554009c095d97269a615a11a3.
//
// Solidity: event CapsuleCreated(address indexed owner, address indexed capsule)
func (_Bootstrap *BootstrapFilterer) ParseCapsuleCreated(log types.Log) (*BootstrapCapsuleCreated, error) {
	event := new(BootstrapCapsuleCreated)
	if err := _Bootstrap.contract.UnpackLog(event, "CapsuleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapClaimPrincipalResultIterator is returned from FilterClaimPrincipalResult and is used to iterate over the raw logs and unpacked data for ClaimPrincipalResult events raised by the Bootstrap contract.
type BootstrapClaimPrincipalResultIterator struct {
	Event *BootstrapClaimPrincipalResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapClaimPrincipalResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapClaimPrincipalResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapClaimPrincipalResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapClaimPrincipalResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapClaimPrincipalResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapClaimPrincipalResult represents a ClaimPrincipalResult event raised by the Bootstrap contract.
type BootstrapClaimPrincipalResult struct {
	Success    bool
	Token      common.Address
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimPrincipalResult is a free log retrieval operation binding the contract event 0xeca0f64c3bab2d9808a4cb927796dd3c69892b2a9ae47cb5a078d6a4405f27f5.
//
// Solidity: event ClaimPrincipalResult(bool indexed success, address indexed token, address indexed withdrawer, uint256 amount)
func (_Bootstrap *BootstrapFilterer) FilterClaimPrincipalResult(opts *bind.FilterOpts, success []bool, token []common.Address, withdrawer []common.Address) (*BootstrapClaimPrincipalResultIterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "ClaimPrincipalResult", successRule, tokenRule, withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapClaimPrincipalResultIterator{contract: _Bootstrap.contract, event: "ClaimPrincipalResult", logs: logs, sub: sub}, nil
}

// WatchClaimPrincipalResult is a free log subscription operation binding the contract event 0xeca0f64c3bab2d9808a4cb927796dd3c69892b2a9ae47cb5a078d6a4405f27f5.
//
// Solidity: event ClaimPrincipalResult(bool indexed success, address indexed token, address indexed withdrawer, uint256 amount)
func (_Bootstrap *BootstrapFilterer) WatchClaimPrincipalResult(opts *bind.WatchOpts, sink chan<- *BootstrapClaimPrincipalResult, success []bool, token []common.Address, withdrawer []common.Address) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "ClaimPrincipalResult", successRule, tokenRule, withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapClaimPrincipalResult)
				if err := _Bootstrap.contract.UnpackLog(event, "ClaimPrincipalResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimPrincipalResult is a log parse operation binding the contract event 0xeca0f64c3bab2d9808a4cb927796dd3c69892b2a9ae47cb5a078d6a4405f27f5.
//
// Solidity: event ClaimPrincipalResult(bool indexed success, address indexed token, address indexed withdrawer, uint256 amount)
func (_Bootstrap *BootstrapFilterer) ParseClaimPrincipalResult(log types.Log) (*BootstrapClaimPrincipalResult, error) {
	event := new(BootstrapClaimPrincipalResult)
	if err := _Bootstrap.contract.UnpackLog(event, "ClaimPrincipalResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapClientChainGatewayLogicUpdatedIterator is returned from FilterClientChainGatewayLogicUpdated and is used to iterate over the raw logs and unpacked data for ClientChainGatewayLogicUpdated events raised by the Bootstrap contract.
type BootstrapClientChainGatewayLogicUpdatedIterator struct {
	Event *BootstrapClientChainGatewayLogicUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapClientChainGatewayLogicUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapClientChainGatewayLogicUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapClientChainGatewayLogicUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapClientChainGatewayLogicUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapClientChainGatewayLogicUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapClientChainGatewayLogicUpdated represents a ClientChainGatewayLogicUpdated event raised by the Bootstrap contract.
type BootstrapClientChainGatewayLogicUpdated struct {
	NewLogic           common.Address
	InitializationData []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterClientChainGatewayLogicUpdated is a free log retrieval operation binding the contract event 0xfef38cb604939fc77b2392a748b93e1d47d9c7220e77266894c4fe56c585564d.
//
// Solidity: event ClientChainGatewayLogicUpdated(address newLogic, bytes initializationData)
func (_Bootstrap *BootstrapFilterer) FilterClientChainGatewayLogicUpdated(opts *bind.FilterOpts) (*BootstrapClientChainGatewayLogicUpdatedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "ClientChainGatewayLogicUpdated")
	if err != nil {
		return nil, err
	}
	return &BootstrapClientChainGatewayLogicUpdatedIterator{contract: _Bootstrap.contract, event: "ClientChainGatewayLogicUpdated", logs: logs, sub: sub}, nil
}

// WatchClientChainGatewayLogicUpdated is a free log subscription operation binding the contract event 0xfef38cb604939fc77b2392a748b93e1d47d9c7220e77266894c4fe56c585564d.
//
// Solidity: event ClientChainGatewayLogicUpdated(address newLogic, bytes initializationData)
func (_Bootstrap *BootstrapFilterer) WatchClientChainGatewayLogicUpdated(opts *bind.WatchOpts, sink chan<- *BootstrapClientChainGatewayLogicUpdated) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "ClientChainGatewayLogicUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapClientChainGatewayLogicUpdated)
				if err := _Bootstrap.contract.UnpackLog(event, "ClientChainGatewayLogicUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClientChainGatewayLogicUpdated is a log parse operation binding the contract event 0xfef38cb604939fc77b2392a748b93e1d47d9c7220e77266894c4fe56c585564d.
//
// Solidity: event ClientChainGatewayLogicUpdated(address newLogic, bytes initializationData)
func (_Bootstrap *BootstrapFilterer) ParseClientChainGatewayLogicUpdated(log types.Log) (*BootstrapClientChainGatewayLogicUpdated, error) {
	event := new(BootstrapClientChainGatewayLogicUpdated)
	if err := _Bootstrap.contract.UnpackLog(event, "ClientChainGatewayLogicUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapDelegateResultIterator is returned from FilterDelegateResult and is used to iterate over the raw logs and unpacked data for DelegateResult events raised by the Bootstrap contract.
type BootstrapDelegateResultIterator struct {
	Event *BootstrapDelegateResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapDelegateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapDelegateResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapDelegateResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapDelegateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapDelegateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapDelegateResult represents a DelegateResult event raised by the Bootstrap contract.
type BootstrapDelegateResult struct {
	Success   bool
	Delegator common.Address
	Delegatee string
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegateResult is a free log retrieval operation binding the contract event 0xdc4fbf33793e5efc89a1c209786465733043b48c000e0d0611a72abe9a871bbc.
//
// Solidity: event DelegateResult(bool indexed success, address indexed delegator, string delegatee, address token, uint256 amount)
func (_Bootstrap *BootstrapFilterer) FilterDelegateResult(opts *bind.FilterOpts, success []bool, delegator []common.Address) (*BootstrapDelegateResultIterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "DelegateResult", successRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapDelegateResultIterator{contract: _Bootstrap.contract, event: "DelegateResult", logs: logs, sub: sub}, nil
}

// WatchDelegateResult is a free log subscription operation binding the contract event 0xdc4fbf33793e5efc89a1c209786465733043b48c000e0d0611a72abe9a871bbc.
//
// Solidity: event DelegateResult(bool indexed success, address indexed delegator, string delegatee, address token, uint256 amount)
func (_Bootstrap *BootstrapFilterer) WatchDelegateResult(opts *bind.WatchOpts, sink chan<- *BootstrapDelegateResult, success []bool, delegator []common.Address) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "DelegateResult", successRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapDelegateResult)
				if err := _Bootstrap.contract.UnpackLog(event, "DelegateResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateResult is a log parse operation binding the contract event 0xdc4fbf33793e5efc89a1c209786465733043b48c000e0d0611a72abe9a871bbc.
//
// Solidity: event DelegateResult(bool indexed success, address indexed delegator, string delegatee, address token, uint256 amount)
func (_Bootstrap *BootstrapFilterer) ParseDelegateResult(log types.Log) (*BootstrapDelegateResult, error) {
	event := new(BootstrapDelegateResult)
	if err := _Bootstrap.contract.UnpackLog(event, "DelegateResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapDepositResultIterator is returned from FilterDepositResult and is used to iterate over the raw logs and unpacked data for DepositResult events raised by the Bootstrap contract.
type BootstrapDepositResultIterator struct {
	Event *BootstrapDepositResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapDepositResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapDepositResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapDepositResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapDepositResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapDepositResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapDepositResult represents a DepositResult event raised by the Bootstrap contract.
type BootstrapDepositResult struct {
	Success   bool
	Token     common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositResult is a free log retrieval operation binding the contract event 0x96a7936e872abc3ec0271aa99dce327d8625f783cb4c455a36a7cd848d9ff1a9.
//
// Solidity: event DepositResult(bool indexed success, address indexed token, address indexed depositor, uint256 amount)
func (_Bootstrap *BootstrapFilterer) FilterDepositResult(opts *bind.FilterOpts, success []bool, token []common.Address, depositor []common.Address) (*BootstrapDepositResultIterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "DepositResult", successRule, tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapDepositResultIterator{contract: _Bootstrap.contract, event: "DepositResult", logs: logs, sub: sub}, nil
}

// WatchDepositResult is a free log subscription operation binding the contract event 0x96a7936e872abc3ec0271aa99dce327d8625f783cb4c455a36a7cd848d9ff1a9.
//
// Solidity: event DepositResult(bool indexed success, address indexed token, address indexed depositor, uint256 amount)
func (_Bootstrap *BootstrapFilterer) WatchDepositResult(opts *bind.WatchOpts, sink chan<- *BootstrapDepositResult, success []bool, token []common.Address, depositor []common.Address) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "DepositResult", successRule, tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapDepositResult)
				if err := _Bootstrap.contract.UnpackLog(event, "DepositResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositResult is a log parse operation binding the contract event 0x96a7936e872abc3ec0271aa99dce327d8625f783cb4c455a36a7cd848d9ff1a9.
//
// Solidity: event DepositResult(bool indexed success, address indexed token, address indexed depositor, uint256 amount)
func (_Bootstrap *BootstrapFilterer) ParseDepositResult(log types.Log) (*BootstrapDepositResult, error) {
	event := new(BootstrapDepositResult)
	if err := _Bootstrap.contract.UnpackLog(event, "DepositResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapDepositThenDelegateResultIterator is returned from FilterDepositThenDelegateResult and is used to iterate over the raw logs and unpacked data for DepositThenDelegateResult events raised by the Bootstrap contract.
type BootstrapDepositThenDelegateResultIterator struct {
	Event *BootstrapDepositThenDelegateResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapDepositThenDelegateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapDepositThenDelegateResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapDepositThenDelegateResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapDepositThenDelegateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapDepositThenDelegateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapDepositThenDelegateResult represents a DepositThenDelegateResult event raised by the Bootstrap contract.
type BootstrapDepositThenDelegateResult struct {
	DelegateSuccess bool
	Delegator       common.Address
	Delegatee       common.Hash
	Token           common.Address
	DelegatedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositThenDelegateResult is a free log retrieval operation binding the contract event 0x7a8415b264ee200190fa69106c4629c425ee8398b8bb4bfec724b8de85cff5f8.
//
// Solidity: event DepositThenDelegateResult(bool indexed delegateSuccess, address indexed delegator, string indexed delegatee, address token, uint256 delegatedAmount)
func (_Bootstrap *BootstrapFilterer) FilterDepositThenDelegateResult(opts *bind.FilterOpts, delegateSuccess []bool, delegator []common.Address, delegatee []string) (*BootstrapDepositThenDelegateResultIterator, error) {

	var delegateSuccessRule []interface{}
	for _, delegateSuccessItem := range delegateSuccess {
		delegateSuccessRule = append(delegateSuccessRule, delegateSuccessItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "DepositThenDelegateResult", delegateSuccessRule, delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapDepositThenDelegateResultIterator{contract: _Bootstrap.contract, event: "DepositThenDelegateResult", logs: logs, sub: sub}, nil
}

// WatchDepositThenDelegateResult is a free log subscription operation binding the contract event 0x7a8415b264ee200190fa69106c4629c425ee8398b8bb4bfec724b8de85cff5f8.
//
// Solidity: event DepositThenDelegateResult(bool indexed delegateSuccess, address indexed delegator, string indexed delegatee, address token, uint256 delegatedAmount)
func (_Bootstrap *BootstrapFilterer) WatchDepositThenDelegateResult(opts *bind.WatchOpts, sink chan<- *BootstrapDepositThenDelegateResult, delegateSuccess []bool, delegator []common.Address, delegatee []string) (event.Subscription, error) {

	var delegateSuccessRule []interface{}
	for _, delegateSuccessItem := range delegateSuccess {
		delegateSuccessRule = append(delegateSuccessRule, delegateSuccessItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "DepositThenDelegateResult", delegateSuccessRule, delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapDepositThenDelegateResult)
				if err := _Bootstrap.contract.UnpackLog(event, "DepositThenDelegateResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositThenDelegateResult is a log parse operation binding the contract event 0x7a8415b264ee200190fa69106c4629c425ee8398b8bb4bfec724b8de85cff5f8.
//
// Solidity: event DepositThenDelegateResult(bool indexed delegateSuccess, address indexed delegator, string indexed delegatee, address token, uint256 delegatedAmount)
func (_Bootstrap *BootstrapFilterer) ParseDepositThenDelegateResult(log types.Log) (*BootstrapDepositThenDelegateResult, error) {
	event := new(BootstrapDepositThenDelegateResult)
	if err := _Bootstrap.contract.UnpackLog(event, "DepositThenDelegateResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bootstrap contract.
type BootstrapInitializedIterator struct {
	Event *BootstrapInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapInitialized represents a Initialized event raised by the Bootstrap contract.
type BootstrapInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bootstrap *BootstrapFilterer) FilterInitialized(opts *bind.FilterOpts) (*BootstrapInitializedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BootstrapInitializedIterator{contract: _Bootstrap.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bootstrap *BootstrapFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BootstrapInitialized) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapInitialized)
				if err := _Bootstrap.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bootstrap *BootstrapFilterer) ParseInitialized(log types.Log) (*BootstrapInitialized, error) {
	event := new(BootstrapInitialized)
	if err := _Bootstrap.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapMessageExecutedIterator is returned from FilterMessageExecuted and is used to iterate over the raw logs and unpacked data for MessageExecuted events raised by the Bootstrap contract.
type BootstrapMessageExecutedIterator struct {
	Event *BootstrapMessageExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapMessageExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapMessageExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapMessageExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapMessageExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapMessageExecuted represents a MessageExecuted event raised by the Bootstrap contract.
type BootstrapMessageExecuted struct {
	Act   uint8
	Nonce uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMessageExecuted is a free log retrieval operation binding the contract event 0xb82f0bf216aec4c87941891caba2e306587945a146feb0cbcd67b55d22b88501.
//
// Solidity: event MessageExecuted(uint8 indexed act, uint64 nonce)
func (_Bootstrap *BootstrapFilterer) FilterMessageExecuted(opts *bind.FilterOpts, act []uint8) (*BootstrapMessageExecutedIterator, error) {

	var actRule []interface{}
	for _, actItem := range act {
		actRule = append(actRule, actItem)
	}

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "MessageExecuted", actRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapMessageExecutedIterator{contract: _Bootstrap.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

// WatchMessageExecuted is a free log subscription operation binding the contract event 0xb82f0bf216aec4c87941891caba2e306587945a146feb0cbcd67b55d22b88501.
//
// Solidity: event MessageExecuted(uint8 indexed act, uint64 nonce)
func (_Bootstrap *BootstrapFilterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *BootstrapMessageExecuted, act []uint8) (event.Subscription, error) {

	var actRule []interface{}
	for _, actItem := range act {
		actRule = append(actRule, actItem)
	}

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "MessageExecuted", actRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapMessageExecuted)
				if err := _Bootstrap.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageExecuted is a log parse operation binding the contract event 0xb82f0bf216aec4c87941891caba2e306587945a146feb0cbcd67b55d22b88501.
//
// Solidity: event MessageExecuted(uint8 indexed act, uint64 nonce)
func (_Bootstrap *BootstrapFilterer) ParseMessageExecuted(log types.Log) (*BootstrapMessageExecuted, error) {
	event := new(BootstrapMessageExecuted)
	if err := _Bootstrap.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the Bootstrap contract.
type BootstrapMessageSentIterator struct {
	Event *BootstrapMessageSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapMessageSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapMessageSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapMessageSent represents a MessageSent event raised by the Bootstrap contract.
type BootstrapMessageSent struct {
	Act       uint8
	PacketId  [32]byte
	Nonce     uint64
	NativeFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0xfba7df935971916019b62d3b0eee5d6989cdb9f8081346465209d2887e4c0968.
//
// Solidity: event MessageSent(uint8 indexed act, bytes32 packetId, uint64 nonce, uint256 nativeFee)
func (_Bootstrap *BootstrapFilterer) FilterMessageSent(opts *bind.FilterOpts, act []uint8) (*BootstrapMessageSentIterator, error) {

	var actRule []interface{}
	for _, actItem := range act {
		actRule = append(actRule, actItem)
	}

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "MessageSent", actRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapMessageSentIterator{contract: _Bootstrap.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0xfba7df935971916019b62d3b0eee5d6989cdb9f8081346465209d2887e4c0968.
//
// Solidity: event MessageSent(uint8 indexed act, bytes32 packetId, uint64 nonce, uint256 nativeFee)
func (_Bootstrap *BootstrapFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *BootstrapMessageSent, act []uint8) (event.Subscription, error) {

	var actRule []interface{}
	for _, actItem := range act {
		actRule = append(actRule, actItem)
	}

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "MessageSent", actRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapMessageSent)
				if err := _Bootstrap.contract.UnpackLog(event, "MessageSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageSent is a log parse operation binding the contract event 0xfba7df935971916019b62d3b0eee5d6989cdb9f8081346465209d2887e4c0968.
//
// Solidity: event MessageSent(uint8 indexed act, bytes32 packetId, uint64 nonce, uint256 nativeFee)
func (_Bootstrap *BootstrapFilterer) ParseMessageSent(log types.Log) (*BootstrapMessageSent, error) {
	event := new(BootstrapMessageSent)
	if err := _Bootstrap.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapOffsetDurationUpdatedIterator is returned from FilterOffsetDurationUpdated and is used to iterate over the raw logs and unpacked data for OffsetDurationUpdated events raised by the Bootstrap contract.
type BootstrapOffsetDurationUpdatedIterator struct {
	Event *BootstrapOffsetDurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapOffsetDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapOffsetDurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapOffsetDurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapOffsetDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapOffsetDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapOffsetDurationUpdated represents a OffsetDurationUpdated event raised by the Bootstrap contract.
type BootstrapOffsetDurationUpdated struct {
	NewOffsetDuration *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOffsetDurationUpdated is a free log retrieval operation binding the contract event 0x33ea0c40bbb70e43d4a933b05cb895a4a5f876debd3b8ed39c3c4d1e016f4ad3.
//
// Solidity: event OffsetDurationUpdated(uint256 newOffsetDuration)
func (_Bootstrap *BootstrapFilterer) FilterOffsetDurationUpdated(opts *bind.FilterOpts) (*BootstrapOffsetDurationUpdatedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "OffsetDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &BootstrapOffsetDurationUpdatedIterator{contract: _Bootstrap.contract, event: "OffsetDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchOffsetDurationUpdated is a free log subscription operation binding the contract event 0x33ea0c40bbb70e43d4a933b05cb895a4a5f876debd3b8ed39c3c4d1e016f4ad3.
//
// Solidity: event OffsetDurationUpdated(uint256 newOffsetDuration)
func (_Bootstrap *BootstrapFilterer) WatchOffsetDurationUpdated(opts *bind.WatchOpts, sink chan<- *BootstrapOffsetDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "OffsetDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapOffsetDurationUpdated)
				if err := _Bootstrap.contract.UnpackLog(event, "OffsetDurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOffsetDurationUpdated is a log parse operation binding the contract event 0x33ea0c40bbb70e43d4a933b05cb895a4a5f876debd3b8ed39c3c4d1e016f4ad3.
//
// Solidity: event OffsetDurationUpdated(uint256 newOffsetDuration)
func (_Bootstrap *BootstrapFilterer) ParseOffsetDurationUpdated(log types.Log) (*BootstrapOffsetDurationUpdated, error) {
	event := new(BootstrapOffsetDurationUpdated)
	if err := _Bootstrap.contract.UnpackLog(event, "OffsetDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bootstrap contract.
type BootstrapOwnershipTransferredIterator struct {
	Event *BootstrapOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapOwnershipTransferred represents a OwnershipTransferred event raised by the Bootstrap contract.
type BootstrapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bootstrap *BootstrapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BootstrapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapOwnershipTransferredIterator{contract: _Bootstrap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bootstrap *BootstrapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BootstrapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapOwnershipTransferred)
				if err := _Bootstrap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bootstrap *BootstrapFilterer) ParseOwnershipTransferred(log types.Log) (*BootstrapOwnershipTransferred, error) {
	event := new(BootstrapOwnershipTransferred)
	if err := _Bootstrap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bootstrap contract.
type BootstrapPausedIterator struct {
	Event *BootstrapPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapPaused represents a Paused event raised by the Bootstrap contract.
type BootstrapPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bootstrap *BootstrapFilterer) FilterPaused(opts *bind.FilterOpts) (*BootstrapPausedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BootstrapPausedIterator{contract: _Bootstrap.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bootstrap *BootstrapFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BootstrapPaused) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapPaused)
				if err := _Bootstrap.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bootstrap *BootstrapFilterer) ParsePaused(log types.Log) (*BootstrapPaused, error) {
	event := new(BootstrapPaused)
	if err := _Bootstrap.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapPeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the Bootstrap contract.
type BootstrapPeerSetIterator struct {
	Event *BootstrapPeerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapPeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapPeerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapPeerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapPeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapPeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapPeerSet represents a PeerSet event raised by the Bootstrap contract.
type BootstrapPeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Bootstrap *BootstrapFilterer) FilterPeerSet(opts *bind.FilterOpts) (*BootstrapPeerSetIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &BootstrapPeerSetIterator{contract: _Bootstrap.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Bootstrap *BootstrapFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *BootstrapPeerSet) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapPeerSet)
				if err := _Bootstrap.contract.UnpackLog(event, "PeerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePeerSet is a log parse operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Bootstrap *BootstrapFilterer) ParsePeerSet(log types.Log) (*BootstrapPeerSet, error) {
	event := new(BootstrapPeerSet)
	if err := _Bootstrap.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapSpawnTimeUpdatedIterator is returned from FilterSpawnTimeUpdated and is used to iterate over the raw logs and unpacked data for SpawnTimeUpdated events raised by the Bootstrap contract.
type BootstrapSpawnTimeUpdatedIterator struct {
	Event *BootstrapSpawnTimeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapSpawnTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapSpawnTimeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapSpawnTimeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapSpawnTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapSpawnTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapSpawnTimeUpdated represents a SpawnTimeUpdated event raised by the Bootstrap contract.
type BootstrapSpawnTimeUpdated struct {
	NewSpawnTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSpawnTimeUpdated is a free log retrieval operation binding the contract event 0xa2d6431b356061bf0720994ce72f8bb72d17ed14a040f92c3600c55b857095db.
//
// Solidity: event SpawnTimeUpdated(uint256 newSpawnTime)
func (_Bootstrap *BootstrapFilterer) FilterSpawnTimeUpdated(opts *bind.FilterOpts) (*BootstrapSpawnTimeUpdatedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "SpawnTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &BootstrapSpawnTimeUpdatedIterator{contract: _Bootstrap.contract, event: "SpawnTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchSpawnTimeUpdated is a free log subscription operation binding the contract event 0xa2d6431b356061bf0720994ce72f8bb72d17ed14a040f92c3600c55b857095db.
//
// Solidity: event SpawnTimeUpdated(uint256 newSpawnTime)
func (_Bootstrap *BootstrapFilterer) WatchSpawnTimeUpdated(opts *bind.WatchOpts, sink chan<- *BootstrapSpawnTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "SpawnTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapSpawnTimeUpdated)
				if err := _Bootstrap.contract.UnpackLog(event, "SpawnTimeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSpawnTimeUpdated is a log parse operation binding the contract event 0xa2d6431b356061bf0720994ce72f8bb72d17ed14a040f92c3600c55b857095db.
//
// Solidity: event SpawnTimeUpdated(uint256 newSpawnTime)
func (_Bootstrap *BootstrapFilterer) ParseSpawnTimeUpdated(log types.Log) (*BootstrapSpawnTimeUpdated, error) {
	event := new(BootstrapSpawnTimeUpdated)
	if err := _Bootstrap.contract.UnpackLog(event, "SpawnTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStakedWithCapsuleIterator is returned from FilterStakedWithCapsule and is used to iterate over the raw logs and unpacked data for StakedWithCapsule events raised by the Bootstrap contract.
type BootstrapStakedWithCapsuleIterator struct {
	Event *BootstrapStakedWithCapsule // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStakedWithCapsuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStakedWithCapsule)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStakedWithCapsule)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStakedWithCapsuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStakedWithCapsuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStakedWithCapsule represents a StakedWithCapsule event raised by the Bootstrap contract.
type BootstrapStakedWithCapsule struct {
	Staker  common.Address
	Capsule common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakedWithCapsule is a free log retrieval operation binding the contract event 0x9baefaaa262c2c2f0c3acf8ee9293d5043ef654d71c2aba83024f1322c72e5ee.
//
// Solidity: event StakedWithCapsule(address indexed staker, address indexed capsule)
func (_Bootstrap *BootstrapFilterer) FilterStakedWithCapsule(opts *bind.FilterOpts, staker []common.Address, capsule []common.Address) (*BootstrapStakedWithCapsuleIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var capsuleRule []interface{}
	for _, capsuleItem := range capsule {
		capsuleRule = append(capsuleRule, capsuleItem)
	}

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "StakedWithCapsule", stakerRule, capsuleRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapStakedWithCapsuleIterator{contract: _Bootstrap.contract, event: "StakedWithCapsule", logs: logs, sub: sub}, nil
}

// WatchStakedWithCapsule is a free log subscription operation binding the contract event 0x9baefaaa262c2c2f0c3acf8ee9293d5043ef654d71c2aba83024f1322c72e5ee.
//
// Solidity: event StakedWithCapsule(address indexed staker, address indexed capsule)
func (_Bootstrap *BootstrapFilterer) WatchStakedWithCapsule(opts *bind.WatchOpts, sink chan<- *BootstrapStakedWithCapsule, staker []common.Address, capsule []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var capsuleRule []interface{}
	for _, capsuleItem := range capsule {
		capsuleRule = append(capsuleRule, capsuleItem)
	}

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "StakedWithCapsule", stakerRule, capsuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStakedWithCapsule)
				if err := _Bootstrap.contract.UnpackLog(event, "StakedWithCapsule", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakedWithCapsule is a log parse operation binding the contract event 0x9baefaaa262c2c2f0c3acf8ee9293d5043ef654d71c2aba83024f1322c72e5ee.
//
// Solidity: event StakedWithCapsule(address indexed staker, address indexed capsule)
func (_Bootstrap *BootstrapFilterer) ParseStakedWithCapsule(log types.Log) (*BootstrapStakedWithCapsule, error) {
	event := new(BootstrapStakedWithCapsule)
	if err := _Bootstrap.contract.UnpackLog(event, "StakedWithCapsule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapUndelegateResultIterator is returned from FilterUndelegateResult and is used to iterate over the raw logs and unpacked data for UndelegateResult events raised by the Bootstrap contract.
type BootstrapUndelegateResultIterator struct {
	Event *BootstrapUndelegateResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapUndelegateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapUndelegateResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapUndelegateResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapUndelegateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapUndelegateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapUndelegateResult represents a UndelegateResult event raised by the Bootstrap contract.
type BootstrapUndelegateResult struct {
	Success     bool
	Undelegator common.Address
	Undelegatee string
	Token       common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUndelegateResult is a free log retrieval operation binding the contract event 0x8bc7dddfddc42a489a24fbcbb5e45de61c15f3228d343959e1b9c176f835531a.
//
// Solidity: event UndelegateResult(bool indexed success, address indexed undelegator, string undelegatee, address token, uint256 amount)
func (_Bootstrap *BootstrapFilterer) FilterUndelegateResult(opts *bind.FilterOpts, success []bool, undelegator []common.Address) (*BootstrapUndelegateResultIterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var undelegatorRule []interface{}
	for _, undelegatorItem := range undelegator {
		undelegatorRule = append(undelegatorRule, undelegatorItem)
	}

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "UndelegateResult", successRule, undelegatorRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapUndelegateResultIterator{contract: _Bootstrap.contract, event: "UndelegateResult", logs: logs, sub: sub}, nil
}

// WatchUndelegateResult is a free log subscription operation binding the contract event 0x8bc7dddfddc42a489a24fbcbb5e45de61c15f3228d343959e1b9c176f835531a.
//
// Solidity: event UndelegateResult(bool indexed success, address indexed undelegator, string undelegatee, address token, uint256 amount)
func (_Bootstrap *BootstrapFilterer) WatchUndelegateResult(opts *bind.WatchOpts, sink chan<- *BootstrapUndelegateResult, success []bool, undelegator []common.Address) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var undelegatorRule []interface{}
	for _, undelegatorItem := range undelegator {
		undelegatorRule = append(undelegatorRule, undelegatorItem)
	}

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "UndelegateResult", successRule, undelegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapUndelegateResult)
				if err := _Bootstrap.contract.UnpackLog(event, "UndelegateResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegateResult is a log parse operation binding the contract event 0x8bc7dddfddc42a489a24fbcbb5e45de61c15f3228d343959e1b9c176f835531a.
//
// Solidity: event UndelegateResult(bool indexed success, address indexed undelegator, string undelegatee, address token, uint256 amount)
func (_Bootstrap *BootstrapFilterer) ParseUndelegateResult(log types.Log) (*BootstrapUndelegateResult, error) {
	event := new(BootstrapUndelegateResult)
	if err := _Bootstrap.contract.UnpackLog(event, "UndelegateResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bootstrap contract.
type BootstrapUnpausedIterator struct {
	Event *BootstrapUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapUnpaused represents a Unpaused event raised by the Bootstrap contract.
type BootstrapUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bootstrap *BootstrapFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BootstrapUnpausedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BootstrapUnpausedIterator{contract: _Bootstrap.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bootstrap *BootstrapFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BootstrapUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapUnpaused)
				if err := _Bootstrap.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bootstrap *BootstrapFilterer) ParseUnpaused(log types.Log) (*BootstrapUnpaused, error) {
	event := new(BootstrapUnpaused)
	if err := _Bootstrap.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapValidatorCommissionUpdatedIterator is returned from FilterValidatorCommissionUpdated and is used to iterate over the raw logs and unpacked data for ValidatorCommissionUpdated events raised by the Bootstrap contract.
type BootstrapValidatorCommissionUpdatedIterator struct {
	Event *BootstrapValidatorCommissionUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapValidatorCommissionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapValidatorCommissionUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapValidatorCommissionUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapValidatorCommissionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapValidatorCommissionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapValidatorCommissionUpdated represents a ValidatorCommissionUpdated event raised by the Bootstrap contract.
type BootstrapValidatorCommissionUpdated struct {
	ValidatorAddress string
	NewRate          *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorCommissionUpdated is a free log retrieval operation binding the contract event 0xca076e65d6f1051e9be83c618630ccf55477aabef0858a2eb2a5c7a69ecd650b.
//
// Solidity: event ValidatorCommissionUpdated(string validatorAddress, uint256 newRate)
func (_Bootstrap *BootstrapFilterer) FilterValidatorCommissionUpdated(opts *bind.FilterOpts) (*BootstrapValidatorCommissionUpdatedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "ValidatorCommissionUpdated")
	if err != nil {
		return nil, err
	}
	return &BootstrapValidatorCommissionUpdatedIterator{contract: _Bootstrap.contract, event: "ValidatorCommissionUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorCommissionUpdated is a free log subscription operation binding the contract event 0xca076e65d6f1051e9be83c618630ccf55477aabef0858a2eb2a5c7a69ecd650b.
//
// Solidity: event ValidatorCommissionUpdated(string validatorAddress, uint256 newRate)
func (_Bootstrap *BootstrapFilterer) WatchValidatorCommissionUpdated(opts *bind.WatchOpts, sink chan<- *BootstrapValidatorCommissionUpdated) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "ValidatorCommissionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapValidatorCommissionUpdated)
				if err := _Bootstrap.contract.UnpackLog(event, "ValidatorCommissionUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorCommissionUpdated is a log parse operation binding the contract event 0xca076e65d6f1051e9be83c618630ccf55477aabef0858a2eb2a5c7a69ecd650b.
//
// Solidity: event ValidatorCommissionUpdated(string validatorAddress, uint256 newRate)
func (_Bootstrap *BootstrapFilterer) ParseValidatorCommissionUpdated(log types.Log) (*BootstrapValidatorCommissionUpdated, error) {
	event := new(BootstrapValidatorCommissionUpdated)
	if err := _Bootstrap.contract.UnpackLog(event, "ValidatorCommissionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapValidatorKeyReplacedIterator is returned from FilterValidatorKeyReplaced and is used to iterate over the raw logs and unpacked data for ValidatorKeyReplaced events raised by the Bootstrap contract.
type BootstrapValidatorKeyReplacedIterator struct {
	Event *BootstrapValidatorKeyReplaced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapValidatorKeyReplacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapValidatorKeyReplaced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapValidatorKeyReplaced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapValidatorKeyReplacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapValidatorKeyReplacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapValidatorKeyReplaced represents a ValidatorKeyReplaced event raised by the Bootstrap contract.
type BootstrapValidatorKeyReplaced struct {
	ValidatorAddress      string
	NewConsensusPublicKey [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorKeyReplaced is a free log retrieval operation binding the contract event 0x859e5bfc563f9796ce88c2845945349aef9e44bb7fdccb47ed0def686cc840d8.
//
// Solidity: event ValidatorKeyReplaced(string validatorAddress, bytes32 newConsensusPublicKey)
func (_Bootstrap *BootstrapFilterer) FilterValidatorKeyReplaced(opts *bind.FilterOpts) (*BootstrapValidatorKeyReplacedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "ValidatorKeyReplaced")
	if err != nil {
		return nil, err
	}
	return &BootstrapValidatorKeyReplacedIterator{contract: _Bootstrap.contract, event: "ValidatorKeyReplaced", logs: logs, sub: sub}, nil
}

// WatchValidatorKeyReplaced is a free log subscription operation binding the contract event 0x859e5bfc563f9796ce88c2845945349aef9e44bb7fdccb47ed0def686cc840d8.
//
// Solidity: event ValidatorKeyReplaced(string validatorAddress, bytes32 newConsensusPublicKey)
func (_Bootstrap *BootstrapFilterer) WatchValidatorKeyReplaced(opts *bind.WatchOpts, sink chan<- *BootstrapValidatorKeyReplaced) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "ValidatorKeyReplaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapValidatorKeyReplaced)
				if err := _Bootstrap.contract.UnpackLog(event, "ValidatorKeyReplaced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorKeyReplaced is a log parse operation binding the contract event 0x859e5bfc563f9796ce88c2845945349aef9e44bb7fdccb47ed0def686cc840d8.
//
// Solidity: event ValidatorKeyReplaced(string validatorAddress, bytes32 newConsensusPublicKey)
func (_Bootstrap *BootstrapFilterer) ParseValidatorKeyReplaced(log types.Log) (*BootstrapValidatorKeyReplaced, error) {
	event := new(BootstrapValidatorKeyReplaced)
	if err := _Bootstrap.contract.UnpackLog(event, "ValidatorKeyReplaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the Bootstrap contract.
type BootstrapValidatorRegisteredIterator struct {
	Event *BootstrapValidatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapValidatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapValidatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapValidatorRegistered represents a ValidatorRegistered event raised by the Bootstrap contract.
type BootstrapValidatorRegistered struct {
	EthAddress         common.Address
	ValidatorAddress   string
	Name               string
	Commission         IValidatorRegistryCommission
	ConsensusPublicKey [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0x5e8c41c1235fda86a52e67bbd0bfe0e3e3d7fa8391087d5cc8e47d607c9b5b7b.
//
// Solidity: event ValidatorRegistered(address ethAddress, string validatorAddress, string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey)
func (_Bootstrap *BootstrapFilterer) FilterValidatorRegistered(opts *bind.FilterOpts) (*BootstrapValidatorRegisteredIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "ValidatorRegistered")
	if err != nil {
		return nil, err
	}
	return &BootstrapValidatorRegisteredIterator{contract: _Bootstrap.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x5e8c41c1235fda86a52e67bbd0bfe0e3e3d7fa8391087d5cc8e47d607c9b5b7b.
//
// Solidity: event ValidatorRegistered(address ethAddress, string validatorAddress, string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey)
func (_Bootstrap *BootstrapFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *BootstrapValidatorRegistered) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "ValidatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapValidatorRegistered)
				if err := _Bootstrap.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRegistered is a log parse operation binding the contract event 0x5e8c41c1235fda86a52e67bbd0bfe0e3e3d7fa8391087d5cc8e47d607c9b5b7b.
//
// Solidity: event ValidatorRegistered(address ethAddress, string validatorAddress, string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey)
func (_Bootstrap *BootstrapFilterer) ParseValidatorRegistered(log types.Log) (*BootstrapValidatorRegistered, error) {
	event := new(BootstrapValidatorRegistered)
	if err := _Bootstrap.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapVaultCreatedIterator is returned from FilterVaultCreated and is used to iterate over the raw logs and unpacked data for VaultCreated events raised by the Bootstrap contract.
type BootstrapVaultCreatedIterator struct {
	Event *BootstrapVaultCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapVaultCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapVaultCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapVaultCreated represents a VaultCreated event raised by the Bootstrap contract.
type BootstrapVaultCreated struct {
	UnderlyingToken common.Address
	Vault           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultCreated is a free log retrieval operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address underlyingToken, address vault)
func (_Bootstrap *BootstrapFilterer) FilterVaultCreated(opts *bind.FilterOpts) (*BootstrapVaultCreatedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "VaultCreated")
	if err != nil {
		return nil, err
	}
	return &BootstrapVaultCreatedIterator{contract: _Bootstrap.contract, event: "VaultCreated", logs: logs, sub: sub}, nil
}

// WatchVaultCreated is a free log subscription operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address underlyingToken, address vault)
func (_Bootstrap *BootstrapFilterer) WatchVaultCreated(opts *bind.WatchOpts, sink chan<- *BootstrapVaultCreated) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "VaultCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapVaultCreated)
				if err := _Bootstrap.contract.UnpackLog(event, "VaultCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVaultCreated is a log parse operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address underlyingToken, address vault)
func (_Bootstrap *BootstrapFilterer) ParseVaultCreated(log types.Log) (*BootstrapVaultCreated, error) {
	event := new(BootstrapVaultCreated)
	if err := _Bootstrap.contract.UnpackLog(event, "VaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapWhitelistTokenAddedIterator is returned from FilterWhitelistTokenAdded and is used to iterate over the raw logs and unpacked data for WhitelistTokenAdded events raised by the Bootstrap contract.
type BootstrapWhitelistTokenAddedIterator struct {
	Event *BootstrapWhitelistTokenAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapWhitelistTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapWhitelistTokenAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapWhitelistTokenAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapWhitelistTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapWhitelistTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapWhitelistTokenAdded represents a WhitelistTokenAdded event raised by the Bootstrap contract.
type BootstrapWhitelistTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelistTokenAdded is a free log retrieval operation binding the contract event 0x3c96ec6d36b4038b363b8d7e8ad3b2fc7204f60317b8d0a4082094b6295fff2c.
//
// Solidity: event WhitelistTokenAdded(address _token)
func (_Bootstrap *BootstrapFilterer) FilterWhitelistTokenAdded(opts *bind.FilterOpts) (*BootstrapWhitelistTokenAddedIterator, error) {

	logs, sub, err := _Bootstrap.contract.FilterLogs(opts, "WhitelistTokenAdded")
	if err != nil {
		return nil, err
	}
	return &BootstrapWhitelistTokenAddedIterator{contract: _Bootstrap.contract, event: "WhitelistTokenAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistTokenAdded is a free log subscription operation binding the contract event 0x3c96ec6d36b4038b363b8d7e8ad3b2fc7204f60317b8d0a4082094b6295fff2c.
//
// Solidity: event WhitelistTokenAdded(address _token)
func (_Bootstrap *BootstrapFilterer) WatchWhitelistTokenAdded(opts *bind.WatchOpts, sink chan<- *BootstrapWhitelistTokenAdded) (event.Subscription, error) {

	logs, sub, err := _Bootstrap.contract.WatchLogs(opts, "WhitelistTokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapWhitelistTokenAdded)
				if err := _Bootstrap.contract.UnpackLog(event, "WhitelistTokenAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelistTokenAdded is a log parse operation binding the contract event 0x3c96ec6d36b4038b363b8d7e8ad3b2fc7204f60317b8d0a4082094b6295fff2c.
//
// Solidity: event WhitelistTokenAdded(address _token)
func (_Bootstrap *BootstrapFilterer) ParseWhitelistTokenAdded(log types.Log) (*BootstrapWhitelistTokenAdded, error) {
	event := new(BootstrapWhitelistTokenAdded)
	if err := _Bootstrap.contract.UnpackLog(event, "WhitelistTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
