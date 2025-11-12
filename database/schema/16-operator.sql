CREATE TABLE operators (
    earnings_addr TEXT NOT NULL PRIMARY KEY,
    approve_addr TEXT NOT NULL,
    operator_meta_info TEXT,
    commission_rate NUMERIC NOT NULL,
    max_commission_rate NUMERIC NOT NULL,
    max_change_rate NUMERIC NOT NULL,
    -- we use ctx.BlockTime() which is in UTC, so drop the TZ
    commission_last_updated TIMESTAMP WITHOUT TIME ZONE
);

-- I made this table a bit separate because it is not used actively yet.
CREATE TABLE client_chain_earning_addresses (
    operator_addr TEXT NOT NULL,
    lz_client_chain_id BIGINT NOT NULL,
    client_chain_earning_addr TEXT NOT NULL,
    PRIMARY KEY (operator_addr, lz_client_chain_id),
    CONSTRAINT fk_operator_addr FOREIGN KEY (operator_addr) REFERENCES operators (earnings_addr)
);

-- operator to avs mapping
CREATE TABLE operator_avs_opt_ins (
    operator_addr TEXT NOT NULL,
    avs_addr TEXT NOT NULL,
    -- can be null for x/dogfood, but not sure why it is tracked because it is an x/avs property
    slash_contract TEXT,
    -- not 0 because height is opted into
    opt_in_height BIGINT NOT NULL,
    -- default is max height
    opt_out_height NUMERIC NOT NULL DEFAULT 18446744073709551615,
    jailed BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (operator_addr, avs_addr),
    CONSTRAINT fk_avs_addr FOREIGN KEY (avs_addr) REFERENCES avs (avs_addr),
    CONSTRAINT fk_operator_addr FOREIGN KEY (operator_addr) REFERENCES operators (earnings_addr)
);

CREATE TABLE consensus_keys (
    operator_addr TEXT NOT NULL,
    chain_id TEXT NOT NULL,
    pubkey TEXT NOT NULL,
    cons_addr TEXT NOT NULL,
    -- optional, visible upon key rotation until pruned and thus can be NULL
    prev_pubkey TEXT,
    prev_cons_addr TEXT,
    -- is_removing represents the situation in which the operator is in the process of opting out
    -- but has not yet completed the unbonding period
    is_removing BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (operator_addr, chain_id),
    CONSTRAINT fk_operator_addr FOREIGN KEY (operator_addr) REFERENCES operators (earnings_addr),
    CONSTRAINT fk_chain_id FOREIGN KEY (chain_id) REFERENCES chain_id_to_avs_addr (chain_id)
);

CREATE TABLE consensus_keys_history (
    chain_id TEXT NOT NULL,
    pubkey TEXT NOT NULL,
    cons_addr TEXT NOT NULL,
    operator_addr TEXT NOT NULL,
    -- the im_height at which the key was added by the operator
    addition_height BIGINT NOT NULL,
    -- the first time a key was added to the validator set
    -- can be null if the key is never added to the validator set
    -- or if the chain_id is not imuachain's
    first_activation_height BIGINT,
    -- the last time a key was seen in the validator set
    -- can be null if the key is never added to the validator set
    -- or if the chain_id is not imuachain's
    last_active_height BIGINT,
    -- can be null if the key removal was not requested
    -- it is tracked for non-imuachain chains as well
    removal_requested_height BIGINT,
    -- the chain does not permit operators to share keys
    -- unless fully unbonded
    PRIMARY KEY (chain_id, pubkey, addition_height),
    CONSTRAINT height_check CHECK (
        (
            first_activation_height IS NULL OR
            addition_height <= first_activation_height
        ) AND
        (
            last_active_height IS NULL OR
            first_activation_height IS NULL OR
            first_activation_height <= last_active_height
        ) AND
        (
            removal_requested_height IS NULL OR
            last_active_height IS NULL OR last_active_height <= removal_requested_height
        )
    )
);

-- needed for Hasura
CREATE OR REPLACE VIEW operator_consensus_status_structure AS
SELECT
    NULL::BIGINT AS first_activation_height,
    NULL::BIGINT AS last_active_height,
    NULL::BOOLEAN AS currently_in_set,
    NULL::TEXT AS last_active_cons_addr;
