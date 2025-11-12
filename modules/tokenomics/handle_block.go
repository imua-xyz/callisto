package tokenomics

import (
	"database/sql"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"

	"github.com/forbole/callisto/v4/types"
	"github.com/imua-xyz/imuachain/utils"
	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"
	oracletypes "github.com/imua-xyz/imuachain/x/oracle/types"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	juno "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
)

// HandleBlock implements BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, _ *tmctypes.ResultBlockResults, _ []*juno.Tx, _ *tmctypes.ResultValidators,
) error {
	log.Debug().Str("module", m.Name()).Int64("height", block.Block.Height).
		Msg(fmt.Sprintf("updating %s", m.Name()))

	err := m.HandleGenesisPoolAirdropRound(block)
	if err != nil {
		return err
	}
	err = m.HandleLiquidityIncentivesAirdropRound(block)
	if err != nil {
		return err
	}
	return nil
}

func (m *Module) HandleGenesisPoolAirdropForStakers(tx *sql.Tx, roundReward sdkmath.LegacyDec, round *types.CommonAirdropRound) error {
	var stakerAirdrop types.GenesisStakerAirdrop
	stakerAirdrop.RewardAmount = sdkmath.LegacyZeroDec().String()
	totalUSDValue := sdkmath.LegacyZeroDec()
	stakerUSDValue := sdkmath.LegacyZeroDec()
	assetPrices := make(map[string]*oracletypes.Price)
	assetDecimals := make(map[string]int)
	totalStakers := 0
	// iterate over all genesis stakers and calculate the USD value
	// the iteration is ordered by staker ID and asset ID, allowing the total USD value of
	// each staker to be calculated in a single pass.
	opFunc := func(sa types.ParsedStakerAsset) error {
		var err error
		if stakerAirdrop.StakerID == "" {
			stakerAirdrop.StakerID = sa.StakerID
		} else if stakerAirdrop.StakerID != sa.StakerID {
			// the airdrop for the previous staker has been fully processed; save it to the database.
			stakerAirdrop.AirdropRound = round.AirdropRound
			stakerAirdrop.USDValue = stakerUSDValue.String()
			if stakerUSDValue.IsPositive() {
				err = m.db.SaveGenesisStakerAirdrop(tx, &stakerAirdrop)
				if err != nil {
					return fmt.Errorf("error saving genesis staker airdrop: %w", err)
				}
				// update the total USD value and total number of stakers
				totalUSDValue.AddMut(stakerUSDValue)
				totalStakers++
			}

			// clear the stakerUSDValue and airdrop info for next staker
			stakerUSDValue = sdkmath.LegacyZeroDec()
			stakerAirdrop = types.GenesisStakerAirdrop{
				StakerID:     sa.StakerID,
				RewardAmount: sdkmath.LegacyZeroDec().String(),
			}
		}

		// get price info for the asset
		var price *oracletypes.Price
		var assetDecimal int
		var ok bool
		price, ok = assetPrices[sa.AssetID]
		if !ok {
			price, err = m.db.GetLatestPriceByAssetID(sa.AssetID)
			if err != nil {
				return fmt.Errorf("error getting latest price by assetID: %w,assetID:%s", err, sa.AssetID)
			}
			assetPrices[sa.AssetID] = price
		}
		// get asset decimal
		assetDecimal, ok = assetDecimals[sa.AssetID]
		if !ok {
			assetDecimal, err = m.db.GetTokenDecimalsByID(sa.AssetID)
			if err != nil {
				return fmt.Errorf("error getting asset decimal: %w,assetID:%s", err, sa.AssetID)
			}
			assetDecimals[sa.AssetID] = assetDecimal
		}
		// calculate the USD value for the stakerID and assetID
		validAssetAmount := sdkmath.MinInt(sa.GenesisDeposited, sa.Deposited)
		assetUSDValue := utils.CalculateUSDValue(validAssetAmount, price.Value, uint32(assetDecimal), price.Decimal)
		stakerUSDValue.AddMut(assetUSDValue)

		return nil
	}
	err := m.db.IterateAirdropStakerAssets(opFunc)
	if err != nil {
		return fmt.Errorf("HandleGenesisPoolAirdropForStakers: error iterating over genesis stakers: %w", err)
	}
	// flush the last staker's data
	if stakerAirdrop.StakerID != "" {
		stakerAirdrop.AirdropRound = round.AirdropRound
		stakerAirdrop.USDValue = stakerUSDValue.String()

		if stakerUSDValue.IsPositive() {
			err = m.db.SaveGenesisStakerAirdrop(tx, &stakerAirdrop)
			if err != nil {
				return fmt.Errorf("error saving final genesis staker airdrop: %w", err)
			}

			totalUSDValue.AddMut(stakerUSDValue)
			totalStakers++
		}
	}
	// update the total USD value and staker number in the input round info.
	round.TotalUSDValue = totalUSDValue.String()
	round.TotalStakers = totalStakers
	if totalStakers == 0 || totalUSDValue.IsZero() {
		// there aren't any valid stakers who can get the airdrop.
		return nil
	}

	// calculate and update the airdrop rewards of all stakers.
	err = m.db.UpdateGenesisAirdropRewardsByRoundSQL(tx, round.AirdropRound, roundReward, totalUSDValue)
	if err != nil {
		return fmt.Errorf("HandleGenesisPoolAirdropForStakers: error updating airdrop reward for all genesis stakers: %w", err)
	}
	return nil
}

