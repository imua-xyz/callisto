package database

import (
	"fmt"

	"github.com/forbole/callisto/v4/types"
	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"
)

// SaveOperatorDetail saves the operator details into the database
func (db *Db) SaveStakerOperatorAssociation(stakerID, operatorAddr string) error {
	stmt := `
	INSERT INTO staker_operator_associations (staker_id, operator_addr)
	VALUES ($1, $2)
	ON CONFLICT (staker_id) DO UPDATE
	SET operator_addr = EXCLUDED.operator_addr;`
	_, err := db.SQL.Exec(stmt, stakerID, operatorAddr)
	if err != nil {
		return fmt.Errorf("failed to save staker operator association: %w", err)
	}
	return nil
}

// DeleteStakerOperatorAssociation deletes the staker operator association from the database.
func (db *Db) DeleteStakerOperatorAssociation(stakerID string) error {
	stmt := `
	DELETE FROM staker_operator_associations WHERE staker_id = $1;`
	_, err := db.SQL.Exec(stmt, stakerID)
	if err != nil {
		return fmt.Errorf("failed to delete staker operator association: %w", err)
	}
	return nil
}

// SaveDelegationState saves the delegation state into the database
func (db *Db) SaveDelegationState(state *types.DelegationState) error {
	stmt := `
	INSERT INTO delegation_states (staker_id, asset_id, operator_addr, undelegatable_share, wait_undelegation_amount)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (staker_id, asset_id, operator_addr) DO UPDATE
	SET undelegatable_share = EXCLUDED.undelegatable_share + delegation_states.undelegatable_share,
	    wait_undelegation_amount = EXCLUDED.wait_undelegation_amount + delegation_states.wait_undelegation_amount;`
	_, err := db.SQL.Exec(
		stmt,
		state.StakerID,
		state.AssetID,
		state.OperatorAddr,
		state.UndelegatableShare,
		state.WaitUndelegationAmount,
	)
	if err != nil {
		return fmt.Errorf("failed to save delegation state: %w", err)
	}
	return nil
}

