package database

import (
	"database/sql"
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"

	"github.com/forbole/callisto/v4/types"
)

func (db *Db) SaveGenesisPoolAirdropRound(tx *sql.Tx, round *types.CommonAirdropRound) error {
	if round.AirdropType != types.GenesisPoolAirdrop {
		return fmt.Errorf("invalid airdrop type:%d", round.AirdropType)
	}
	stmt := `
INSERT INTO genesis_pool_airdrop_rounds (
    airdrop_round,
    block_height,
    total_stakers,
    total_usd_value,
    total_reward_amount,
    distributed_stakers,
    round_duration,
    round_start_at,
    round_end_at,
    created_at,
    is_completed
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT (airdrop_round) DO NOTHING;`

	execFn := db.SQL.Exec
	if tx != nil {
		execFn = tx.Exec
	}
	_, err := execFn(stmt,
		round.AirdropRound,
		round.BlockHeight,
		round.TotalStakers,
		round.TotalUSDValue,
		round.TotalRewardAmount,
		round.DistributedStakers,
		round.RoundDuration,
		round.RoundStartAt,
		round.RoundEndAt,
		round.CreatedAt,
		round.IsCompleted,
	)
	if err != nil {
		return fmt.Errorf("failed to save genesis pool airdrop round %d: %w", round.AirdropRound, err)
	}
	return nil
}

func (db *Db) UpdateGenesisPoolAirdropMetrics(
	tx *sql.Tx,
	airdropRound int,
	totalStakers int,
	totalUSDValue string,
) error {
	stmt := `
UPDATE genesis_pool_airdrop_rounds
SET
    total_stakers = $1,
    total_usd_value = $2
WHERE airdrop_round = $3;`

	execFn := db.SQL.Exec
	if tx != nil {
		execFn = tx.Exec
	}

	_, err := execFn(stmt, totalStakers, totalUSDValue, airdropRound)
	if err != nil {
		return fmt.Errorf(
			"failed to update genesis pool airdrop metrics, round %d: %w",
			airdropRound, err,
		)
	}

	return nil
}

func (db *Db) GetGenesisPoolAirdropRound(roundIndex int) (*types.CommonAirdropRound, error) {
	stmt := `
SELECT
    airdrop_round,
    block_height,
    total_stakers,
    total_usd_value,
    total_reward_amount,
    distributed_stakers,
    round_duration,
    round_start_at,
    round_end_at,
    created_at,
    is_completed
FROM genesis_pool_airdrop_rounds
WHERE airdrop_round = $1;`

	var round types.CommonAirdropRound
	err := db.SQL.QueryRow(stmt, roundIndex).Scan(
		&round.AirdropRound,
		&round.BlockHeight,
		&round.TotalStakers,
		&round.TotalUSDValue,
		&round.TotalRewardAmount,
		&round.DistributedStakers,
		&round.RoundDuration,
		&round.RoundStartAt,
		&round.RoundEndAt,
		&round.CreatedAt,
		&round.IsCompleted,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("airdrop round %d not found", roundIndex)
		}
		return nil, fmt.Errorf("failed to query airdrop round %d: %w", roundIndex, err)
	}
	round.AirdropType = types.GenesisPoolAirdrop
	return &round, nil
}

func (db *Db) GetLatestGenesisPoolAirdropRound() (*types.CommonAirdropRound, error) {
	stmt := `
SELECT
    airdrop_round,
    block_height,
    total_stakers,
    total_usd_value,
    total_reward_amount,
    distributed_stakers,
    round_duration,
    round_start_at,
    round_end_at,
    created_at,
    is_completed
FROM genesis_pool_airdrop_rounds
ORDER BY airdrop_round DESC
LIMIT 1;`

	var round types.CommonAirdropRound
	err := db.SQL.QueryRow(stmt).Scan(
		&round.AirdropRound,
		&round.BlockHeight,
		&round.TotalStakers,
		&round.TotalUSDValue,
		&round.TotalRewardAmount,
		&round.DistributedStakers,
		&round.RoundDuration,
		&round.RoundStartAt,
		&round.RoundEndAt,
		&round.CreatedAt,
		&round.IsCompleted,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Don't return an error since there may be no round info initially.
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query latest genesis pool airdrop round: %w", err)
	}
	round.AirdropType = types.GenesisPoolAirdrop
	return &round, nil
}

