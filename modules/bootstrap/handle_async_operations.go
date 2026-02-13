package bootstrap

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	sdkmath "cosmossdk.io/math"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forbole/callisto/v4/modules/bootstrap/bootstrap_binding"
	"github.com/forbole/callisto/v4/types"
	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"
	"github.com/rs/zerolog/log"
)

func (m *Module) updateStatesAfterDepositAndClaiming(stakerAddr, assetAddr common.Address, blockHeight int64) error {
	_, assetID, err := m.updateStakerAsset(stakerAddr, assetAddr, blockHeight)
	if err != nil {
		return err
	}

	// update the total deposit amount in asset states
	assetDepositAmount, err := m.bootstrapSession.DepositsByToken(assetAddr)
	if err != nil {
		return err
	}

	usdValue := sdkmath.LegacyZeroDec()
	priceStr, err := m.database.GetBootstrapTokenPrice(assetID)
	if err != nil {
		log.Err(err).Msg("call updateStatesAfterDepositAndClaiming")
		// Using zero as the USD value; continue handling other assets without returning
	} else {
		// get token info
		bootstrapTokenState, err := m.database.GetBootstrapToken(assetID)
		if err != nil {
			return err
		}
		// calculate the total USD value of this asset
		priceDec, err := sdkmath.LegacyNewDecFromStr(priceStr)
		if err != nil {
			log.Err(err).Str("dbPrice", priceStr).Msg("failed to parse the db price to a big legacyDec")
			// don't return to continue addressing the other assets
		} else {
			divisor := sdkmath.NewIntWithDecimal(1, int(bootstrapTokenState.Decimals)) // #nosec G115
			usdValue = priceDec.MulInt(sdkmath.NewIntFromBigInt(assetDepositAmount)).QuoInt(divisor)
		}

	}

	return m.database.UpdateBootstrapTokenAmountAndUSDValue(assetID, assetDepositAmount.String(), usdValue.String())
}