// GetAllStakersFromDelegationStates returns all unique staker IDs from the delegation_states table.
func (db *Db) GetAllStakersFromDelegationStates() ([]string, error) {
	stmt := `
	SELECT DISTINCT staker_id 
	FROM delegation_states;`

	rows, err := db.SQL.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("failed to query all staker_ids from delegation_states: %w", err)
	}
	defer rows.Close()

	var stakerIDs []string
	for rows.Next() {
		var stakerID string
		if err := rows.Scan(&stakerID); err != nil {
			return nil, fmt.Errorf("failed to scan staker_id: %w", err)
		}
		stakerIDs = append(stakerIDs, stakerID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return stakerIDs, nil
}

// AppendStakerToOperatorAsset appends the staker to the operator + asset combination.
// It does nothing if the staker is already in the database.
func (db *Db) AppendStakerToOperatorAsset(stakerID, operatorAddr, assetID string) error {
	stmt := `
	INSERT INTO operator_asset_stakers (operator_addr, asset_id, staker_id)
	VALUES ($1, $2, $3)
	ON CONFLICT (operator_addr, asset_id, staker_id) DO NOTHING;`
	_, err := db.SQL.Exec(stmt, operatorAddr, assetID, stakerID)
	if err != nil {
		return fmt.Errorf("failed to append staker to operator asset: %w", err)
	}
	return nil
}

// GetStakersByOperatorAsset gets the stakers for the operator + asset combination.
func (db *Db) GetStakersByOperatorAsset(operatorAddr, assetID string) ([]string, error) {
	stmt := `
	SELECT staker_id 
	FROM operator_asset_stakers 
	WHERE operator_addr = $1 AND asset_id = $2;`

	rows, err := db.SQL.Query(stmt, operatorAddr, assetID)
	if err != nil {
		return nil, fmt.Errorf("failed to query stakers by operator and asset: %w", err)
	}
	defer rows.Close()

	var stakerIDs []string
	for rows.Next() {
		var stakerID string
		if err := rows.Scan(&stakerID); err != nil {
			return nil, fmt.Errorf("failed to scan staker_id: %w", err)
		}
		stakerIDs = append(stakerIDs, stakerID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return stakerIDs, nil
}

// RemoveStakerFromOperatorAsset removes the staker from the operator + asset combination.
// It does nothing if the staker is not in the database.
func (db *Db) RemoveStakerFromOperatorAsset(stakerID, operatorAddr, assetID string) error {
	stmt := `
	DELETE FROM operator_asset_stakers WHERE operator_addr = $1 AND asset_id = $2 AND staker_id = $3;`
	_, err := db.SQL.Exec(stmt, operatorAddr, assetID, stakerID)
	if err != nil {
		return fmt.Errorf("failed to remove staker from operator asset: %w", err)
	}
	return nil
}

// DeleteAllStakersFromOperatorAsset deletes all stakers from the operator + asset combination.
func (db *Db) DeleteAllStakersFromOperatorAsset(operatorAddr, assetID string) error {
	stmt := `
	DELETE FROM operator_asset_stakers WHERE operator_addr = $1 AND asset_id = $2;`
	_, err := db.SQL.Exec(stmt, operatorAddr, assetID)
	if err != nil {
		return fmt.Errorf("failed to delete all stakers from operator asset: %w", err)
	}
	return nil
}

// SaveUndelegationRecord saves the undelegation record into the database. In case of a conflict,
// it updates the actual completed amount and the hold count.
func (db *Db) SaveUndelegationRecord(record *types.UndelegationRecord) error {
	stmt := `
    INSERT INTO undelegation_records (
        record_id, staker_id, asset_id, operator_addr, tx_hash, 
        block_number, completed_epoch_identifier, completed_epoch_number, 
        undelegation_id, amount, actual_completed_amount, hold_count
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
    ON CONFLICT (record_id) DO UPDATE
    SET 
        actual_completed_amount = EXCLUDED.actual_completed_amount,
        hold_count = EXCLUDED.hold_count;`

	_, err := db.SQL.Exec(stmt,
		record.RecordID, record.StakerID, record.AssetID, record.OperatorAddr,
		record.TxHash, record.BlockNumber, record.CompletedEpochIdentifier,
		record.CompletedEpochNumber, record.UndelegationID, record.Amount,
		record.ActualCompletedAmount, record.HoldCount)
	if err != nil {
		return fmt.Errorf("failed to save undelegation record: %w", err)
	}
	return nil
}

// UpdateUndelegationRecordHoldCount updates the hold count of the undelegation record.
func (db *Db) UpdateUndelegationRecordHoldCount(recordID, holdCount string) error {
	stmt := `
	UPDATE undelegation_records SET hold_count = $1 WHERE record_id = $2;`
	_, err := db.SQL.Exec(stmt, holdCount, recordID)
	if err != nil {
		return fmt.Errorf("failed to update undelegation record hold count: %w", err)
	}
	return nil
}

// MatureUndelegationRecord matures the undelegation record. It sets the height of maturity
// and the actual completed amount. The actual completed amount may be different from the amount
// that was undelegated, as it may include some slashing impact. Remember that slashing is
// applied first to the undelegation, and then to the delegation.
func (db *Db) MatureUndelegationRecord(recordID string, amount string, height int64) error {
	stmt := `
	UPDATE undelegation_records
	SET maturity_height = $1,
		actual_completed_amount = $2
	WHERE record_id = $3;`
	_, err := db.SQL.Exec(stmt, height, amount, recordID)
	if err != nil {
		return fmt.Errorf("failed to mature undelegation record: %w", err)
	}
	return nil
}

// SlashUndelegationRecord slashes the undelegation record. It updates the actual completed amount
func (db *Db) SlashUndelegationRecord(recordID, postSlashingAmount, slashedAmount string) error {
	stmt := `
	UPDATE undelegation_records
	SET actual_completed_amount = $1
	WHERE record_id = $2;`
	_, err := db.SQL.Exec(stmt, postSlashingAmount, recordID)
	if err != nil {
		return fmt.Errorf("failed to slash undelegation record: %w", err)
	}
	// update the lifetime slashed amount for the staker asset - this is only from undelegation
	// slashing, not from active delegation slashing.
	stakerID, assetID, err := db.GetStakerIDAssetIDFromUndelegationRecord(recordID)
	if err != nil {
		return fmt.Errorf("failed to get staker ID and asset ID from undelegation record: %w", err)
	}
	if assetID == assetstypes.ImuachainAssetID {
		stmt = `
		UPDATE im_asset_delegations
		SET lifetime_slashed = lifetime_slashed + $1,
			pending_undelegation = pending_undelegation - $1
		WHERE staker_id = $2;`
		_, err = db.SQL.Exec(stmt, slashedAmount, stakerID)
		if err != nil {
			return fmt.Errorf("failed to update im asset delegation lifetime slashed: %w", err)
		}
	} else {
		stmt = `
		UPDATE staker_assets
		SET lifetime_slashed = lifetime_slashed + $1,
			pending_undelegation = pending_undelegation - $1
		WHERE staker_id = $2 AND asset_id = $3;`
		_, err = db.SQL.Exec(stmt, slashedAmount, stakerID, assetID)
		if err != nil {
			return fmt.Errorf("failed to update staker asset lifetime slashed: %w", err)
		}
	}
	return nil
}

// GetStakerIDAssetIDFromUndelegationRecord gets the staker ID and asset ID from the undelegation record.
func (db *Db) GetStakerIDAssetIDFromUndelegationRecord(recordID string) (string, string, error) {
	stmt := `
	SELECT staker_id, asset_id FROM undelegation_records WHERE record_id = $1;`
	var stakerID, assetID string
	err := db.SQL.QueryRow(stmt, recordID).Scan(&stakerID, &assetID)
	return stakerID, assetID, err
}

// AccumulateImAssetDelegation accumulates the im asset delegation amounts into the database.
// It adds the new values to any existing values for delegated and pending_undelegation amounts.

func (db *Db) AccumulateImAssetDelegation(delegation *types.ImAssetDelegation) error {
	stmt := `
	INSERT INTO im_asset_delegations (staker_id, delegated, pending_undelegation)
	VALUES ($1, $2, $3)
	ON CONFLICT (staker_id) DO UPDATE
	SET delegated = im_asset_delegations.delegated + EXCLUDED.delegated,
		pending_undelegation = im_asset_delegations.pending_undelegation + EXCLUDED.pending_undelegation;`
	_, err := db.SQL.Exec(stmt,
		delegation.StakerID, delegation.Delegated, delegation.PendingUndelegation,
	)
	if err != nil {
		return fmt.Errorf("failed to accumulate im asset delegation: %w", err)
	}
	return nil
}

// SlashImAssetDelegation slashes the im asset delegation. It updates the lifetime slashed amount
// and the delegated amount.
func (db *Db) SlashImAssetDelegation(stakerID, amount string) error {
	stmt := `
	UPDATE im_asset_delegations
	SET lifetime_slashed = im_asset_delegations.lifetime_slashed + $2,
		delegated = im_asset_delegations.delegated - $2
	WHERE staker_id = $1;`
	_, err := db.SQL.Exec(stmt, stakerID, amount)
	if err != nil {
		return fmt.Errorf("failed to slash im asset delegation: %w", err)
	}
	return nil
}

// UndelegateImAsset undelegates an amount from the im asset delegation. As a result of this
// undelegation, the amount is added to the pending_undelegation and subtracted from the delegated amount.
func (db *Db) UndelegateImAsset(stakerID, amount string) error {
	stmt := `
	UPDATE im_asset_delegations
	SET pending_undelegation = im_asset_delegations.pending_undelegation + $2,
		delegated = im_asset_delegations.delegated - $3
	WHERE staker_id = $1;`
	_, err := db.SQL.Exec(stmt, stakerID, amount)
	if err != nil {
		return fmt.Errorf("failed to undelegate im asset: %w", err)
	}
	return nil
}

// MatureImAssetUndelegation matures the im asset undelegation. It updates the pending_undelegation
// amount.
func (db *Db) MatureImAssetUndelegation(stakerID, amount string) error {
	stmt := `
	UPDATE im_asset_delegations
	SET pending_undelegation = im_asset_delegations.pending_undelegation - $2
	WHERE staker_id = $1;`
	_, err := db.SQL.Exec(stmt, stakerID, amount)
	if err != nil {
		return fmt.Errorf("failed to mature im asset undelegation: %w", err)
	}
	return nil
}