-- given an operator acc address and chain id,
-- it returns the first activation height, last active height,
-- whether the operator is currently in the validator set,
-- and the last active consensus address.
CREATE OR REPLACE FUNCTION get_operator_consensus_status(
    in_operator_addr TEXT,
    in_chain_id TEXT
)
RETURNS SETOF operator_consensus_status_structure
LANGUAGE plpgsql
STABLE
AS $$
BEGIN
    RETURN QUERY
    WITH relevant_keys AS (
        SELECT *
        FROM consensus_keys_history h
        WHERE h.operator_addr = in_operator_addr
          AND h.chain_id = in_chain_id
          AND h.first_activation_height IS NOT NULL
    ),
    first_height AS (
        SELECT MIN(rk.first_activation_height) AS fah
        FROM relevant_keys rk
    ),
    last_height AS (
        SELECT MAX(rk.last_active_height) AS lah
        FROM relevant_keys rk
    ),
    active_keys AS (
        SELECT rk.cons_addr
        FROM relevant_keys rk
        WHERE rk.last_active_height IS NULL
        ORDER BY rk.first_activation_height DESC
        LIMIT 1
    ),
    fallback_last_active AS (
        SELECT rk.cons_addr
        FROM relevant_keys rk
        WHERE rk.last_active_height IS NOT NULL
        ORDER BY rk.last_active_height DESC
        LIMIT 1
    )
    SELECT
        (SELECT fah FROM first_height) AS first_activation_height,
        (SELECT lah FROM last_height) AS last_active_height,
        EXISTS (
            SELECT 1 FROM relevant_keys WHERE last_active_height IS NULL
        ) AS currently_in_set,
        COALESCE(
            (SELECT cons_addr FROM active_keys),
            (SELECT cons_addr FROM fallback_last_active)
        ) AS last_active_cons_addr
    GROUP BY in_operator_addr, in_chain_id;
END;
$$;

CREATE OR REPLACE VIEW operator_from_consensus_address_structure AS
SELECT
    NULL::TEXT AS operator_addr,
    NULL::JSONB AS meta_info;
-- given a consensus address, chain id, and height,
-- it returns the operator address and meta info.
CREATE OR REPLACE FUNCTION get_operator_from_consensus_address(
    in_cons_addr TEXT,
    in_chain_id TEXT,
    in_height BIGINT
)
RETURNS SETOF operator_from_consensus_address_structure -- Returns our tracked view's type
LANGUAGE plpgsql
STABLE
AS $$
BEGIN
    RETURN QUERY
    SELECT
        ckh.operator_addr,
        o.meta_info
    FROM consensus_keys_history ckh
    JOIN operators o ON o.earnings_addr = ckh.operator_addr
    WHERE ckh.cons_addr = in_cons_addr
      AND ckh.chain_id = in_chain_id
      AND ckh.first_activation_height IS NOT NULL
      AND (
          ckh.first_activation_height <= in_height
          AND (ckh.last_active_height IS NULL OR in_height <= ckh.last_active_height)
      )
    LIMIT 1;
END;
$$;


-- no correlation with NN-delegation.sql because the staker is not captured below.
CREATE TABLE operator_usd_values (
    operator_addr TEXT NOT NULL,
    avs_addr TEXT NOT NULL,
    self_usd_value NUMERIC NOT NULL,
    total_usd_value NUMERIC NOT NULL,
    -- for ease of query, we store the other usd value = total_usd_value - self_usd_value
    other_usd_value NUMERIC NOT NULL,
    -- 0 if self_usd_value < min_self_delegation for avs_addr
    active_usd_value NUMERIC NOT NULL,
    PRIMARY KEY (operator_addr, avs_addr),
    CONSTRAINT fk_operator_addr FOREIGN KEY (operator_addr) REFERENCES operators (earnings_addr),
    CONSTRAINT fk_avs_addr FOREIGN KEY (avs_addr) REFERENCES avs (avs_addr)
);

CREATE TABLE avs_usd_values (
    avs_addr TEXT NOT NULL PRIMARY KEY,
    usd_value NUMERIC NOT NULL,
    CONSTRAINT fk_avs_addr FOREIGN KEY (avs_addr) REFERENCES avs (avs_addr)
);

-- TODO slash states?
