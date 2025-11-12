package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	dbtypes "github.com/forbole/callisto/v4/database/types"
	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"
	distrtypes "github.com/imua-xyz/imuachain/x/feedistribution/types"

	"github.com/forbole/callisto/v4/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lib/pq"
)

// SaveCommunityPool stores the community pool coins for a specific AVS at a given height.
// If a record for the AVS already exists, it updates it only if the new height is greater or equal.
func (db *Db) SaveCommunityPool(avsAddr string, coins sdk.DecCoins, height int64) error {
	query := `
INSERT INTO community_pool(avs_addr, coins, height)
VALUES ($1, $2, $3)
ON CONFLICT (avs_addr) DO UPDATE
    SET coins = excluded.coins,
        height = excluded.height
WHERE community_pool.height <= excluded.height;
`

	_, err := db.SQL.Exec(query, strings.ToLower(avsAddr), pq.Array(dbtypes.NewDbDecCoins(coins)), height)
	if err != nil {
		return fmt.Errorf("error while storing community pool for AVS %s: %w", avsAddr, err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveDistributionParams allows to store the given distribution parameters inside the database
func (db *Db) SaveDistributionParams(params *types.DistributionParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling params: %s", err)
	}

	stmt := `
INSERT INTO distribution_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
      	height = excluded.height
WHERE distribution_params.height <= excluded.height`
	_, err = db.SQL.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing distribution params: %s", err)
	}

	return nil
}

// SaveAVSRewardAsset saves or updates an AVS reward asset record in the database.
// reward_pool_withdrawn_total and reward_pool_debt are computed in SQL.
func (db *Db) SaveAVSRewardAsset(asset *types.AVSRewardAsset) error {
	stmt := `
INSERT INTO avs_reward_assets (
	avs_addr,
	asset_id,
	name,
	symbol,
	address,
	decimals,
	layer_zero_chain_id,
	imuachain_index,
	meta_info,
	reward_pool_balance,
	reward_pool_total,
	reward_allocation_total,
	reward_pool_withdrawn_total,
	reward_pool_debt
) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9,
	$10, $11, $12,
	($11::numeric - $10::numeric),  -- reward_pool_withdrawn_total
	($12::numeric - $11::numeric)   -- reward_pool_debt
)
ON CONFLICT (avs_addr, asset_id) DO UPDATE
SET name = EXCLUDED.name,
	symbol = EXCLUDED.symbol,
	address = EXCLUDED.address,
	decimals = EXCLUDED.decimals,
	layer_zero_chain_id = EXCLUDED.layer_zero_chain_id,
	imuachain_index = EXCLUDED.imuachain_index,
	meta_info = EXCLUDED.meta_info,
	reward_pool_balance = EXCLUDED.reward_pool_balance,
	reward_pool_total = EXCLUDED.reward_pool_total,
	reward_allocation_total = EXCLUDED.reward_allocation_total,
	reward_pool_withdrawn_total = EXCLUDED.reward_pool_total - EXCLUDED.reward_pool_balance,
	reward_pool_debt = EXCLUDED.reward_allocation_total - EXCLUDED.reward_pool_total;`

	_, err := db.SQL.Exec(stmt,
		asset.AVSAddr,  // $1
		asset.AssetID,  // $2
		asset.Name,     // $3
		asset.Symbol,   // $4
		asset.Address,  // $5
		asset.Decimals, // $6
		// #nosec G115
		int64(asset.LayerZeroChainID), // $7
		// #nosec G115
		int64(asset.ImuaChainIndex),          // $8
		asset.MetaInfo,                       // $9
		asset.RewardPoolBalance.String(),     // $10
		asset.RewardPoolTotal.String(),       // $11
		asset.RewardAllocationTotal.String(), // $12
	)
	if err != nil {
		return fmt.Errorf("failed to save avs reward asset: %w", err)
	}
	return nil
}

// GetAVSAssetInfo retrieves the asset information of a given AVS reward asset.
func (db *Db) GetAVSAssetInfo(avsAddr string, assetID string) (*assetstypes.AssetInfo, error) {
	stmt := `
SELECT
	name,
	symbol,
	address,
	decimals,
	layer_zero_chain_id,
	imuachain_index,
	meta_info
FROM avs_reward_assets
WHERE avs_addr = $1 AND asset_id = $2
LIMIT 1;`

	row := db.SQL.QueryRow(stmt, avsAddr, assetID)

	var info assetstypes.AssetInfo
	err := row.Scan(
		&info.Name,
		&info.Symbol,
		&info.Address,
		&info.Decimals,
		&info.LayerZeroChainID,
		&info.ImuaChainIndex,
		&info.MetaInfo,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("asset not found for AVS (%s, asset_id %s)", avsAddr, assetID)
		}
		return nil, fmt.Errorf("failed to query AVS asset info (%s, asset_id %s): %w", avsAddr, assetID, err)
	}

	return &info, nil
}