func (m *Module) calculateAirdropRound(block *tmctypes.ResultBlock, airdropType types.AirdropType) (*types.GeneralRoundInfo, error) {
	if block.Block.Height == 1 {
		// We skip the first block because the block time at height 1 might be later than
		// the genesis time in the DB, which can be caused by precision loss when saving
		// block times to the DB. This will result in an error being logged at height 1.
		return nil, nil
	}
	oneMinute := int64(time.Minute)
	oneDay := int64(24 * time.Hour)
	oneYear := 365 * oneDay
	var interval, duration int64
	var latestAirdropRound func() (*types.CommonAirdropRound, error)
	switch airdropType {
	case types.GenesisPoolAirdrop:
		interval = m.cfg.GenesisPoolAirdropInterval * oneMinute
		duration = m.cfg.GenesisPoolAirdropDuration * oneMinute
		latestAirdropRound = m.db.GetLatestGenesisPoolAirdropRound
	case types.LiquidityIncentivesAirdrop:
		interval = m.cfg.LiquidityIncentiveAirdropInterval.MulInt64(oneYear).TruncateInt64()
		duration = m.cfg.LiquidityIncentiveAirdropDuration * oneMinute
		latestAirdropRound = m.db.GetLatestLiquidityIncentivesAirdropRound
	default:
		return nil, fmt.Errorf("invalid airdrop type:%d", airdropType)
	}

	if interval == 0 || duration == 0 {
		// do nothing when the interval or duration is zero.
		log.Info().Msg("the interval or duration is zero")
		return nil, nil
	}

	latestAirdropRoundInfo, err := latestAirdropRound()
	if err != nil {
		return nil, err
	}

	var roundStart time.Time
	var latestRoundID int
	// get the genesis time
	genesis, err := m.db.GetGenesis()
	if err != nil {
		return nil, fmt.Errorf("error getting genesis: %w", err)
	}
	blockTime := block.Block.Time
	if latestAirdropRoundInfo == nil {
		roundStart = genesis.Time
	} else {
		roundStart = latestAirdropRoundInfo.RoundEndAt
		latestRoundID = latestAirdropRoundInfo.AirdropRound
	}
	if !roundStart.Before(blockTime) {
		return nil, fmt.Errorf("round start time:%s isn't before block time:%s,latestRoundID:%d", roundStart, blockTime, latestRoundID)
	}

	// calculate the round id by the genesis and block time.
	airdropEndTime := genesis.Time.Add(time.Duration(duration))
	if !roundStart.Before(airdropEndTime) {
		// Do nothing because the airdrop has already ended.
		return nil, nil
	}
	roundEnd := blockTime
	roundDur := int64(roundEnd.Sub(roundStart))
	if roundDur < interval {
		if !roundEnd.Before(airdropEndTime) {
			// Use airdropEndTime to ensure the total airdrop duration does not exceed
			// the configured duration, which would otherwise cause incorrect reward
			// calculations in the final round.
			roundEnd = airdropEndTime
		} else {
			switch airdropType {
			case types.GenesisPoolAirdrop:
				// Do nothing because the airdrop round is not due yet.
				return nil, nil
			case types.LiquidityIncentivesAirdrop:
				roundStartDurFromGenesis := int64(roundStart.Sub(genesis.Time))
				roundEndDurFromGenesis := int64(roundEnd.Sub(genesis.Time))
				if roundEndDurFromGenesis/oneYear == roundStartDurFromGenesis/oneYear+1 {
					// The round will cross a calendar-year boundary here, but yearly reward ratios differ.
					// To avoid incorrect reward calculations, we truncate such rounds by setting the round
					// end time to the end of the earlier year so that no round spans multiple years.
					roundEnd = genesis.Time.Add(time.Duration(roundEndDurFromGenesis / oneYear * oneYear))
				} else if roundEndDurFromGenesis/oneYear == roundStartDurFromGenesis/oneYear {
					// Do nothing because the airdrop round is not due yet.
					return nil, nil
				} else {
					return nil, fmt.Errorf("invalid start and end duration since enesis,roundStartDurFromGenesis:%d,roundEndDurFromGenesis:%d", roundStartDurFromGenesis, roundEndDurFromGenesis)
				}
			default:
				return nil, fmt.Errorf("invalid airdrop type:%d", airdropType)
			}
		}
	}

	// the roundID will start from 1.
	roundID := latestRoundID + 1
	// recalculate the round duration using updated roundEnd
	roundDur = int64(roundEnd.Sub(roundStart))

	return &types.GeneralRoundInfo{
		RoundID:         roundID,
		RoundDur:        roundDur,
		RoundStart:      roundStart,
		RoundEnd:        roundEnd,
		GenesisTime:     genesis.Time,
		PreRound:        latestAirdropRoundInfo,
		AirdropDur:      duration,
		AirdropInterval: interval,
	}, nil
}

