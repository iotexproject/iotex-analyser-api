package apiservice

import (
	"math"
	"math/big"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
)

type HermesVotingResult struct {
	ID                        uint64
	EpochNumber               uint64
	DelegateName              string
	OperatorAddress           string
	RewardAddress             string
	StakingAddress            string
	TotalWeightedVotes        string
	SelfStaking               string
	BlockRewardPercentage     float64
	EpochRewardPercentage     float64
	FoundationBonusPercentage float64
}

type HermesAggregateVoting struct {
	ID             uint64
	EpochNumber    uint64
	CandidateName  string
	VoterAddress   string
	NativeFlag     bool
	AggregateVotes string
}

type HermesDistributionPlan struct {
	TotalWeightedVotes        *big.Int
	StakingAddress            string
	BlockRewardPercentage     float64
	EpochRewardPercentage     float64
	FoundationBonusPercentage float64
}

type HermesDistributionSource struct {
	BlockReward     *big.Int
	EpochReward     *big.Int
	FoundationBonus *big.Int
}

type RewardDistribution api.RewardDistribution

// DelegateHermesDistribution defines the Hermes reward distributions for each delegate
type DelegateHermesDistribution struct {
	DelegateName        string
	Distributions       []*RewardDistribution
	StakingIotexAddress string
	VoterCount          uint64
	WaiveServiceFee     bool
	Refund              string
}

type AccountReward struct {
	ID              uint64
	BlockHeight     uint64
	EpochNumber     uint64
	RewardAddress   string
	ActionHash      string
	CandidateName   string
	BlockReward     string
	EpochReward     string
	FoundationBonus string
}

type HermesDistribution api.HermesDistribution

func distributionPlanByRewardAddress(startEpoch uint64, endEpoch uint64, rewardAddress string) (map[string]map[uint64]*HermesDistributionPlan, error) {

	db := db.DB()
	var ids []HermesVotingResult
	if err := db.Table("hermes_voting_results").Where("epoch_number >= ?  AND epoch_number <= ? AND reward_address= ?", startEpoch, endEpoch, rewardAddress).Find(&ids).Error; err != nil {
		return nil, err
	}

	return parseDistributionPlanFromVotingResult(ids)
}

func parseDistributionPlanFromVotingResult(rows []HermesVotingResult) (map[string]map[uint64]*HermesDistributionPlan, error) {

	if len(rows) == 0 {
		return nil, errors.New("records empty")
	}

	distributePlanMap := make(map[string]map[uint64]*HermesDistributionPlan)
	for _, row := range rows {
		if _, ok := distributePlanMap[row.DelegateName]; !ok {
			distributePlanMap[row.DelegateName] = make(map[uint64]*HermesDistributionPlan)
		}
		planMap := distributePlanMap[row.DelegateName]
		totalWeightedVotes, err := stringToBigInt(row.TotalWeightedVotes)
		if err != nil {
			return nil, errors.New("failed to convert string to big int")
		}
		planMap[row.EpochNumber] = &HermesDistributionPlan{
			BlockRewardPercentage:     row.BlockRewardPercentage,
			EpochRewardPercentage:     row.EpochRewardPercentage,
			FoundationBonusPercentage: row.FoundationBonusPercentage,
			StakingAddress:            row.StakingAddress,
			TotalWeightedVotes:        totalWeightedVotes,
		}
	}
	return distributePlanMap, nil
}

// stringToBigInt transforms a string to big int
func stringToBigInt(estr string) (*big.Int, error) {
	ret, ok := big.NewInt(0).SetString(estr, 10)
	if !ok {
		return nil, errors.New("failed to parse string to big int")
	}
	return ret, nil
}

func accountRewards(delegateMap map[uint64][]string) (map[string]map[uint64]*HermesDistributionSource, error) {

	db := db.DB()
	var rows []AccountReward
	var minEpoch, maxEpoch uint64
	minEpoch = math.MaxUint64
	maxEpoch = 0
	for k := range delegateMap {
		if k >= maxEpoch {
			maxEpoch = k
		}
		if k <= minEpoch {
			minEpoch = k
		}
	}

	if err := db.Table("hermes_account_rewards").Where("epoch_number >= ?  AND epoch_number <= ?", minEpoch, maxEpoch).Find(&rows).Error; err != nil {
		return nil, err
	}

	//map[delegate]map[epoch]*HermesDistributionSource
	accountRewardsMap := make(map[string]map[uint64]*HermesDistributionSource)
	for _, row := range rows {
		exist := false
		for _, v := range delegateMap[row.EpochNumber] {
			if row.CandidateName == v {
				exist = true
				break
			}
		}
		if !exist {
			continue
		}
		if _, ok := accountRewardsMap[row.CandidateName]; !ok {
			accountRewardsMap[row.CandidateName] = make(map[uint64]*HermesDistributionSource, 0)
		}
		rewardsMap := accountRewardsMap[row.CandidateName]
		blockReward, err := stringToBigInt(row.BlockReward)
		if err != nil {
			return nil, errors.New("failed to covert string to big int")
		}
		epochReward, err := stringToBigInt(row.EpochReward)
		if err != nil {
			return nil, errors.New("failed to covert string to big int")
		}
		foundationBonus, err := stringToBigInt(row.FoundationBonus)
		if err != nil {
			return nil, errors.New("failed to covert string to big int")
		}
		rewardsMap[row.EpochNumber] = &HermesDistributionSource{
			BlockReward:     blockReward,
			EpochReward:     epochReward,
			FoundationBonus: foundationBonus,
		}
	}
	return accountRewardsMap, nil
}

