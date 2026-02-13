package slashing

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cometbft/cometbft/types"
	"github.com/rs/zerolog/log"
)

// HandleBlock implements BlockModule.
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, results *tmctypes.ResultBlockResults, _ []*juno.Tx, vals *tmctypes.ResultValidators,
) error {
	log.Debug().Str("module", m.Name()).Int64("height", block.Block.Height).
		Msg(fmt.Sprintf("updating %s", m.Name()))
	// Update the signing infos
	err := m.updateSigningInfo(block.Block.Height)
	if err != nil {
		return fmt.Errorf("error while updating signing info: %s", err)
	}
	err = m.handleUptime(block.Block)
	if err != nil {
		return fmt.Errorf("error while handling uptime: %s", err)
	}
	return nil
}

// updateSigningInfo reads from the LCD the current staking pool and stores its value inside the database
func (m *Module) updateSigningInfo(height int64) error {
	signingInfos, err := m.getSigningInfos(height)
	if err != nil {
		return err
	}

	return m.db.SaveValidatorsSigningInfos(signingInfos)
}

// handleUptime records the signing (or lack thereof) for
// each validator who should have signed the block.
func (m *Module) handleUptime(block *types.Block) error {
	commit := block.LastCommit
	// for the initial height, there is no last commit.
	if commit == nil {
		return nil
	}
	effectiveHeight := block.Height - 1
	// int32 <= int across all platforms
	size := int32(commit.Size()) // #nosec G115
	for i := int32(0); i < size; i++ {
		// iterate only through validators who should have signed
		vote := commit.GetVote(i)
		signature := commit.Signatures[i]
		// determine uptime
		signed := signature.BlockIDFlag != types.BlockIDFlagAbsent
		address := sdk.ConsAddress(vote.ValidatorAddress.Bytes())
		// save the uptime
		err := m.db.SaveBlockUptime(
			effectiveHeight, address.String(), signed,
		)
		if err != nil {
			return fmt.Errorf("error while saving uptime: %s", err)
		}
	}
	return nil
}
