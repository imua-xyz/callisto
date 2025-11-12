package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"

	"github.com/forbole/callisto/v4/types"
	"github.com/rs/zerolog/log"
)

// SaveAssetsParams allows to store the given params inside the database
func (db *Db) SaveAssetsParams(params *types.AssetsParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling assets params: %s", err)
	}

	stmt := `
INSERT INTO assets_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
        height = excluded.height
WHERE assets_params.height <= excluded.height`

	_, err = db.SQL.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing assets params: %s", err)
	}

	return nil
}

// SaveClientChain inserts or updates a client chain record in the database
func (db *Db) SaveOrUpdateClientChain(chain *types.ClientChain) error {
	stmt := `
INSERT INTO client_chains (name, meta_info, chain_id, imuachain_index, finalization_blocks, layer_zero_chain_id, signature_type, address_length)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (layer_zero_chain_id) DO UPDATE
	SET name = EXCLUDED.name,
		meta_info = EXCLUDED.meta_info,
		chain_id = EXCLUDED.chain_id,
		finalization_blocks = EXCLUDED.finalization_blocks,
		layer_zero_chain_id = EXCLUDED.layer_zero_chain_id,
		signature_type = EXCLUDED.signature_type,
		address_length = EXCLUDED.address_length;`

	_, err := db.SQL.Exec(stmt,
		chain.Name,
		chain.MetaInfo,
		chain.ChainId,
		chain.ImuaChainIndex,
		chain.FinalizationBlocks,
		chain.LayerZeroChainID,
		chain.SignatureType,
		chain.AddressLength,
	)
	if err != nil {
		return fmt.Errorf("failed to save client chain: %w", err)
	}
	return nil
}

// SaveAssetsToken saves a token record into the database. Once added, only the
// metadata may be altered by the blockchain.
func (db *Db) SaveAssetsToken(token *types.AssetsToken) error {
	// Q. drop the total deposit amount or retain?
	// A. slashing is applied to staker level, not the deposit amount.
	//    so it is a good thing to retain the total deposit amount.
	stmt := `
INSERT INTO assets_tokens (asset_id, name, symbol, address, decimals, layer_zero_chain_id, imuachain_index, meta_info, staking_total_amount)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT (asset_id) DO UPDATE
SET name = EXCLUDED.name,
    symbol = EXCLUDED.symbol,
    address = EXCLUDED.address,
    decimals = EXCLUDED.decimals,
    layer_zero_chain_id = EXCLUDED.layer_zero_chain_id,
    imuachain_index = EXCLUDED.imuachain_index,
    meta_info = EXCLUDED.meta_info,
	staking_total_amount = EXCLUDED.staking_total_amount;`
	_, err := db.SQL.Exec(stmt,
		token.AssetID,
		token.Name,
		token.Symbol,
		token.Address,
		token.Decimals,
		token.LayerZeroChainID,
		token.ImuaChainIndex,
		token.MetaInfo,
		token.Amount,
	)
	if err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}
	log.Debug().Msgf("saved token %s", token.AssetID)
	return nil
}

// GetTokenDecimalsByID retrieves the decimals of a token by its asset ID.
func (db *Db) GetTokenDecimalsByID(assetID string) (int, error) {
	stmt := `
SELECT decimals
FROM assets_tokens
WHERE asset_id = $1;`

	var decimals int
	err := db.SQL.QueryRow(stmt, assetID).Scan(&decimals)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("no decimals found for asset_id %s", assetID)
		}
		return 0, fmt.Errorf("failed to retrieve decimals for asset_id %s: %w", assetID, err)
	}

	return decimals, nil
}

// UpdateAssetMetadata updates the metadata for an asset based on its assetID.
func (db *Db) UpdateAssetMetadata(
	assetID string, newMetaInfo string,
) error {
	stmt := `
UPDATE tokens
SET meta_info = $1
WHERE asset_id = $2;`
	// no conflict can happen above since it is update
	// we don't check for existence because again, we are not a business logic implementation
	_, err := db.SQL.Exec(stmt, newMetaInfo, assetID)
	if err != nil {
		return fmt.Errorf(
			"failed to update metadata for asset_id %s: %w",
			assetID, err,
		)
	}
	return nil
}

// UpdateStakingTotalAmount updates the total staking amount for an asset based on its assetID.
func (db *Db) UpdateStakingTotalAmount(
	assetID string, newAmount string,
) error {
	stmt := `
UPDATE tokens
SET staking_total_amount = $1
WHERE asset_id = $2;`
	_, err := db.SQL.Exec(stmt, newAmount, assetID)
	if err != nil {
		return fmt.Errorf(
			"failed to update staking total amount for asset_id %s: %w",
			assetID, err,
		)
	}
	return nil
}

