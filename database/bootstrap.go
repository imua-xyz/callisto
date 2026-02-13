package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"

	"github.com/forbole/callisto/v4/types"
)

func (db *Db) SaveBootstrapValidator(v *types.BootstrapValidator) error {
	stmt := `
INSERT INTO bootstrap_validator (
    validator_eth_addr,
    validator_im_addr,
    validator_name,
    consensus_pub_key,
    commission_rate,
    max_commission_rate,
    max_change_rate,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (validator_eth_addr) DO UPDATE
SET validator_im_addr   = EXCLUDED.validator_im_addr,
    validator_name      = EXCLUDED.validator_name,
    consensus_pub_key   = EXCLUDED.consensus_pub_key,
    commission_rate     = EXCLUDED.commission_rate,
    max_commission_rate = EXCLUDED.max_commission_rate,
    max_change_rate     = EXCLUDED.max_change_rate,
    updated_at          = EXCLUDED.updated_at;`

	_, err := db.SQL.Exec(
		stmt,
		v.ValidatorEthAddress,
		v.ValidatorIMAddress,
		v.ValidatorName,
		v.ConsensusPubKey,
		v.Rate,
		v.MaxRate,
		v.MaxChangeRate,
		v.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap validator: %w", err)
	}

	return nil
}

// UpdateCommissionRate updates the commission rate for a bootstrap validator identified
// by validatorAddr. It also updates the updated_at timestamp to the current time.
func (db *Db) UpdateCommissionRate(validatorAddr string, newRate string) error {
	stmt := `
UPDATE bootstrap_validator
SET commission_rate     = $1,
    updated_at          = $2
WHERE validator_im_addr = $3;
`

	_, err := db.SQL.Exec(
		stmt,
		newRate,
		time.Now(),
		validatorAddr,
	)
	if err != nil {
		return fmt.Errorf("failed to update commission rate for %s: %w", validatorAddr, err)
	}
	return nil
}

// UpdateConsensusPubKey updates the consensus public key for a bootstrap validator
// identified by validatorIMAddr. The updated_at timestamp is also set to the current time.
func (db *Db) UpdateConsensusPubKey(validatorIMAddr string, newPubKey string) error {
	stmt := `
UPDATE bootstrap_validator
SET consensus_pub_key = $1,
    updated_at        = $2
WHERE validator_im_addr = $3;
`

	_, err := db.SQL.Exec(
		stmt,
		newPubKey,
		time.Now(),
		validatorIMAddr,
	)
	if err != nil {
		return fmt.Errorf("failed to update consensus pub key for %s: %w", validatorIMAddr, err)
	}
	return nil
}

func (db *Db) SaveBootstrapClientChain(c *types.BootstrapClientChain) error {
	stmt := `
INSERT INTO bootstrap_client_chains (
    name, meta_info, layer_zero_chain_id, updated_at
) VALUES ($1, $2, $3, now())
ON CONFLICT (layer_zero_chain_id) DO UPDATE
SET name       = EXCLUDED.name,
    meta_info  = EXCLUDED.meta_info,
    updated_at = EXCLUDED.updated_at;`

	_, err := db.SQL.Exec(stmt,
		c.Name,
		c.MetaInfo,
		c.LZChainID,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap client chain: %w", err)
	}
	return nil
}

func (db *Db) GetBootstrapClientChain(layerZeroChainID uint64) (*types.BootstrapClientChain, error) {
	stmt := `
SELECT name, meta_info, layer_zero_chain_id
FROM bootstrap_client_chains
WHERE layer_zero_chain_id = $1
LIMIT 1;`

	row := db.SQL.QueryRow(stmt, layerZeroChainID)

	var c types.BootstrapClientChain
	err := row.Scan(
		&c.Name,
		&c.MetaInfo,
		&c.LZChainID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no bootstrap client chain found for layer_zero_chain_id %d", layerZeroChainID)
		}
		return nil, fmt.Errorf("failed to get bootstrap client chain: %w", err)
	}

	return &c, nil
}

