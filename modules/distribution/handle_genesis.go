package distribution

import (
	"encoding/json"
	"fmt"

	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"

	tmtypes "github.com/cometbft/cometbft/types"

	"github.com/forbole/callisto/v4/types"

	distrtypes "github.com/imua-xyz/imuachain/x/feedistribution/types"
	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "distribution").Msg("parsing genesis")

	// Read the genesis state
	var genState distrtypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[distrtypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while reading distribution genesis data: %s", err)
	}

	// Save the params
	err = m.db.SaveDistributionParams(types.NewDistributionParams(genState.Params, doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis distribution params: %s", err)
	}

	// save reward assets
	for _, avsRewardAssets := range genState.AllAvsRewardAssets {
		for _, rewardAsset := range avsRewardAssets.AvsRewardAssets {
			err := m.db.SaveAVSRewardAsset(&types.AVSRewardAsset{
				AVSAddr: avsRewardAssets.Avs,
				AssetID: rewardAsset.AssetBasicInfo.AssetID(),
				AssetInfo: assetstypes.AssetInfo{
					Name:             rewardAsset.AssetBasicInfo.Name,
					Symbol:           rewardAsset.AssetBasicInfo.Symbol,
					Address:          rewardAsset.AssetBasicInfo.Address,
					Decimals:         rewardAsset.AssetBasicInfo.Decimals,
					LayerZeroChainID: rewardAsset.AssetBasicInfo.LayerZeroChainID,
					ImuaChainIndex:   rewardAsset.AssetBasicInfo.ImuaChainIndex,
					MetaInfo:         rewardAsset.AssetBasicInfo.MetaInfo,
				},
				AVSRewardAssetState: distrtypes.AVSRewardAssetState{
					RewardPoolBalance:     rewardAsset.RewardAssetState.RewardPoolBalance,
					RewardPoolTotal:       rewardAsset.RewardAssetState.RewardPoolTotal,
					RewardAllocationTotal: rewardAsset.RewardAssetState.RewardAllocationTotal,
				},
			})
			if err != nil {
				return fmt.Errorf("error while storing genesis avs reward assets,avs: %s, assetID:%s, err:%s", avsRewardAssets.Avs, rewardAsset.AssetBasicInfo.AssetID(), err)
			}
		}
	}
	return nil
}
