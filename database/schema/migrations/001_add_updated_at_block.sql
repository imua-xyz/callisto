-- Migration: Add updated_at_block to bootstrap tables for optimistic update invalidation
-- This column stores the block height when a record was last updated,
-- allowing the frontend to invalidate optimistic updates once the backend has caught up.

-- Add updated_at_block to bootstrap_staker_assets
ALTER TABLE bootstrap_staker_assets 
ADD COLUMN IF NOT EXISTS updated_at_block BIGINT;

-- Add updated_at_block to bootstrap_delegation_states
ALTER TABLE bootstrap_delegation_states 
ADD COLUMN IF NOT EXISTS updated_at_block BIGINT;

-- Create indexes for efficient querying
CREATE INDEX IF NOT EXISTS idx_bootstrap_staker_assets_block 
ON bootstrap_staker_assets (updated_at_block);

CREATE INDEX IF NOT EXISTS idx_bootstrap_delegations_block 
ON bootstrap_delegation_states (updated_at_block);

-- Update existing records with NULL (will be populated on next update)
-- No action needed as NULL is the default
