CREATE TABLE distribution_params
(
    one_row_id BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    params     JSONB   NOT NULL,
    height     BIGINT  NOT NULL,
    CHECK (one_row_id = TRUE)
);
CREATE INDEX distribution_params_height_index ON distribution_params (height);

CREATE TABLE avs_reward_assets
(
    avs_addr                    TEXT    NOT NULL,
    asset_id                    TEXT    NOT NULL,
    name                        TEXT    NOT NULL,
    symbol                      TEXT    NOT NULL,
    address                     TEXT    NOT NULL,
    decimals                    INT     NOT NULL,
    layer_zero_chain_id         BIGINT  NOT NULL,
    imuachain_index             BIGINT  NOT NULL,
    meta_info                   TEXT,
    reward_pool_balance         NUMERIC NOT NULL DEFAULT 0,
    reward_pool_total           NUMERIC NOT NULL DEFAULT 0,
    reward_allocation_total     NUMERIC NOT NULL DEFAULT 0,
    -- derived value via subtraction; only kept for speed
    reward_pool_withdrawn_total NUMERIC NOT NULL DEFAULT 0,
    -- derived value via subtraction; only kept for speed
    reward_pool_debt            NUMERIC NOT NULL DEFAULT 0,
    PRIMARY KEY (avs_addr, asset_id),
    -- relational constraint
    CONSTRAINT fk_avs_addr FOREIGN KEY (avs_addr) REFERENCES avs (avs_addr),
    CONSTRAINT chk_withdrawn_total CHECK (reward_pool_total = reward_pool_balance + reward_pool_withdrawn_total),
    CONSTRAINT chk_debt CHECK (reward_allocation_total = reward_pool_total + reward_pool_debt),
    CONSTRAINT fk_layer_zero_chain_id FOREIGN KEY (layer_zero_chain_id) REFERENCES client_chains (layer_zero_chain_id)
);
CREATE INDEX idx_reward_assets_avs_addr ON avs_reward_assets (avs_addr);
CREATE INDEX idx_reward_assets_asset_id ON avs_reward_assets (asset_id);
CREATE INDEX idx_reward_assets_avs_symbol ON avs_reward_assets (avs_addr, symbol);

CREATE TYPE DEC_COIN AS
    (
    denom TEXT,
    amount TEXT
    );

CREATE TABLE community_pool
(
    avs_addr TEXT PRIMARY KEY NOT NULL,
    coins    DEC_COIN[] NOT NULL,
    height   BIGINT           NOT NULL,

    CONSTRAINT fk_avs_addr FOREIGN KEY (avs_addr) REFERENCES avs (avs_addr)
);

CREATE TABLE avs_reward_params
(
    avs_addr                TEXT PRIMARY KEY NOT NULL,
    custom_reward_inflation BOOLEAN          NOT NULL DEFAULT FALSE,
    custom_operator_ratio   BOOLEAN          NOT NULL DEFAULT FALSE,
    height                  BIGINT           NOT NULL,

    CONSTRAINT fk_avs_addr FOREIGN KEY (avs_addr) REFERENCES avs (avs_addr)
);


-- Define a composite type for OperatorRewardProportion
CREATE TYPE OPERATOR_REWARD_PROPORTION AS
    (
    operator_addr TEXT, -- Operator address (sdk.AccAddress, case preserved)
    reward_proportion NUMERIC -- Reward proportion (cosmos.Dec, stored as NUMERIC)
    );

-- Table for storing AVS reward distribution for a specific epoch
CREATE TABLE avs_reward_distribution
(
    -- Address of the AVS, must be lowercase and unique
    avs_addr                    TEXT PRIMARY KEY NOT NULL,

    -- Epoch identifier for the following epoch numbers.
    epoch_identifier            TEXT,

    -- List of rewards for one epoch (array of DecCoin)
    rewards                     DEC_COIN[] NOT NULL,

    -- Reward proportions of all eligible operators in this epoch
    operator_reward_proportions OPERATOR_REWARD_PROPORTION[] NOT NULL,

    -- Epoch number when rewards were calculated
    rewards_epoch_number        BIGINT           NOT NULL,

    -- Epoch number when operator proportions were updated
    proportions_epoch_number    BIGINT           NOT NULL,

    CONSTRAINT fk_avs_addr FOREIGN KEY (avs_addr) REFERENCES avs (avs_addr)
);

-- Table to store rewards of operators per AVS
CREATE TABLE operator_rewards
(
    -- Address of the operator (sdk.AccAddress)
    operator_addr        TEXT NOT NULL,

    -- Address of the AVS, stored in lowercase
    avs_addr             TEXT NOT NULL,

    -- Array of total AVS rewards (denom + amount)
    total_rewards        DEC_COIN[] NOT NULL,

    -- Array of total accumulated commission (denom + amount)
    total_commission     DEC_COIN[] NOT NULL,

    -- Array of the withdrawn commission (denom + amount)
    withdrawn_commission DEC_COIN[] NOT NULL,

    -- Array of total staker rewards (denom + amount)
    -- (total_rewards - total_commission): derived value via subtraction; only kept for speed.
    total_staker_rewards DEC_COIN[] NOT NULL,

    -- Array of the remaining commission (denom + amount)
    -- (total_commission - withdrawn_commission): derived value via subtraction; only kept for speed.
    remaining_commission DEC_COIN[] NOT NULL,

    -- Composite key ensures uniqueness per operator per AVS
    PRIMARY KEY (operator_addr, avs_addr),

    CONSTRAINT fk_operator_addr FOREIGN KEY (operator_addr) REFERENCES operators (earnings_addr),
    CONSTRAINT fk_avs_addr FOREIGN KEY (avs_addr) REFERENCES avs (avs_addr)
);