func (db *Db) UpdateAVSRewardAssetMetadata(
	avsAddr string, assetID string, newMetaInfo string,
) error {
	stmt := `
UPDATE avs_reward_assets
SET meta_info = $1
WHERE avs_addr = $2 AND asset_id = $3;`
	// no conflict can happen above since it is update
	// we don't check for existence because again, we are not a business logic implementation
	_, err := db.SQL.Exec(stmt, newMetaInfo, avsAddr, assetID)
	if err != nil {
		return fmt.Errorf(
			"failed to update metadata for AVS reward asset with avs_addr %s and asset_id %s: %w",
			avsAddr, assetID, err,
		)
	}
	return nil
}

func (db *Db) UpdateAVSRewardPool(
	avsAddr, assetID, rewardPoolBalance, rewardPoolTotal, rewardAllocationTotal string,
) error {
	stmt := `
UPDATE avs_reward_assets
SET reward_pool_balance = $1,
	reward_pool_total = $2,
	reward_allocation_total = $3,
	reward_pool_withdrawn_total = ($2::numeric - $1::numeric),
	reward_pool_debt = ($3::numeric - $2::numeric)
WHERE avs_addr = $4 AND asset_id = $5;`

	_, err := db.SQL.Exec(
		stmt,
		rewardPoolBalance,     // $1
		rewardPoolTotal,       // $2
		rewardAllocationTotal, // $3
		avsAddr,               // $4
		assetID,               // $5
	)
	if err != nil {
		return fmt.Errorf("failed to update AVS reward pool for avs_addr %s and asset_id %s: %w", avsAddr, assetID, err)
	}
	return nil
}

// SaveAVSRewardParams saves or updates AVS reward parameters in the database.
func (db *Db) SaveAVSRewardParams(params *types.AVSRewardParams) error {
	stmt := `
INSERT INTO avs_reward_params (
	avs_addr,
	custom_reward_inflation,
	custom_operator_ratio,
	height
) VALUES (
	$1, $2, $3, $4
)
ON CONFLICT (avs_addr) DO UPDATE
SET custom_reward_inflation = EXCLUDED.custom_reward_inflation,
	custom_operator_ratio = EXCLUDED.custom_operator_ratio,
	height = EXCLUDED.height;`

	_, err := db.SQL.Exec(stmt,
		params.AVSAddr,               // $1
		params.CustomRewardInflation, // $2
		params.CustomOperatorRatio,   // $3
		params.Height,                // $4
	)
	if err != nil {
		return fmt.Errorf("failed to save avs reward params: %w", err)
	}
	return nil
}

// SaveAVSRewardDistribution inserts or updates a single AVS reward distribution record in the database.
// The rewards field is stored as JSONB, and operator_reward_proportions is stored as a Postgres array of composite types.
func (db *Db) SaveAVSRewardDistribution(dist *types.AVSRewardDistribution) error {
	stmt := `
INSERT INTO avs_reward_distribution (
    avs_addr,
    epoch_identifier,
    rewards,
    operator_reward_proportions,
    rewards_epoch_number,
    proportions_epoch_number
) VALUES (
    $1, $2, $3, $4, $5, $6
)
ON CONFLICT (avs_addr) DO UPDATE SET
    epoch_identifier = EXCLUDED.EXCLUDED,
    rewards = EXCLUDED.rewards,
    operator_reward_proportions = EXCLUDED.operator_reward_proportions,
    rewards_epoch_number = EXCLUDED.rewards_epoch_number,
    proportions_epoch_number = EXCLUDED.proportions_epoch_number;`

	// OperatorRewardProportions implements driver.Valuer,
	// so we pass it directly to Exec and let the driver handle serialization.
	_, err := db.SQL.Exec(stmt,
		dist.EpochIdentifier,
		dist.AVSAddr,
		pq.Array(dbtypes.NewDbDecCoins(dist.Rewards)),
		dbtypes.NewDbOperatorRewardProportions(dist.OperatorRewardProportions),
		dist.RewardsEpochNumber,
		dist.ProportionsEpochNumber,
	)
	if err != nil {
		return fmt.Errorf("failed to save AVS reward distribution: %w", err)
	}
	return nil
}

