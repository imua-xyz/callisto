package bootstrap

import (
	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forbole/callisto/v4/types"
	"github.com/go-co-op/gocron"
	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"
	operatorkeeper "github.com/imua-xyz/imuachain/x/operator/keeper"
	aggregatorv3 "github.com/imua-xyz/price-feeder/fetcher/chainlink/aggregatorv3"
	"github.com/rs/zerolog/log"
	"github.com/xrpscan/xrpl-go"
)

func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "bootstrap").Msg("setting up periodic tasks")

	// Schedule a cron job to run.
	if _, err := scheduler.Every(m.Config.ETHUpdateInterval).Minutes().WaitForSchedule().Do(func() {
		if err := m.refetchETHStates(); err != nil {
			log.Error().Err(err).Str("module", "bootstrap").Msg("failed to refetch ETH states")
		}
	}); err != nil {
		return fmt.Errorf("failed to set up the periodic ETH states refetch operation: %s", err)
	}

	if _, err := scheduler.Every(m.Config.BTCUpdateInterval).Minutes().Do(func() {
		if err := m.refetchBTCStates(); err != nil {
			log.Error().Err(err).Str("module", "bootstrap").Msg("failed to refetch BTC states")
		}
	}); err != nil {
		return fmt.Errorf("failed to set up the periodic BTC states refetch operation: %s", err)
	}

	if _, err := scheduler.Every(m.Config.XRPUpdateInterval).Minutes().Do(func() {
		if err := m.refetchXRPStates(); err != nil {
			log.Error().Err(err).Str("module", "bootstrap").Msg("failed to refetch XRP states")
		}
	}); err != nil {
		return fmt.Errorf("failed to set up the periodic XRP states refetch operation: %s", err)
	}

	if _, err := scheduler.Every(m.Config.PriceUpdateInterval).Minutes().WaitForSchedule().Do(func() {
		if err := m.updatePricesAndTVL(); err != nil {
			log.Error().Err(err).Str("module", "bootstrap").Msg("failed to update prices and TVL")
		}
	}); err != nil {
		return fmt.Errorf("failed to set up the periodic prices update operation: %s", err)
	}
	return nil
}

