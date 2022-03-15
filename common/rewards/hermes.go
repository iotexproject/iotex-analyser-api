package rewards

import (
	"math/big"

	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
)

type RewardDistribution struct {
	VoterEthAddress   string
	VoterIotexAddress string
	Amount            string
}

type TotalWeight struct {
	EpochNumber        uint64
	TotalWeightedVotes string
}

type AggregateVoting struct {
	EpochNumber    uint64
	CandidateName  string
	VoterAddress   string
	NativeFlag     bool
	AggregateVotes string
}

type EpochFoundationReward struct {
	EpochNumber     uint64
	EpochReward     string
	FoundationBonus string
	BlockReward     string
}

// rewardToSplit gets the reward to split from the given delegate from start epoch to end epoch
func rewardsToSplit(startEpoch uint64, endEpoch uint64, delegateName string, percentage int, includeBlockReward, includeFoundationBonus bool) (map[uint64]*big.Int, error) {
	var rewards []*EpochFoundationReward
	db := db.DB()
	if err := db.Table("hermes_account_rewards").Where("epoch_number >= ?  AND epoch_number <= ? AND candidate_name= ?", startEpoch, endEpoch, delegateName).Scan(&rewards).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	distrRewardMap := make(map[uint64]*big.Int)

	for _, reward := range rewards {
		rewardToSplit, _ := new(big.Int).SetString(reward.EpochReward, 10)
		if includeBlockReward {
			num, _ := new(big.Int).SetString(reward.BlockReward, 10)
			rewardToSplit.Add(rewardToSplit, num)
		}
		if includeFoundationBonus {
			num, _ := new(big.Int).SetString(reward.FoundationBonus, 10)
			rewardToSplit.Add(rewardToSplit, num)
		}
		distrRewardMap[reward.EpochNumber] = rewardToSplit.Mul(rewardToSplit, big.NewInt(int64(percentage))).Div(rewardToSplit, big.NewInt(100))
	}

	return distrRewardMap, nil
}

func totalWeightedVotes(startEpoch uint64, endEpoch uint64, delegateName string) (map[uint64]*big.Int, error) {
	var weights []*TotalWeight
	db := db.DB()
	if err := db.Table("hermes_voting_results").Where("epoch_number >= ?  AND epoch_number <= ? AND delegate_name= ?", startEpoch, endEpoch, delegateName).Scan(&weights).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	totalVotesMap := make(map[uint64]*big.Int)
	for _, row := range weights {
		votes, ok := new(big.Int).SetString(row.TotalWeightedVotes, 10)
		if !ok {
			return nil, errors.New("failed to covert string to big int")
		}
		totalVotesMap[row.EpochNumber] = votes
	}
	return totalVotesMap, nil
}

func voterVotes(startEpoch uint64, endEpoch uint64, delegateName string) (map[uint64]map[string]*big.Int, error) {
	var rows []*AggregateVoting
	db := db.DB()
	if err := db.Table("hermes_aggregate_votings").Where("epoch_number >= ?  AND epoch_number <= ? AND candidate_name= ?", startEpoch, endEpoch, delegateName).Scan(&rows).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	epochToVoters := make(map[uint64]map[string]*big.Int)
	for _, row := range rows {

		if _, ok := epochToVoters[row.EpochNumber]; !ok {
			epochToVoters[row.EpochNumber] = make(map[string]*big.Int)
		}
		votes, ok := new(big.Int).SetString(row.AggregateVotes, 10)
		if !ok {
			return nil, errors.New("failed to covert string to big int")
		}
		if val, ok := epochToVoters[row.EpochNumber][row.VoterAddress]; !ok {
			epochToVoters[row.EpochNumber][row.VoterAddress] = votes
		} else {
			val.Add(val, votes)
		}
	}

	return epochToVoters, nil
}

func GetBookkeeping(startEpoch uint64, epochCount uint64, delegateName string, percentage int, includeBlockReward, includeFoundationBonus bool) (map[string]*big.Int, error) {
	endEpoch := startEpoch + epochCount - 1

	distrRewardMap, err := rewardsToSplit(startEpoch, endEpoch, delegateName, percentage, includeBlockReward, includeFoundationBonus)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get reward distribution map")
	}
	delegateTotalVotesMap, err := totalWeightedVotes(startEpoch, endEpoch, delegateName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get delegate total weighted votes")
	}
	epochToVotersMap, err := voterVotes(startEpoch, endEpoch, delegateName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get voters map")
	}
	voterAddrToReward := make(map[string]*big.Int)
	for epoch, distrReward := range distrRewardMap {
		totalWeightedVotes, ok := delegateTotalVotesMap[epoch]
		if !ok {
			return nil, errors.Errorf("Missing delegate total weighted votes information on epoch %d", epoch)
		}
		if totalWeightedVotes.Sign() == 0 {
			continue
		}
		votersInfo, ok := epochToVotersMap[epoch]
		if !ok {
			return nil, errors.Errorf("Missing voters' weighted votes information on epoch %d", epoch)
		}
		for voterAddr, weightedVotes := range votersInfo {
			amount := new(big.Int).Set(distrReward)
			amount = amount.Mul(amount, weightedVotes).Div(amount, totalWeightedVotes)
			if _, ok := voterAddrToReward[voterAddr]; !ok {
				voterAddrToReward[voterAddr] = big.NewInt(0)
			}
			voterAddrToReward[voterAddr].Add(voterAddrToReward[voterAddr], amount)
		}
	}
	return voterAddrToReward, nil
}