func (db *Db) SaveBootstrapToken(t *types.BootstrapTokenState) error {
	stmt := `
INSERT INTO bootstrap_tokens (
    asset_id, name, symbol, address, decimals,
    layer_zero_chain_id, staking_total_amount, total_usd_value, updated_at
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
ON CONFLICT (asset_id) DO UPDATE
SET name                 = EXCLUDED.name,
    symbol               = EXCLUDED.symbol,
    address              = EXCLUDED.address,
    decimals             = EXCLUDED.decimals,
    layer_zero_chain_id  = EXCLUDED.layer_zero_chain_id,
    staking_total_amount = EXCLUDED.staking_total_amount,
    total_usd_value      = EXCLUDED.total_usd_value,
    updated_at           = EXCLUDED.updated_at;`

	_, err := db.SQL.Exec(stmt,
		t.AssetID,
		t.Name,
		t.Symbol,
		strings.ToLower(t.Address),
		t.Decimals,
		t.LZChainID,
		t.StakingTotalAmount,
		t.TotalUSDValue,
		t.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap token: %w", err)
	}
	return nil
}

func (db *Db) GetBootstrapToken(assetID string) (*types.BootstrapTokenState, error) {
	stmt := `
SELECT asset_id, name, symbol, address, decimals,
       layer_zero_chain_id, staking_total_amount, total_usd_value, updated_at
FROM bootstrap_tokens
WHERE asset_id = $1
LIMIT 1;`

	var t types.BootstrapTokenState
	err := db.SQL.QueryRow(stmt, assetID).Scan(
		&t.AssetID,
		&t.Name,
		&t.Symbol,
		&t.Address,
		&t.Decimals,
		&t.LZChainID,
		&t.StakingTotalAmount,
		&t.TotalUSDValue,
		&t.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no bootstrap token found for asset_id %s", assetID)
		}
		return nil, fmt.Errorf("failed to get bootstrap token: %w", err)
	}

	return &t, nil
}

func (db *Db) ListBootstrapTokens() ([]*types.BootstrapTokenState, error) {
	stmt := `
SELECT asset_id, name, symbol, address, decimals,
       layer_zero_chain_id, staking_total_amount,total_usd_value, updated_at
FROM bootstrap_tokens;`

	rows, err := db.SQL.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("failed to query bootstrap tokens: %w", err)
	}
	defer rows.Close()

	var tokens []*types.BootstrapTokenState
	for rows.Next() {
		t := new(types.BootstrapTokenState)
		err := rows.Scan(
			&t.AssetID,
			&t.Name,
			&t.Symbol,
			&t.Address,
			&t.Decimals,
			&t.LZChainID,
			&t.StakingTotalAmount,
			&t.TotalUSDValue,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan bootstrap token: %w", err)
		}
		tokens = append(tokens, t)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return tokens, nil
}

func (db *Db) UpdateBootstrapTokenAmountAndUSDValue(assetID string, amount, usdValue string) error {
	stmt := `
UPDATE bootstrap_tokens
SET staking_total_amount = $1,
    total_usd_value = $2,
    updated_at = $3
WHERE asset_id = $4;`

	_, err := db.SQL.Exec(stmt,
		amount,
		usdValue,
		time.Now(),
		assetID,
	)
	if err != nil {
		return fmt.Errorf("failed to update staking_total_amount for asset_id=%s: %w", assetID, err)
	}
	return nil
}

func (db *Db) UpdateBootstrapTokenUSDValue(assetID string, usdValue string) error {
	stmt := `
UPDATE bootstrap_tokens
SET total_usd_value = $1,
    updated_at      = $2
WHERE asset_id = $3;`

	_, err := db.SQL.Exec(stmt,
		usdValue,
		time.Now(),
		assetID,
	)
	if err != nil {
		return fmt.Errorf("failed to update total_usd_value for asset_id=%s: %w", assetID, err)
	}
	return nil
}

