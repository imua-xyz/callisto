package bootstrap

import (
	"fmt"
	"math/big"
	"time"

	"github.com/forbole/callisto/v4/types"
	"github.com/rs/zerolog/log"
)

// loadExistingBindings loads existing address bindings from database into memory
func (m *Module) loadExistingBindings() error {
	// Load BTC bindings
	btcBindings, err := m.database.GetAddressBindings("BTC")
	if err != nil {
		return fmt.Errorf("failed to load BTC address bindings: %w", err)
	}

	for _, binding := range btcBindings {
		m.btcAddressMappings[binding.SourceAddr] = binding.TargetAddr
	}

	// Load XRP bindings
	xrpBindings, err := m.database.GetAddressBindings("XRP")
	if err != nil {
		return fmt.Errorf("failed to load XRP address bindings: %w", err)
	}

	for _, binding := range xrpBindings {
		m.xrpAddressMappings[binding.SourceAddr] = binding.TargetAddr
	}

	return nil
}

// RunAdditionalOperations implements modules.AdditionalOperationsModule
func (m *Module) RunAdditionalOperations() error {
	// Save default client chains from the config
	for _, clientChain := range m.Config.ClientChainInfos {
		_, err := m.database.GetBootstrapClientChain(clientChain.LZChainID)
		if err != nil {
			err = m.database.SaveBootstrapClientChain(&types.BootstrapClientChain{
				Name:      clientChain.Name,
				MetaInfo:  clientChain.MetaInfo,
				LZChainID: clientChain.LZChainID,
			})
			if err != nil {
				return err
			}
		}
	}
	// save default staking tokens from the config
	for _, stakingToken := range m.Config.StakingTokenInfos {
		_, err := m.database.GetBootstrapToken(stakingToken.AssetID)
		if err != nil {
			err = m.database.SaveBootstrapToken(&types.BootstrapTokenState{
				BootstrapToken:     stakingToken,
				StakingTotalAmount: big.NewInt(0).String(),
				TotalUSDValue:      big.NewInt(0).String(),
				UpdatedAt:          time.Now(),
			})
			if err != nil {
				return err
			}
		}
	}
	if err := m.refetchETHStates(); err != nil {
		return err
	}
	if err := m.updatePricesAndTVL(); err != nil {
		return err
	}
	// Initialize address bindings from database
	if err := m.loadExistingBindings(); err != nil {
		return fmt.Errorf("failed to load existing address bindings: %w", err)
	}
	log.Info().Str("module", "bootstrap").Msg("Complete Ethereum state sync on startup.")
	return nil
}
