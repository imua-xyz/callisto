package consensus

import (
	"fmt"

	"github.com/forbole/juno/v5/types"

	"github.com/rs/zerolog/log"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
)

// HandleBlock implements modules.Module
func (m *Module) HandleBlock(
	b *tmctypes.ResultBlock, _ *tmctypes.ResultBlockResults, _ []*types.Tx, _ *tmctypes.ResultValidators,
) error {
	log.Debug().Str("module", m.Name()).Int64("height", b.Block.Height).
		Msg(fmt.Sprintf("updating %s", m.Name()))
	err := m.updateBlockTimeFromGenesis(b)
	if err != nil {
		log.Error().Str("module", "consensus").Int64("height", b.Block.Height).
			Err(err).Msg("error while updating block time from genesis")
	}

	return nil
}

// updateBlockTimeFromGenesis insert average block time from genesis
func (m *Module) updateBlockTimeFromGenesis(block *tmctypes.ResultBlock) error {
	genesis, err := m.db.GetGenesis()
	if err != nil {
		return fmt.Errorf("error while getting genesis: %s", err)
	}
	if genesis == nil {
		return fmt.Errorf("genesis table is empty")
	}

	newBlockTime := block.Block.Time.Sub(genesis.Time).Seconds() / float64(block.Block.Height-genesis.InitialHeight)
	return m.db.SaveAverageBlockTimeGenesis(newBlockTime, block.Block.Height)
}
