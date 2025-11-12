package tokenomics

import (
	"fmt"

	sdkmath "cosmossdk.io/math"

	"gopkg.in/yaml.v3"
)

type IndexerLegacyDec struct {
	sdkmath.LegacyDec
}

func (d *IndexerLegacyDec) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err == nil {
		dec, err := sdkmath.LegacyNewDecFromStr(s)
		if err != nil {
			return err
		}
		*d = IndexerLegacyDec{dec}
		return nil
	}
	// Try to unmarshal as a number, fallback if string fails
	var f float64
	if err := unmarshal(&f); err == nil {
		dec, err := sdkmath.LegacyNewDecFromStr(fmt.Sprintf("%.18f", f))
		if err != nil {
			return err
		}
		*d = IndexerLegacyDec{dec}
		return nil
	}
	return fmt.Errorf("failed to unmarshal IndexerLegacyDec from YAML")
}

// Config defines all the configurable parameters for the airdrop, incentive, and tokenomics system.
type Config struct {
	// GenesisSupply defines the total token supply at genesis (integer, in the smallest unit).
	GenesisSupply int64 `json:"genesis_supply" yaml:"genesis_supply"`

	// ------------------------------------------------------------
	// Genesis Pool Configuration (measured in minutes)
	// Although the configuration is typically defined in days, it is measured in minutes here
	// for finer granularity and easier local testing
	// ------------------------------------------------------------
	// GenesisPoolRatio defines the proportion of genesis supply allocated to the genesis pool.
	// For example: 0.03 means 3% of the genesis supply.
	GenesisPoolRatio IndexerLegacyDec `json:"genesis_pool_ratio" yaml:"genesis_pool_ratio"`

	// GenesisPoolAirdropDuration defines how long the genesis pool airdrop lasts after TGE.
	// For example: 90*24*60 -> lasts for 90 days.
	GenesisPoolAirdropDuration int64 `json:"genesis_pool_airdrop_duration" yaml:"genesis_pool_airdrop_duration"`

	// GenesisPoolAirdropInterval defines how often to execute a genesis pool airdrop.
	// For example: 7*24*60 -> every 7 days (weekly).
	GenesisPoolAirdropInterval int64 `json:"genesis_pool_airdrop_interval" yaml:"genesis_pool_airdrop_interval"`

	// ------------------------------------------------------------
	// Liquidity Incentive Configuration (measured in minutes)
	// Although the configuration is typically defined in years, it is measured in minutes here
	// for finer granularity and easier local testing
	// ------------------------------------------------------------

	// LiquidityIncentiveRatios defines the annual ratios of liquidity incentives to be distributed.
	// Each entry represents the proportion of the genesis supply allocated for that year.
	// Example: [0.02, 0.01, 0.005] means:
	//   Year 1 -> 2%
	//   Year 2 -> 1%
	//   Year 3 -> 0.5%
	LiquidityIncentiveRatios []IndexerLegacyDec `json:"liquidity_incentive_ratios" yaml:"liquidity_incentive_ratios"`

	// LiquidityIncentiveAirdropDuration defines how long the liquidity incentive airdrop lasts.
	// Example:
	// 	20*365*24*60  -> lasts for 20 years
	LiquidityIncentiveAirdropDuration int64 `json:"liquidity_incentive_airdrop_duration" yaml:"liquidity_incentive_airdrop_duration"`

	// LiquidityIncentiveAirdropInterval defines how often (in years) to execute a liquidity airdrop.
	// It should be a divisor of one(year) to prevent airdrop rounds from crossing over into the next
	// year, which could cause reward calculation errors.
	// Example:
	//	 0.13 -> invalid value
	//   0.25 -> every quarter
	//   0.5  -> every half year
	//   1.0  -> annually
	LiquidityIncentiveAirdropInterval IndexerLegacyDec `json:"liquidity_incentive_airdrop_interval" yaml:"liquidity_incentive_airdrop_interval"`
}

