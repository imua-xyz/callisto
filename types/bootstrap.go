package types

import (
	"time"
)

type BootstrapValidator struct {
	ValidatorEthAddress string
	ValidatorIMAddress  string
	ValidatorName       string
	ConsensusPubKey     string
	Rate                string
	MaxRate             string
	MaxChangeRate       string
	UpdatedAt           time.Time
}

type BootstrapClientChain struct {
	Name      string `yaml:"name"`
	MetaInfo  string `yaml:"meta_info"`
	LZChainID uint64 `yaml:"lz_chain_id"`
}

type BootstrapToken struct {
	AssetID   string `yaml:"asset_id"`
	Name      string `yaml:"name"`
	Symbol    string `yaml:"symbol"`
	Address   string `yaml:"address"`
	Decimals  uint8  `yaml:"decimals"`
	LZChainID uint64 `yaml:"lz_chain_id"`
}
type BootstrapTokenState struct {
	BootstrapToken
	StakingTotalAmount string
	TotalUSDValue      string
	UpdatedAt          time.Time
}

type OracleFeed struct {
	AssetID    string `yaml:"asset_id"`
	OracleAddr string `yaml:"oracle_addr"` // contract address of Chainlink aggregator
}

type BootstrapStakerAsset struct {
	StakerID       string
	AssetID        string
	Deposited      string
	Withdrawable   string
	Delegated      string
	UpdatedAt      time.Time
	UpdatedAtBlock int64 // Block height when this record was last updated (for optimistic update invalidation)
}

type BootstrapDelegationState struct {
	StakerID       string
	AssetID        string
	OperatorAddr   string
	Delegated      string
	UpdatedAt      time.Time
	UpdatedAtBlock int64 // Block height when this record was last updated (for optimistic update invalidation)
}

type BootstrapOperatorAsset struct {
	OperatorAddr string
	AssetID      string
	TotalAmount  string
	SelfAmount   string
	OtherAmount  string
	UpdatedAt    time.Time
}

// BTC transaction related types
type BTCTx struct {
	TxID    string    `json:"txid"`
	Status  BTCStatus `json:"status"`
	Vin     []BTCVin  `json:"vin"`
	Vout    []BTCVout `json:"vout"`
	TxIndex int64     // Transaction index within block (filled when needed)
	// Parsed OP_RETURN data for bootstrap transactions
	ImuachainAddress string `json:"imuachain_address,omitempty"`
	ValidatorAddress string `json:"validator_address,omitempty"`
}

type BTCStatus struct {
	Confirmed   bool   `json:"confirmed"`
	BlockHeight int64  `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   int64  `json:"block_time"`
}

type BTCVin struct {
	Prevout BTCPrevout `json:"prevout"`
}

type BTCPrevout struct {
	ScriptPubKeyAddr string `json:"scriptpubkey_address"`
}

type BTCVout struct {
	ScriptPubKey     string `json:"scriptpubkey"`
	ScriptPubKeyType string `json:"scriptpubkey_type"`
	ScriptPubKeyAddr string `json:"scriptpubkey_address"`
	Value            int64  `json:"value"`
}

type BTCBlockTip struct {
	Height int64 `json:"height"`
}

// XRP transaction related types
type XRPTransaction struct {
	Hash        string  `json:"hash"`
	LedgerIndex int64   `json:"ledger_index"`
	Date        int64   `json:"date"`
	Validated   bool    `json:"validated"`
	Tx          XRPTx   `json:"tx"`
	Meta        XRPMeta `json:"meta"`
	// Parsed memo data for bootstrap transactions
	ImuachainAddress string `json:"imuachain_address,omitempty"`
	ValidatorAddress string `json:"validator_address,omitempty"`
}

type XRPTx struct {
	TransactionType string      `json:"TransactionType"`
	Account         string      `json:"Account"`
	Destination     string      `json:"Destination,omitempty"`
	Amount          interface{} `json:"Amount"`
	Fee             string      `json:"Fee"`
	Sequence        int64       `json:"Sequence"`
	Memos           []XRPMemo   `json:"Memos,omitempty"`
	DestinationTag  int64       `json:"DestinationTag,omitempty"`
}

type XRPMemo struct {
	Memo XRPMemoData `json:"Memo"`
}

type XRPMemoData struct {
	MemoType   string `json:"MemoType"`
	MemoData   string `json:"MemoData"`
	MemoFormat string `json:"MemoFormat"`
}

type XRPMeta struct {
	TransactionResult string `json:"TransactionResult"`
	TransactionIndex  int64  `json:"TransactionIndex"`
	DeliveredAmount   string `json:"delivered_amount,omitempty"`
}

type XRPCurrentLedger struct {
	LedgerIndex int64 `json:"ledger_index"`
}

// Incremental scanning related types
type ScanState struct {
	ChainType  string    `db:"chain_type"`
	LastHeight int64     `db:"last_height"`
	LastHash   string    `db:"last_hash"`
	SafeHeight int64     `db:"safe_height"`
	UpdatedAt  time.Time `db:"updated_at"`
	CreatedAt  time.Time `db:"created_at"`
}

type ProcessedTransaction struct {
	ChainType   string    `db:"chain_type"`
	TxHash      string    `db:"tx_hash"`
	BlockHeight int64     `db:"block_height"`
	ProcessedAt time.Time `db:"processed_at"`
}

// Address binding related types
type AddressBinding struct {
	ChainType  string    `db:"chain_type"`
	SourceAddr string    `db:"source_addr"`
	TargetAddr string    `db:"target_addr"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// Transaction batch processing
type TransactionBatch struct {
	ID           string    `db:"id"`
	ChainType    string    `db:"chain_type"`
	StartHeight  int64     `db:"start_height"`
	EndHeight    int64     `db:"end_height"`
	Status       string    `db:"status"`
	TotalTxs     int       `db:"total_txs"`
	ProcessedTxs int       `db:"processed_txs"`
	FailedTxs    int       `db:"failed_txs"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// Transaction retry configuration
type RetryConfig struct {
	MaxRetries      int           `json:"max_retries"`
	InitialDelay    time.Duration `json:"initial_delay"`
	MaxDelay        time.Duration `json:"max_delay"`
	BackoffMultiple float64       `json:"backoff_multiple"`
}
