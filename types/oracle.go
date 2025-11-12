package types

import (
	"fmt"

	oracletypes "github.com/imua-xyz/imuachain/x/oracle/types"
)

// OracleParams represents the x/oracle parameters
type OracleParams struct {
	oracletypes.Params
	Height int64
}

// NewOracleParams allows to build a new OracleParams instance
func NewOracleParams(params oracletypes.Params, height int64) *OracleParams {
	return &OracleParams{
		Params: params,
		Height: height,
	}
}

type OracleTokenConfig struct {
	TokenID     string
	NextRoundID string
}

// NewOracleTokenConfig allows to build a new OracleTokenConfig instance
func NewOracleTokenConfig(
	prices oracletypes.Prices,
) *OracleTokenConfig {
	return &OracleTokenConfig{
		TokenID:     fmt.Sprintf("%d", prices.TokenID),
		NextRoundID: fmt.Sprintf("%d", prices.NextRoundID),
	}
}

// NewOracleTokenConfigFromStr allows to build a new OracleTokenConfig instance
// from the given string values
func NewOracleTokenConfigFromStr(
	tokenID string,
	nextRoundID string,
) *OracleTokenConfig {
	return &OracleTokenConfig{
		TokenID:     tokenID,
		NextRoundID: nextRoundID,
	}
}

type OraclePriceHistory struct {
	TokenID        string
	RoundID        string
	Price          string
	PriceDecimals  string
	PriceTimestamp string
}

// NewOraclePriceHistory allows to build a new OraclePriceHistory instance
func NewOraclePriceHistory(
	oracleTokenConfig *OracleTokenConfig,
	priceTimeRound *oracletypes.PriceTimeRound,
) *OraclePriceHistory {
	return &OraclePriceHistory{
		TokenID:        oracleTokenConfig.TokenID,
		RoundID:        fmt.Sprintf("%d", priceTimeRound.RoundID),
		Price:          priceTimeRound.Price,
		PriceDecimals:  fmt.Sprintf("%d", priceTimeRound.Decimal),
		PriceTimestamp: priceTimeRound.Timestamp,
	}
}

// NewOraclePriceHistoryFromStr allows to build a new OraclePriceHistory instance
// from the given string values
func NewOraclePriceHistoryFromStr(
	tokenID string,
	roundID string,
	price string,
	priceDecimals string,
	priceTimestamp string,
) *OraclePriceHistory {
	return &OraclePriceHistory{
		TokenID:        tokenID,
		RoundID:        roundID,
		Price:          price,
		PriceDecimals:  priceDecimals,
		PriceTimestamp: priceTimestamp,
	}
}
