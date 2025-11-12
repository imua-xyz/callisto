package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"
	distrtypes "github.com/imua-xyz/imuachain/x/feedistribution/types"
)

// DistributionParams represents the parameters of the x/distribution module
type DistributionParams struct {
	distrtypes.Params
	Height int64
}

// NewDistributionParams allows to build a new DistributionParams instance
func NewDistributionParams(params distrtypes.Params, height int64) *DistributionParams {
	return &DistributionParams{
		Params: params,
		Height: height,
	}
}

type AVSRewardAsset struct {
	AVSAddr string // avs_addr
	AssetID string // asset_id
	assetstypes.AssetInfo
	distrtypes.AVSRewardAssetState
}

type AVSRewardParams struct {
	AVSAddr string
	distrtypes.AVSRewardParam
	Height int64
}

type AVSRewardDistribution struct {
	AVSAddr         string
	EpochIdentifier string
	distrtypes.AVSRewardDistribution
}

type StakerRewards struct {
	StakerID string
	AVSAddr  string
	distrtypes.StakerClaimedRewards
	ClaimedRewards   sdk.DecCoins
	UnclaimedRewards sdk.DecCoins
	TotalRewards     sdk.DecCoins
	RewardUpdateTime time.Time
}