func (db *Db) UpdateBootstrapTokenInTx(tx *sql.Tx, assetID string, stakingDelta string) error {
	tokenInfo, err := db.GetBootstrapToken(assetID)
	if err != nil {
		return err
	}
	initialStakingAmount, ok := sdkmath.NewIntFromString(tokenInfo.StakingTotalAmount)
	if !ok {
		return fmt.Errorf("UpdateBootstrapTokenInTx: failed to parse total staking amount from string:%s", tokenInfo.StakingTotalAmount)
	}
	deltaAmount, ok := sdkmath.NewIntFromString(stakingDelta)
	if !ok {
		return fmt.Errorf("UpdateBootstrapTokenInTx: failed to parse stakingDelta staking amount from string:%s", stakingDelta)
	}
	newStakingAmount := initialStakingAmount.Add(deltaAmount)

	usdValue := sdkmath.LegacyZeroDec()
	priceStr, err := db.GetBootstrapTokenPrice(assetID)
	if err != nil {
		log.Err(err).Str("assetID", assetID).Msg("UpdateBootstrapTokenInTx: get token price from database")
		// Using zero as the USD value; continue handling other assets without returning
	} else {
		// calculate the total USD value of this asset
		priceDec := sdkmath.LegacyMustNewDecFromStr(priceStr)
		divisor := sdkmath.NewIntWithDecimal(1, int(tokenInfo.Decimals)) // #nosec G115
		usdValue = priceDec.MulInt(newStakingAmount).QuoInt(divisor)
	}

	// update the total staking amount and USD value in the bootstrap token state.
	var ownTx bool
	var retErr error
	if tx == nil {
		var err error
		tx, err = db.SQL.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		ownTx = true
		defer func() {
			if ownTx {
				if rbErr := tx.Rollback(); rbErr != nil {
					retErr = errors.Join(retErr, fmt.Errorf("failed to rollback transaction: %w", rbErr))
				}
			}
		}()
	}

	stmt := `
UPDATE bootstrap_tokens
SET staking_total_amount = $1,
    total_usd_value = $2,
    updated_at = $3
WHERE asset_id = $4;`

	_, err = tx.Exec(stmt,
		newStakingAmount.String(),
		usdValue.String(),
		time.Now(),
		assetID,
	)
	if err != nil {
		return fmt.Errorf("failed to update bootstrap asset: %w", err)
	}
	if !ownTx {
		return nil
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	ownTx = false
	return retErr
}

func (db *Db) SaveBootstrapStakerAsset(a *types.BootstrapStakerAsset) error {
	stmt := `
INSERT INTO bootstrap_staker_assets (
    staker_id, asset_id, deposited, withdrawable, delegated, updated_at, updated_at_block
) VALUES ($1,$2,$3,$4,$5,$6,$7)
ON CONFLICT (staker_id, asset_id) DO UPDATE
SET deposited        = EXCLUDED.deposited,
    withdrawable     = EXCLUDED.withdrawable,
    delegated        = EXCLUDED.delegated,
    updated_at       = EXCLUDED.updated_at,
    updated_at_block = EXCLUDED.updated_at_block;`

	_, err := db.SQL.Exec(stmt,
		a.StakerID,
		a.AssetID,
		a.Deposited,
		a.Withdrawable,
		a.Delegated,
		a.UpdatedAt,
		a.UpdatedAtBlock,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap staker asset: %w", err)
	}
	return nil
}

// SaveBootstrapStakerAssetInTx saves a bootstrap staker asset within an existing transaction
func (db *Db) SaveBootstrapStakerAssetInTx(tx *sql.Tx, a *types.BootstrapStakerAsset) (retErr error) {
	var ownTx bool
	if tx == nil {
		var err error
		tx, err = db.SQL.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		ownTx = true
		defer func() {
			if ownTx {
				if rbErr := tx.Rollback(); rbErr != nil {
					retErr = errors.Join(retErr, fmt.Errorf("failed to rollback transaction: %w", rbErr))
				}
			}
		}()
	}
	stmt := `
INSERT INTO bootstrap_staker_assets (
    staker_id, asset_id, deposited, withdrawable, delegated, updated_at, updated_at_block
) VALUES ($1,$2,$3,$4,$5,$6,$7)
ON CONFLICT (staker_id, asset_id) DO UPDATE
SET deposited        = bootstrap_staker_assets.deposited + EXCLUDED.deposited,
    withdrawable     = bootstrap_staker_assets.withdrawable + EXCLUDED.withdrawable,
    delegated        = bootstrap_staker_assets.delegated + EXCLUDED.delegated,
    updated_at       = EXCLUDED.updated_at,
    updated_at_block = EXCLUDED.updated_at_block;`

	_, err := tx.Exec(stmt,
		a.StakerID,
		a.AssetID,
		a.Deposited,
		a.Withdrawable,
		a.Delegated,
		a.UpdatedAt,
		a.UpdatedAtBlock,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap staker asset: %w", err)
	}

	if !ownTx {
		return nil
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	ownTx = false
	return nil
}

func (db *Db) BootstrapStakerAssetExists(stakerID, assetID string) (bool, error) {
	stmt := `
SELECT EXISTS(
    SELECT 1
    FROM bootstrap_staker_assets
    WHERE staker_id = $1 AND asset_id = $2
);`

	var exists bool
	err := db.SQL.QueryRow(stmt, stakerID, assetID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check bootstrap staker asset existence: %w", err)
	}
	return exists, nil
}

func (db *Db) ClaimBootstrapStakerAsset(stakerID, assetID string, claimAmount int64, updatedAt time.Time, blockHeight int64) error {
	stmt := `
UPDATE bootstrap_staker_assets
SET
    deposited        = deposited - $3,
    withdrawable     = withdrawable - $3,
    updated_at       = $4,
    updated_at_block = $5
WHERE staker_id = $1
  AND asset_id  = $2
  AND withdrawable >= $3
  AND deposited >= $3;`

	res, err := db.SQL.Exec(stmt,
		stakerID,
		assetID,
		claimAmount,
		updatedAt,
		blockHeight,
	)
	if err != nil {
		return fmt.Errorf("failed to claim bootstrap staker asset: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check claim rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("claim failed: either record does not exist or insufficient balance")
	}

	return nil
}

func (db *Db) SaveBootstrapDelegationState(d *types.BootstrapDelegationState) error {
	stmt := `
INSERT INTO bootstrap_delegation_states (
    staker_id, asset_id, operator_addr, delegated, updated_at, updated_at_block
) VALUES ($1,$2,$3,$4,$5,$6)
ON CONFLICT (staker_id, asset_id, operator_addr) DO UPDATE
SET delegated        = EXCLUDED.delegated,
    updated_at       = EXCLUDED.updated_at,
    updated_at_block = EXCLUDED.updated_at_block;`

	_, err := db.SQL.Exec(stmt,
		d.StakerID,
		d.AssetID,
		d.OperatorAddr,
		d.Delegated,
		d.UpdatedAt,
		d.UpdatedAtBlock,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap delegation state: %w", err)
	}
	return nil
}

// SaveBootstrapDelegationStateInTx saves a bootstrap delegation state within an existing transaction
func (db *Db) SaveBootstrapDelegationStateInTx(tx *sql.Tx, d *types.BootstrapDelegationState) (retErr error) {
	var ownTx bool
	if tx == nil {
		var err error
		tx, err = db.SQL.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		ownTx = true
		defer func() {
			if ownTx {
				if rbErr := tx.Rollback(); rbErr != nil {
					retErr = errors.Join(retErr, fmt.Errorf("failed to rollback transaction: %w", rbErr))
				}
			}
		}()
	}
	stmt := `
INSERT INTO bootstrap_delegation_states (
    staker_id, asset_id, operator_addr, delegated, updated_at, updated_at_block
) VALUES ($1,$2,$3,$4,$5,$6)
ON CONFLICT (staker_id, asset_id, operator_addr) DO UPDATE
SET delegated        = bootstrap_delegation_states.delegated + EXCLUDED.delegated,
    updated_at       = EXCLUDED.updated_at,
    updated_at_block = EXCLUDED.updated_at_block;`

	_, err := tx.Exec(stmt,
		d.StakerID,
		d.AssetID,
		d.OperatorAddr,
		d.Delegated,
		d.UpdatedAt,
		d.UpdatedAtBlock,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap delegation state: %w", err)
	}

	if !ownTx {
		return nil
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	ownTx = false
	return nil
}

func (db *Db) BootstrapDelegationExists(stakerID, assetID, operatorAddr string) (bool, error) {
	stmt := `
SELECT EXISTS(
    SELECT 1
    FROM bootstrap_delegation_states
    WHERE staker_id = $1 AND asset_id = $2 AND operator_addr = $3
);`

	var exists bool
	err := db.SQL.QueryRow(stmt, stakerID, assetID, operatorAddr).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check bootstrap delegation state existence: %w", err)
	}
	return exists, nil
}

func (db *Db) SaveBootstrapOperatorAsset(o *types.BootstrapOperatorAsset) error {
	tx, err := db.SQL.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err := db.SaveBootstrapOperatorAssetInTx(tx, o); err != nil {
		return err
	}

	return tx.Commit()
}

// SaveBootstrapOperatorAssetInTx saves a bootstrap operator asset within an existing transaction
func (db *Db) SaveBootstrapOperatorAssetInTx(tx *sql.Tx, o *types.BootstrapOperatorAsset) error {
	stmt := `
INSERT INTO bootstrap_operator_assets (
    operator_addr, asset_id, total_amount, self_amount, other_amount, updated_at
) VALUES ($1,$2,$3,$4,$5,$6)
ON CONFLICT (operator_addr, asset_id) DO UPDATE
SET total_amount = EXCLUDED.total_amount,
    self_amount  = EXCLUDED.self_amount,
    other_amount = EXCLUDED.other_amount,
    updated_at   = EXCLUDED.updated_at;`

	_, err := tx.Exec(stmt,
		o.OperatorAddr,
		o.AssetID,
		o.TotalAmount,
		o.SelfAmount,
		o.OtherAmount,
		o.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap operator asset: %w", err)
	}
	return nil
}

func (db *Db) GetBootstrapOperatorAsset(operatorAddr, assetID string) (*types.BootstrapOperatorAsset, error) {
	stmt := `
SELECT operator_addr, asset_id, total_amount, self_amount, other_amount, updated_at
FROM bootstrap_operator_assets
WHERE operator_addr = $1 AND asset_id = $2
LIMIT 1;`

	var o types.BootstrapOperatorAsset
	err := db.SQL.QueryRow(stmt, operatorAddr, assetID).Scan(
		&o.OperatorAddr,
		&o.AssetID,
		&o.TotalAmount,
		&o.SelfAmount,
		&o.OtherAmount,
		&o.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no bootstrap operator asset found for operator %s and asset %s", operatorAddr, assetID)
		}
		return nil, fmt.Errorf("failed to get bootstrap operator asset: %w", err)
	}

	return &o, nil
}

func (db *Db) UpdateBootstrapOperatorAssetInTx(tx *sql.Tx, assetID, operatorAddr, stakerIMAddr string, delegationDelta string) error {
	operatorAsset, err := db.GetBootstrapOperatorAsset(operatorAddr, assetID)
	if err != nil {
		operatorAsset = &types.BootstrapOperatorAsset{
			AssetID:      assetID,
			OperatorAddr: operatorAddr,
			TotalAmount:  sdkmath.ZeroInt().String(),
			SelfAmount:   sdkmath.ZeroInt().String(),
			OtherAmount:  sdkmath.ZeroInt().String(),
		}
	}
	totalAmount, ok := sdkmath.NewIntFromString(operatorAsset.TotalAmount)
	if !ok {
		return fmt.Errorf("UpdateBootstrapOperatorAssetInTx: failed to parse total amount from string:%s", operatorAsset.TotalAmount)
	}
	selfAmount, ok := sdkmath.NewIntFromString(operatorAsset.SelfAmount)
	if !ok {
		return fmt.Errorf("UpdateBootstrapOperatorAssetInTx: failed to parse self amount from string:%s", operatorAsset.SelfAmount)
	}
	otherAmount, ok := sdkmath.NewIntFromString(operatorAsset.OtherAmount)
	if !ok {
		return fmt.Errorf("UpdateBootstrapOperatorAssetInTx: failed to parse other amount from string:%s", operatorAsset.OtherAmount)
	}
	deltaAmount, ok := sdkmath.NewIntFromString(delegationDelta)
	if !ok {
		return fmt.Errorf("UpdateBootstrapOperatorAssetInTx: failed to parse delegationDelta staking amount from string:%s", delegationDelta)
	}

	totalAmount = totalAmount.Add(deltaAmount)
	operatorAccAddr, err := sdk.AccAddressFromBech32(operatorAddr)
	if err != nil {
		return err
	}
	stakerEVMAddr := common.HexToAddress(stakerIMAddr)
	if stakerEVMAddr == common.Address(operatorAccAddr) {
		selfAmount = selfAmount.Add(deltaAmount)
	} else {
		otherAmount = otherAmount.Add(deltaAmount)
	}

	var ownTx bool
	var retErr error
	if tx == nil {
		var err error
		tx, err = db.SQL.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		ownTx = true
		defer func() {
			if ownTx {
				if rbErr := tx.Rollback(); rbErr != nil {
					retErr = errors.Join(retErr, fmt.Errorf("failed to rollback transaction: %w", rbErr))
				}
			}
		}()
	}
	err = db.SaveBootstrapOperatorAssetInTx(tx, &types.BootstrapOperatorAsset{
		AssetID:      assetID,
		OperatorAddr: operatorAddr,
		TotalAmount:  totalAmount.String(),
		SelfAmount:   selfAmount.String(),
		OtherAmount:  otherAmount.String(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return err
	}

	if !ownTx {
		return nil
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	ownTx = false
	return retErr
}

// Incremental scanning state management functions
// GetScanState retrieves the scanning state for a specific chain type
func (db *Db) GetScanState(chainType string) (*types.ScanState, error) {
	stmt := `SELECT chain_type, last_height, last_hash, safe_height, updated_at, created_at
             FROM bootstrap_scan_state WHERE chain_type = $1`

	var state types.ScanState
	err := db.SQL.QueryRow(stmt, chainType).Scan(
		&state.ChainType,
		&state.LastHeight,
		&state.LastHash,
		&state.SafeHeight,
		&state.UpdatedAt,
		&state.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // no row yet — expected on first run or after DB reset
		}
		return nil, fmt.Errorf("failed to get scan state for %s: %w", chainType, err)
	}

	return &state, nil
}

// UpdateScanState updates or inserts the scanning state for a specific chain type
func (db *Db) UpdateScanState(state *types.ScanState) error {
	stmt := `
INSERT INTO bootstrap_scan_state (chain_type, last_height, last_hash, safe_height, updated_at)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (chain_type) DO UPDATE
SET last_height = EXCLUDED.last_height,
    last_hash   = EXCLUDED.last_hash,
    safe_height = EXCLUDED.safe_height,
    updated_at  = EXCLUDED.updated_at`

	_, err := db.SQL.Exec(stmt,
		state.ChainType,
		state.LastHeight,
		state.LastHash,
		state.SafeHeight,
		state.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update scan state for %s: %w", state.ChainType, err)
	}

	return nil
}

// IsTransactionProcessed checks if a transaction has already been processed
func (db *Db) IsTransactionProcessed(chainType, txHash string) (bool, error) {
	stmt := `SELECT EXISTS(SELECT 1 FROM bootstrap_processed_transactions
                          WHERE chain_type = $1 AND tx_hash = $2)`

	var exists bool
	err := db.SQL.QueryRow(stmt, chainType, txHash).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if transaction is processed: %w", err)
	}

	return exists, nil
}

// MarkTransactionProcessed marks a transaction as processed
func (db *Db) MarkTransactionProcessed(chainType, txHash string, blockHeight int64) error {
	stmt := `INSERT INTO bootstrap_processed_transactions (chain_type, tx_hash, block_height)
             VALUES ($1, $2, $3)
             ON CONFLICT (chain_type, tx_hash) DO NOTHING`

	_, err := db.SQL.Exec(stmt, chainType, txHash, blockHeight)
	if err != nil {
		return fmt.Errorf("failed to mark transaction as processed: %w", err)
	}

	return nil
}

// Address binding management functions
// GetAddressBindings retrieves all address bindings for a specific chain type
func (db *Db) GetAddressBindings(chainType string) ([]types.AddressBinding, error) {
	stmt := `SELECT chain_type, source_addr, target_addr, created_at, updated_at
             FROM bootstrap_address_bindings WHERE chain_type = $1
             ORDER BY created_at ASC`

	rows, err := db.SQL.Query(stmt, chainType)
	if err != nil {
		return nil, fmt.Errorf("failed to get address bindings for %s: %w", chainType, err)
	}
	defer rows.Close()

	var bindings []types.AddressBinding
	for rows.Next() {
		var binding types.AddressBinding
		err := rows.Scan(
			&binding.ChainType,
			&binding.SourceAddr,
			&binding.TargetAddr,
			&binding.CreatedAt,
			&binding.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan address binding: %w", err)
		}
		bindings = append(bindings, binding)
	}

	return bindings, nil
}

// SaveAddressBinding saves or updates an address binding
func (db *Db) SaveAddressBinding(binding *types.AddressBinding) error {
	stmt := `
INSERT INTO bootstrap_address_bindings (chain_type, source_addr, target_addr, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (chain_type, source_addr) DO UPDATE
SET target_addr = EXCLUDED.target_addr,
    updated_at  = EXCLUDED.updated_at`

	_, err := db.SQL.Exec(
		stmt,
		binding.ChainType,
		binding.SourceAddr,
		binding.TargetAddr,
		binding.CreatedAt,
		binding.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save address binding: %w", err)
	}

	return nil
}

// GetAddressBinding retrieves a specific address binding
func (db *Db) GetAddressBinding(chainType, sourceAddr string) (*types.AddressBinding, error) {
	stmt := `SELECT chain_type, source_addr, target_addr, created_at, updated_at
             FROM bootstrap_address_bindings
             WHERE chain_type = $1 AND source_addr = $2`

	var binding types.AddressBinding
	err := db.SQL.QueryRow(stmt, chainType, sourceAddr).Scan(
		&binding.ChainType,
		&binding.SourceAddr,
		&binding.TargetAddr,
		&binding.CreatedAt,
		&binding.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found, not an error
		}
		return nil, fmt.Errorf("failed to get address binding for %s/%s: %w", chainType, sourceAddr, err)
	}

	return &binding, nil
}

// CheckTargetAddressBinding checks if a target address is already bound to a different source address
func (db *Db) CheckTargetAddressBinding(chainType, targetAddr, excludeSourceAddr string) (*types.AddressBinding, error) {
	stmt := `SELECT chain_type, source_addr, target_addr, created_at, updated_at
             FROM bootstrap_address_bindings
             WHERE chain_type = $1 AND target_addr = $2 AND source_addr != $3`

	var binding types.AddressBinding
	err := db.SQL.QueryRow(stmt, chainType, targetAddr, excludeSourceAddr).Scan(
		&binding.ChainType,
		&binding.SourceAddr,
		&binding.TargetAddr,
		&binding.CreatedAt,
		&binding.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found, not an error
		}
		return nil, fmt.Errorf("failed to check target address binding for %s/%s: %w", chainType, targetAddr, err)
	}

	return &binding, nil
}

// DeleteAddressBinding removes an address binding
func (db *Db) DeleteAddressBinding(chainType, sourceAddr string) error {
	stmt := `DELETE FROM bootstrap_address_bindings WHERE chain_type = $1 AND source_addr = $2`

	result, err := db.SQL.Exec(stmt, chainType, sourceAddr)
	if err != nil {
		return fmt.Errorf("failed to delete address binding: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check deletion result: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("address binding not found for %s/%s", chainType, sourceAddr)
	}

	return nil
}

// Transaction support methods

// WithTransaction executes a function within a database transaction
func (db *Db) WithTransaction(fn func(tx *sql.Tx) error) error {
	tx, err := db.SQL.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("failed to rollback transaction after error %v: %w", err, rollbackErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// MarkTransactionProcessedInTx marks a transaction as processed within an existing transaction
func (db *Db) MarkTransactionProcessedInTx(tx *sql.Tx, chainType, txHash string, blockHeight int64) error {
	stmt := `
INSERT INTO bootstrap_processed_transactions (chain_type, tx_hash, block_height, processed_at)
VALUES ($1, $2, $3, NOW())
ON CONFLICT (chain_type, tx_hash) DO NOTHING`

	_, err := tx.Exec(stmt, chainType, txHash, blockHeight)
	if err != nil {
		return fmt.Errorf("failed to mark transaction as processed: %w", err)
	}

	return nil
}

// GetLastProcessedTransaction gets the last processed transaction ID for a specific chain type
// Used for pagination when fetching transactions from address API
func (db *Db) GetLastProcessedTransaction(chainType string) (string, int64, error) {
	stmt := `SELECT tx_hash, block_height FROM bootstrap_processed_transactions
             WHERE chain_type = $1
             ORDER BY block_height DESC, processed_at DESC
             LIMIT 1`

	var txHash string
	var blockHeight sql.NullInt64
	err := db.SQL.QueryRow(stmt, chainType).Scan(&txHash, &blockHeight)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", 0, nil // No transactions processed yet
		}
		return "", 0, fmt.Errorf("failed to get last processed transaction for %s: %w", chainType, err)
	}

	return txHash, blockHeight.Int64, nil
}

func (db *Db) OperatorAssetExists(operatorAddr string, assetID string) (bool, error) {
	stmt := `
SELECT EXISTS (
    SELECT 1
    FROM bootstrap_operator_assets
    WHERE operator_addr = $1 AND asset_id = $2
);`

	var exists bool
	err := db.SQL.QueryRow(stmt, operatorAddr, assetID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check bootstrap operator asset existence: %w", err)
	}
	return exists, nil
}

func (db *Db) SaveBootstrapTokenPrice(assetID string, price string) error {
	stmt := `
INSERT INTO bootstrap_token_prices (asset_id, price, updated_at)
VALUES ($1, $2, now())
ON CONFLICT (asset_id) DO UPDATE
SET price      = EXCLUDED.price,
    updated_at = EXCLUDED.updated_at;
`

	_, err := db.SQL.Exec(stmt,
		assetID,
		price,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap token price: %w", err)
	}
	return nil
}

func (db *Db) GetBootstrapTokenPrice(assetID string) (string, error) {
	var price string

	stmt := `
SELECT price
FROM bootstrap_token_prices
WHERE asset_id = $1
LIMIT 1;
`

	err := db.SQL.QueryRow(stmt, assetID).Scan(&price)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no price found for asset_id %s", assetID)
		}
		return "", fmt.Errorf("failed to get bootstrap token price: %w", err)
	}

	return price, nil
}

func (db *Db) SaveBootstrapStatistics(tvl string) error {
	stmt := `
INSERT INTO bootstrap_statistics (one_row_id, tvl, updated_at)
VALUES (TRUE, $1, now())
ON CONFLICT (one_row_id) DO UPDATE
SET tvl        = EXCLUDED.tvl,
    updated_at = EXCLUDED.updated_at;
`

	_, err := db.SQL.Exec(stmt,
		tvl,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap statistics: %w", err)
	}
	return nil
}