func (m *Module) HandleGenesisPoolAirdropRound(block *tmctypes.ResultBlock) error {
	if m.cfg.GenesisSupply == 0 || m.cfg.GenesisPoolRatio.IsZero() {
		// do nothing when the genesis supply or the genesis pool ratio is zero.
		log.Info().Msg("the genesis supply or the genesis pool ratio is zero")
		return nil
	}

	generalRoundInfo, err := m.calculateAirdropRound(block, types.GenesisPoolAirdrop)
	if err != nil {
		return err
	}

	if generalRoundInfo == nil {
		return nil
	}

	roundDurMinute := sdkmath.LegacyNewDec(generalRoundInfo.RoundDur).QuoTruncate(sdkmath.LegacyNewDec(int64(time.Minute))).String()
	round := &types.CommonAirdropRound{
		AirdropType:   types.GenesisPoolAirdrop,
		AirdropRound:  generalRoundInfo.RoundID,
		BlockHeight:   block.Block.Height,
		RoundDuration: roundDurMinute,
		RoundStartAt:  generalRoundInfo.RoundStart,
		RoundEndAt:    generalRoundInfo.RoundEnd,
		CreatedAt:     block.Block.Time,
		TotalUSDValue: sdkmath.LegacyZeroDec().String(),
	}
	// calculate the total reward amount for this round
	genesisSupplyInt := sdkmath.NewInt(m.cfg.GenesisSupply)

	roundReward := m.cfg.GenesisPoolRatio.MulInt(genesisSupplyInt).MulInt64(generalRoundInfo.RoundDur).QuoTruncate(sdkmath.LegacyNewDec(generalRoundInfo.AirdropDur))
	round.TotalRewardAmount = roundReward.String()

	// Since the airdrop calculation depends on the state of the previous round, we must
	// ensure the atomicity of an airdrop round.
	tx, err := m.db.SQL.Begin()
	if err != nil {
		return fmt.Errorf("error begining db tx: %w", err)
	}
	// The rollback will succeed if an error occurs; otherwise, it won’t revert any state changes
	// because the transaction has already been committed.
	defer tx.Rollback()

	err = m.db.SaveGenesisPoolAirdropRound(tx, round)
	if err != nil {
		return fmt.Errorf("error saving genesis pool airdrop round: %w", err)
	}

	// calculate and address the airdrop for all genesis stakers.
	err = m.HandleGenesisPoolAirdropForStakers(tx, roundReward, round)
	if err != nil {
		return fmt.Errorf("error addressing genesis pool airdrop for all stakers: %w", err)
	}

	// update he staker number and total USD value in the round info.
	// the input parameters will be set in the `HandleGenesisPoolAirdropForStakers` function above.
	err = m.db.UpdateGenesisPoolAirdropMetrics(tx, round.AirdropRound, round.TotalStakers, round.TotalUSDValue)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing tx: %w", err)
	}
	return nil
}

