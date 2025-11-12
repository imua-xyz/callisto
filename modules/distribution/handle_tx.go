package distribution

import (
	"fmt"
	"strconv"

	sdkmath "cosmossdk.io/math"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"
	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"
	distrtypes "github.com/imua-xyz/imuachain/x/feedistribution/types"

	"github.com/forbole/callisto/v4/types"
)

// HandleTx implements modules.TransactionModule
func (m *Module) HandleTx(tx *juno.Tx) error {
	// same logic as above; only triggered by tx
	if err := m.handleNewRewardTokenEvents(tx.Events); err != nil {
		return fmt.Errorf("error while handling new token events: %s", err)
	}
	if err := m.handleUpdateRewardTokenEvents(tx.Events); err != nil {
		return fmt.Errorf("error while handling updated token events: %s", err)
	}
	if err := m.handleUpdateRewardTokenStates(tx.Events); err != nil {
		return fmt.Errorf("error while handling updated staking total amount: %s", err)
	}
	if err := m.handleAVSRewardParams(tx.Events, tx.Height); err != nil {
		return fmt.Errorf("error while handling updated staking total amount: %s", err)
	}
	if err := m.handleAVSRewardDistribution(tx.Events); err != nil {
		return fmt.Errorf("error while handling AVS reward distribution events: %s", err)
	}
	if err := m.handleAVSEpochRewardSet(tx.Events); err != nil {
		return fmt.Errorf("error while handling AVS eppoch reward set events: %s", err)
	}
	if err := m.handleAVSOperatorRewardProportionsSet(tx.Events); err != nil {
		return fmt.Errorf("error while handling AVS operator reward prportions set events: %s", err)
	}
	if err := m.handleWithdrawCommissionFromAVS(tx.Events); err != nil {
		return fmt.Errorf("error while handl events about withdrawing commission from AVS: %s", err)
	}
	return nil
}

// handleNewRewardTokenEvents filters, parses and indexes the new reward token events.
func (m *Module) handleNewRewardTokenEvents(events []abci.Event) error {
	events = juno.FindEventsByType(events, distrtypes.EventTypeNewAVSRewardAsset)
	for _, event := range events {
		avsAddr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAvsAddress)
		if err != nil {
			return fmt.Errorf("error while getting AVS address: %s", err)
		}
		assetID, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAssetID)
		if err != nil {
			return fmt.Errorf("error while getting asset ID: %s", err)
		}
		name, err := juno.FindAttributeByKey(event, assetstypes.AttributeKeyName)
		if err != nil {
			return fmt.Errorf("error while getting token name: %s", err)
		}
		symbol, err := juno.FindAttributeByKey(event, assetstypes.AttributeKeySymbol)
		if err != nil {
			return fmt.Errorf("error while getting token symbol: %s", err)
		}
		address, err := juno.FindAttributeByKey(event, assetstypes.AttributeKeyAddress)
		if err != nil {
			return fmt.Errorf("error while getting token address: %s", err)
		}
		decimalsAttr, err := juno.FindAttributeByKey(event, assetstypes.AttributeKeyDecimals)
		if err != nil {
			return fmt.Errorf("error while getting token decimals: %s", err)
		}
		lzIDAttr, err := juno.FindAttributeByKey(event, assetstypes.AttributeKeyLZID)
		if err != nil {
			return fmt.Errorf("error while getting lzID: %s", err)
		}
		metaInfo, err := juno.FindAttributeByKey(event, assetstypes.AttributeKeyMetaInfo)
		if err != nil {
			return fmt.Errorf("error while getting token meta info: %s", err)
		}
		imuachainIndexAttr, err := juno.FindAttributeByKey(event, assetstypes.AttributeKeyImuachainIndex)
		if err != nil {
			return fmt.Errorf("error while getting imuachain index: %s", err)
		}

		// Parse int fields from string
		decimals, err := strconv.Atoi(decimalsAttr.Value)
		if err != nil {
			return fmt.Errorf("invalid decimals value: %s", decimalsAttr.Value)
		}
		lzID, err := strconv.ParseUint(lzIDAttr.Value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid layerZeroChainID: %s", lzIDAttr.Value)
		}
		imuachainIndex, err := strconv.ParseUint(imuachainIndexAttr.Value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid imuachainIndex: %s", imuachainIndexAttr.Value)
		}

		err = m.db.SaveAVSRewardAsset(&types.AVSRewardAsset{
			AVSAddr: avsAddr.Value,
			AssetID: assetID.Value,
			AssetInfo: assetstypes.AssetInfo{
				Name:             name.Value,
				Symbol:           symbol.Value,
				Address:          address.Value,
				Decimals:         uint32(decimals),
				LayerZeroChainID: lzID,
				ImuaChainIndex:   imuachainIndex,
				MetaInfo:         metaInfo.Value,
			},
			AVSRewardAssetState: distrtypes.AVSRewardAssetState{
				RewardPoolBalance:     sdkmath.LegacyZeroDec(), // 0 for reward asset creation
				RewardPoolTotal:       sdkmath.LegacyZeroDec(),
				RewardAllocationTotal: sdkmath.LegacyZeroDec(),
			},
		})
		if err != nil {
			return fmt.Errorf("error while storing new avs reward asset, avs: %s, assetID: %s, err: %s", avsAddr.Value, assetID.Value, err)
		}
	}
	return nil
}