func (db *Db) MarkGenesisPoolAirdropRoundCompleted(roundIndex int) error {
	stmt := `
UPDATE genesis_pool_airdrop_rounds
SET is_completed = TRUE
WHERE airdrop_round = $1;`

	result, err := db.SQL.Exec(stmt, roundIndex)
	if err != nil {
		return fmt.Errorf("failed to mark airdrop round %d as completed: %w", roundIndex, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for round %d: %w", roundIndex, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no airdrop round found with index %d", roundIndex)
	}

	return nil
}

func (db *Db) SaveGenesisStakerAirdrop(tx *sql.Tx, airdrop *types.GenesisStakerAirdrop) error {
	stmt := `
INSERT INTO genesis_staker_airdrops (
    staker_id,
    airdrop_round,
    usd_value,
    reward_amount,
    is_distributed,
    distributed_at
) VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (staker_id, airdrop_round) DO NOTHING;`

	execFn := db.SQL.Exec
	if tx != nil {
		execFn = tx.Exec
	}

	_, err := execFn(stmt,
		airdrop.StakerID,
		airdrop.AirdropRound,
		airdrop.USDValue,
		airdrop.RewardAmount,
		airdrop.IsDistributed,
		airdrop.DistributedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save genesis staker airdrop (%s, round %d): %w",
			airdrop.StakerID, airdrop.AirdropRound, err)
	}
	return nil
}

func (db *Db) MarkGenesisStakerAirdropDistributed(stakerID string, roundIndex int) error {
	stmt := `
UPDATE genesis_staker_airdrops
SET is_distributed = TRUE,
    distributed_at = NOW()
WHERE staker_id = $1 AND airdrop_round = $2;`

	result, err := db.SQL.Exec(stmt, stakerID, roundIndex)
	if err != nil {
		return fmt.Errorf("failed to mark airdrop as distributed for staker %s, round %d: %w", stakerID, roundIndex, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for staker %s, round %d: %w", stakerID, roundIndex, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no record found for staker %s in round %d", stakerID, roundIndex)
	}

	return nil
}

func (db *Db) UpdateGenesisAirdropRewardsByRoundSQL(
	tx *sql.Tx,
	round int,
	roundReward sdkmath.LegacyDec,
	totalUSDValue sdkmath.LegacyDec,
) error {
	execFn := db.SQL.Exec
	if tx != nil {
		execFn = tx.Exec
	}

	stmt := `
        UPDATE genesis_staker_airdrops
        SET reward_amount = ROUND(usd_value * $2 / $3, 18)
        WHERE airdrop_round = $1;
    `

	_, err := execFn(stmt, round, roundReward.String(), totalUSDValue.String())
	if err != nil {
		return fmt.Errorf("failed to update airdrop rewards for round %d: %w", round, err)
	}

	return nil
}

func (db *Db) SaveLiquidityAirdropRound(tx *sql.Tx, round *types.CommonAirdropRound) error {
	if round.AirdropType != types.LiquidityIncentivesAirdrop {
		return fmt.Errorf("invalid airdrop type:%d", round.AirdropType)
	}
	stmt := `
INSERT INTO liquidity_incentives_airdrop_rounds (
    airdrop_round,
    block_height,
    total_stakers,
    total_native_imua_rewards,
    total_reward_amount,
    distributed_stakers,
    round_duration,
    round_start_at,
    round_end_at,
    created_at,
    is_completed
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT (airdrop_round) DO NOTHING;`

	execFn := db.SQL.Exec
	if tx != nil {
		execFn = tx.Exec
	}
	_, err := execFn(stmt,
		round.AirdropRound,
		round.BlockHeight,
		round.TotalStakers,
		round.TotalNativeIMUARewards,
		round.TotalRewardAmount,
		round.DistributedStakers,
		round.RoundDuration,
		round.RoundStartAt,
		round.RoundEndAt,
		round.CreatedAt,
		round.IsCompleted,
	)
	if err != nil {
		return fmt.Errorf("failed to save liquidity incentive airdrop round %d: %w", round.AirdropRound, err)
	}
	return nil
}

func (db *Db) UpdateLiquidityAirdropMetrics(
	tx *sql.Tx,
	airdropRound int,
	totalStakers int,
	totalNativeIMUARewards string,
) error {

	stmt := `
UPDATE liquidity_incentives_airdrop_rounds
SET
    total_stakers = $1,
    total_native_imua_rewards = $2
WHERE airdrop_round = $3;`

	execFn := db.SQL.Exec
	if tx != nil {
		execFn = tx.Exec
	}

	_, err := execFn(stmt, totalStakers, totalNativeIMUARewards, airdropRound)
	if err != nil {
		return fmt.Errorf(
			"failed to update liquidity incentive airdrop metrics, round %d: %w",
			airdropRound, err,
		)
	}

	return nil
}