func (m *Module) updateStatesAfterDelegationChange(stakerAddr, assetAddr common.Address, validatorAddr string, blockHeight int64) error {
	// update the states of staker assets
	stakerID, assetID, err := m.updateStakerAsset(stakerAddr, assetAddr, blockHeight)
	if err != nil {
		return err
	}

	delegationAmount, err := m.bootstrapSession.Delegations(stakerAddr, validatorAddr, assetAddr)
	if err != nil {
		return err
	}

	// update the delegation states
	err = m.database.SaveBootstrapDelegationState(&types.BootstrapDelegationState{
		StakerID:       stakerID,
		AssetID:        assetID,
		OperatorAddr:   validatorAddr,
		Delegated:      delegationAmount.String(),
		UpdatedAt:      time.Now(),
		UpdatedAtBlock: blockHeight, // Block height for optimistic update invalidation
	})
	if err != nil {
		return err
	}

	// update the states of operator assets
	operatorAmount, err := m.bootstrapSession.DelegationsByValidator(validatorAddr, assetAddr)
	if err != nil {
		return err
	}
	validatorCount, err := m.bootstrapSession.GetValidatorsCount()
	if err != nil {
		return err
	}

	var validatorETHAddr common.Address
	for i := int64(0); i < validatorCount.Int64(); i++ {
		tmpValidatorETHAddr, err := m.bootstrapSession.RegisteredValidators(big.NewInt(i))
		if err != nil {
			return err
		}
		tmpValidatorAddr, err := m.bootstrapSession.EthToImAddress(tmpValidatorETHAddr)
		if err != nil {
			return err
		}
		if tmpValidatorAddr == validatorAddr {
			validatorETHAddr = tmpValidatorETHAddr
			break
		}
	}
	if validatorETHAddr == (common.Address{}) {
		return fmt.Errorf("can't find the validator in the registered list")
	}

	// get the self delegation amount
	selfDelegation, err := m.bootstrapSession.Delegations(validatorETHAddr, validatorAddr, assetAddr)
	if err != nil {
		return err
	}

	err = m.database.SaveBootstrapOperatorAsset(&types.BootstrapOperatorAsset{
		OperatorAddr: validatorAddr,
		AssetID:      assetID,
		TotalAmount:  operatorAmount.String(),
		SelfAmount:   selfDelegation.String(),
		OtherAmount:  big.NewInt(0).Sub(operatorAmount, selfDelegation).String(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}

// RunAsyncOperations implements modules.AsyncOperationsModule
func (m *Module) RunAsyncOperations() {
	for {
		redial, err := m.SubscribeBootstrapEvents()
		if err != nil && redial {
			log.Err(err).Msg("redial the ETH websocket RPC due to an error")
			if m.EthWSClient != nil {
				m.EthWSClient.Close()
			}
			bootstrapAddr := common.HexToAddress(m.Config.BootstrapAddr)
			// redial the websocket RPC
			for {
				websocketRC, err := rpc.DialContext(context.Background(), m.Config.ETHWebsocket)
				if err != nil {
					log.Err(err).Msg("failed to dial ETH websocket RPC")
					// redial the RPC after sleeping 1 minute
					time.Sleep(time.Minute)
					continue
				}
				ethWSClient := ethclient.NewClient(websocketRC)
				// create the filterer to subscribe all related events
				bootstrapFilterer, err := bootstrap_binding.NewBootstrapFilterer(bootstrapAddr, ethWSClient)
				if err != nil {
					panic(fmt.Errorf("failed to new bootstrap filterer,err:%s", err))
				}
				m.EthWSClient = ethWSClient
				m.bootstrapFilterer = bootstrapFilterer
				break
			}
		} else {
			panic(fmt.Errorf("failed to subscribe bootstrap events,err:%s", err))
		}
	}
}

func (m *Module) SubscribeBootstrapEvents() (bool, error) {
	commonWatchCtx := &bind.WatchOpts{
		Context: m.ctx,
	}
	// create event channels for validator
	newValidatorCh := make(chan *bootstrap_binding.BootstrapValidatorRegistered)
	commissionUpdatedCh := make(chan *bootstrap_binding.BootstrapValidatorCommissionUpdated)
	keyReplaceCh := make(chan *bootstrap_binding.BootstrapValidatorKeyReplaced)

	newValidatorSub, err := m.bootstrapFilterer.WatchValidatorRegistered(commonWatchCtx, newValidatorCh)
	if err != nil {
		return false, fmt.Errorf("failed to watch validator registeration,err:%s", err)
	}
	defer newValidatorSub.Unsubscribe()

	commissionSub, err := m.bootstrapFilterer.WatchValidatorCommissionUpdated(commonWatchCtx, commissionUpdatedCh)
	if err != nil {
		return false, fmt.Errorf("failed to watch commission update,err:%s", err)
	}
	defer commissionSub.Unsubscribe()

	keyReplaceSub, err := m.bootstrapFilterer.WatchValidatorKeyReplaced(commonWatchCtx, keyReplaceCh)
	if err != nil {
		return false, fmt.Errorf("failed to watch key replace,err:%s", err)
	}
	defer keyReplaceSub.Unsubscribe()

	// create event channels for whitelist assets
	newAssetCh := make(chan *bootstrap_binding.BootstrapWhitelistTokenAdded)
	newAssetSub, err := m.bootstrapFilterer.WatchWhitelistTokenAdded(commonWatchCtx, newAssetCh)
	if err != nil {
		return false, fmt.Errorf("failed to watch whitelist token addition,err:%s", err)
	}
	defer newAssetSub.Unsubscribe()

	// create event channels for deposit, claim, delegation and undelegation
	depositCh := make(chan *bootstrap_binding.BootstrapDepositResult)
	depositSub, err := m.bootstrapFilterer.WatchDepositResult(commonWatchCtx, depositCh, nil, nil, nil)
	if err != nil {
		return false, fmt.Errorf("failed to watch token deposit,err:%s", err)
	}
	defer depositSub.Unsubscribe()

	claimCh := make(chan *bootstrap_binding.BootstrapClaimPrincipalResult)
	claimSub, err := m.bootstrapFilterer.WatchClaimPrincipalResult(commonWatchCtx, claimCh, nil, nil, nil)
	if err != nil {
		return false, fmt.Errorf("failed to watch token claim,err:%s", err)
	}
	defer claimSub.Unsubscribe()

	delegationCh := make(chan *bootstrap_binding.BootstrapDelegateResult)
	delegationSub, err := m.bootstrapFilterer.WatchDelegateResult(commonWatchCtx, delegationCh, nil, nil)
	if err != nil {
		return false, fmt.Errorf("failed to watch token delegation,err:%s", err)
	}
	defer delegationSub.Unsubscribe()

	undelegationCh := make(chan *bootstrap_binding.BootstrapUndelegateResult)
	undelegationSub, err := m.bootstrapFilterer.WatchUndelegateResult(commonWatchCtx, undelegationCh, nil, nil)
	if err != nil {
		return false, fmt.Errorf("failed to watch token undelegation,err:%s", err)
	}
	defer undelegationSub.Unsubscribe()

	for {
		select {
		case err := <-newValidatorSub.Err():
			log.Err(err).Msg("new validator subscription error")
			return true, err
		case err := <-commissionSub.Err():
			log.Err(err).Msg("commission update subscription error")
			return true, err
		case err := <-keyReplaceSub.Err():
			log.Err(err).Msg("key replace subscription error")
			return true, err
		case err := <-newAssetSub.Err():
			log.Err(err).Msg("whitelist token addition subscription error")
			return true, err
		case err := <-depositSub.Err():
			log.Err(err).Msg("token deposit subscription error")
			return true, err
		case err := <-claimSub.Err():
			log.Err(err).Msg("token claim subscription error")
			return true, err
		case err := <-delegationSub.Err():
			log.Err(err).Msg("token delegation subscription error")
			return true, err
		case err := <-undelegationSub.Err():
			log.Err(err).Msg("token undelegation subscription error")
			return true, err
		case e := <-newValidatorCh:
			// save the new validator
			err := m.database.SaveBootstrapValidator(&types.BootstrapValidator{
				ValidatorEthAddress: e.EthAddress.String(),
				ValidatorIMAddress:  e.ValidatorAddress,
				ValidatorName:       e.Name,
				ConsensusPubKey:     hexutil.Encode(e.ConsensusPublicKey[:]),
				Rate:                e.Commission.Rate.String(),
				MaxRate:             e.Commission.MaxRate.String(),
				MaxChangeRate:       e.Commission.MaxChangeRate.String(),
				UpdatedAt:           time.Now(),
			})
			if err != nil {
				log.Err(err).Msg("failed to saving the new validator")
			}
		case e := <-commissionUpdatedCh:
			err = m.database.UpdateCommissionRate(e.ValidatorAddress, e.NewRate.String())
			if err != nil {
				log.Err(err).Str("validatorAddr", e.ValidatorAddress).Msg("failed to update the commission rate")
			}
		case e := <-keyReplaceCh:
			err = m.database.UpdateConsensusPubKey(e.ValidatorAddress, hexutil.Encode(e.NewConsensusPublicKey[:]))
			if err != nil {
				log.Err(err).Str("validatorAddr", e.ValidatorAddress).Msg("failed to replace the consensus key")
			}
		case e := <-newAssetCh:
			tokenCount, err := m.bootstrapSession.GetWhitelistedTokensCount()
			if err != nil {
				log.Err(err).Msg("failed to get the count of whitelisted tokens")
				continue
			}
			// get the token info by iterating all indexes
			for i := int64(0); i < tokenCount.Int64(); i++ {
				tokenInfo, err := m.bootstrapSession.GetWhitelistedTokenAtIndex(big.NewInt(i))
				if err != nil {
					log.Err(err).Int64("index", i).Msg("failed to get the count of whitelisted tokens")
					continue
				}
				if tokenInfo.TokenAddress == e.Token {
					_, assetID := assetstypes.GetStakerIDAndAssetID(m.Config.ETHLZChainID, nil, e.Token[:])
					err = m.database.SaveBootstrapToken(&types.BootstrapTokenState{
						BootstrapToken: types.BootstrapToken{
							AssetID:   assetID,
							Address:   e.Token.String(),
							Name:      tokenInfo.Name,
							Symbol:    tokenInfo.Symbol,
							Decimals:  tokenInfo.Decimals,
							LZChainID: m.Config.ETHLZChainID,
						},
						StakingTotalAmount: big.NewInt(0).String(),
						TotalUSDValue:      big.NewInt(0).String(),
						UpdatedAt:          time.Now(),
					})
					if err != nil {
						log.Err(err).Str("token", e.Token.String()).Msg("failed to save the whitelist asset")
					}
					break
				}
			}
		case e := <-depositCh:
			if e.Success {
				err = m.updateStatesAfterDepositAndClaiming(e.Depositor, e.Token, int64(e.Raw.BlockNumber))
				if err != nil {
					log.Err(err).Str("depositor", e.Depositor.String()).Str("token", e.Token.String()).Str("amount", e.Amount.String()).Msg("failed to handle the deposit event")
				}
			}
		case e := <-claimCh:
			if e.Success {
				err = m.updateStatesAfterDepositAndClaiming(e.Withdrawer, e.Token, int64(e.Raw.BlockNumber))
				if err != nil {
					log.Err(err).Str("withdrawer", e.Withdrawer.String()).Str("token", e.Token.String()).Str("amount", e.Amount.String()).Msg("failed to handle the claim event")
				}
			}
		case e := <-delegationCh:
			if e.Success {
				err := m.updateStatesAfterDelegationChange(e.Delegator, e.Token, e.Delegatee, int64(e.Raw.BlockNumber))
				if err != nil {
					log.Err(err).Str("delegator", e.Delegator.String()).Str("token", e.Token.String()).Str("validator", e.Delegatee).Str("amount", e.Amount.String()).Msg("failed to handle the delegation event")
				}
			}
		case e := <-undelegationCh:
			if e.Success {
				err := m.updateStatesAfterDelegationChange(e.Undelegator, e.Token, e.Undelegatee, int64(e.Raw.BlockNumber))
				if err != nil {
					log.Err(err).Str("undelegator", e.Undelegator.String()).Str("token", e.Token.String()).Str("validator", e.Undelegatee).Str("amount", e.Amount.String()).Msg("failed to handle the undelegation event")
				}
			}
		}
	}
}
