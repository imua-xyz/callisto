CREATE TABLE IF NOT EXISTS bootstrap_validator
(
    validator_eth_addr  TEXT    NOT NULL PRIMARY KEY,
    validator_im_addr   TEXT    NOT NULL UNIQUE,
    validator_name      TEXT    NOT NULL UNIQUE,
    consensus_pub_key   TEXT    NOT NULL UNIQUE,
    commission_rate     NUMERIC NOT NULL,
    max_commission_rate NUMERIC NOT NULL,
    max_change_rate     NUMERIC NOT NULL,
    updated_at          TIMESTAMP WITHOUT TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_bootstrap_validator_im_addr ON bootstrap_validator (validator_im_addr);

CREATE TABLE IF NOT EXISTS bootstrap_client_chains
(
    name                TEXT NOT NULL,
    meta_info           TEXT NOT NULL,
    layer_zero_chain_id BIGINT PRIMARY KEY,
    updated_at          TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE IF NOT EXISTS bootstrap_tokens
(
    -- generated for ease; not required to be part of the schema
    asset_id             TEXT PRIMARY KEY,
    name                 TEXT    NOT NULL,
    symbol               TEXT    NOT NULL,
    address              TEXT    NOT NULL CHECK (address = lower(address)),
    decimals             INT     NOT NULL,
    layer_zero_chain_id  BIGINT  NOT NULL,
    staking_total_amount NUMERIC NOT NULL DEFAULT 0,
    total_usd_value      NUMERIC NOT NULL DEFAULT 0,
    updated_at           TIMESTAMP WITHOUT TIME ZONE,
    -- relational constraint
    CONSTRAINT fk_layer_zero_chain_id FOREIGN KEY (layer_zero_chain_id) REFERENCES bootstrap_client_chains (layer_zero_chain_id)
);

-- The view is required because of Hasura.
CREATE OR REPLACE VIEW total_tvl_structure AS
SELECT NULL::NUMERIC AS total;

CREATE OR REPLACE FUNCTION get_total_tvl()
RETURNS SETOF total_tvl_structure
LANGUAGE plpgsql STABLE AS $$
BEGIN
RETURN QUERY
SELECT COALESCE(SUM(total_usd_value), 0) AS total
FROM bootstrap_tokens;
END;
$$;

-- needed for Hasura
CREATE OR REPLACE VIEW max_usd_value_token_structure AS
SELECT
    NULL::TEXT    AS asset_id,
    NULL::TEXT    AS name,
    NULL::TEXT    AS symbol,
    NULL::NUMERIC AS total_usd_value,
    NULL::NUMERIC AS staking_total_amount;

CREATE OR REPLACE FUNCTION get_max_usd_value_token()
RETURNS SETOF max_usd_value_token_structure
LANGUAGE plpgsql
STABLE
AS $$
BEGIN
RETURN QUERY
SELECT t.asset_id, t.name, t.symbol, t.total_usd_value, t.staking_total_amount
FROM bootstrap_tokens t
ORDER BY t.total_usd_value DESC
    LIMIT 1;
END;
$$;

CREATE TABLE bootstrap_token_prices
(
    asset_id   TEXT PRIMARY KEY,
    price      NUMERIC NOT NULL DEFAULT 0,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT fk_asset_id FOREIGN KEY (asset_id) REFERENCES bootstrap_tokens (asset_id)
);


CREATE TABLE IF NOT EXISTS bootstrap_staker_assets
(
    staker_id       TEXT    NOT NULL,
    asset_id        TEXT    NOT NULL,
    deposited       NUMERIC NOT NULL DEFAULT 0,
    withdrawable    NUMERIC NOT NULL DEFAULT 0,
    delegated       NUMERIC NOT NULL DEFAULT 0,
    updated_at      TIMESTAMP WITHOUT TIME ZONE,
    updated_at_block BIGINT,  -- Block height when this record was last updated (for optimistic update invalidation)
    PRIMARY KEY (staker_id, asset_id),
    CONSTRAINT chk_total CHECK (deposited = withdrawable + delegated),
    CONSTRAINT fk_asset_id FOREIGN KEY (asset_id) REFERENCES bootstrap_tokens (asset_id)
);

CREATE INDEX IF NOT EXISTS idx_bootstrap_deposits_staker_id ON bootstrap_staker_assets (staker_id);
CREATE INDEX IF NOT EXISTS idx_bootstrap_deposits_asset_id ON bootstrap_staker_assets (asset_id);
CREATE INDEX IF NOT EXISTS idx_bootstrap_staker_assets_block ON bootstrap_staker_assets (updated_at_block);

-- The view is required because of Hasura.
CREATE OR REPLACE VIEW active_staker_count_structure AS
SELECT NULL::BIGINT AS count;

CREATE OR REPLACE FUNCTION get_active_staker_count()
RETURNS SETOF active_staker_count_structure
LANGUAGE plpgsql STABLE AS $$
BEGIN
RETURN QUERY
SELECT COUNT(DISTINCT staker_id) AS count
FROM bootstrap_staker_assets
WHERE deposited > 0;
END;
$$;



CREATE TABLE IF NOT EXISTS bootstrap_delegation_states
(
    staker_id        TEXT    NOT NULL,
    asset_id         TEXT    NOT NULL,
    operator_addr    TEXT    NOT NULL,
    delegated        NUMERIC NOT NULL DEFAULT 0,
    updated_at       TIMESTAMP WITHOUT TIME ZONE,
    updated_at_block BIGINT,  -- Block height when this record was last updated (for optimistic update invalidation)
    PRIMARY KEY (staker_id, asset_id, operator_addr),
    CONSTRAINT fk_operator FOREIGN KEY (operator_addr) REFERENCES bootstrap_validator (validator_im_addr),
    CONSTRAINT fk_asset_id FOREIGN KEY (asset_id) REFERENCES bootstrap_tokens (asset_id),
    CONSTRAINT fk_staker_asset FOREIGN KEY (staker_id, asset_id) REFERENCES bootstrap_staker_assets (staker_id, asset_id)
);

CREATE INDEX IF NOT EXISTS idx_bootstrap_delegations_staker_id ON bootstrap_delegation_states (staker_id);
CREATE INDEX IF NOT EXISTS idx_bootstrap_delegations_asset_id ON bootstrap_delegation_states (asset_id);
CREATE INDEX IF NOT EXISTS idx_bootstrap_delegations_block ON bootstrap_delegation_states (updated_at_block);

CREATE TABLE IF NOT EXISTS bootstrap_operator_assets
(
    operator_addr TEXT    NOT NULL,
    asset_id      TEXT    NOT NULL,
    total_amount  NUMERIC NOT NULL,
    self_amount   NUMERIC NOT NULL DEFAULT 0,
    other_amount  NUMERIC NOT NULL DEFAULT 0,
    updated_at    TIMESTAMP WITHOUT TIME ZONE,
    PRIMARY KEY (operator_addr, asset_id),
    CONSTRAINT fk_asset_id FOREIGN KEY (asset_id) REFERENCES bootstrap_tokens (asset_id),
    CONSTRAINT fk_operator FOREIGN KEY (operator_addr) REFERENCES bootstrap_validator (validator_im_addr),
    CONSTRAINT chk_total_amount CHECK (total_amount = self_amount + other_amount)
);

CREATE INDEX IF NOT EXISTS idx_bootstrap_operator_assets_operator ON bootstrap_operator_assets (operator_addr);
CREATE INDEX IF NOT EXISTS idx_bootstrap_operator_assets_asset_id ON bootstrap_operator_assets (asset_id);

-- Incremental scanning support tables
CREATE TABLE IF NOT EXISTS bootstrap_scan_state
(
    chain_type     TEXT PRIMARY KEY, -- 'BTC' or 'XRP'
    last_height    BIGINT NOT NULL,  -- BTC: block height, XRP: ledger index
    last_hash      TEXT,             -- Last processed block/ledger hash (for reorg detection)
    safe_height    BIGINT NOT NULL,  -- Safe confirmed height (considering reorgs)
    updated_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS bootstrap_processed_transactions
(
    chain_type      TEXT NOT NULL,   -- 'BTC' or 'XRP'
    tx_hash         TEXT NOT NULL,   -- Transaction hash
    block_height    BIGINT NOT NULL, -- Block/ledger height
    processed_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (chain_type, tx_hash)
);

CREATE INDEX IF NOT EXISTS idx_processed_tx_height ON bootstrap_processed_transactions (chain_type, block_height);
CREATE INDEX IF NOT EXISTS idx_processed_tx_processed_at ON bootstrap_processed_transactions (processed_at);

-- Address binding tables for 1:1 mapping between external chains and Imuachain addresses
CREATE TABLE IF NOT EXISTS bootstrap_address_bindings
(
    chain_type    TEXT NOT NULL, -- 'BTC' or 'XRP'
    source_addr   TEXT NOT NULL, -- BTC/XRP address (normalized: lowercase, trimmed)
    target_addr   TEXT NOT NULL, -- Imuachain address (normalized: lowercase, trimmed)
    created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (chain_type, source_addr),
    -- Ensure 1:1 mapping - each target address can only bind to one source address per chain
    CONSTRAINT uq_target_per_chain UNIQUE (chain_type, target_addr)
);

CREATE INDEX IF NOT EXISTS idx_address_bindings_target ON bootstrap_address_bindings (target_addr);
CREATE INDEX IF NOT EXISTS idx_address_bindings_created ON bootstrap_address_bindings (created_at);

CREATE TABLE bootstrap_statistics
(
    one_row_id BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    tvl        NUMERIC NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    CHECK (one_row_id)
);