func (db *Db) GetLatestLiquidityIncentivesAirdropRound() (*types.CommonAirdropRound, error) {
	stmt := `
SELECT
    airdrop_round,
    block_height,
    total_stakers,
    total_native_imua_rewards,
    total_reward_amount,
    distributed_stakers,
    round_duration,
    round_start_at,
    round_end_at,
    created_at,
    is_completed
FROM liquidity_incentives_airdrop_rounds
ORDER BY airdrop_round DESC
LIMIT 1;`

	var round types.CommonAirdropRound
	err := db.SQL.QueryRow(stmt).Scan(
		&round.AirdropRound,
		&round.BlockHeight,
		&round.TotalStakers,
		&round.TotalNativeIMUARewards,
		&round.TotalRewardAmount,
		&round.DistributedStakers,
		&round.RoundDuration,
		&round.RoundStartAt,
		&round.RoundEndAt,
		&round.CreatedAt,
		&round.IsCompleted,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Don't return an error since there may be no round info initially.
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query latest liquidity incentives airdrop round: %w", err)
	}
	round.AirdropType = types.LiquidityIncentivesAirdrop
	return &round, nil
}

func (db *Db) SaveLiquidityStakerAirdrop(tx *sql.Tx, airdrop *types.LiquidityStakerAirdrop) error {
	stmt := `
INSERT INTO liquidity_incentives_staker_airdrops (
    staker_id,
    airdrop_round,
    outstanding_rewards,
    withdrawn_rewards,
    claimed_rewards,
    unclaimed_rewards,
    total_rewards,
	round_native_rewards,   
    airdrop_reward_amount,
    is_distributed,
    distributed_at,
    created_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
ON CONFLICT (staker_id, airdrop_round) DO NOTHING;`

	execFn := db.SQL.Exec
	if tx != nil {
		execFn = tx.Exec
	}
	_, err := execFn(stmt,
		airdrop.StakerID,
		airdrop.AirdropRound,
		airdrop.OutstandingRewards,
		airdrop.WithdrawnRewards,
		airdrop.ClaimedRewards,
		airdrop.UnclaimedRewards,
		airdrop.TotalRewards,
		airdrop.RoundNativeRewards,
		airdrop.AirdropRewardAmount,
		airdrop.IsDistributed,
		airdrop.DistributedAt,
		airdrop.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save liquidity incentives staker airdrop (%s, round %d): %w",
			airdrop.StakerID, airdrop.AirdropRound, err)
	}
	return nil
}

func (db *Db) GetLiquidityStakerAirdrop(stakerID string, airdropRound int) (*types.LiquidityStakerAirdrop, error) {
	stmt := `
SELECT 
	staker_id,
	airdrop_round,
	outstanding_rewards,
	withdrawn_rewards,
	claimed_rewards,
	unclaimed_rewards,
	total_rewards,
	round_native_rewards,
	airdrop_reward_amount,
	is_distributed,
	distributed_at,
	created_at
FROM liquidity_incentives_staker_airdrops
WHERE staker_id = $1 AND airdrop_round = $2;
`
	row := db.SQL.QueryRow(stmt, stakerID, airdropRound)

	var airdrop types.LiquidityStakerAirdrop
	err := row.Scan(
		&airdrop.StakerID,
		&airdrop.AirdropRound,
		&airdrop.OutstandingRewards,
		&airdrop.WithdrawnRewards,
		&airdrop.ClaimedRewards,
		&airdrop.UnclaimedRewards,
		&airdrop.TotalRewards,
		&airdrop.RoundNativeRewards,
		&airdrop.AirdropRewardAmount,
		&airdrop.IsDistributed,
		&airdrop.DistributedAt,
		&airdrop.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get liquidity incentives staker airdrop (%s, round %d): %w",
			stakerID, airdropRound, err)
	}

	return &airdrop, nil
}

func (db *Db) UpdateLiquidityAirdropRewardsByRoundSQL(
	tx *sql.Tx,
	round int,
	roundReward, totalRoundNativeRewards sdkmath.LegacyDec,
) error {
	execFn := db.SQL.Exec
	if tx != nil {
		execFn = tx.Exec
	}

	stmt := `
		UPDATE liquidity_incentives_staker_airdrops
		SET airdrop_reward_amount = ROUND(round_native_rewards * $2 / $3, 18)
		WHERE airdrop_round = $1;
	`

	_, err := execFn(stmt, round, roundReward.String(), totalRoundNativeRewards.String())
	if err != nil {
		return fmt.Errorf("failed to update liquidity airdrop rewards for round %d: %w", round, err)
	}

	return nil
}
