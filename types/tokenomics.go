package types

import "time"

type AirdropType int

const (
	GenesisPoolAirdrop AirdropType = iota
	LiquidityIncentivesAirdrop
)

type GeneralRoundInfo struct {
	RoundID         int
	RoundDur        int64
	RoundStart      time.Time
	RoundEnd        time.Time
	GenesisTime     time.Time
	PreRound        *CommonAirdropRound
	AirdropDur      int64
	AirdropInterval int64
}

// GenesisStakerAirdrop represents the reward state for a single staker in a specific airdrop round.
type GenesisStakerAirdrop struct {
	StakerID      string     // Unique identifier of the staker
	AirdropRound  int        // Airdrop round index
	USDValue      string     // Total USD value of the staker's assets at snapshot time (NUMERIC)
	RewardAmount  string     // Allocated reward amount (NUMERIC)
	IsDistributed bool       // Whether the reward has been distributed
	DistributedAt *time.Time // When the reward was distributed (nullable)
}

type CommonAirdropRound struct {
	AirdropType            AirdropType // type of the airdrop
	AirdropRound           int         // Index of the airdrop round
	BlockHeight            int64       // Snapshot block height for this round
	TotalStakers           int         // Total number of stakers eligible in this round
	TotalUSDValue          string      // Total USD value of all eligible stakers, used for genesis pool airdrop
	TotalNativeIMUARewards string      // Total native IMUA rewards aggregated across stakers, used for liquidity incentives airdrop
	TotalRewardAmount      string      // Total reward amount allocated for this round
	RoundDuration          string      // Round duration(minute) for this round
	RoundStartAt           time.Time   // Timestamp when this round started
	RoundEndAt             time.Time   // Timestamp when this round ended
	CreatedAt              time.Time   // Timestamp when this record was created
	DistributedStakers     int         // Number of stakers who have received their rewards
	IsCompleted            bool        // Indicates whether the reward distribution is finished
}

type LiquidityStakerAirdrop struct {
	StakerID            string     // Unique identifier of the staker
	AirdropRound        int        // Airdrop round index
	OutstandingRewards  string     // Amount of outstanding rewards (NUMERIC)
	WithdrawnRewards    string     // Amount of withdrawn rewards (NUMERIC)
	ClaimedRewards      string     // Amount of claimed rewards (NUMERIC, derived)
	UnclaimedRewards    string     // Amount of unclaimed rewards (NUMERIC)
	TotalRewards        string     // Total reward amount (NUMERIC, derived)
	RoundNativeRewards  string     // Total round native reward amount (NUMERIC, derived)
	AirdropRewardAmount string     // The airdrop reward amount allocated to the staker (NUMERIC)
	IsDistributed       bool       // Whether the reward has been distributed
	DistributedAt       *time.Time // When the reward was distributed (nullable)
	CreatedAt           time.Time  // Record creation timestamp
}
