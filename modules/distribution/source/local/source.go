package local

import (
	"fmt"
	"strings"

	distrtypes "github.com/imua-xyz/imuachain/x/feedistribution/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v5/node/local"

	distrsource "github.com/forbole/callisto/v4/modules/distribution/source"
)

var _ distrsource.Source = &Source{}

// Source implements distrsource.Source reading the data from a local node
type Source struct {
	*local.Source
	q distrtypes.QueryServer
}

func NewSource(source *local.Source, keeper distrtypes.QueryServer) *Source {
	return &Source{
		Source: source,
		q:      keeper,
	}
}

// AVSCommunityPool implements distrsource.Source
func (s Source) AVSCommunityPool(height int64, avsAddr string) (sdk.DecCoins, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.AVSCommunityPool(sdk.WrapSDKContext(ctx), &distrtypes.AVSRequest{Avs: avsAddr})
	if err != nil {
		return nil, err
	}

	return res.FeePool.CommunityPool, nil
}

// Params implements distrsource.Source
func (s Source) Params(height int64) (distrtypes.Params, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return distrtypes.Params{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Params(sdk.WrapSDKContext(ctx), &distrtypes.QueryParamsRequest{})
	if err != nil {
		return distrtypes.Params{}, err
	}

	return res.Params, nil
}

func (s Source) StakerAVSClaimedRewards(height int64, stakerID, avs string) (*distrtypes.StakerClaimedRewards, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}
	res, err := s.q.StakerClaimedRewards(
		sdk.WrapSDKContext(ctx),
		&distrtypes.QueryStakerClaimedRewardsRequest{
			StakerId: stakerID,
			Avs:      avs,
		},
	)

	if err != nil {
		if strings.Contains(err.Error(), distrtypes.ErrNoKeyInTheStore.Error()) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return res.StakerClaimedRewards, nil
}

func (s Source) StakerAVSUnclaimedRewards(height int64, stakerID, avs string) (sdk.DecCoins, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}
	res, err := s.q.StakerUnclaimedRewards(
		sdk.WrapSDKContext(ctx),
		&distrtypes.QueryStakerUnclaimedRewardsRequest{StakerId: stakerID},
	)
	if err != nil {
		return nil, err
	}

	return distrtypes.CommonAVSRewards(res.Rewards).RewardsOf(avs), nil
}
