package remote

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v5/node/remote"
	distrtypes "github.com/imua-xyz/imuachain/x/feedistribution/types"
	"strings"

	distrsource "github.com/forbole/callisto/v4/modules/distribution/source"
)

var _ distrsource.Source = &Source{}

// Source implements distrsource.Source querying the data from a remote node
type Source struct {
	*remote.Source
	distrClient distrtypes.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, distrClient distrtypes.QueryClient) *Source {
	return &Source{
		Source:      source,
		distrClient: distrClient,
	}
}

// AVSCommunityPool implements distrsource.Source
func (s Source) AVSCommunityPool(height int64, avsAddr string) (sdk.DecCoins, error) {
	res, err := s.distrClient.AVSCommunityPool(
		remote.GetHeightRequestContext(s.Ctx, height),
		&distrtypes.AVSRequest{Avs: avsAddr},
	)
	if err != nil {
		return nil, err
	}

	return res.FeePool.CommunityPool, nil
}

// Params implements distrsource.Source
func (s Source) Params(height int64) (distrtypes.Params, error) {
	res, err := s.distrClient.Params(
		remote.GetHeightRequestContext(s.Ctx, height),
		&distrtypes.QueryParamsRequest{},
	)
	if err != nil {
		return distrtypes.Params{}, err
	}

	return res.Params, nil
}

func (s Source) StakerAVSClaimedRewards(height int64, stakerID, avs string) (*distrtypes.StakerClaimedRewards, error) {
	res, err := s.distrClient.StakerClaimedRewards(
		remote.GetHeightRequestContext(s.Ctx, height),
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
	res, err := s.distrClient.StakerUnclaimedRewards(
		remote.GetHeightRequestContext(s.Ctx, height),
		&distrtypes.QueryStakerUnclaimedRewardsRequest{StakerId: stakerID},
	)
	if err != nil {
		return nil, err
	}

	return distrtypes.CommonAVSRewards(res.Rewards).RewardsOf(avs), nil
}