// handleUpdateRewardTokenEvents filters, parses and indexes the updated reward token events.
func (m *Module) handleUpdateRewardTokenEvents(events []abci.Event) error {
	events = juno.FindEventsByType(events, distrtypes.EventTypeUpdatedRewardAssetMetaInfo)
	for _, event := range events {
		avsAddr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAvsAddress)
		if err != nil {
			return fmt.Errorf("error while getting AVS address: %s", err)
		}
		assetID, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAssetID)
		if err != nil {
			return fmt.Errorf("error while getting asset ID: %s", err)
		}
		// only metaInfo can be changed after reward token registration
		metaInfo, err := juno.FindAttributeByKey(event, assetstypes.AttributeKeyMetaInfo)
		if err != nil {
			return fmt.Errorf("error while getting reward token meta info: %s", err)
		}
		if err := m.db.UpdateAVSRewardAssetMetadata(avsAddr.Value, assetID.Value, metaInfo.Value); err != nil {
			return fmt.Errorf("error while updating rewarad token metadata: %s", err)
		}
	}
	return nil
}

// handleUpdateRewardTokenStates filters, parses and indexes the updated AVS reward pool.
// events.
func (m *Module) handleUpdateRewardTokenStates(events []abci.Event) error {
	events = juno.FindEventsByType(events, distrtypes.EventTypeUpdatedAVSRewardAsset)
	for _, event := range events {
		avsAddr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAvsAddress)
		if err != nil {
			return fmt.Errorf("error while getting AVS address: %s", err)
		}
		assetID, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAssetID)
		if err != nil {
			return fmt.Errorf("error while getting asset ID: %s", err)
		}
		rewardPoolBalance, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyRewardPoolBalance)
		if err != nil {
			return fmt.Errorf("error while getting the reward pool balance: %s", err)
		}
		rewardPoolTotal, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyRewardPoolTotal)
		if err != nil {
			return fmt.Errorf("error while getting the reward pool total: %s", err)
		}
		rewardAllocationTotal, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyRewardAllocationTotal)
		if err != nil {
			return fmt.Errorf("error while getting the reward allocation total: %s", err)
		}
		err = m.db.UpdateAVSRewardPool(avsAddr.Value, assetID.Value, rewardPoolBalance.Value, rewardPoolTotal.Value, rewardAllocationTotal.Value)
		if err != nil {
			return fmt.Errorf("error while updating avs reward pool, avs:%s, assetID:%s, err:%s", avsAddr.Value, assetID.Value, err)
		}
	}
	return nil
}

