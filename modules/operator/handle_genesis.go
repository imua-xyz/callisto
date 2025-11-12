package operator

import (
	"encoding/json"
	"fmt"
	"github.com/imua-xyz/imuachain/utils"

	tmtypes "github.com/cometbft/cometbft/types"
	junotypes "github.com/forbole/juno/v5/types"

	"github.com/forbole/callisto/v4/types"

	keytypes "github.com/imua-xyz/imuachain/types/keys"
	operatortypes "github.com/imua-xyz/imuachain/x/operator/types"
	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", m.Name()).Msg("parsing genesis")
	var genState operatortypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[operatortypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while reading operator genesis data: %s", err)
	}
	// genesis state is made up of:
	// - operator
	for _, operatorDetail := range genState.Operators {
		if err := m.db.SaveOperatorDetail(types.NewOperator(&operatorDetail.OperatorInfo)); err != nil {
			return fmt.Errorf("error while saving operator detail: %s", err)
		}
	}
	// - operator consensus keys for chain ids
	for _, record := range genState.OperatorRecords {
		addr := record.OperatorAddress
		for _, detail := range record.Chains {
			wrappedKey := keytypes.NewWrappedConsKeyFromHex(detail.ConsensusKey)
			consPubKey, err := junotypes.ConvertValidatorPubKeyToBech32String(wrappedKey.ToTmKey())
			if err != nil {
				return fmt.Errorf("error while converting validator pubkey to bech32 string: %s", err)
			}
			consAddress := wrappedKey.ToConsAddr().String()
			if err := m.db.SaveOperatorConsKey(addr, detail.ChainID, consPubKey, consAddress); err != nil {
				return fmt.Errorf("error while saving operator cons key: %s", err)
			}
			if err := m.db.SaveConsensusKeyAddition(
				addr, detail.ChainID, consPubKey, consAddress, doc.InitialHeight,
			); err != nil {
				return fmt.Errorf("error while saving consensus key addition: %s", err)
			}
		}
	}
	// - operator opted in state
	for _, data := range genState.OptStates {
		keys, err := utils.ParseJoinedKeyWithCount([]byte(data.Key), 2)
		if err != nil {
			return fmt.Errorf("failed to parse joined key: %w", err)
		}
		operatorAddr, avsAddr := keys[0], keys[1]
		opted := types.NewOpted(
			operatorAddr, avsAddr, &data.OptInfo,
		)
		if err := m.db.SaveOptedState(opted); err != nil {
			return fmt.Errorf("error while saving operator opt state: %s", err)
		}
	}
	// - operator usd values
	for _, usdValue := range genState.OperatorUSDValues {
		parsed, err := utils.ParseJoinedKeyWithCount([]byte(usdValue.Key), 2)
		if err != nil {
			return fmt.Errorf("error while parsing operator usd value: %s", err)
		}
		avsAddr, operatorAddr := parsed[0], parsed[1]
		operatorUSDValue := types.NewOperatorUSDValue(
			operatorAddr, avsAddr,
			&usdValue.OptedUSDValue,
		)
		if err := m.db.SaveOperatorUSDValue(operatorUSDValue); err != nil {
			return fmt.Errorf("error while saving operator usd value: %s", err)
		}
	}
	// - avs usd values
	for _, usdValue := range genState.AVSUSDValues {
		avsUsdValue := types.NewAvsUSDValueFromStr(usdValue.AVSAddr, usdValue.Value.String())
		if err := m.db.SaveAvsUSDValue(avsUsdValue); err != nil {
			return fmt.Errorf("error while saving avs usd value: %s", err)
		}
	}

	// - TODO slash states (skipped for now)

	// - prev consensus keys
	// TODO: is this even worth tracking?
	// for any given operator, it represents the key from which they will switch
	// to another within this epoch. it is useful, for example, if you want to
	// track the key to the operator and find signing details.
	for _, prev := range genState.PreConsKeys {
		parsed, err := utils.ParseJoinedKeyWithCount([]byte(prev.Key), 2)
		if err != nil {
			return fmt.Errorf("error while parsing prev cons key: %s", err)
		}
		chainId, operatorAddr := parsed[0], parsed[1]
		wrappedKey := keytypes.NewWrappedConsKeyFromHex(prev.ConsensusKey)
		if err := m.db.SaveOperatorPrevConsKey(
			chainId, operatorAddr, wrappedKey.ToHex(), wrappedKey.ToConsAddr().String(),
		); err != nil {
			return fmt.Errorf("error while saving prev cons key: %s", err)
		}
	}
	// - consensus key removals
	for _, removal := range genState.OperatorKeyRemovals {
		parsed, err := utils.ParseJoinedKeyWithCount([]byte(removal.Key), 2)
		if err != nil {
			return fmt.Errorf("error while parsing operator key removal: %s", err)
		}
		// NOTE the reversed order
		operatorAddr, chainId := parsed[0], parsed[1]
		if err := m.db.MarkOperatorKeyRemoval(chainId, operatorAddr); err != nil {
			return fmt.Errorf("error while removing operator cons key: %s", err)
		}
		if err := m.db.SetConsensusKeyRemovalRequested(chainId, operatorAddr, doc.InitialHeight); err != nil {
			return fmt.Errorf("error while setting consensus key removal requested: %s", err)
		}
	}

	// the constraint can be added only after the genesis.
	// because we populate `operator_assets` under x/assets
	// then we populate `operators` under x/operator, which is performed after x/assets.
	// therefore, any such constraint can only be added after the data is populated.
	// for a live chain, we populate an operator's creation event first, post which,
	// any of their assets are populated (whether via a deposit or a delegation).
	// so the population works as expected.
	if err := m.db.AddOperatorAssetConstraint(); err != nil {
		return fmt.Errorf("error while adding operator asset constraint: %s", err)
	}
	return nil
}
