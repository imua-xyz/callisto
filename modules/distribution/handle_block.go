package distribution

import (
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"
	distrtypes "github.com/imua-xyz/imuachain/x/feedistribution/types"
	"github.com/rs/zerolog/log"
)

// HandleBlock implements BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, res *tmctypes.ResultBlockResults, _ []*juno.Tx, _ *tmctypes.ResultValidators,
) error {
	log.Debug().Str("module", m.Name()).Int64("height", block.Block.Height).
		Msg(fmt.Sprintf("updating %s", m.Name()))
	// the events about allocating rewards to operator are emitted during EndBlock
	if err := m.handleAllocateRewardsToOperator(res.EndBlockEvents); err != nil {
		return fmt.Errorf("error while handling events for allocating rewards to operator: %s", err)
	}
	return nil
}

// handleAllocateRewardsToOperator filters, parses, and stores events emitted
// when rewards are allocated to an operator.
func (m *Module) handleAllocateRewardsToOperator(events []abci.Event) error {
	events = juno.FindEventsByType(events, distrtypes.EventTypeAllocateRewardsToOperator)

	for _, event := range events {
		// Extract attributes
		avsAddrAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAvsAddress)
		if err != nil {
			return fmt.Errorf("failed to get AVS address: %w", err)
		}

		operatorAddrAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyOperator)
		if err != nil {
			return fmt.Errorf("failed to get operator address: %w", err)
		}

		totalRewardAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyOperatorTotalReward)
		if err != nil {
			return fmt.Errorf("failed to get operator total reward: %w", err)
		}

		commissionAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyOperatorCommission)
		if err != nil {
			return fmt.Errorf("failed to get operator commission: %w", err)
		}

		// Parse rewards and commission
		totalReward, err := sdk.ParseDecCoins(totalRewardAttr.Value)
		if err != nil {
			return fmt.Errorf("failed to parse operator total reward: %w", err)
		}

		commission, err := sdk.ParseDecCoins(commissionAttr.Value)
		if err != nil {
			return fmt.Errorf("failed to parse operator commission: %w", err)
		}

		// Upsert into operator_rewards table
		err = m.db.UpsertOperatorRewards(
			operatorAddrAttr.Value,
			avsAddrAttr.Value,
			totalReward,
			commission,
		)
		if err != nil {
			return fmt.Errorf("failed to upsert operator rewards for %s/%s: %w",
				operatorAddrAttr.Value, avsAddrAttr.Value, err)
		}
	}

	return nil
}