// handleAVSRewardParams filters, parses and indexes the updated AVS reward parameters.
// events.
func (m *Module) handleAVSRewardParams(events []abci.Event, height int64) error {
	events = juno.FindEventsByType(events, distrtypes.EventTypeAVSRewardParamSet)
	for _, event := range events {
		avsAddrAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAvsAddress)
		if err != nil {
			return fmt.Errorf("error while getting AVS address: %s", err)
		}
		paramAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAVSRewardParam)
		if err != nil {
			return fmt.Errorf("error while getting AVS reward param string: %s", err)
		}

		params, err := distrtypes.ParseAVSRewardParams(paramAttr.Value)
		if err != nil {
			return fmt.Errorf("error while parsing AVS reward param string: %s", err)
		}

		if err := m.db.SaveAVSRewardParams(&types.AVSRewardParams{
			AVSAddr: avsAddrAttr.Value,
			AVSRewardParam: distrtypes.AVSRewardParam{
				CustomRewardInflation: params.CustomRewardInflation,
				CustomOperatorRatio:   params.CustomOperatorRatio,
			},
			Height: height,
		}); err != nil {
			return fmt.Errorf("error while saving AVS reward params for avs %s: %s", avsAddrAttr.Value, err)
		}
	}
	return nil
}

// handleAVSRewardDistribution filters, parses, and saves AVS reward distribution events.
func (m *Module) handleAVSRewardDistribution(events []abci.Event) error {
	events = juno.FindEventsByType(events, distrtypes.EventTypeAVSRewardDistributionSet)

	for _, event := range events {
		avsAddrAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAvsAddress)
		if err != nil {
			return fmt.Errorf("failed to get AVS address: %w", err)
		}

		epochIdentifierAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyEpochIdentifier)
		if err != nil {
			return fmt.Errorf("failed to get epoch identifier: %w", err)
		}

		epochNumberAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyEpochNumber)
		if err != nil {
			return fmt.Errorf("failed to get epoch number: %w", err)
		}

		rewardsAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyEpochRewards)
		if err != nil {
			return fmt.Errorf("failed to get rewards: %w", err)
		}

		operatorPropsAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyOperatorProportions)
		if err != nil {
			return fmt.Errorf("failed to get operator proportions: %w", err)
		}

		// Parse values
		epochNumber, err := strconv.ParseInt(epochNumberAttr.Value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid epoch number: %w", err)
		}

		rewards, err := sdk.ParseDecCoins(rewardsAttr.Value)
		if err != nil {
			return fmt.Errorf("failed to parse rewards: %w", err)
		}

		operatorProps, err := distrtypes.ParseOperatorRewardProportions(operatorPropsAttr.Value)
		if err != nil {
			return fmt.Errorf("failed to parse operator proportions: %w", err)
		}

		// Save the parsed distribution to DB
		err = m.db.SaveAVSRewardDistribution(&types.AVSRewardDistribution{
			AVSAddr:         avsAddrAttr.Value,
			EpochIdentifier: epochIdentifierAttr.Value,
			AVSRewardDistribution: distrtypes.AVSRewardDistribution{
				Rewards:                   rewards,
				OperatorRewardProportions: operatorProps,
				RewardsEpochNumber:        epochNumber,
				ProportionsEpochNumber:    epochNumber,
			},
		})
		if err != nil {
			return fmt.Errorf("failed to save AVS reward distribution for %s: %w", avsAddrAttr.Value, err)
		}
	}

	return nil
}