func (m *Module) HandleLiquidityIncentivesAirdropRound(block *tmctypes.ResultBlock) error {
	if m.cfg.GenesisSupply == 0 {
		// do nothing when the genesis supply is zero.
		log.Info().Msg("the genesis supply is zero")
		return nil
	}
	generalRoundInfo, err := m.calculateAirdropRound(block, types.LiquidityIncentivesAirdrop)
	if err != nil {
		return err
	}
	if generalRoundInfo == nil {
		return nil
	}

	roundDurMinute := sdkmath.LegacyNewDec(generalRoundInfo.RoundDur).QuoTruncate(sdkmath.LegacyNewDec(int64(time.Minute))).String()
	round := &types.CommonAirdropRound{
		AirdropType:            types.LiquidityIncentivesAirdrop,
		AirdropRound:           generalRoundInfo.RoundID,
		BlockHeight:            block.Block.Height,
		RoundDuration:          roundDurMinute,
		RoundStartAt:           generalRoundInfo.RoundStart,
		RoundEndAt:             generalRoundInfo.RoundEnd,
		CreatedAt:              block.Block.Time,
		TotalNativeIMUARewards: sdkmath.LegacyZeroDec().String(),
	}
	// calculate the total reward amount for this round
	genesisSupplyInt := sdkmath.NewInt(m.cfg.GenesisSupply)
	oneYear := 365 * 24 * time.Hour
	var rewardRatioIndex int
	if generalRoundInfo.PreRound != nil {
		// The start time of this round should be the end of the previous round, and each round is always created at the end.
		rewardRatioIndex = int(generalRoundInfo.PreRound.CreatedAt.Sub(generalRoundInfo.GenesisTime) / oneYear)
	}
	rewardRatiosLength := len(m.cfg.LiquidityIncentiveRatios)
	var rewardRatio sdkmath.LegacyDec
	if rewardRatioIndex < rewardRatiosLength {
		rewardRatio = m.cfg.LiquidityIncentiveRatios[rewardRatioIndex].LegacyDec
	} else {
		rewardRatio = m.cfg.LiquidityIncentiveRatios[rewardRatiosLength-1].LegacyDec
	}
	if rewardRatio.IsZero() {
		// do nothing when the reward ratio is zero.
		log.Info().Msg("the reward ratio is zero")
		return nil
	}
	roundReward := rewardRatio.MulInt(genesisSupplyInt).MulInt64(generalRoundInfo.RoundDur).QuoTruncate(sdkmath.LegacyNewDec(int64(oneYear)))
	round.TotalRewardAmount = roundReward.String()

	// Since the airdrop calculation depends on the state of the previous round, we must
	// ensure the atomicity of an airdrop round.
	tx, err := m.db.SQL.Begin()
	if err != nil {
		return fmt.Errorf("error begining db tx: %w", err)
	}
	// The rollback will succeed if an error occurs; otherwise, it won’t revert any state changes
	// because the transaction has already been committed.
	defer tx.Rollback()

	err = m.db.SaveLiquidityAirdropRound(tx, round)
	if err != nil {
		return fmt.Errorf("error saving liquidity incentive airdrop round: %w", err)
	}

	// calculate and address the airdrop for all stakers.
	err = m.HandleLiquidityIncentiveAirdropForStakers(tx, roundReward, round, generalRoundInfo.PreRound)
	if err != nil {
		return fmt.Errorf("error addressing liquidity incentive airdrop for all stakers: %w", err)
	}

	// update he staker number and total native IMUA rewards in the round info.
	// the input parameters will be set in the `HandleLiquidityIncentiveAirdropForStakers` function above.
	err = m.db.UpdateLiquidityAirdropMetrics(tx, round.AirdropRound, round.TotalStakers, round.TotalNativeIMUARewards)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing tx: %w", err)
	}
	return nil
}

