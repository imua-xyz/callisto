package avs

import (
	"encoding/json"
	"fmt"
	"github.com/imua-xyz/imuachain/utils"
	"strings"

	tmtypes "github.com/cometbft/cometbft/types"

	avstypes "github.com/imua-xyz/imuachain/x/avs/types"
	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", m.Name()).Msg("parsing genesis")

	// step 1: chainID to avs address mapping
	// TODO: compare against handleDogfoodAvsCreationEvents and why does it not work?
	// repeats are permitted by these SaveXXX functions, so it's ok to add them manually
	// if they exist or not in the genesis.
	chainID := utils.ChainIDWithoutRevision(doc.ChainID)
	avsAddr := utils.GenerateAVSAddress(chainID)
	if err := m.db.SaveAvsAddr(avsAddr); err != nil {
		return fmt.Errorf("error while saving avs address: %s", err)
	}
	if err := m.db.SaveChainIdToAvsAddr(chainID, avsAddr); err != nil {
		return fmt.Errorf("error while saving chain id to avs address: %s", err)
	}

	// step 2: as-is genesis data
	var state avstypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[avstypes.ModuleName], &state)
	if err != nil {
		return fmt.Errorf("error while reading avs genesis data: %s", err)
	}
	// TODO: handle this completely
	for _, avs := range state.AvsInfos {
		avs.AvsAddress = strings.ToLower(avs.AvsAddress)
		if err := m.db.SaveAvsAddr(avs.AvsAddress); err != nil {
			return fmt.Errorf("error while saving avs address: %s", err)
		}
	}
	for _, elem := range state.ChainIdInfos {
		elem.AvsAddress = strings.ToLower(elem.AvsAddress)
		if err := m.db.SaveChainIdToAvsAddr(elem.ChainId, elem.AvsAddress); err != nil {
			return fmt.Errorf("error while saving chain id to avs address: %s", err)
		}
	}
	return nil
}