// handleAVSEpochRewardSet filters, parses, and saves AVS epoch reward set events.
func (m *Module) handleAVSEpochRewardSet(events []abci.Event) error {
	events = juno.FindEventsByType(events, distrtypes.EventTypeAVSEpochRewardSet)

	for _, event := range events {
		// Extract attributes
		avsAddrAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAvsAddress)
		if err != nil {
			return fmt.Errorf("failed to get AVS address: %w", err)
		}

		epochIdentifierAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyEpochIdentifier)
		if err != nil {
			return fmt.Errorf("failed to get epoch identifier: %w", err)
		}

		epochNumberAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyEpochNumber)
		if err != nil {
			return fmt.Errorf("failed to get epoch number: %w", err)
		}

		rewardsAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyEpochRewards)
		if err != nil {
			return fmt.Errorf("failed to get rewards: %w", err)
		}

		// Parse attributes
		epochNumber, err := strconv.ParseInt(epochNumberAttr.Value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid epoch number: %w", err)
		}

		rewards, err := sdk.ParseDecCoins(rewardsAttr.Value)
		if err != nil {
			return fmt.Errorf("failed to parse rewards: %w", err)
		}

		// Save to DB
		err = m.db.UpsertAVSEpochRewards(
			avsAddrAttr.Value, epochIdentifierAttr.Value,
			epochNumber, rewards)
		if err != nil {
			return fmt.Errorf("failed to save AVS epoch reward for %s: %w", avsAddrAttr.Value, err)
		}
	}

	return nil
}

// handleAVSOperatorRewardProportionsSet filters, parses, and saves AVS operator reward proportions set events.
func (m *Module) handleAVSOperatorRewardProportionsSet(events []abci.Event) error {
	events = juno.FindEventsByType(events, distrtypes.EventTypeAVSRewardProportionsSet)

	for _, event := range events {
		// Extract attributes
		avsAddrAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAvsAddress)
		if err != nil {
			return fmt.Errorf("failed to get AVS address: %w", err)
		}

		epochIdentifierAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyEpochIdentifier)
		if err != nil {
			return fmt.Errorf("failed to get epoch identifier: %w", err)
		}

		epochNumberAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyEpochNumber)
		if err != nil {
			return fmt.Errorf("failed to get epoch number: %w", err)
		}

		operatorPropsAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyOperatorProportions)
		if err != nil {
			return fmt.Errorf("failed to get operator reward proportions: %w", err)
		}

		// Parse attributes
		epochNumber, err := strconv.ParseInt(epochNumberAttr.Value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid epoch number: %w", err)
		}

		operatorProps, err := distrtypes.ParseOperatorRewardProportions(operatorPropsAttr.Value)
		if err != nil {
			return fmt.Errorf("failed to parse operator proportions: %w", err)
		}

		// Save to DB
		err = m.db.UpsertOperatorRewardProportions(
			avsAddrAttr.Value, epochIdentifierAttr.Value,
			operatorProps, epochNumber)
		if err != nil {
			return fmt.Errorf("failed to save AVS operator reward proportions for %s: %w", avsAddrAttr.Value, err)
		}
	}

	return nil
}

// handleWithdrawCommissionFromAVS filters, parses, and stores events emitted
// when an operator withdraws commission from AVS.
func (m *Module) handleWithdrawCommissionFromAVS(events []abci.Event) error {
	events = juno.FindEventsByType(events, distrtypes.EventTypeWithdrawCommissionFromAVS)

	for _, event := range events {
		// Extract attributes
		operatorAddrAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyOperator)
		if err != nil {
			return fmt.Errorf("failed to get operator address: %w", err)
		}

		avsAddrAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAvsAddress)
		if err != nil {
			return fmt.Errorf("failed to get AVS address: %w", err)
		}

		withdrawnDecCoins, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyWithdrawDecCoinsFromAVS)
		if err != nil {
			return fmt.Errorf("failed to get the withdrawn decCoins: %w", err)
		}

		// Parse DecCoins
		withdrawnCommission, err := sdk.ParseDecCoins(withdrawnDecCoins.Value)
		if err != nil {
			return fmt.Errorf("failed to parse the withdraw amounts: %w", err)
		}

		// update withdrawn commission in the database
		err = m.db.WithdrawOperatorCommission(
			operatorAddrAttr.Value,
			avsAddrAttr.Value,
			withdrawnCommission,
		)
		if err != nil {
			return fmt.Errorf("failed to save withdrawn commission for operator %s and avs %s: %w",
				operatorAddrAttr.Value, avsAddrAttr.Value, err)
		}
	}

	return nil
}
