CREATE TABLE genesis_pool_airdrop_rounds
(
    airdrop_round       INTEGER PRIMARY KEY,             -- Airdrop round index
    block_height        BIGINT    NOT NULL,              -- Snapshot block height
    total_stakers       INTEGER   NOT NULL,              -- Total number of stakers in this round
    total_usd_value     NUMERIC   NOT NULL,              -- Total USD value aggregated across stakers
    total_reward_amount NUMERIC   NOT NULL,              -- Total reward allocated for this round
    distributed_stakers INTEGER   NOT NULL DEFAULT 0,    -- Number of stakers who have received rewards
    round_duration      NUMERIC   NOT NULL,              -- the duration(minute) for this round
    round_start_at      TIMESTAMP NOT NULL,              -- round start timestamp derived from the block timestamp
    round_end_at        TIMESTAMP NOT NULL,              -- round end timestamp derived from the block timestamp
    created_at          TIMESTAMP NOT NULL,              -- Creation timestamp derived from the block timestamp
    is_completed        BOOLEAN   NOT NULL DEFAULT FALSE -- Whether this round's distribution is finished
);

-- These are the airdrop states for genesis stakers, used to distribute the genesis pool reward.
CREATE TABLE genesis_staker_airdrops
(
    staker_id      TEXT    NOT NULL,               -- The unique identifier of the staker
    airdrop_round  INTEGER NOT NULL,               -- The airdrop round (e.g., weekly interval index)
    usd_value      NUMERIC NOT NULL,               -- The staker's total asset value in USD at snapshot time
    reward_amount  NUMERIC NOT NULL,               -- The reward amount allocated to the staker
    is_distributed BOOLEAN NOT NULL DEFAULT FALSE, -- Whether the reward has been distributed
    distributed_at TIMESTAMP,                      -- Timestamp when the reward was distributed (nullable)

    PRIMARY KEY (staker_id, airdrop_round),
    CONSTRAINT fk_airdrop_round FOREIGN KEY (airdrop_round) REFERENCES genesis_pool_airdrop_rounds (airdrop_round)
);
CREATE INDEX idx_genesis_airdrops_round ON genesis_staker_airdrops (airdrop_round);
CREATE INDEX idx_genesis_airdrops_staker ON genesis_staker_airdrops (staker_id);

CREATE TABLE liquidity_incentives_airdrop_rounds
(
    airdrop_round             INTEGER PRIMARY KEY,             -- Airdrop round index
    block_height              BIGINT    NOT NULL,              -- Snapshot block height
    total_stakers             INTEGER   NOT NULL,              -- Total number of stakers in this round
    total_native_imua_rewards NUMERIC   NOT NULL,              -- Total native IMUA rewards aggregated across stakers
    total_reward_amount       NUMERIC   NOT NULL,              -- Total reward allocated for this round
    distributed_stakers       INTEGER   NOT NULL DEFAULT 0,    -- Number of stakers who have received rewards
    round_duration            NUMERIC   NOT NULL,              -- the duration(minute) for this round
    round_start_at            TIMESTAMP NOT NULL,              -- round start timestamp derived from the block timestamp
    round_end_at              TIMESTAMP NOT NULL,              -- round end timestamp derived from the block timestamp
    created_at                TIMESTAMP NOT NULL,              -- Creation timestamp derived from the block timestamp
    is_completed              BOOLEAN   NOT NULL DEFAULT FALSE -- Whether this round's distribution is finished
);

-- Table to store IMUA reward snapshots and the calculated liquidity incentives for stakers.
-- The reward snapshots are used to calculate the liquidity incentive airdrop.
CREATE TABLE liquidity_incentives_staker_airdrops
(
    -- ID of the staker
    staker_id             TEXT      NOT NULL,

    -- Airdrop round index
    airdrop_round         INTEGER   NOT NULL,

    -- amount of outstanding rewards
    outstanding_rewards   NUMERIC   NOT NULL,

    -- amount of withdrawn rewards
    withdrawn_rewards     NUMERIC   NOT NULL,

    -- amount of claimed rewards
    -- derived value via addition; only kept for speed (claimed_rewards = outstanding_rewards + withdrawn_rewards)
    claimed_rewards       NUMERIC   NOT NULL,

    -- amount of unclaimed rewards
    unclaimed_rewards     NUMERIC   NOT NULL,

    -- amount of total rewards
    -- derived value via addition; only kept for speed (total_rewards = claimed_rewards + unclaimed_rewards)
    -- This reward comes from the native inflation on the IMUA chain — the portion minted and distributed as
    -- validator rewards. It does not include any airdrop rewards. We use this amount to calculate the airdrop.
    total_rewards         NUMERIC   NOT NULL,

    -- All five rewards above are native rewards since TGE.
    -- To calculate the airdrop, we need the native rewards for the current round,
    -- which can be calculated as:
    --   round_native_rewards(currentRound) = total_rewards(currentRound) - total_rewards(previousRound)
    -- since the airdrop amount is proportional to the native rewards earned from the inflation.
    -- airdrop_reward_amount =
    --   liquidity_incentives_airdrop_rounds.total_reward_amount * round_native_rewards
    --   / liquidity_incentives_airdrop_rounds.total_native_imua_rewards
    round_native_rewards  NUMERIC   NOT NULL,

    -- The airdrop reward amount allocated to the staker
    airdrop_reward_amount NUMERIC   NOT NULL,

    -- Whether the reward has been distributed
    is_distributed        BOOLEAN   NOT NULL DEFAULT FALSE,

    -- Timestamp when the reward was distributed (nullable)
    distributed_at        TIMESTAMP,

    -- Record creation timestamp
    created_at            TIMESTAMP NOT NULL DEFAULT now(),

    -- Composite key ensures uniqueness per staker per airdrop round
    PRIMARY KEY (staker_id, airdrop_round),
    CONSTRAINT fk_airdrop_round FOREIGN KEY (airdrop_round) REFERENCES liquidity_incentives_airdrop_rounds (airdrop_round)
);
CREATE INDEX idx_liquidity_airdrops_round ON liquidity_incentives_staker_airdrops (airdrop_round);
CREATE INDEX idx_liquidity_airdrops_staker ON liquidity_incentives_staker_airdrops (staker_id);
