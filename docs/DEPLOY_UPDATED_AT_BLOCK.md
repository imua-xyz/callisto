# Deployment Guide: `updated_at_block` Feature for Optimistic Updates

This guide provides step-by-step instructions for deploying the `updated_at_block` feature to the callisto GraphQL service using the existing docker-compose infrastructure.

## Table of Contents

1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [Local Testing](#local-testing)
4. [Production Deployment](#production-deployment)
5. [Verification](#verification)
6. [Rollback Procedure](#rollback-procedure)
7. [Troubleshooting](#troubleshooting)

---

## Overview

### What Changed

| Component | Changes |
|-----------|---------|
| **Database Schema** | Added `updated_at_block` column to `bootstrap_staker_assets` and `bootstrap_delegation_states` tables |
| **Go Types** | Added `UpdatedAtBlock` field to `BootstrapStakerAsset` and `BootstrapDelegationState` structs |
| **Database Operations** | Updated all CRUD operations to read/write the new column |
| **Module Logic** | Updated event handlers and periodic operations to pass block heights |
| **Hasura Metadata** | Added `updated_at_block` to GraphQL select permissions |

### Files Changed

```
database/schema/22-bootstrap.sql                                              # Schema definition
database/schema/migrations/001_add_updated_at_block.sql                       # Migration script
types/bootstrap.go                                                            # Go types
database/bootstrap.go                                                         # Database operations
modules/bootstrap/handle_periodic_operations.go                               # Periodic tasks
modules/bootstrap/handle_async_operations.go                                  # Event handlers
hasura/metadata/databases/bdjuno/tables/public_bootstrap_staker_assets.yaml   # Hasura permissions
hasura/metadata/databases/bdjuno/tables/public_bootstrap_delegation_states.yaml
```

---

## Prerequisites

### Required Tools

```bash
# Verify Docker is installed and running
docker --version
docker-compose --version

# Verify Go is installed (for building)
go version  # Go 1.21+
```

### Environment Setup

Create a `.env` file in the callisto root directory:

```bash
cd /path/to/callisto

cat > .env << 'EOF'
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
BDJUNO_USER=bdjuno
BDJUNO_PASSWORD=bdjuno_password
RESET_DB=false
RESET_METADATA=false
HASURA_ADMIN_SECRET=your_admin_secret
EOF
```

> **Important:** Use the same `HASURA_ADMIN_SECRET` value when making API calls to Hasura.
> Load the variables with `source .env` before running curl commands.

### Parser database configuration

The Callisto parser (**imuad_callisto_parser**) does **not** read the database URL from `.env` or from Docker Compose environment. It reads it from the config file at **`callisto-config/config.yaml`** (relative to the directory where you run `docker compose`).

Before starting or redeploying, **verify that `database.url` in `callisto-config/config.yaml` is correct**:

- **Host** must be the Postgres service name: `db` (so the parser connects to the database container on the Compose network). Do **not** use `127.0.0.1` or `localhost` inside the parser config, or the parser will try to connect to itself and fail.
- **User** and **password** must match the `bdjuno` role and password created by the setup container (i.e. the same as `BDJUNO_USER` and `BDJUNO_PASSWORD` in your `.env`).
- **Database name** is typically `bdjuno`.

Example (align with your `.env`):

```yaml
database:
  url: "postgresql://bdjuno:bdjuno_password@db:5432/bdjuno?sslmode=disable&search_path=public"
```

A wrong or outdated `database.url` (e.g. old password, or `127.0.0.1`/wrong port) will cause the parser to fail with `pq: password authentication failed for user "bdjuno"` or connection errors. See [PARSER_DATABASE_CONFIG.md](PARSER_DATABASE_CONFIG.md) for details.

---

## Local Testing

### Option A: Fresh Install (No Existing Data)

If you don't have existing data or want to start fresh:

```bash
cd /path/to/callisto

# Set RESET_DB=true to apply fresh schema
export RESET_DB=true

# Start all services (schema will be applied automatically)
docker-compose up -d

# Verify services are running
docker-compose ps
```

The `setup` service will automatically run all SQL files in `database/schema/` including the updated `22-bootstrap.sql`.

---

### Option B: Existing Database (Preserve Data)

If you have existing data that needs to be preserved:

#### Step 1: Start Database Service

```bash
cd /path/to/callisto

# Ensure RESET_DB is false to preserve data
export RESET_DB=false

# Start only the database
docker-compose up -d db

# Wait for database to be ready
sleep 5

# Verify database is running
docker exec imuad_callisto_db pg_isready -U postgres
```

#### Step 2: Run Migration

```bash
# Run the migration script
docker exec -i imuad_callisto_db psql -U bdjuno -d bdjuno < database/schema/migrations/001_add_updated_at_block.sql
```

Expected output:
```
ALTER TABLE
ALTER TABLE
CREATE INDEX
CREATE INDEX
```

#### Step 3: Verify Migration

```bash
# Check columns were added
docker exec imuad_callisto_db psql -U bdjuno -d bdjuno -c "\d bootstrap_staker_assets" | grep updated_at_block
docker exec imuad_callisto_db psql -U bdjuno -d bdjuno -c "\d bootstrap_delegation_states" | grep updated_at_block
```

Expected output:
```
 updated_at_block | bigint |           |          |
```

#### Step 4: Rebuild and Start Callisto

```bash
# Rebuild callisto with new code
docker-compose build callisto

# Start all services
docker-compose up -d
```

#### Step 5: Apply Hasura Metadata

```bash
# Restart hasura-cli to apply metadata
docker-compose up hasura-cli
```

Or manually via Hasura Console:
1. Open http://localhost:8080/console
2. Go to **Settings** → **Metadata Actions** → **Reload Metadata**

---

### Verify Local Setup

#### 1. Check All Services Are Running

```bash
docker-compose ps
```

Expected output:
```
NAME                        STATUS
imuad_callisto_db           Up
imuad_callisto_hasura       Up
imuad_callisto_parser       Up
```

#### 2. Test GraphQL Query

```bash
# Use the admin secret from your .env file
curl -s -X POST \
  -H "Content-Type: application/json" \
  -H "X-Hasura-Admin-Secret: $HASURA_ADMIN_SECRET" \
  -d '{"query": "{ bootstrap_staker_assets(limit: 1) { staker_id updated_at_block } }"}' \
  http://localhost:8080/v1/graphql | jq
```

> **Note:** Load your .env first with `source .env` or `export $(cat .env | xargs)`

Expected response (with or without data):
```json
{
  "data": {
    "bootstrap_staker_assets": []
  }
}
```

If you get an error about `updated_at_block` field not existing, reload Hasura metadata.

#### 3. Check Callisto Logs

```bash
docker logs imuad_callisto_parser --tail 50 -f
```

---

## Production Deployment

### Pre-Deployment Checklist

- [ ] Backup the database
- [ ] **Verify `database.url` in `callisto-config/config.yaml`** — host must be `db`, user/password must match `BDJUNO_USER`/`BDJUNO_PASSWORD` in `.env` (see [Parser database configuration](#parser-database-configuration))
- [ ] Verify the migration script works in staging
- [ ] Schedule maintenance window (if needed)
- [ ] Notify stakeholders

### Deployment Steps

#### 1. Backup Database (Important!)

```bash
# SSH to production server
ssh user@production-server

cd /path/to/callisto

# Create backup
docker exec imuad_callisto_db pg_dump -U bdjuno -d bdjuno > backup_$(date +%Y%m%d_%H%M%S).sql
```

#### 2. Pull Latest Code

```bash
git pull origin main
```

#### 3. Run Database Migration

```bash
# Run migration
docker exec -i imuad_callisto_db psql -U bdjuno -d bdjuno < database/schema/migrations/001_add_updated_at_block.sql
```

#### 4. Rebuild and Deploy Callisto

```bash
# Rebuild with new code
docker-compose build callisto

# Restart callisto service (minimal downtime)
docker-compose up -d callisto
```

#### 5. Apply Hasura Metadata

```bash
# Apply updated metadata
docker-compose up hasura-cli

# Or reload via API (load .env first)
source .env
curl -X POST \
  -H "X-Hasura-Admin-Secret: $HASURA_ADMIN_SECRET" \
  http://localhost:8080/v1/metadata \
  -d '{"type": "reload_metadata", "args": {}}'
```

#### 6. Verify Deployment

```bash
# Check service health
docker-compose ps

# Check logs for errors
docker logs imuad_callisto_parser --tail 100

# Load env and test GraphQL endpoint
source .env
curl -s -X POST \
  -H "Content-Type: application/json" \
  -H "X-Hasura-Admin-Secret: $HASURA_ADMIN_SECRET" \
  -d '{"query": "{ bootstrap_staker_assets(limit: 1) { staker_id updated_at_block } }"}' \
  http://localhost:8080/v1/graphql
```

---

## Verification

### Database Verification

```bash
# Connect to database
docker exec -it imuad_callisto_db psql -U bdjuno -d bdjuno

# Check schema
\d bootstrap_staker_assets
\d bootstrap_delegation_states

# Check indexes
\di idx_bootstrap_staker_assets_block
\di idx_bootstrap_delegations_block

# Check data (after some transactions are processed)
SELECT staker_id, asset_id, updated_at, updated_at_block 
FROM bootstrap_staker_assets 
WHERE updated_at_block IS NOT NULL 
LIMIT 5;
```

### GraphQL Verification

Test both tables expose the new field:

```graphql
# Via Hasura Console (http://localhost:8080/console) or curl

query TestUpdatedAtBlock {
  bootstrap_staker_assets(limit: 5) {
    staker_id
    asset_id
    deposited
    withdrawable
    delegated
    updated_at
    updated_at_block
  }
  
  bootstrap_delegation_states(limit: 5) {
    staker_id
    asset_id
    operator_addr
    delegated
    updated_at
    updated_at_block
  }
}
```

### End-to-End Verification

1. Perform a staking operation on the frontend
2. Check that `updated_at_block` is populated after the transaction is processed:

```bash
docker exec imuad_callisto_db psql -U bdjuno -d bdjuno -c \
  "SELECT staker_id, updated_at_block FROM bootstrap_staker_assets ORDER BY updated_at DESC LIMIT 5;"
```

---

## Rollback Procedure

### If Issues Occur After Deployment

#### 1. Rollback Callisto Service

```bash
# Stop current service
docker-compose stop callisto

# Checkout previous version
git checkout <previous-commit-hash>

# Rebuild and restart
docker-compose build callisto
docker-compose up -d callisto
```

#### 2. Rollback Hasura Metadata

```bash
# Checkout previous metadata
git checkout <previous-commit-hash> -- hasura/metadata/

# Reapply metadata
docker-compose up hasura-cli
```

#### 3. Database Rollback (Optional)

The new columns don't break existing functionality, so database rollback is usually unnecessary. If needed:

```bash
docker exec -i imuad_callisto_db psql -U bdjuno -d bdjuno << 'EOF'
DROP INDEX IF EXISTS idx_bootstrap_staker_assets_block;
DROP INDEX IF EXISTS idx_bootstrap_delegations_block;
ALTER TABLE bootstrap_staker_assets DROP COLUMN IF EXISTS updated_at_block;
ALTER TABLE bootstrap_delegation_states DROP COLUMN IF EXISTS updated_at_block;
EOF
```

#### 4. Restore from Backup (Last Resort)

```bash
# Stop services
docker-compose down

# Restore database
docker-compose up -d db
sleep 5
docker exec -i imuad_callisto_db psql -U postgres -d postgres -c "DROP DATABASE IF EXISTS bdjuno;"
docker exec -i imuad_callisto_db psql -U postgres -d postgres -c "CREATE DATABASE bdjuno;"
docker exec -i imuad_callisto_db psql -U bdjuno -d bdjuno < backup_YYYYMMDD_HHMMSS.sql

# Restart services
docker-compose up -d
```

---

## Troubleshooting

### Issue: Database Connection Refused

**Symptom:** Cannot connect to PostgreSQL

**Solution:**
```bash
# Check if database is running
docker-compose ps db

# Start database if not running
docker-compose up -d db

# Check logs
docker logs imuad_callisto_db
```

### Issue: Migration Fails with "column already exists"

**Symptom:** `ERROR: column "updated_at_block" of relation "bootstrap_staker_assets" already exists`

**Solution:** This is fine - the migration uses `IF NOT EXISTS`. The column was already added.

### Issue: `field 'bootstrap_staker_assets' not found in type: 'query_root'`

**Meaning:** Hasura’s loaded metadata doesn’t expose that table — the metadata in `hasura/metadata/` hasn’t been applied (or apply failed).

**Solution — re-apply metadata:**

```bash
cd /path/to/callisto
source .env

# Option A: Run hasura-cli container to apply metadata
docker-compose run --rm hasura-cli bash -c "
  apt-get update -qq && apt-get install -y -qq curl
  curl -L https://github.com/hasura/graphql-engine/releases/download/v2.37.0/cli-hasura-linux-amd64 -o /usr/local/bin/hasura
  chmod +x /usr/local/bin/hasura
  cd /hasura && hasura metadata apply --endpoint http://hasura:8080 --admin-secret \"$HASURA_ADMIN_SECRET\"
"
```

If you have Hasura CLI installed locally:

```bash
cd hasura
hasura metadata apply --endpoint http://localhost:8080 --admin-secret "$HASURA_ADMIN_SECRET"
```

Then retry your GraphQL query.

### Issue: GraphQL Field Not Found (e.g. `updated_at_block`)

**Symptom:** `field "updated_at_block" not found in type`

**Solution:**
```bash
# Load env variables first
source .env

# Reload Hasura metadata (after updating metadata YAML)
curl -X POST \
  -H "X-Hasura-Admin-Secret: $HASURA_ADMIN_SECRET" \
  http://localhost:8080/v1/metadata \
  -d '{"type": "reload_metadata", "args": {}}'

# Or re-apply metadata (see above)
```

### Issue: updated_at_block is Always NULL

**Symptom:** Records show NULL for `updated_at_block`

**Cause:** Records were created before migration or callisto hasn't processed new transactions yet.

**Solution:**
- Wait for new transactions to be processed
- Or trigger a manual refetch by restarting callisto:
  ```bash
  docker-compose restart callisto
  ```

### Issue: Callisto Fails to Start

**Symptom:** Callisto container keeps restarting

**Solution:**
```bash
# Check logs
docker logs imuad_callisto_parser --tail 100

# Common issues:
# - Database not ready: wait and retry
# - Config error: check config.yaml
# - Build error: rebuild with `docker-compose build callisto`
```

---

## Appendix

### Docker Compose Services

| Service | Container Name | Purpose |
|---------|---------------|---------|
| `db` | `imuad_callisto_db` | PostgreSQL database |
| `setup` | `imuad_callisto_db_setup` | Runs schema SQL files |
| `callisto` | `imuad_callisto_parser` | Main indexer service |
| `hasura` | `imuad_callisto_hasura` | GraphQL engine |
| `hasura-cli` | `imuad_callisto_hasura_cli` | Applies Hasura metadata |

### Useful Commands

```bash
# View all logs
docker-compose logs -f

# Restart all services
docker-compose restart

# Stop all services
docker-compose down

# Rebuild everything from scratch
docker-compose down -v
export RESET_DB=true
docker-compose up -d

# Access database shell
docker exec -it imuad_callisto_db psql -U bdjuno -d bdjuno

# Access Hasura console (use HASURA_ADMIN_SECRET from .env to login)
open http://localhost:8080/console
```

### Block Height Sources by Chain

| Chain | Source | Field Used |
|-------|--------|------------|
| ETH (Periodic Refetch) | `EthHTTPClient.BlockNumber()` | Current ETH block |
| ETH (Event Subscription) | `event.Raw.BlockNumber` | Event's block number |
| BTC | `tx.Status.BlockHeight` | Bitcoin block height |
| XRP | `tx.LedgerIndex` | XRP ledger index |

### Related Frontend Changes

Ensure the frontend (`exocore-frontend`) includes:
- `updated_at_block` field in GraphQL queries
- Optimistic update invalidation logic

See frontend files:
- `lib/graphql/queries.ts`
- `lib/graphql/schema.ts`
- `stores/optimisticCacheStore.ts`