// SaveStakerAsset saves a staker asset record in the database.
func (db *Db) SaveStakerAsset(data *types.StakerAsset) error { // No lastUpdatedHeight here
	// The genesis_deposit is only set on insert and left unchanged on update to
	// avoid being affected by deposits after mainnet launch.
	stmt := `
INSERT INTO staker_assets (
    staker_id, asset_id,
    deposited, withdrawable, pending_undelegation,
    delegated, lifetime_slashed, genesis_deposit
) 
VALUES (
    $1, $2,
    $3, $4, $5,
    ($3::numeric - $4::numeric - $5::numeric),
    0, $6
)
ON CONFLICT (staker_id, asset_id) DO UPDATE
SET deposited = EXCLUDED.deposited, 
    withdrawable = EXCLUDED.withdrawable,
    pending_undelegation = EXCLUDED.pending_undelegation,
    delegated = EXCLUDED.deposited - EXCLUDED.withdrawable - EXCLUDED.pending_undelegation - staker_assets.lifetime_slashed,
    lifetime_slashed = staker_assets.lifetime_slashed;`

	_, err := db.SQL.Exec(
		stmt,
		data.StakerID,            // $1
		data.AssetID,             // $2
		data.Deposited,           // $3
		data.Withdrawable,        // $4
		data.PendingUndelegation, // $5
		data.GenesisDeposited,    // $6
		// Only 5 arguments passed to Exec
	)
	if err != nil {
		return fmt.Errorf("failed to save staker asset: %w", err)
	}
	return nil
}

func (db *Db) IterateAirdropStakerAssets(opFunc func(sa types.ParsedStakerAsset) error) error {
	stmt := `
	SELECT staker_id, asset_id, deposited, genesis_deposit
	FROM staker_assets
	WHERE genesis_deposit > 0 AND deposited >= genesis_deposit
	ORDER BY staker_id, asset_id;`

	rows, err := db.SQL.Query(stmt)
	if err != nil {
		return fmt.Errorf("failed to query genesis deposits: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sa types.ParsedStakerAsset
		var depositedStr, genesisDepositedStr string
		err := rows.Scan(
			&sa.StakerID,
			&sa.AssetID,
			&depositedStr,
			&genesisDepositedStr,
		)
		if err != nil {
			return fmt.Errorf("failed to scan staker asset: %w", err)
		}

		depositedInt, ok := sdkmath.NewIntFromString(depositedStr)
		if !ok {
			return fmt.Errorf("invalid deposited: %s", depositedStr)
		}
		genesisDepositInt, ok := sdkmath.NewIntFromString(genesisDepositedStr)
		if !ok {
			return fmt.Errorf("invalid genesis deposit: %s", genesisDepositedStr)
		}
		sa.Deposited = depositedInt
		sa.GenesisDeposited = genesisDepositInt

		if err := opFunc(sa); err != nil {
			return fmt.Errorf("opFunc failed for staker_id %s, asset_id %s: %w", sa.StakerID, sa.AssetID, err)
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("row iteration error: %w", err)
	}

	return nil
}

// GetDelegatedAmount returns the delegated amount for a given staker and asset.
func (db *Db) GetDelegatedAmount(stakerID, assetID string) (sdkmath.Int, error) {
	stmt := `
	SELECT delegated FROM staker_assets WHERE staker_id = $1 AND asset_id = $2;`
	var delegatedAmount string
	err := db.SQL.QueryRow(stmt, stakerID, assetID).Scan(&delegatedAmount)
	if err != nil {
		return sdkmath.Int{}, fmt.Errorf("failed to get delegated amount: %w", err)
	}
	delegatedAmountInt, ok := sdkmath.NewIntFromString(delegatedAmount)
	if !ok {
		return sdkmath.Int{}, fmt.Errorf("failed to convert delegated amount to int: %s", delegatedAmount)
	}
	return delegatedAmountInt, nil
}

// SlashStakerDelegation slashes the staker delegation. It updates the lifetime slashed amount
// and the delegated amount.
func (db *Db) SlashStakerDelegation(stakerID, assetID, slashedAmount string) error {
	stmt := `
	UPDATE staker_assets
	SET lifetime_slashed = lifetime_slashed + $1,
	    delegated = delegated - $1
	WHERE staker_id = $2 AND asset_id = $3;`
	_, err := db.SQL.Exec(stmt, slashedAmount, stakerID, assetID)
	if err != nil {
		return fmt.Errorf("failed to accumulate staker lifetime slashing: %w", err)
	}
	return nil
}

// SaveOperatorAsset saves an operator asset record in the database,
// ensuring other_share is derived as total_share - self_share.
// This function is in `assets.go` because it is triggered by events in `x/assets`.
func (db *Db) SaveOperatorAsset(data *types.OperatorAsset) error {
	stmt := `
INSERT INTO operator_assets (
    operator_addr, asset_id,
    total_amount, pending_undelegation_amount,
    total_share, self_share,
    other_share
)
VALUES (
    $1, $2, $3, $4, $5, $6, 
    ($5::numeric - $6::numeric) 
)
ON CONFLICT (operator_addr, asset_id) DO UPDATE
SET total_amount = EXCLUDED.total_amount,
    pending_undelegation_amount = EXCLUDED.pending_undelegation_amount,
    total_share = EXCLUDED.total_share,
    self_share = EXCLUDED.self_share,
    other_share = (EXCLUDED.total_share - EXCLUDED.self_share);`

	_, err := db.SQL.Exec(
		stmt,
		data.OperatorAddress,           // $1
		data.AssetID,                   // $2
		data.TotalAmount,               // $3 (string, PG converts to NUMERIC for column)
		data.PendingUndelegationAmount, // $4 (string, PG converts to NUMERIC for column)
		data.TotalShare,                // $5 (string, used in $5::numeric)
		data.SelfShare,                 // $6 (string, used in $6::numeric)
	)
	if err != nil {
		// It's useful to log the inputs if an error occurs
		return fmt.Errorf("failed to save operator asset (operator: %s, asset: %s, total_share: %s, self_share: %s): %w",
			data.OperatorAddress, data.AssetID, data.TotalShare, data.SelfShare, err)
	}
	return nil
}