func weightedVotesBySearchPairs(delegateMap map[uint64][]string) (map[string]map[uint64]map[string]*big.Int, error) {
	db := db.DB()
	var rows []HermesAggregateVoting
	var minEpoch, maxEpoch uint64
	minEpoch = math.MaxUint64
	maxEpoch = 0
	for k := range delegateMap {
		if k >= maxEpoch {
			maxEpoch = k
		}
		if k <= minEpoch {
			minEpoch = k
		}
	}

	if err := db.Table("hermes_aggregate_votings").Where("epoch_number >= ?  AND epoch_number <= ?", minEpoch, maxEpoch).Find(&rows).Error; err != nil {
		return nil, err
	}

	voterVotesMap := make(map[string]map[uint64]map[string]*big.Int)
	for _, row := range rows {
		exist := false
		for _, v := range delegateMap[row.EpochNumber] {
			if row.CandidateName == v {
				exist = true
				break
			}
		}
		if !exist {
			continue
		}
		if _, ok := voterVotesMap[row.CandidateName]; !ok {
			voterVotesMap[row.CandidateName] = make(map[uint64]map[string]*big.Int)
		}
		epochVoterMap := voterVotesMap[row.CandidateName]
		if _, ok := epochVoterMap[row.EpochNumber]; !ok {
			epochVoterMap[row.EpochNumber] = make(map[string]*big.Int)
		}
		voterMap := epochVoterMap[row.EpochNumber]

		weightedVotesInt, errs := stringToBigInt(row.AggregateVotes)
		if errs != nil {
			return nil, errors.Wrap(errs, "failed to convert to big int")
		}
		if val, ok := voterMap[row.VoterAddress]; !ok {
			voterMap[row.VoterAddress] = weightedVotesInt
		} else {
			val.Add(val, weightedVotesInt)
		}
	}
	return voterVotesMap, nil
}

func calculatedDistributedReward(distributePlan *HermesDistributionPlan, rewards *HermesDistributionSource) (*big.Int, error) {
	blockRewardPercentage := distributePlan.BlockRewardPercentage
	epochRewardPercentage := distributePlan.EpochRewardPercentage
	foundationBonusPercentage := distributePlan.FoundationBonusPercentage
	distrReward := big.NewInt(0)
	if blockRewardPercentage > 0 {
		distrBlockReward := new(big.Int).Set(rewards.BlockReward)
		distrBlockReward.Mul(distrBlockReward, big.NewInt(int64(blockRewardPercentage*100))).Div(distrBlockReward, big.NewInt(10000))
		distrReward.Add(distrReward, distrBlockReward)
	}
	if epochRewardPercentage > 0 {
		distrEpochReward := new(big.Int).Set(rewards.EpochReward)
		distrEpochReward.Mul(distrEpochReward, big.NewInt(int64(epochRewardPercentage*100))).Div(distrEpochReward, big.NewInt(10000))
		distrReward.Add(distrReward, distrEpochReward)
	}
	if foundationBonusPercentage > 0 {
		distrFoundationBonus := new(big.Int).Set(rewards.FoundationBonus)
		distrFoundationBonus.Mul(distrFoundationBonus, big.NewInt(int64(foundationBonusPercentage*100))).Div(distrFoundationBonus, big.NewInt(10000))
		distrReward.Add(distrReward, distrFoundationBonus)
	}
	return distrReward, nil
}

// convertVoterDistributionMapToList converts voter reward distribution map to list
func convertVoterDistributionMapToList(voterAddrToReward map[string]*big.Int) ([]*RewardDistribution, error) {
	rewardDistribution := make([]*RewardDistribution, 0)
	for ioAddress, rewardAmount := range voterAddrToReward {
		voterAddr, _ := address.FromString(ioAddress)
		rewardDistribution = append(rewardDistribution, &RewardDistribution{
			VoterEthAddress:   voterAddr.Hex(),
			VoterIotexAddress: ioAddress,
			Amount:            rewardAmount.String(),
		})
	}
	return rewardDistribution, nil
}