func (c *Config) Validate() error {
	// The total supply at genesis must be positive.
	if c.GenesisSupply < 0 {
		return fmt.Errorf("genesis supply must be positive, got %d", c.GenesisSupply)
	}

	// The genesis pool ratio must be between 0 and 1.
	if c.GenesisPoolRatio.IsNegative() || c.GenesisPoolRatio.GT(sdkmath.LegacyOneDec()) {
		return fmt.Errorf("genesis pool ratio must be between 0 and 1, got %s", c.GenesisPoolRatio)
	}

	// The genesis pool airdrop duration must not be negative.
	if c.GenesisPoolAirdropDuration < 0 {
		return fmt.Errorf("genesis pool airdrop duration must not be negative, got %d", c.GenesisPoolAirdropDuration)
	}

	// The airdrop interval must not be less than 0 and less than or equal to the total duration.
	if c.GenesisPoolAirdropInterval < 0 || c.GenesisPoolAirdropInterval > c.GenesisPoolAirdropDuration {
		return fmt.Errorf("invalid genesis pool airdrop interval: %d (must be >= 0 and <= duration %d)",
			c.GenesisPoolAirdropInterval, c.GenesisPoolAirdropDuration)
	}

	// Each ratio must be within the range [0, 1].
	for i, ratio := range c.LiquidityIncentiveRatios {
		if ratio.IsNegative() || ratio.GT(sdkmath.LegacyOneDec()) {
			return fmt.Errorf("invalid liquidity incentive ratio at index %d: %s (must be between 0 and 1)", i, ratio)
		}
	}

	// The total duration of the liquidity incentive program must not be negative.
	if c.LiquidityIncentiveAirdropDuration < 0 {
		return fmt.Errorf("liquidity incentive airdrop duration must not be negative, got %d", c.LiquidityIncentiveAirdropDuration)
	}

	// The airdrop interval must be within [0, 1].
	if c.LiquidityIncentiveAirdropInterval.IsNegative() || c.LiquidityIncentiveAirdropInterval.GT(sdkmath.LegacyOneDec()) {
		return fmt.Errorf("invalid liquidity incentive airdrop interval: %s (must be >= 0 and <= 1)",
			c.LiquidityIncentiveAirdropInterval)
	}

	// The airdrop interval must evenly divide one year to prevent cross-year rounding errors.
	if !sdkmath.LegacyOneDec().Quo(c.LiquidityIncentiveAirdropInterval.LegacyDec).IsInteger() {
		return fmt.Errorf("liquidity incentive airdrop interval %s must evenly divide one year", c.LiquidityIncentiveAirdropInterval)
	}
	return nil
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		GenesisSupply:              314159265,                                                   // from tokenomics documentation
		GenesisPoolRatio:           IndexerLegacyDec{sdkmath.LegacyMustNewDecFromStr("0.0300")}, // 3% of genesis supply
		GenesisPoolAirdropDuration: 90 * 24 * 60,                                                // 90 days after TGE
		GenesisPoolAirdropInterval: 7 * 24 * 60,                                                 // every 7 days (weekly)

		LiquidityIncentiveRatios: []IndexerLegacyDec{
			{sdkmath.LegacyMustNewDecFromStr("0.0200")},
			{sdkmath.LegacyMustNewDecFromStr("0.0100")},
			{sdkmath.LegacyMustNewDecFromStr("0.0050")},
			{sdkmath.LegacyMustNewDecFromStr("0.0025")},
			{sdkmath.LegacyMustNewDecFromStr("0.0013")},
		},
		LiquidityIncentiveAirdropDuration: 20 * 365 * 24 * 60,                                        // 20 years
		LiquidityIncentiveAirdropInterval: IndexerLegacyDec{sdkmath.LegacyMustNewDecFromStr("0.25")}, // every quarter (90 days ≈ 0.25 years)
	}
}

func ParseConfig(bz []byte) (*Config, error) {
	type T struct {
		Config *Config `yaml:"tokenomics"`
	}
	var cfg T
	err := yaml.Unmarshal(bz, &cfg)
	if err != nil {
		return nil, fmt.Errorf("parse tokenomics config: %w", err)
	}
	if cfg.Config == nil {
		return DefaultConfig(), nil
	}

	err = cfg.Config.Validate()
	if err != nil {
		return nil, err
	}
	return cfg.Config, nil
}