func (m *Module) updateStakerAsset(stakerAddr, assetAddr common.Address) (string, string, error) {
	totalDepositAmount, err := m.bootstrapSession.TotalDepositAmounts(stakerAddr, assetAddr)
	if err != nil {
		return "", "", err
	}

	withdrawableAmount, err := m.bootstrapSession.WithdrawableAmounts(stakerAddr, assetAddr)
	if err != nil {
		return "", "", err
	}
	if totalDepositAmount.Cmp(withdrawableAmount) < 0 {
		return "", "", fmt.Errorf("total deposit amount:%s is less than withdrawable amount:%s", totalDepositAmount, withdrawableAmount)
	}
	delegationAmount := big.NewInt(0).Sub(totalDepositAmount, withdrawableAmount)
	stakerID, assetID := assetstypes.GetStakerIDAndAssetID(m.Config.ETHLZChainID, stakerAddr[:], assetAddr[:])

	stakerAssetExist, err := m.database.BootstrapStakerAssetExists(stakerID, assetID)
	if err != nil {
		return "", "", err
	}
	if !stakerAssetExist && totalDepositAmount.Cmp(big.NewInt(0)) == 0 {
		// In the refetch case, do nothing since the staker has no deposit for this asset.
		return "", "", nil
	}
	err = m.database.SaveBootstrapStakerAsset(&types.BootstrapStakerAsset{
		StakerID:     stakerID,
		AssetID:      assetID,
		Deposited:    totalDepositAmount.String(),
		Withdrawable: withdrawableAmount.String(),
		Delegated:    delegationAmount.String(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return "", "", err
	}
	return stakerID, assetID, nil
}

func (m *Module) refetchETHStates() error {
	log.Debug().Str("module", "bootstrap").Str("refetching", "ETH states").
		Msg("refetching ETH states")

	// refetch all validators
	validatorCount, err := m.bootstrapSession.GetValidatorsCount()
	if err != nil {
		return err
	}
	validatorIMAddresses := make([]string, validatorCount.Int64())
	validatorETHAddresses := make([]common.Address, validatorCount.Int64())
	for i := int64(0); i < validatorCount.Int64(); i++ {
		validatorEthAddr, err := m.bootstrapSession.RegisteredValidators(big.NewInt(i))
		if err != nil {
			return fmt.Errorf("failed to call RegisteredValidators,index:%d,err:%s", i, err)
		}
		validatorIMAddr, err := m.bootstrapSession.EthToImAddress(validatorEthAddr)
		if err != nil {
			return fmt.Errorf("failed to call EthToImAddress,validatorEthAddr:%s,err:%s", validatorEthAddr, err)
		}
		validatorInfo, err := m.bootstrapSession.Validators(validatorIMAddr)
		if err != nil {
			return fmt.Errorf("failed to call Validators,validatorIMAddr:%s,err:%s", validatorIMAddr, err)
		}
		err = m.database.SaveBootstrapValidator(&types.BootstrapValidator{
			ValidatorEthAddress: validatorEthAddr.String(),
			ValidatorIMAddress:  validatorIMAddr,
			ValidatorName:       validatorInfo.Name,
			ConsensusPubKey:     hexutil.Encode(validatorInfo.ConsensusPublicKey[:]),
			Rate:                validatorInfo.Commission.Rate.String(),
			MaxRate:             validatorInfo.Commission.MaxRate.String(),
			MaxChangeRate:       validatorInfo.Commission.MaxChangeRate.String(),
			UpdatedAt:           time.Now(),
		})
		if err != nil {
			return err
		}
		validatorIMAddresses[i] = validatorIMAddr
		validatorETHAddresses[i] = validatorEthAddr
	}

	// refetch all staking assets
	tokenCount, err := m.bootstrapSession.GetWhitelistedTokensCount()
	if err != nil {
		return err
	}
	stakingAssets := make([]common.Address, tokenCount.Int64())
	stakingAssetIDs := make([]string, tokenCount.Int64())
	// get the token info by iterating all indexes
	for i := int64(0); i < tokenCount.Int64(); i++ {
		tokenInfo, err := m.bootstrapSession.GetWhitelistedTokenAtIndex(big.NewInt(i))
		if err != nil {
			return fmt.Errorf("failed to call GetWhitelistedTokenAtIndex,index:%d,err:%s", i, err)
		}

		_, assetID := assetstypes.GetStakerIDAndAssetID(m.Config.ETHLZChainID, nil, tokenInfo.TokenAddress[:])
		assetDepositAmount, err := m.bootstrapSession.DepositsByToken(tokenInfo.TokenAddress)
		if err != nil {
			return fmt.Errorf("failed to call DepositsByToken,tokenAddr:%s,err:%s", tokenInfo.TokenAddress, err)
		}

		usdValue := sdkmath.LegacyZeroDec()
		priceStr, err := m.database.GetBootstrapTokenPrice(assetID)
		if err != nil {
			log.Err(err).Str("assetID", assetID).Msg("refetchETHStates: get token price from database")
			// Using zero as the USD value; continue handling other assets without returning
		} else {
			// calculate the total USD value of this asset
			priceDec, err := sdkmath.LegacyNewDecFromStr(priceStr)
			if err != nil {
				log.Err(err).Str("dbPrice", priceStr).Msg("failed to parse the db price to a big legacyDec")
				// don't return to continue addressing the other assets
				break
			}
			divisor := sdkmath.NewIntWithDecimal(1, int(tokenInfo.Decimals)) // #nosec G115
			usdValue = priceDec.MulInt(sdkmath.NewIntFromBigInt(assetDepositAmount)).QuoInt(divisor)
		}

		err = m.database.SaveBootstrapToken(&types.BootstrapTokenState{
			BootstrapToken: types.BootstrapToken{
				AssetID:   assetID,
				Address:   tokenInfo.TokenAddress.String(),
				Name:      tokenInfo.Name,
				Symbol:    tokenInfo.Symbol,
				Decimals:  tokenInfo.Decimals,
				LZChainID: m.Config.ETHLZChainID,
			},
			UpdatedAt:          time.Now(),
			StakingTotalAmount: assetDepositAmount.String(),
			TotalUSDValue:      usdValue.String(),
		})
		if err != nil {
			return err
		}
		stakingAssets[i] = tokenInfo.TokenAddress
		stakingAssetIDs[i] = assetID
	}

	// refetch all staker assets
	depositorCount, err := m.bootstrapSession.GetDepositorsCount()
	if err != nil {
		return fmt.Errorf("failed to call GetDepositorsCount,err:%s", err)
	}
	for i := int64(0); i < depositorCount.Int64(); i++ {
		depositer, err := m.bootstrapSession.Depositors(big.NewInt(i))
		if err != nil {
			return fmt.Errorf("failed to call Depositors,index:%d,err:%s", i, err)
		}
		for _, assetAddr := range stakingAssets {
			stakerID, assetID, err := m.updateStakerAsset(depositer, assetAddr)
			if err != nil {
				return fmt.Errorf("failed to call Depositors,index:%d,err:%s", i, err)
			}
			if stakerID != "" && assetID != "" {
				for _, validator := range validatorIMAddresses {
					delegationAmount, err := m.bootstrapSession.Delegations(depositer, validator, assetAddr)
					if err != nil {
						return err
					}
					delegationExist, err := m.database.BootstrapDelegationExists(stakerID, assetID, validator)
					if err != nil {
						return err
					}
					if !delegationExist && delegationAmount.Cmp(big.NewInt(0)) == 0 {
						// Skip the validator if the delegation hasn't been saved in the database and the fetched amount
						// is zero. This avoids saving delegations that don't exist. Since there is no flag indicating
						// whether a delegation exists, the bootstrap map will always return zero and nil error when the
						// delegation does not exist.
						// With this approach, it is possible that a delegation which was delegated in the bootstrap
						// but later fully undelegated—and whose actions were not captured by events—will not be
						// saved to the database. This is acceptable, because such a staker effectively has no delegated assets.
						// The database only records the current latest state and does not store operation history.
						continue
					}
					// update the delegation states
					err = m.database.SaveBootstrapDelegationState(&types.BootstrapDelegationState{
						StakerID:     stakerID,
						AssetID:      assetID,
						OperatorAddr: validator,
						Delegated:    delegationAmount.String(),
						UpdatedAt:    time.Now(),
					})
					if err != nil {
						return err
					}
				}
			}
		}
	}

	// fetch all validator assets
	for i, validator := range validatorIMAddresses {
		for j, assetAddr := range stakingAssets {
			validatorAssetAmount, err := m.bootstrapSession.DelegationsByValidator(validator, assetAddr)
			if err != nil {
				return err
			}
			operatorAssetExist, err := m.database.OperatorAssetExists(validator, stakingAssetIDs[j])
			if err != nil {
				return err
			}
			if !operatorAssetExist && validatorAssetAmount.Cmp(big.NewInt(0)) == 0 {
				continue
			}
			// get the self delegation amount
			selfDelegation, err := m.bootstrapSession.Delegations(validatorETHAddresses[i], validator, assetAddr)
			if err != nil {
				return err
			}
			err = m.database.SaveBootstrapOperatorAsset(&types.BootstrapOperatorAsset{
				OperatorAddr: validator,
				AssetID:      stakingAssetIDs[j],
				TotalAmount:  validatorAssetAmount.String(),
				SelfAmount:   selfDelegation.String(),
				OtherAmount:  big.NewInt(0).Sub(validatorAssetAmount, selfDelegation).String(),
				UpdatedAt:    time.Now(),
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Helper functions for safe data parsing and validation
// safeStringExtract safely extracts string value from map
func safeStringExtract(data map[string]interface{}, key string) (string, error) {
	value, exists := data[key]
	if !exists {
		return "", fmt.Errorf("missing key: %s", key)
	}
	str, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("invalid type for key %s: expected string, got %T", key, value)
	}
	return str, nil
}

// validateBTCAddressBinding validates and resolves BTC address binding with 1-1 mapping rules
// Returns the imuachain address to use (may be different from input if conflict resolved)
// Returns empty string if binding is invalid and transaction should be rejected
func (m *Module) validateBTCAddressBinding(senderAddr, parsedImuachainAddr, txid string) (string, error) {
	if senderAddr == "" || parsedImuachainAddr == "" {
		return "", fmt.Errorf("empty address in binding validation")
	}

	// Normalize addresses for consistent comparison
	senderAddr = strings.ToLower(strings.TrimSpace(senderAddr))
	parsedImuachainAddr = strings.ToLower(strings.TrimSpace(parsedImuachainAddr))

	// First, check if sender is already bound to an imuachain address in memory
	if existingImuachainAddr, exists := m.btcAddressMappings[senderAddr]; exists {
		if existingImuachainAddr != parsedImuachainAddr {
			log.Warn().Str("txid", txid).
				Str("bitcoin_addr", senderAddr).
				Str("existing_binding", existingImuachainAddr).
				Str("parsed_binding", parsedImuachainAddr).
				Msg("BTC sender already bound to different imuachain address, using existing binding")
			return existingImuachainAddr, nil
		}
		// Address already bound correctly
		return existingImuachainAddr, nil
	}

	// Check if parsed imuachain address is already bound to another sender in memory
	for existingSender, existingImuachain := range m.btcAddressMappings {
		if existingImuachain == parsedImuachainAddr && existingSender != senderAddr {
			log.Warn().Str("txid", txid).
				Str("parsed_imuachain_addr", parsedImuachainAddr).
				Str("existing_sender", existingSender).
				Str("new_sender", senderAddr).
				Msg("BTC imuachain address already bound to different sender, rejecting transaction")
			return "", fmt.Errorf("imuachain address %s already bound to BTC address %s", parsedImuachainAddr, existingSender)
		}
	}

	// Double-check with database for consistency (in case memory was cleared)
	// Check if sender is already bound in database
	existingSenderBinding, err := m.database.GetAddressBinding("BTC", senderAddr)
	if err != nil {
		return "", fmt.Errorf("error checking sender address binding in database: %w", err)
	}
	if existingSenderBinding != nil {
		if existingSenderBinding.TargetAddr != parsedImuachainAddr {
			log.Warn().Str("txid", txid).
				Str("bitcoin_addr", senderAddr).
				Str("existing_binding", existingSenderBinding.TargetAddr).
				Str("parsed_binding", parsedImuachainAddr).
				Time("existing_created_at", existingSenderBinding.CreatedAt).
				Msg("BTC sender already bound to different imuachain address in database, using existing binding")
			// Update memory with the correct binding
			m.btcAddressMappings[senderAddr] = existingSenderBinding.TargetAddr
			return existingSenderBinding.TargetAddr, nil
		}
		// Already bound correctly, update memory
		m.btcAddressMappings[senderAddr] = existingSenderBinding.TargetAddr
		return existingSenderBinding.TargetAddr, nil
	}

	// Check if parsed imuachain address is already bound to another sender in database
	existingTargetBinding, err := m.database.CheckTargetAddressBinding("BTC", parsedImuachainAddr, senderAddr)
	if err != nil {
		return "", fmt.Errorf("error checking target address binding in database: %w", err)
	}
	if existingTargetBinding != nil {
		log.Warn().Str("txid", txid).
			Str("parsed_imuachain_addr", parsedImuachainAddr).
			Str("existing_sender", existingTargetBinding.SourceAddr).
			Str("new_sender", senderAddr).
			Time("existing_created_at", existingTargetBinding.CreatedAt).
			Msg("BTC imuachain address already bound to different sender in database, rejecting transaction")
		return "", fmt.Errorf("imuachain address %s already bound to BTC address %s", parsedImuachainAddr, existingTargetBinding.SourceAddr)
	}

	// No conflicts found, establish new binding
	m.btcAddressMappings[senderAddr] = parsedImuachainAddr

	// Save to database
	binding := &types.AddressBinding{
		ChainType:  "BTC",
		SourceAddr: senderAddr,
		TargetAddr: parsedImuachainAddr,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := m.database.SaveAddressBinding(binding); err != nil {
		// Remove from memory if database save failed
		delete(m.btcAddressMappings, senderAddr)
		return "", fmt.Errorf("failed to save address binding: %w", err)
	}

	log.Info().Str("txid", txid).
		Str("bitcoin_addr", senderAddr).
		Str("imuachain_addr", parsedImuachainAddr).
		Msg("established new BTC address binding")
	return parsedImuachainAddr, nil
}

// validateXRPAddressBinding validates and resolves XRP address binding with 1-1 mapping rules
// Returns the imuachain address to use (may be different from input if conflict resolved)
// Returns empty string if binding is invalid and transaction should be rejected
func (m *Module) validateXRPAddressBinding(senderAddr, parsedImuachainAddr, txHash string) (string, error) {
	if senderAddr == "" || parsedImuachainAddr == "" {
		return "", fmt.Errorf("empty address in binding validation")
	}

	// Normalize addresses for consistent comparison
	senderAddr = strings.ToLower(strings.TrimSpace(senderAddr))
	parsedImuachainAddr = strings.ToLower(strings.TrimSpace(parsedImuachainAddr))

	// First, check if sender is already bound to an imuachain address in memory
	if existingImuachainAddr, exists := m.xrpAddressMappings[senderAddr]; exists {
		if existingImuachainAddr != parsedImuachainAddr {
			log.Warn().Str("hash", txHash).
				Str("xrp_addr", senderAddr).
				Str("existing_binding", existingImuachainAddr).
				Str("parsed_binding", parsedImuachainAddr).
				Msg("XRP sender already bound to different imuachain address, using existing binding")
			return existingImuachainAddr, nil
		}
		// Address already bound correctly
		return existingImuachainAddr, nil
	}

	// Check if parsed imuachain address is already bound to another sender in memory
	for existingSender, existingImuachain := range m.xrpAddressMappings {
		if existingImuachain == parsedImuachainAddr && existingSender != senderAddr {
			log.Warn().Str("hash", txHash).
				Str("parsed_imuachain_addr", parsedImuachainAddr).
				Str("existing_sender", existingSender).
				Str("new_sender", senderAddr).
				Msg("XRP imuachain address already bound to different sender, rejecting transaction")
			return "", fmt.Errorf("imuachain address %s already bound to XRP address %s", parsedImuachainAddr, existingSender)
		}
	}

	// Double-check with database for consistency (in case memory was cleared)
	// Check if sender is already bound in database
	existingSenderBinding, err := m.database.GetAddressBinding("XRP", senderAddr)
	if err != nil {
		return "", fmt.Errorf("error checking sender address binding in database: %w", err)
	}
	if existingSenderBinding != nil {
		if existingSenderBinding.TargetAddr != parsedImuachainAddr {
			log.Warn().Str("hash", txHash).
				Str("xrp_addr", senderAddr).
				Str("existing_binding", existingSenderBinding.TargetAddr).
				Str("parsed_binding", parsedImuachainAddr).
				Time("existing_created_at", existingSenderBinding.CreatedAt).
				Msg("XRP sender already bound to different imuachain address in database, using existing binding")
			// Update memory with the correct binding
			m.xrpAddressMappings[senderAddr] = existingSenderBinding.TargetAddr
			return existingSenderBinding.TargetAddr, nil
		}
		// Already bound correctly, update memory
		m.xrpAddressMappings[senderAddr] = existingSenderBinding.TargetAddr
		return existingSenderBinding.TargetAddr, nil
	}

	// Check if parsed imuachain address is already bound to another sender in database
	existingTargetBinding, err := m.database.CheckTargetAddressBinding("XRP", parsedImuachainAddr, senderAddr)
	if err != nil {
		return "", fmt.Errorf("error checking target address binding in database: %w", err)
	}
	if existingTargetBinding != nil {
		log.Warn().Str("hash", txHash).
			Str("parsed_imuachain_addr", parsedImuachainAddr).
			Str("existing_sender", existingTargetBinding.SourceAddr).
			Str("new_sender", senderAddr).
			Time("existing_created_at", existingTargetBinding.CreatedAt).
			Msg("XRP imuachain address already bound to different sender in database, rejecting transaction")
		return "", fmt.Errorf("imuachain address %s already bound to XRP address %s", parsedImuachainAddr, existingTargetBinding.SourceAddr)
	}

	// No conflicts found, establish new binding
	m.xrpAddressMappings[senderAddr] = parsedImuachainAddr

	// Save to database
	binding := &types.AddressBinding{
		ChainType:  "XRP",
		SourceAddr: senderAddr,
		TargetAddr: parsedImuachainAddr,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := m.database.SaveAddressBinding(binding); err != nil {
		// Remove from memory if database save failed
		delete(m.xrpAddressMappings, senderAddr)
		return "", fmt.Errorf("failed to save address binding: %w", err)
	}

	log.Info().Str("hash", txHash).
		Str("xrp_addr", senderAddr).
		Str("imuachain_addr", parsedImuachainAddr).
		Msg("established new XRP address binding")
	return parsedImuachainAddr, nil
}

// processTransactionWithRetry executes a transaction function with retry logic for database errors
func (m *Module) processTransactionWithRetry(chainType string, txFunc func() error) error {
	const maxRetries = 3

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		err := txFunc()
		if err == nil {
			return nil // Success
		}

		lastErr = err

		// Check if error is retryable
		if !isRetryableDatabaseError(err) {
			// Non-retryable error, fail immediately
			return err
		}

		if attempt < maxRetries-1 {
			// Exponential backoff: 1s, 2s, 4s...
			backoff := time.Duration(1<<uint(attempt)) * time.Second
			log.Warn().
				Err(err).
				Str("chain", chainType).
				Int("attempt", attempt+1).
				Dur("backoff", backoff).
				Msg("database transaction failed, retrying")

			time.Sleep(backoff)
		}
	}

	return fmt.Errorf("transaction failed after %d retries: %w", maxRetries, lastErr)
}

// isRetryableDatabaseError checks if a database error should be retried
func isRetryableDatabaseError(err error) bool {
	if err == nil {
		return false
	}

	errStr := strings.ToLower(err.Error())

	// Deadlock errors
	if strings.Contains(errStr, "deadlock") ||
		strings.Contains(errStr, "lock wait timeout") {
		return true
	}

	// Connection errors
	if strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "connection reset") ||
		strings.Contains(errStr, "connection lost") ||
		strings.Contains(errStr, "broken pipe") {
		return true
	}

	// Serialization failures (PostgreSQL)
	if strings.Contains(errStr, "serialization failure") ||
		strings.Contains(errStr, "could not serialize") {
		return true
	}

	// Lock acquisition timeouts
	if strings.Contains(errStr, "lock acquisition") ||
		strings.Contains(errStr, "timeout") {
		return true
	}

	return false
}

// sortBTCTransactions sorts BTC transactions by block height and transaction index
func sortBTCTransactions(txs []types.BTCTx) {
	sort.Slice(txs, func(i, j int) bool {
		if txs[i].Status.BlockHeight != txs[j].Status.BlockHeight {
			return txs[i].Status.BlockHeight < txs[j].Status.BlockHeight
		}
		// Use TxID for deterministic order if same block
		return txs[i].TxID < txs[j].TxID
	})
}

// sortXRPTransactions sorts XRP transactions by ledger index and transaction index
func sortXRPTransactions(txs []types.XRPTransaction) {
	sort.Slice(txs, func(i, j int) bool {
		if txs[i].LedgerIndex != txs[j].LedgerIndex {
			return txs[i].LedgerIndex < txs[j].LedgerIndex
		}
		return txs[i].Meta.TransactionIndex < txs[j].Meta.TransactionIndex
	})
}

// BTC vault transaction fetching and processing
func (m *Module) refetchBTCStates() error {
	// Get current block height for confirmation calculations
	currentHeight, err := m.getBTCCurrentBlockHeight()
	if err != nil {
		return fmt.Errorf("error getting current BTC block height: %w", err)
	}

	log.Info().Int64("current_height", currentHeight).
		Int("min_confirmations", m.Config.BTCMinConfirmations).
		Msg("starting BTC deposit transaction processing")

	// Get all transactions from vault address (matches TypeScript getNewerConfirmedTxs logic)
	transactions, err := m.getConfirmedVaultTransactions()
	if err != nil {
		return fmt.Errorf("error getting BTC vault transactions: %w", err)
	}

	if len(transactions) == 0 {
		log.Info().Msg("no transactions found for vault address")
		return nil
	}

	log.Info().Int("total_transactions", len(transactions)).
		Msg("fetched transactions from vault address")

	// Filter transactions with sufficient confirmations (matches TypeScript safelyFinalizedTxs logic)
	safelyFinalizedTxs := m.filterSafelyFinalizedTransactions(transactions, currentHeight)
	if len(safelyFinalizedTxs) == 0 {
		log.Info().Msg("no transactions with enough confirmations found")
		return nil
	}

	// Sort by block height and tx index (matches TypeScript sorting logic)
	sortBTCTransactions(safelyFinalizedTxs)

	log.Info().Int("safely_finalized_and_validated_txs", len(safelyFinalizedTxs)).
		Msg("filtered and validated transactions with sufficient confirmations")

	processedCount := 0

	// Process transactions sequentially (matches TypeScript processing loop)
	// All transactions are already validated and have parsed OP_RETURN data
	for _, tx := range safelyFinalizedTxs {
		// Process transaction (matches TypeScript saveCrossChainTx logic)
		// Address data is already parsed and available in tx struct
		if err := m.processBTCTxWithTransaction(tx); err != nil {
			log.Err(err).Str("txid", tx.TxID).Msg("error processing BTC transaction")
			continue
		}

		processedCount++
	}

	// Log processing statistics (matches TypeScript logging)
	log.Info().
		Int("total_transactions", len(transactions)).
		Int("safely_finalized_and_validated", len(safelyFinalizedTxs)).
		Int("processed", processedCount).
		Int64("block_height_tip", currentHeight).
		Msg("BTC deposit transaction processing completed")

	return nil
}

// filterSafelyFinalizedTransactions filters transactions with sufficient confirmations and validates them
// Matches TypeScript safelyFinalizedTxs filtering logic with added validation
func (m *Module) filterSafelyFinalizedTransactions(transactions []types.BTCTx, currentHeight int64) []types.BTCTx {
	var safelyFinalized []types.BTCTx

	for _, tx := range transactions {
		// Check if transaction has sufficient confirmations
		if tx.Status.BlockHeight <= currentHeight &&
			currentHeight-tx.Status.BlockHeight+1 >= int64(m.Config.BTCMinConfirmations) {

			// Validate transaction and parse OP_RETURN data in one step (matches TypeScript logic)
			validTx := tx // Create a copy for modification
			err := m.isValidDepositTransaction(&validTx)
			if err != nil {
				log.Debug().Err(err).Str("txid", tx.TxID).Msg("BTC transaction validation failed, skipping")
				continue
			}

			// Transaction is now validated and parsed, add to result
			safelyFinalized = append(safelyFinalized, validTx)
		}
	}

	return safelyFinalized
}

// normalizeAddress normalizes an address by trimming spaces and converting to lowercase
func normalizeAddress(addr string) string {
	return strings.ToLower(strings.TrimSpace(addr))
}

// isValidDepositTransaction validates transaction format and parses OP_RETURN data
// Note: Confirmation checks are handled by filterSafelyFinalizedTransactions
// Returns error if validation fails, nil if transaction is valid for bootstrap
// Parsed data is stored directly in the tx struct
func (m *Module) isValidDepositTransaction(tx *types.BTCTx) error {
	vaultAddr := normalizeAddress(m.Config.BTCVaultAddr)

	// Transaction is already confirmed and finalized by filterSafelyFinalizedTransactions
	// Check if it's from vault (should not be) - matches TypeScript isFromVault check
	for _, vin := range tx.Vin {
		if normalizeAddress(vin.Prevout.ScriptPubKeyAddr) == vaultAddr {
			return fmt.Errorf("BTC transaction %s is from vault (invalid for bootstrap)", tx.TxID)
		}
	}

	// Check if there's exactly one output to vault with minimum amount - matches TypeScript vaultOutputs check
	vaultOutputCount := 0
	for _, vout := range tx.Vout {
		if normalizeAddress(vout.ScriptPubKeyAddr) == vaultAddr &&
			vout.Value >= m.Config.BTCMinAmount {
			vaultOutputCount++
		}
	}
	if vaultOutputCount == 0 {
		return fmt.Errorf("BTC transaction %s has no valid vault output (minimum %d satoshi)", tx.TxID, m.Config.BTCMinAmount)
	}
	if vaultOutputCount > 1 {
		return fmt.Errorf("BTC transaction %s has multiple vault outputs (%d), expected exactly 1", tx.TxID, vaultOutputCount)
	}

	// Find OP_RETURN output - matches TypeScript opReturnOutputs check
	var opReturnOutput *types.BTCVout
	opReturnCount := 0
	for i, vout := range tx.Vout {
		if vout.ScriptPubKeyType == "op_return" {
			opReturnCount++
			if opReturnCount > 1 {
				return fmt.Errorf("BTC transaction %s has multiple OP_RETURN outputs (%d), expected exactly 1", tx.TxID, opReturnCount)
			}
			// Use index instead of pointer to avoid address reuse issues
			opReturnOutput = &tx.Vout[i]
		}
	}
	if opReturnCount == 0 {
		return fmt.Errorf("BTC transaction %s missing OP_RETURN output", tx.TxID)
	}

	// Parse OP_RETURN data (matches TypeScript parseOpReturnDataInline)
	opReturnData, err := m.parseOpReturnData(opReturnOutput.ScriptPubKey)
	if err != nil {
		return fmt.Errorf("BTC transaction %s OP_RETURN parsing failed: %w", tx.TxID, err)
	}

	// Validator address should always be present with new parsing logic
	if opReturnData.ValidatorAddress == "" {
		return fmt.Errorf("BTC transaction %s missing validator address in OP_RETURN", tx.TxID)
	}

	// Validate validator is registered for bootstrap
	// TODO: Temporarily commented out for testing - no available bootstrap contract with valid validators
	/*
		isRegistered, err := m.isValidatorRegistered(tx.ValidatorAddress)
		if err != nil {
			return fmt.Errorf("BTC transaction %s validator registration check failed: %w", tx.TxID, err)
		}

		if !isRegistered {
			return fmt.Errorf("BTC transaction %s validator %s is not registered for bootstrap", tx.TxID, tx.ValidatorAddress)
		}
	*/

	// Find the sender address from transaction inputs
	senderAddr := ""
	for _, vin := range tx.Vin {
		if vin.Prevout.ScriptPubKeyAddr != "" {
			senderAddr = vin.Prevout.ScriptPubKeyAddr
			break
		}
	}
	if senderAddr == "" {
		return fmt.Errorf("BTC transaction %s has no sender address found", tx.TxID)
	}

	// Validate address binding and get the correct imuachain address to use
	correctImuachainAddr, err := m.validateBTCAddressBinding(senderAddr, opReturnData.ImuachainAddressHex, tx.TxID)
	if err != nil {
		return fmt.Errorf("BTC transaction %s address binding validation failed: %w", tx.TxID, err)
	}

	// Set parsed data directly in the transaction struct
	tx.ImuachainAddress = correctImuachainAddr
	tx.ValidatorAddress = opReturnData.ValidatorAddress

	return nil
}

// isValidValidatorAddress validates if a string is a valid validator address (bech32 format with 'im' prefix)
func (m *Module) isValidValidatorAddress(address string) bool {
	// Basic format check: should start with 'im1' and have appropriate length
	if !strings.HasPrefix(address, "im1") || len(address) < 10 {
		return false
	}

	// Check if it contains only alphanumeric characters (basic bech32 validation)
	for _, char := range address {
		if !((char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')) {
			return false
		}
	}

	return true
}

// OpReturnData represents parsed OP_RETURN data (internal struct for parsing)
type OpReturnData struct {
	ImuachainAddressHex string
	ValidatorAddress    string
}

// parseOpReturnData parses and validates OP_RETURN data from Bitcoin transaction output
// Required format: 6a3d{20 bytes imuachain}{41 bytes validator} (imua + validator addresses)
// All bootstrap transactions must include validator information
func (m *Module) parseOpReturnData(scriptPubKey string) (*OpReturnData, error) {
	// Check if it starts with OP_RETURN prefix
	if !strings.HasPrefix(scriptPubKey, "6a") {
		return nil, fmt.Errorf("invalid OP_RETURN prefix: expected '6a', got %s", scriptPubKey[:2])
	}

	// Extract length byte and data
	if len(scriptPubKey) < 6 {
		return nil, fmt.Errorf("OP_RETURN script too short: %d bytes", len(scriptPubKey))
	}

	lengthHex := scriptPubKey[2:4]
	length, err := strconv.ParseInt(lengthHex, 16, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse OP_RETURN length: %s", err)
	}

	hexOpReturnData := scriptPubKey[4:]

	// Validate data length matches declared length
	if len(hexOpReturnData) != int(length)*2 {
		return nil, fmt.Errorf("OP_RETURN data length mismatch: declared %d bytes, got %d hex chars", length, len(hexOpReturnData))
	}

	// Only support extended format with validator information
	if length != 61 {
		return nil, fmt.Errorf("unsupported OP_RETURN data length: expected 61 bytes (IMUA + validator), got %d bytes", length)
	}

	// Extended format: IMUA address (20 bytes) + validator address (41 bytes)
	imuachainAddressHex := strings.ToLower("0x" + hexOpReturnData[:40])

	// Validate IMUA address format
	if !isValidEthereumAddress(imuachainAddressHex) {
		return nil, fmt.Errorf("invalid IMUA address format: %s", imuachainAddressHex)
	}

	validatorAddressHex := hexOpReturnData[40:]

	// Decode validator address from hex
	validatorBytes, err := hex.DecodeString(validatorAddressHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode validator address from hex: %s", err)
	}

	// Check if the bytes contain only printable ASCII characters (bech32 addresses should be ASCII)
	for _, b := range validatorBytes {
		if b < 32 || b > 126 {
			return nil, fmt.Errorf("validator address contains non-printable character: %d", b)
		}
	}

	validatorAddress := string(validatorBytes)

	// Additional format validation - bech32 addresses should only contain alphanumeric characters
	for _, char := range validatorAddress {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return nil, fmt.Errorf("validator address contains invalid character: %c", char)
		}
	}

	// Validate the validator address format (should be bech32 with 'im1' prefix and correct length)
	if len(validatorAddress) != 41 || !strings.HasPrefix(validatorAddress, "im1") {
		return nil, fmt.Errorf("invalid validator address format: expected 41 chars starting with 'im1', got %d chars starting with %s", len(validatorAddress), validatorAddress[:3])
	}

	return &OpReturnData{
		ImuachainAddressHex: imuachainAddressHex,
		ValidatorAddress:    validatorAddress,
	}, nil
}

// isValidEthereumAddress validates if a string is a valid Ethereum address format
func isValidEthereumAddress(address string) bool {
	if len(address) != 42 {
		return false
	}
	if !strings.HasPrefix(address, "0x") {
		return false
	}
	for _, char := range address[2:] {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')) {
			return false
		}
	}
	return true
}

// getBTCCurrentBlockHeight gets the current BTC block height from Esplora API
func (m *Module) getBTCCurrentBlockHeight() (int64, error) {
	url := fmt.Sprintf("%s/api/blocks/tip/height", m.Config.BTCRPC)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create request with context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch block height: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var height int64
	if err := json.NewDecoder(resp.Body).Decode(&height); err != nil {
		return 0, fmt.Errorf("failed to decode response: %s", err)
	}

	if height <= 0 {
		return 0, fmt.Errorf("invalid block height: %d", height)
	}

	return height, nil
}

// getConfirmedVaultTransactions replicates the TypeScript getConfirmedTransactions logic
// Gets the last processed transaction ID from database for pagination
func (m *Module) getConfirmedVaultTransactions() ([]types.BTCTx, error) {
	var allTxs []types.BTCTx

	startTxID, startHeight, err := m.database.GetLastProcessedTransaction("BTC")
	if err != nil {
		return nil, fmt.Errorf("failed to get last processed transaction: %w", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	ctx := context.Background()

	var startIndex int64
	if startTxID != "" && startHeight > 0 {
		idx, err := m.getBTCTransactionIndex(ctx, client, startTxID)
		if err != nil {
			log.Warn().Str("txid", startTxID).Err(err).
				Msg("failed to fetch last processed tx index; scanning full history")
			startTxID = ""
			startHeight = 0
		} else {
			startIndex = idx
		}
	}

	vaultAddress := normalizeAddress(m.Config.BTCVaultAddr)
	// Start with empty cursor to fetch from newest transactions first
	// cursor will be updated to paginate towards older transactions
	cursor := ""
	pageCount := 0
	stop := false

	log.Info().Str("vault_address", vaultAddress).
		Str("target_txid", startTxID).
		Int64("target_height", startHeight).
		Msg("starting BTC vault transaction fetch from newest")

	for !stop {
		txs, err := m.fetchBTCTxsPage(ctx, client, vaultAddress, cursor)
		if err != nil {
			return nil, err
		}
		if len(txs) == 0 {
			break
		}

		pageCount++

		newTxs, reachedOlder := m.filterNewBTCTxs(ctx, client, txs, startHeight, startIndex)
		allTxs = append(allTxs, newTxs...)

		if reachedOlder {
			break
		}

		cursor = txs[len(txs)-1].TxID
	}

	sortBTCTransactions(allTxs)

	log.Info().Int("total_confirmed_txs", len(allTxs)).
		Int("pages_fetched", pageCount).
		Msg("completed BTC deposit processing")

	return allTxs, nil
}

func (m *Module) fetchBTCTxsPage(ctx context.Context, client *http.Client, vaultAddress, cursor string) ([]types.BTCTx, error) {
	var url string
	if cursor != "" {
		url = fmt.Sprintf("%s/api/address/%s/txs/chain/%s", m.Config.BTCRPC, vaultAddress, cursor)
	} else {
		url = fmt.Sprintf("%s/api/address/%s/txs", m.Config.BTCRPC, vaultAddress)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, url, nil)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create BTC tx request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to fetch BTC transactions: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		cancel()
		return nil, fmt.Errorf("BTC tx API returned status %d", resp.StatusCode)
	}

	var txs []types.BTCTx
	if err := json.NewDecoder(resp.Body).Decode(&txs); err != nil {
		resp.Body.Close()
		cancel()
		return nil, fmt.Errorf("failed to decode BTC transactions: %w", err)
	}
	resp.Body.Close()
	cancel()

	return txs, nil
}

func (m *Module) filterNewBTCTxs(ctx context.Context, client *http.Client, txs []types.BTCTx, startHeight, startIndex int64) ([]types.BTCTx, bool) {
	filtered := make([]types.BTCTx, 0, len(txs))

	// If no target (startHeight == 0), collect all confirmed transactions
	if startHeight == 0 {
		for _, tx := range txs {
			// Skip early if not confirmed. Unconfirmed txs do not have block height.
			if !tx.Status.Confirmed {
				continue
			}

			// Ensure block height is present for confirmed txs only
			if tx.Status.BlockHeight == 0 {
				if err := m.ensureBTCTxBlockHeight(ctx, client, &tx); err != nil {
					// Missing block height on confirmed tx is unexpected but non-fatal; skip quietly
					log.Debug().Err(err).Str("txid", tx.TxID).Msg("missing BTC tx block height for confirmed tx, skipping")
					continue
				}
			}

			if _, err := m.populateBTCTxIndex(ctx, client, &tx); err != nil {
				log.Err(err).Str("txid", tx.TxID).Msg("failed to fetch BTC tx index, skipping")
				continue
			}

			filtered = append(filtered, tx)
		}
		return filtered, false
	}

	// Filter transactions newer than the target (startHeight, startIndex)
	// We're paginating from newest to oldest, so we collect until we reach the target
	for _, tx := range txs {
		// Skip unconfirmed transactions before any extra calls
		if !tx.Status.Confirmed {
			continue
		}

		if tx.Status.BlockHeight == 0 {
			if err := m.ensureBTCTxBlockHeight(ctx, client, &tx); err != nil {
				// Height missing: treat as non-fatal and skip quietly
				log.Debug().Err(err).Str("txid", tx.TxID).Msg("missing BTC tx block height for confirmed tx, skipping")
				continue
			}
		}

		// If we reached a block older than target, stop pagination
		if tx.Status.BlockHeight < startHeight {
			return filtered, true
		}

		// Get transaction index for proper comparison
		txIndex, err := m.populateBTCTxIndex(ctx, client, &tx)
		if err != nil {
			log.Err(err).Str("txid", tx.TxID).Msg("failed to fetch BTC tx index, skipping")
			continue
		}

		// Include transactions from newer blocks
		if tx.Status.BlockHeight > startHeight {
			filtered = append(filtered, tx)
		} else if tx.Status.BlockHeight == startHeight {
			// For same block, only include if transaction index is higher (newer)
			if txIndex > startIndex {
				filtered = append(filtered, tx)
			} else if txIndex == startIndex {
				// Reached exactly the target transaction, stop here
				return filtered, true
			}
		}
	}

	return filtered, false
}

func (m *Module) getBTCTxDetails(ctx context.Context, client *http.Client, txid string) (*types.BTCTx, error) {
	reqCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s/api/tx/%s", m.Config.BTCRPC, txid)
	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create tx detail request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tx detail: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tx detail API returned status %d", resp.StatusCode)
	}

	var tx types.BTCTx
	if err := json.NewDecoder(resp.Body).Decode(&tx); err != nil {
		return nil, fmt.Errorf("failed to decode tx detail: %w", err)
	}

	return &tx, nil
}

func (m *Module) getBTCTransactionIndex(ctx context.Context, client *http.Client, txid string) (int64, error) {
	reqCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s/api/tx/%s/merkle-proof", m.Config.BTCRPC, txid)
	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create merkle proof request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch merkle proof: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("merkle proof API returned status %d", resp.StatusCode)
	}

	var proof struct {
		Pos int64 `json:"pos"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&proof); err != nil {
		return 0, fmt.Errorf("failed to decode merkle proof: %w", err)
	}

	return proof.Pos, nil
}

func (m *Module) ensureBTCTxBlockHeight(ctx context.Context, client *http.Client, tx *types.BTCTx) error {
	if tx.Status.BlockHeight != 0 {
		return nil
	}

	detail, err := m.getBTCTxDetails(ctx, client, tx.TxID)
	if err != nil {
		return err
	}

	if detail.Status.BlockHeight == 0 {
		return fmt.Errorf("tx detail missing block height")
	}

	tx.Status.BlockHeight = detail.Status.BlockHeight
	return nil
}

func (m *Module) populateBTCTxIndex(ctx context.Context, client *http.Client, tx *types.BTCTx) (int64, error) {
	idx, err := m.getBTCTransactionIndex(ctx, client, tx.TxID)
	if err != nil {
		return 0, err
	}

	tx.TxIndex = idx
	return idx, nil
}

// processBTCTxWithTransaction processes a single BTC transaction with pre-parsed address data
// Address data is already available in the tx struct, validation has been done earlier
func (m *Module) processBTCTxWithTransaction(tx types.BTCTx) error {
	return m.processTransactionWithRetry("BTC", func() error {
		return m.database.WithTransaction(func(dbTx *sql.Tx) error {
			// 1. Save business data (using address data from tx struct)
			if err := m.saveBTCTransaction(dbTx, tx); err != nil {
				return fmt.Errorf("failed to save BTC transaction data: %w", err)
			}

			// 2. Mark as processed last (atomicity guarantee)
			// Note: ON CONFLICT DO NOTHING in the database handles concurrent processing
			if err := m.database.MarkTransactionProcessedInTx(dbTx, "BTC", tx.TxID, tx.Status.BlockHeight); err != nil {
				return fmt.Errorf("failed to mark BTC transaction as processed: %w", err)
			}

			return nil
		})
	})
}

// saveBTCTransaction saves BTC transaction data to database within a transaction
func (m *Module) saveBTCTransaction(dbTx *sql.Tx, tx types.BTCTx) error {
	// Find vault output
	var vaultOutput *types.BTCVout
	for _, vout := range tx.Vout {
		if normalizeAddress(vout.ScriptPubKeyAddr) ==
			normalizeAddress(m.Config.BTCVaultAddr) {
			vaultOutput = &vout
			break
		}
	}

	if vaultOutput == nil {
		return fmt.Errorf("vault output not found")
	}

	// Create staker ID
	stakerID := tx.ImuachainAddress + "_0x1" // BTC chain ID = 1

	// Save staker asset
	stakerAsset := &types.BootstrapStakerAsset{
		StakerID:     stakerID,
		AssetID:      VirtualAddress + "_0x1", // BTC asset ID: virtualAddress + chainID
		Deposited:    strconv.FormatInt(vaultOutput.Value, 10),
		Withdrawable: "0", // All stakes must be delegated
		Delegated:    strconv.FormatInt(vaultOutput.Value, 10),
		UpdatedAt:    time.Now(),
	}

	if err := m.database.SaveBootstrapStakerAssetInTx(dbTx, stakerAsset); err != nil {
		return fmt.Errorf("failed to save staker asset: %s", err)
	}

	// update the related token states
	if err := m.database.UpdateBootstrapTokenInTx(dbTx, stakerAsset.AssetID, stakerAsset.Deposited); err != nil {
		return fmt.Errorf("failed to update the state of bootstrap token: %s", err)
	}

	// Save delegation state
	delegationState := &types.BootstrapDelegationState{
		StakerID:     stakerID,
		AssetID:      VirtualAddress + "_0x1",
		OperatorAddr: tx.ValidatorAddress,
		Delegated:    strconv.FormatInt(vaultOutput.Value, 10),
		UpdatedAt:    time.Now(),
	}

	if err := m.database.SaveBootstrapDelegationStateInTx(dbTx, delegationState); err != nil {
		return fmt.Errorf("failed to save delegation state: %s", err)
	}

	// update the related operator asset states
	if err := m.database.UpdateBootstrapOperatorAssetInTx(dbTx, delegationState.AssetID, delegationState.OperatorAddr, tx.ImuachainAddress, delegationState.Delegated); err != nil {
		return fmt.Errorf("failed to update the state of operator asset: %s", err)
	}

	log.Info().
		Str("txid", tx.TxID).
		Str("staker", stakerID).
		Str("validator", tx.ValidatorAddress).
		Int64("amount", vaultOutput.Value).
		Msg("processed BTC bootstrap transaction")

	return nil
}

// XRP vault transaction fetching and processing
func (m *Module) refetchXRPStates() error {
	log.Debug().Str("module", "bootstrap").Str("refetching", "XRP states").
		Msg("refetching XRP states")

	// Pre-flight XRP connection health check to avoid first request failure
	if err := m.XrpClient.Ping([]byte("PING")); err != nil {
		log.Warn().Str("module", "bootstrap").Err(err).Msg("XRP ping failed, attempting reconnect")
		if recErr := m.recreateXRPClient(); recErr != nil {
			return fmt.Errorf("failed to reconnect XRP client after ping failure: %w", recErr)
		}
	}

	// Get current ledger index
	currentLedger, err := m.getXRPCurrentLedger()
	if err != nil {
		return fmt.Errorf("error getting current XRP ledger: %s", err)
	}

	// Get or initialize scan state
	scanState, err := m.database.GetScanState("XRP")
	if err != nil || scanState == nil {
		// First time scanning or record not found, initialize with start ledger
		startLedger := m.Config.XRPStartLedger
		if startLedger == 0 {
			// Start from current ledger minus confirmation depth for safety
			startLedger = currentLedger - int64(m.Config.XRPMinConfirmations)
			if startLedger < 0 {
				startLedger = 0
			}
		}

		scanState = &types.ScanState{
			ChainType:  "XRP",
			LastHeight: startLedger,
			SafeHeight: startLedger,
			UpdatedAt:  time.Now(),
		}

		if err != nil {
			log.Info().Int64("start_ledger", startLedger).Err(err).Msg("initializing XRP scan state due to error")
		} else {
			log.Info().Int64("start_ledger", startLedger).Msg("initializing XRP scan state (no existing record)")
		}
	}

	// Calculate safe current ledger (with confirmation buffer)
	safeCurrentLedger := currentLedger - int64(m.Config.XRPMinConfirmations)
	if safeCurrentLedger <= scanState.SafeHeight {
		log.Debug().Int64("safe_ledger", safeCurrentLedger).Int64("last_safe", scanState.SafeHeight).
			Msg("no new confirmed XRP ledgers to scan")
		return nil
	}

	log.Info().Int64("from_ledger", scanState.SafeHeight).Int64("to_ledger", safeCurrentLedger).
		Msg("starting incremental XRP scan")

	// Get transactions for ledger range (incremental scan)
	// Transactions are pre-validated, confirmed, and include parsed memo data
	transactions, err := m.getXRPVaultTransactionsFromLedger(scanState.SafeHeight+1, safeCurrentLedger)
	if err != nil {
		return fmt.Errorf("error getting XRP vault transactions from ledger %d to %d: %s",
			scanState.SafeHeight+1, safeCurrentLedger, err)
	}

	// Sort transactions by ledger index and order for consistent processing
	sortXRPTransactions(transactions)
	processedCount := 0

	// Process transactions that are already validated and confirmed
	for _, tx := range transactions {
		if err := m.processXRPTxWithTransaction(tx); err != nil {
			log.Err(err).Str("hash", tx.Hash).Msg("error processing XRP transaction")
			continue
		}

		processedCount++
	}

	// Log processing statistics
	totalTxs := len(transactions)
	log.Info().
		Int("total_txs", totalTxs).
		Int("processed", processedCount).
		Int64("from_ledger", scanState.SafeHeight).
		Int64("to_ledger", safeCurrentLedger).
		Msg("XRP transaction processing statistics")

	// Update scan state
	newScanState := &types.ScanState{
		ChainType:  "XRP",
		LastHeight: currentLedger,
		SafeHeight: safeCurrentLedger,
		UpdatedAt:  time.Now(),
	}

	if err := m.database.UpdateScanState(newScanState); err != nil {
		return fmt.Errorf("failed to update XRP scan state: %s", err)
	}

	log.Info().Int("processed_count", processedCount).
		Int64("from_ledger", scanState.SafeHeight).
		Int64("to_ledger", safeCurrentLedger).
		Msg("completed incremental XRP scan")

	return nil
}

// getXRPCurrentLedger gets the current XRP ledger index using xrpl-go client
func (m *Module) getXRPCurrentLedger() (int64, error) {
	request := map[string]interface{}{
		"command":      "ledger",
		"ledger_index": "validated",
		"binary":       false,
		"api_version":  2,
	}

	log.Debug().Str("module", "bootstrap").Interface("request", request).Msg("sending XRP ledger request")

	responseMap, err := m.xrpRequestWithReconnect(request)
	if err != nil {
		return 0, fmt.Errorf("failed to request current ledger: %w", err)
	}

	log.Debug().Str("module", "bootstrap").Interface("response", responseMap).Msg("received XRP ledger response")

	// Parse response to extract ledger index
	result, ok := responseMap["result"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("invalid response result format")
	}

	// First try to get ledger_index from result level (this is usually a number)
	if ledgerIndexValue, exists := result["ledger_index"]; exists {
		switch v := ledgerIndexValue.(type) {
		case float64:
			return int64(v), nil
		case int:
			return int64(v), nil
		case int64:
			return v, nil
		case string:
			// Try to parse string to int64
			if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
				return parsed, nil
			}
		}
	}

	// Fallback to ledger.ledger_index (this might be a string)
	ledger, ok := result["ledger"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("invalid ledger format")
	}

	ledgerIndex, ok := ledger["ledger_index"]
	if !ok {
		return 0, fmt.Errorf("ledger_index not found in either location")
	}

	switch v := ledgerIndex.(type) {
	case float64:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case string:
		// Try to parse string to int64
		if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
			return parsed, nil
		}
		return 0, fmt.Errorf("failed to parse ledger_index string: %s", v)
	default:
		return 0, fmt.Errorf("invalid ledger_index type: %T", v)
	}
}

// xrpRequestWithReconnect performs an XRPL request and, on failure, attempts a single reconnect then retries once.
func (m *Module) xrpRequestWithReconnect(request map[string]interface{}) (map[string]interface{}, error) {
	response, err := m.XrpClient.Request(request)
	if err == nil {
		return response, nil
	}

	log.Error().Err(err).Str("module", "bootstrap").Msg("XRP request failed, attempting reconnect")
	if recErr := m.recreateXRPClient(); recErr != nil {
		return nil, fmt.Errorf("xrp request failed and reconnect failed: %w", err)
	}
	log.Warn().Str("module", "bootstrap").Msg("retrying XRP request after reconnect")
	response, err = m.XrpClient.Request(request)
	if err != nil {
		return nil, fmt.Errorf("xrp request failed after reconnect: %w", err)
	}
	return response, nil
}

// recreateXRPClient rebuilds the XRP client connection and verifies connectivity
func (m *Module) recreateXRPClient() error {
	log.Warn().Str("module", "bootstrap").Msg("recreating XRP client due to connection issue")
	client := xrpl.NewClient(xrpl.ClientConfig{URL: m.Config.XRPRPC})
	if err := client.Ping([]byte("PING")); err != nil {
		log.Error().Err(err).Str("module", "bootstrap").Msg("failed to ping XRPL after recreating client")
		return err
	}
	m.XrpClient = client
	log.Info().Str("module", "bootstrap").Msg("successfully reconnected XRP client")
	return nil
}

// parseXRPTransactionFromAccountTx safely parses XRP transaction data from account_tx response
// This function also performs bootstrap validation and memo parsing to avoid duplicate processing
func (m *Module) parseXRPTransactionFromAccountTx(txData map[string]interface{}) (*types.XRPTransaction, error) {
	// The transaction data is nested under "tx_json" field and metadata under "meta"
	txField, hasTx := txData["tx_json"]
	metaField, hasMeta := txData["meta"]

	if !hasTx || !hasMeta {
		return nil, fmt.Errorf("missing tx_json or meta field in account_tx response")
	}

	txJSON, ok := txField.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid tx_json field type")
	}

	metaJSON, ok := metaField.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid meta field type")
	}

	// Extract basic transaction info - hash might be in txJSON or top level txData
	var hash string
	var err error

	// Try to get hash from txJSON first
	if hashValue, exists := txJSON["hash"]; exists {
		if hashStr, ok := hashValue.(string); ok {
			hash = hashStr
		}
	}

	// If not found in txJSON, try top level txData
	if hash == "" {
		hash, err = safeStringExtract(txData, "hash")
		if err != nil {
			return nil, fmt.Errorf("failed to extract hash from either location: %s", err)
		}
	}

	// Get ledger_index from txJSON
	ledgerIndexFloat, ok := txJSON["ledger_index"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid ledger_index type")
	}
	ledgerIndex := int64(ledgerIndexFloat)

	// Get date from txJSON or txData
	var date int64
	if dateFloat, ok := txJSON["date"].(float64); ok {
		date = int64(dateFloat)
	} else if dateFloat, ok := txData["date"].(float64); ok {
		date = int64(dateFloat)
	}

	// Check if validated
	validated, _ := txData["validated"].(bool)
	if !validated {
		return nil, fmt.Errorf("transaction not validated")
	}

	// Extract transaction details with early validation
	transactionType, err := safeStringExtract(txJSON, "TransactionType")
	if err != nil {
		return nil, fmt.Errorf("failed to extract TransactionType: %s", err)
	}
	if transactionType != "Payment" {
		log.Debug().Str("module", "bootstrap").Str("hash", hash).
			Str("type", transactionType).Msg("skipping non-payment transaction")
		return nil, nil // Not a payment transaction, skip early
	}

	account, err := safeStringExtract(txJSON, "Account")
	if err != nil {
		return nil, fmt.Errorf("failed to extract Account: %s", err)
	}

	// Validate account format
	if !strings.HasPrefix(account, "r") || len(account) < 25 || len(account) > 35 {
		return nil, fmt.Errorf("invalid XRP account format: %s", account)
	}
	if account == m.Config.XRPVaultAddr {
		log.Debug().Str("module", "bootstrap").Str("hash", hash).
			Str("account", account).Msg("skipping self-transfer from vault")
		return nil, nil // Self-transfer from vault, skip early
	}

	// Extract and validate destination
	var destination string
	if dest, ok := txJSON["Destination"].(string); ok {
		if dest != m.Config.XRPVaultAddr {
			log.Debug().Str("module", "bootstrap").Str("hash", hash).
				Str("destination", dest).Str("expected", m.Config.XRPVaultAddr).
				Msg("skipping transaction not sent to vault")
			return nil, nil // Not sent to our vault, skip early
		}
		destination = dest
	} else {
		log.Debug().Str("module", "bootstrap").Str("hash", hash).
			Msg("skipping transaction with no destination")
		return nil, nil // No destination specified, skip
	}

	// Extract and validate destination tag
	var destTagInt int64
	if destTag, ok := txJSON["DestinationTag"].(float64); ok {
		destTagInt = int64(destTag)
		if destTagInt != m.Config.XRPDestinationTag {
			return nil, nil // Wrong destination tag, skip early
		}
	} else {
		return nil, nil // No destination tag specified, skip
	}

	// Extract meta information
	transactionResult, err := safeStringExtract(metaJSON, "TransactionResult")
	if err != nil {
		return nil, fmt.Errorf("failed to extract TransactionResult: %s", err)
	}
	if transactionResult != "tesSUCCESS" {
		return nil, nil // Transaction not successful, skip early
	}

	transactionIndexFloat, ok := metaJSON["TransactionIndex"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid TransactionIndex type")
	}
	transactionIndex := int64(transactionIndexFloat)

	// Create transaction struct
	tx := &types.XRPTransaction{
		Hash:        hash,
		LedgerIndex: ledgerIndex,
		Date:        date,
		Validated:   validated,
		Tx: types.XRPTx{
			TransactionType: transactionType,
			Account:         account,
			Destination:     destination,
			DestinationTag:  destTagInt,
		},
		Meta: types.XRPMeta{
			TransactionResult: transactionResult,
			TransactionIndex:  transactionIndex,
		},
	}

	// Use DeliverMax if available, otherwise delivered_amount from meta, otherwise Amount
	if deliverMax, ok := txJSON["DeliverMax"]; ok {
		tx.Tx.Amount = deliverMax
	} else if deliveredAmount, ok := metaJSON["delivered_amount"].(string); ok && deliveredAmount != "" {
		tx.Tx.Amount = deliveredAmount
	} else if amount, ok := txJSON["Amount"]; ok {
		tx.Tx.Amount = amount
	}

	if fee, ok := txJSON["Fee"].(string); ok {
		tx.Tx.Fee = fee
	}

	// Parse memos safely
	if memosInterface, ok := txJSON["Memos"].([]interface{}); ok {
		memos := make([]types.XRPMemo, 0, len(memosInterface))
		for _, memoInterface := range memosInterface {
			memoData, ok := memoInterface.(map[string]interface{})
			if !ok {
				continue
			}
			memo, ok := memoData["Memo"].(map[string]interface{})
			if !ok {
				continue
			}

			xrpMemo := types.XRPMemo{}
			if memoType, ok := memo["MemoType"].(string); ok {
				xrpMemo.Memo.MemoType = memoType
			}
			if memoDataStr, ok := memo["MemoData"].(string); ok {
				xrpMemo.Memo.MemoData = memoDataStr
			}
			if memoFormat, ok := memo["MemoFormat"].(string); ok {
				xrpMemo.Memo.MemoFormat = memoFormat
			}

			memos = append(memos, xrpMemo)
		}
		tx.Tx.Memos = memos
	}

	// Perform detailed validation: memo parsing and amount checks
	if err := m.validateAndParseBootstrapXRPTx(tx); err != nil {
		return nil, fmt.Errorf("bootstrap XRP transaction validation failed: %w", err)
	}

	return tx, nil
}

// processXRPTxWithTransaction processes a single XRP transaction within a database transaction
// This function assumes the transaction has already been validated and memo data parsed
func (m *Module) processXRPTxWithTransaction(tx types.XRPTransaction) error {
	// Ensure 1-1 address binding at processing time (after transactions are sorted by ledger_index)
	correctImuachainAddr, err := m.validateXRPAddressBinding(tx.Tx.Account, tx.ImuachainAddress, tx.Hash)
	if err != nil {
		return err
	}
	tx.ImuachainAddress = correctImuachainAddr

	return m.processTransactionWithRetry("XRP", func() error {
		return m.database.WithTransaction(func(dbTx *sql.Tx) error {
			// Save business data using the pre-parsed address fields
			if err := m.saveXRPTransaction(dbTx, tx); err != nil {
				return fmt.Errorf("failed to save XRP transaction data: %w", err)
			}

			// Mark as processed last (atomicity guarantee)
			// Note: ON CONFLICT DO NOTHING in the database handles concurrent processing
			if err := m.database.MarkTransactionProcessedInTx(dbTx, "XRP", tx.Hash, tx.LedgerIndex); err != nil {
				return fmt.Errorf("failed to mark XRP transaction as processed: %w", err)
			}

			return nil
		})
	})
}

// validateAndParseBootstrapXRPTx validates XRP transaction memo and amount (basic checks already done)
// Returns error if the transaction is invalid for bootstrap, nil if valid
func (m *Module) validateAndParseBootstrapXRPTx(tx *types.XRPTransaction) error {
	// Must be XRP payment (not token)
	amountStr, ok := tx.Tx.Amount.(string)
	if !ok {
		return fmt.Errorf("invalid amount type: expected string for XRP payment")
	}

	// Check minimum amount
	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse XRP amount %s: %w", amountStr, err)
	}

	if amount < m.Config.XRPMinAmount {
		return fmt.Errorf("XRP amount %d below minimum required %d", amount, m.Config.XRPMinAmount)
	}

	// Must have memo with validator info
	if len(tx.Tx.Memos) == 0 {
		return fmt.Errorf("XRP transaction missing memo data for bootstrap")
	}

	// Parse memo data
	memoData, err := parseXRPMemo(tx.Tx.Memos)
	if err != nil {
		return fmt.Errorf("failed to parse XRP memo: %w", err)
	}

	// Validate validator address
	if !m.isValidValidatorAddress(memoData.ValidatorAddress) {
		return fmt.Errorf("invalid validator address in XRP memo: %s", memoData.ValidatorAddress)
	}

	// Check if validator is registered
	// TODO: Temporarily commented out for testing - no available bootstrap contract with valid validators
	/*
		isRegistered, err := m.isValidatorRegistered(memoData.ValidatorAddress)
		if err != nil {
			return fmt.Errorf("failed to check validator registration for %s: %w", memoData.ValidatorAddress, err)
		}

		if !isRegistered {
			return fmt.Errorf("validator %s is not registered for bootstrap", memoData.ValidatorAddress)
		}
	*/

	// Set parsed addresses in the transaction struct (binding is established during processing stage)
	tx.ImuachainAddress = memoData.ImuachainAddress
	tx.ValidatorAddress = memoData.ValidatorAddress

	return nil
}

// XRPMemoData represents parsed memo data
type XRPMemoData struct {
	ImuachainAddress string
	ValidatorAddress string
}

// parseXRPMemo parses memo data from XRP transaction
func parseXRPMemo(memos []types.XRPMemo) (*XRPMemoData, error) {
	for _, memo := range memos {
		// Validate MemoType is "Description" (hex: 4465736372697074696F6E)
		if memo.Memo.MemoType != "4465736372697074696F6E" {
			continue
		}

		// Decode memo data
		buffer, err := hex.DecodeString(memo.Memo.MemoData)
		if err != nil {
			continue
		}

		// Validate minimum length (40 chars eth hex + 41 chars validator = 81 chars)
		if len(buffer) < 81 {
			continue
		}

		// The buffer is a concatenated string: eth_hex(40) + validator_ascii(41)
		bufferStr := string(buffer)

		// Extract ethereum address (first 40 characters as hex)
		ethHex := bufferStr[:40]
		imuachainAddress := "0x" + ethHex

		if !common.IsHexAddress(imuachainAddress) {
			continue
		}

		// Extract validator address (remaining 41 characters)
		validatorAddress := bufferStr[40:]

		// Basic validation of validator address format
		if len(validatorAddress) != 41 || !strings.HasPrefix(validatorAddress, "im1") {
			continue
		}

		// Validate validator address characters (basic bech32 format)
		validChars := true
		for _, char := range validatorAddress {
			if !((char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')) {
				validChars = false
				break
			}
		}
		if !validChars {
			continue
		}

		return &XRPMemoData{
			ImuachainAddress: strings.ToLower(imuachainAddress),
			ValidatorAddress: validatorAddress,
		}, nil
	}

	return nil, fmt.Errorf("no valid memo data found")
}

// saveXRPTransaction saves XRP transaction data to database within a transaction
func (m *Module) saveXRPTransaction(dbTx *sql.Tx, tx types.XRPTransaction) error {
	// Get amount
	amountStr, ok := tx.Tx.Amount.(string)
	if !ok {
		return fmt.Errorf("invalid amount type")
	}

	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse amount: %s", err)
	}

	// Create staker ID
	stakerID := tx.ImuachainAddress + "_0x2" // XRP chain ID = 2

	// Save staker asset
	stakerAsset := &types.BootstrapStakerAsset{
		StakerID:     stakerID,
		AssetID:      VirtualAddress + "_0x2", // XRP asset ID: virtualAddress + chainID
		Deposited:    amountStr,
		Withdrawable: "0", // All stakes must be delegated
		Delegated:    amountStr,
		UpdatedAt:    time.Now(),
	}

	if err := m.database.SaveBootstrapStakerAssetInTx(dbTx, stakerAsset); err != nil {
		return fmt.Errorf("failed to save staker asset: %s", err)
	}

	// update the related token states
	if err := m.database.UpdateBootstrapTokenInTx(dbTx, stakerAsset.AssetID, stakerAsset.Deposited); err != nil {
		return fmt.Errorf("failed to update the state of bootstrap token: %s", err)
	}

	// Save delegation state
	delegationState := &types.BootstrapDelegationState{
		StakerID:     stakerID,
		AssetID:      VirtualAddress + "_0x2",
		OperatorAddr: tx.ValidatorAddress,
		Delegated:    amountStr,
		UpdatedAt:    time.Now(),
	}

	if err := m.database.SaveBootstrapDelegationStateInTx(dbTx, delegationState); err != nil {
		return fmt.Errorf("failed to save delegation state: %s", err)
	}

	// update the related operator asset states
	if err := m.database.UpdateBootstrapOperatorAssetInTx(dbTx, delegationState.AssetID, delegationState.OperatorAddr, tx.ImuachainAddress, delegationState.Delegated); err != nil {
		return fmt.Errorf("failed to update the state of operator asset: %s", err)
	}

	log.Info().
		Str("hash", tx.Hash).
		Str("staker", stakerID).
		Str("validator", tx.ValidatorAddress).
		Int64("amount", amount).
		Msg("processed XRP bootstrap transaction")

	return nil
}

// getXRPVaultTransactionsFromLedger gets XRP vault transactions from a specific ledger range using efficient account_tx method
func (m *Module) getXRPVaultTransactionsFromLedger(fromLedger, toLedger int64) ([]types.XRPTransaction, error) {
	log.Info().Int64("from_ledger", fromLedger).Int64("to_ledger", toLedger).
		Str("vault_address", m.Config.XRPVaultAddr).
		Msg("fetching XRP vault transactions using account_tx method")

	var allTxs []types.XRPTransaction
	var marker interface{} // For pagination

	for {
		// Build account_tx request
		request := map[string]interface{}{
			"command":          "account_tx",
			"account":          m.Config.XRPVaultAddr,
			"binary":           false,
			"api_version":      2,
			"ledger_index_min": fromLedger,
			"ledger_index_max": toLedger,
			"limit":            200, // Maximum allowed by rippled
		}

		// Add pagination marker if available
		if marker != nil {
			request["marker"] = marker
		}

		log.Debug().Interface("request", request).Msg("sending account_tx request")

		responseMap, err := m.xrpRequestWithReconnect(request)
		if err != nil {
			return nil, fmt.Errorf("failed to request account transactions: %w", err)
		}

		// Parse response
		result, ok := responseMap["result"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid response result format")
		}

		// Check if the request was successful
		if status, ok := result["status"].(string); ok && status != "success" {
			// If no transactions found, it's not an error
			if status == "actNotFound" {
				log.Info().Str("vault_address", m.Config.XRPVaultAddr).
					Msg("no transactions found for vault address (account not found)")
				break
			}
			return nil, fmt.Errorf("account_tx request failed with status: %s", status)
		}

		// Extract transactions
		transactions, ok := result["transactions"].([]interface{})
		if !ok {
			log.Info().Msg("no transactions field in response")
			break
		}

		// Process transactions
		for _, txData := range transactions {
			txMap, ok := txData.(map[string]interface{})
			if !ok {
				log.Debug().Msg("skipping invalid transaction data")
				continue
			}

			// Parse the transaction with validation
			tx, err := m.parseXRPTransactionFromAccountTx(txMap)
			if err != nil {
				log.Debug().Err(err).Msg("skipping invalid XRP transaction")
				continue
			}
			if tx == nil {
				// Transaction parsed but failed bootstrap validation
				continue
			}

			// Filter by ledger range (double-check since the API should already filter)
			if tx.LedgerIndex >= fromLedger && tx.LedgerIndex <= toLedger {
				allTxs = append(allTxs, *tx)
			}
		}

		// Check for pagination marker
		if nextMarker, exists := result["marker"]; exists {
			marker = nextMarker
			log.Debug().Interface("marker", marker).
				Int("txs_processed", len(transactions)).
				Msg("continuing with next page of transactions")

			// Add small delay between requests to be API-friendly
			time.Sleep(100 * time.Millisecond)
		} else {
			// No more pages
			break
		}
	}

	log.Info().Int("total_txs", len(allTxs)).
		Int64("from_ledger", fromLedger).
		Int64("to_ledger", toLedger).
		Msg("completed efficient XRP vault transaction fetching")

	return allTxs, nil
}

type CoinbaseResponse struct {
	Data struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"data"`
}

func GetSpotPrice(symbol string) (string, error) {
	url := fmt.Sprintf("https://api.coinbase.com/v2/prices/%s-USD/spot", symbol)

	//nolint:gosec // G107: url is trusted
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result CoinbaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result.Data.Amount, nil
}

func (m *Module) updatePricesAndTVL() error {
	log.Debug().Str("module", "bootstrap").Str("refetching", "prices and TVL").
		Msg("refetching prices and TVL")
	tokens, err := m.database.ListBootstrapTokens()
	if err != nil {
		return err
	}

	oracleFeedsMap := make(map[string]common.Address, len(m.Config.TokenOracleFeeds))
	for _, tokenOracleFeed := range m.Config.TokenOracleFeeds {
		oracleFeedsMap[tokenOracleFeed.AssetID] = common.HexToAddress(tokenOracleFeed.OracleAddr)
	}
	totalTVL := sdkmath.LegacyZeroDec()

	for _, t := range tokens {
		stakingAmountInt, ok := sdkmath.NewIntFromString(t.StakingTotalAmount)
		if !ok {
			log.Error().Str("stakingTotalAmount", t.StakingTotalAmount).Msg("failed to parse the staking amount to a big int")
			stakingAmountInt = sdkmath.ZeroInt()
		}

		oracleFeedAddr, ok := oracleFeedsMap[t.AssetID]
		if !ok {
			price, err := GetSpotPrice(t.Symbol)
			if err != nil {
				log.Err(err).Str("module", "bootstrap").Str("assetID", t.AssetID).Str("name", t.Name).Str("symbol", t.Symbol).
					Msg("failed to get the asset price from coinbase")
				totalTVL.AddMut(sdkmath.LegacyMustNewDecFromStr(t.TotalUSDValue))
			} else {
				// update the price
				err = m.database.SaveBootstrapTokenPrice(t.AssetID, price)
				if err != nil {
					return err
				}
				// calculate the total USD value of this asset
				priceDec, err := sdkmath.LegacyNewDecFromStr(price)
				if err != nil {
					log.Err(err).Str("binancePrice", price).Msg("failed to parse the coinbase price to a big legacyDec")
					// don't return to continue addressing the other assets
					continue
				}
				divisor := sdkmath.NewIntWithDecimal(1, int(t.Decimals)) // #nosec G115
				usdValue := priceDec.MulInt(stakingAmountInt).QuoInt(divisor)
				totalTVL.AddMut(usdValue)
				err = m.database.UpdateBootstrapTokenUSDValue(t.AssetID, usdValue.String())
				if err != nil {
					return err
				}
			}
		} else {
			// fetch the price from ChainLink
			aggregatorContract, err := aggregatorv3.NewAggregatorV3Interface(oracleFeedAddr, m.FeederEthHTTPClient)
			if err != nil {
				return err
			}
			aggregatorSession := aggregatorv3.AggregatorV3InterfaceSession{
				Contract: aggregatorContract,
				CallOpts: bind.CallOpts{Context: m.ctx},
			}
			roundData, err := aggregatorSession.LatestRoundData()
			if err != nil {
				return err
			}
			decimals, err := aggregatorSession.Decimals()
			if err != nil {
				return err
			}
			divisor := sdkmath.NewIntWithDecimal(1, int(decimals)) // #nosec G115
			priceDec := sdk.NewDecFromBigInt(roundData.Answer).QuoInt(divisor)
			// update the price
			err = m.database.SaveBootstrapTokenPrice(t.AssetID, priceDec.String())
			if err != nil {
				return err
			}
			// calculate the total USD value of this asset
			usdValue := operatorkeeper.CalculateUSDValue(stakingAmountInt, sdkmath.NewIntFromBigInt(roundData.Answer), uint32(t.Decimals), decimals)
			totalTVL.AddMut(usdValue)
			err = m.database.UpdateBootstrapTokenUSDValue(t.AssetID, usdValue.String())
			if err != nil {
				return err
			}
		}
	}

	// save the total TVL
	return m.database.SaveBootstrapStatistics(totalTVL.String())
}