// UpsertAVSEpochRewards inserts or updates the rewards, rewards_epoch_number,
// and epoch_identifier for a given AVS address in the database. If the record
// does not exist, it inserts a new row with default values for other fields.
func (db *Db) UpsertAVSEpochRewards(
	avsAddr string,
	epochIdentifier string,
	rewardsEpochNumber int64,
	rewards sdk.DecCoins,
) error {
	stmt := `
		UPDATE avs_reward_distribution
		SET
			rewards = $1,
			rewards_epoch_number = $2,
			epoch_identifier = $3
		WHERE avs_addr = $4
	`

	result, err := db.SQL.Exec(stmt,
		pq.Array(dbtypes.NewDbDecCoins(rewards)),
		rewardsEpochNumber,
		epochIdentifier,
		avsAddr,
	)
	if err != nil {
		return fmt.Errorf("failed to update rewards: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		stmt = `
			INSERT INTO avs_reward_distribution (
				avs_addr, epoch_identifier,
				rewards, rewards_epoch_number
			) VALUES ($1, $2, $3, $4)
		`
		_, err = db.SQL.Exec(
			stmt,
			avsAddr,
			epochIdentifier,
			pq.Array(dbtypes.NewDbDecCoins(rewards)),
			rewardsEpochNumber,
		)
		if err != nil {
			return fmt.Errorf("failed to insert new rewards record: %w", err)
		}
	}

	return nil
}

// UpsertOperatorRewardProportions inserts or updates the operator reward proportions
// and proportions epoch number for a given AVS address in the database.
// It updates only the operator_reward_proportions and proportions_epoch_number fields,
// leaving other fields unchanged. If no record exists for the given avsAddr,
// it inserts a new record with only these fields set, other fields left as default.
func (db *Db) UpsertOperatorRewardProportions(
	avsAddr string,
	epochIdentifier string,
	operatorRewardProportions []distrtypes.OperatorRewardProportion,
	proportionsEpochNumber int64,
) error {
	stmt := `
		UPDATE avs_reward_distribution
		SET
			operator_reward_proportions = $1,
			proportions_epoch_number = $2,
			epoch_identifier = $3
		WHERE avs_addr = $4
	`

	result, err := db.SQL.Exec(stmt,
		dbtypes.NewDbOperatorRewardProportions(operatorRewardProportions),
		proportionsEpochNumber,
		epochIdentifier,
		avsAddr,
	)
	if err != nil {
		return fmt.Errorf("failed to update operator reward proportions: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		// Insert a new record with only operator_reward_proportions and proportions_epoch_number,
		// other fields will be set to their default values.
		stmt = `
			INSERT INTO avs_reward_distribution (
				avs_addr,epoch_identifier,
				operator_reward_proportions,
				proportions_epoch_number
			) VALUES ($1, $2, $3, $4)
		`
		_, err = db.SQL.Exec(
			stmt, avsAddr, epochIdentifier,
			dbtypes.NewDbOperatorRewardProportions(operatorRewardProportions),
			proportionsEpochNumber)
		if err != nil {
			return fmt.Errorf("failed to insert new operator reward proportions record: %w", err)
		}
	}

	return nil
}

// UpsertOperatorRewards inserts or updates the operator_rewards record for a given operator and AVS.
// If the record exists, it adds the newly accumulated rewards and commission to the existing totals.
// If the record does not exist, it inserts a new row with the provided values.
// It also updates derived fields: total_staker_rewards and remaining_commission.
func (db *Db) UpsertOperatorRewards(
	operatorAddr string,
	avsAddr string,
	addedRewards sdk.DecCoins,
	addedCommission sdk.DecCoins,
) error {
	var existingTotalRewards, existingTotalCommission, existingWithdrawn sdk.DecCoins

	// Step 1: Query existing record (if any) from the database
	query := `
		SELECT total_rewards, total_commission, withdrawn_commission
		FROM operator_rewards
		WHERE operator_addr = $1 AND avs_addr = $2
	`
	row := db.SQL.QueryRow(query, operatorAddr, avsAddr)

	var dbTotalRewards, dbTotalCommission, dbWithdrawnCommission dbtypes.DbDecCoins
	err := row.Scan(&dbTotalRewards, &dbTotalCommission, &dbWithdrawnCommission)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check existing operator_rewards: %w", err)
	}

	// Step 2: If a record exists, update it with accumulated values
	if err == nil {
		existingTotalRewards = dbTotalRewards.ToDecCoins()
		existingTotalCommission = dbTotalCommission.ToDecCoins()
		existingWithdrawn = dbWithdrawnCommission.ToDecCoins()

		// Add the new values to the existing ones
		newTotalRewards := existingTotalRewards.Add(addedRewards...)
		newTotalCommission := existingTotalCommission.Add(addedCommission...)

		// Calculate derived fields
		totalStakerRewards := newTotalRewards.Sub(newTotalCommission)
		remainingCommission := newTotalCommission.Sub(existingWithdrawn)

		// Perform the update
		stmt := `
			UPDATE operator_rewards
			SET
				total_rewards = $1,
				total_commission = $2,
				total_staker_rewards = $3,
				remaining_commission = $4
			WHERE operator_addr = $5 AND avs_addr = $6
		`
		_, err = db.SQL.Exec(stmt,
			pq.Array(dbtypes.NewDbDecCoins(newTotalRewards)),
			pq.Array(dbtypes.NewDbDecCoins(newTotalCommission)),
			pq.Array(dbtypes.NewDbDecCoins(totalStakerRewards)),
			pq.Array(dbtypes.NewDbDecCoins(remainingCommission)),
			operatorAddr,
			avsAddr,
		)
		if err != nil {
			return fmt.Errorf("failed to update operator_rewards: %w", err)
		}
		return nil
	}

	// Step 3: If no record exists, insert a new one
	totalStakerRewards := addedRewards.Sub(addedCommission)
	stmt := `
		INSERT INTO operator_rewards (
			operator_addr,
			avs_addr,
			total_rewards,
			total_commission,
			withdrawn_commission,
			total_staker_rewards,
			remaining_commission
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	empty := dbtypes.NewDbDecCoins(nil)

	_, err = db.SQL.Exec(stmt,
		operatorAddr,
		avsAddr,
		pq.Array(dbtypes.NewDbDecCoins(addedRewards)),
		pq.Array(dbtypes.NewDbDecCoins(addedCommission)),
		pq.Array(empty), // withdrawn_commission is 0 initially
		pq.Array(dbtypes.NewDbDecCoins(totalStakerRewards)),
		pq.Array(dbtypes.NewDbDecCoins(addedCommission)), // remaining = commission - 0
	)
	if err != nil {
		return fmt.Errorf("failed to insert operator_rewards: %w", err)
	}

	return nil
}

// WithdrawOperatorCommission subtracts a withdrawn commission from the operator's record
// and updates the withdrawn_commission and remaining_commission accordingly.
func (db *Db) WithdrawOperatorCommission(
	operatorAddr string,
	avsAddr string,
	subCommission sdk.DecCoins,
) error {
	// Query existing commission and withdrawn_commission
	stmt := `
		SELECT total_commission, withdrawn_commission
		FROM operator_rewards
		WHERE operator_addr = $1 AND avs_addr = $2
	`

	var dbTotalCommission, dbWithdrawnCommission dbtypes.DbDecCoins
	err := db.SQL.QueryRow(stmt, operatorAddr, avsAddr).Scan(
		&dbTotalCommission, &dbWithdrawnCommission,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("operator_rewards not found for operator %s and AVS %s", operatorAddr, avsAddr)
		}
		return fmt.Errorf("failed to query operator_rewards: %w", err)
	}

	// Update withdrawn_commission and remaining_commission
	newWithdrawnCommission := dbWithdrawnCommission.ToDecCoins().Add(subCommission...)
	newRemainingCommission, isNegative := dbTotalCommission.ToDecCoins().SafeSub(newWithdrawnCommission)
	if isNegative {
		return fmt.Errorf("failed to calculate remaining commission, totalCommission:%s, subCommission:%s", dbTotalCommission.ToDecCoins(), newWithdrawnCommission)
	}
	// Update in database
	updateStmt := `
		UPDATE operator_rewards
		SET withdrawn_commission = $1,
		    remaining_commission = $2
		WHERE operator_addr = $3 AND avs_addr = $4
	`

	_, err = db.SQL.Exec(
		updateStmt,
		pq.Array(dbtypes.NewDbDecCoins(newWithdrawnCommission)),
		pq.Array(dbtypes.NewDbDecCoins(newRemainingCommission)),
		operatorAddr, avsAddr)
	if err != nil {
		return fmt.Errorf("failed to update operator_rewards: %w", err)
	}

	return nil
}