func (m *Module) HandleLiquidityIncentiveAirdropForStakers(tx *sql.Tx, roundReward sdkmath.LegacyDec, round, preRound *types.CommonAirdropRound) error {
	// update the rewards for all stakers
	// get all stakers
	stakers, err := m.db.GetAllStakersFromDelegationStates()
	if err != nil {
		return err
	}
	totalRoundNativeRewards := sdkmath.LegacyZeroDec()
	totalStakers := 0
	// Iterate over all stakers to calculate their total rewards for the current round,
	// and save the rewards snapshot to the database.
	// The snapshot is created by fetching reward states from the RPC, and is later used
	// to calculate the liquidity incentive airdrop amount.
	//
	// TODO: Consider potential performance issues—processing all stakers
	// may put pressure on both the indexer and the source full node.
	// This affects only the indexer itself and the node it fetches data from.
	for _, stakerID := range stakers {
		// get claimed dogfood rewards
		claimedRewards, err := m.source.StakerAVSClaimedRewards(round.BlockHeight, stakerID, m.dogfoodAddr)
		if err != nil {
			return fmt.Errorf("failed to get claimed dogfood rewards for staker:%s, err:%w", stakerID, err)
		}
		outstandingReward := sdkmath.LegacyZeroDec()
		withdrawnReward := sdkmath.LegacyZeroDec()
		totalClaimedReward := sdkmath.LegacyZeroDec()
		if claimedRewards != nil {
			outstandingReward = claimedRewards.OutstandingRewards.AmountOf(assetstypes.ImuachainAssetDenom)
			withdrawnReward = claimedRewards.WithdrawnRewards.AmountOf(assetstypes.ImuachainAssetDenom)
			totalClaimedReward = outstandingReward.Add(withdrawnReward)
		}
		// get unclaimed dogfood rewards
		unclaimedRewards, err := m.source.StakerAVSUnclaimedRewards(round.BlockHeight, stakerID, m.dogfoodAddr)
		if err != nil {
			return fmt.Errorf("failed to get unclaimed dogfood rewards for staker:%s,err:%w", stakerID, err)
		}
		unclaimedReward := unclaimedRewards.AmountOf(assetstypes.ImuachainAssetDenom)

		totalReward := totalClaimedReward.Add(unclaimedReward)
		roundNativeRewards := totalReward
		if preRound != nil {
			preStakerAirdrop, err := m.db.GetLiquidityStakerAirdrop(stakerID, preRound.AirdropRound)
			if err != nil {
				return fmt.Errorf("failed to get the staker liquidity airdrop of previous round,err:%s", err)
			}
			if preStakerAirdrop != nil {
				roundNativeRewards = totalReward.Sub(sdkmath.LegacyMustNewDecFromStr(preStakerAirdrop.TotalRewards))
			}
		}
		if roundNativeRewards.IsNegative() {
			return fmt.Errorf("negative round native rewards for staker:%s,round:%d,reward:%s", stakerID, round.AirdropRound, roundNativeRewards)
		} else if roundNativeRewards.IsZero() {
			// do nothing for stakers without any native rewards.
			continue
		}

		// save reward snapshot
		err = m.db.SaveLiquidityStakerAirdrop(tx, &types.LiquidityStakerAirdrop{
			StakerID:            stakerID,
			AirdropRound:        round.AirdropRound,
			OutstandingRewards:  outstandingReward.String(),
			WithdrawnRewards:    withdrawnReward.String(),
			ClaimedRewards:      totalClaimedReward.String(),
			UnclaimedRewards:    unclaimedReward.String(),
			TotalRewards:        totalReward.String(),
			RoundNativeRewards:  roundNativeRewards.String(),
			AirdropRewardAmount: sdkmath.LegacyZeroDec().String(),
			CreatedAt:           time.Now(),
		})
		if err != nil {
			return fmt.Errorf("failed to save liquidity incentive airdrop for staker: %s, round: %d, err: %w", stakerID, round.AirdropRound, err)
		}
		totalRoundNativeRewards.AddMut(roundNativeRewards)
		totalStakers++
	}
	// update the total native IMUA rewards and staker number in the input round info.
	round.TotalNativeIMUARewards = totalRoundNativeRewards.String()
	round.TotalStakers = totalStakers

	if totalStakers == 0 || totalRoundNativeRewards.IsZero() {
		// there aren't any valid stakers who can get the airdrop.
		return nil
	}
	// calculate and update the airdrop rewards of all stakers.
	err = m.db.UpdateLiquidityAirdropRewardsByRoundSQL(tx, round.AirdropRound, roundReward, totalRoundNativeRewards)
	if err != nil {
		return fmt.Errorf("HandleLiquidityIncentiveAirdropForStakers: error updating airdrop reward for all stakers: %w", err)
	}
	return nil
}
