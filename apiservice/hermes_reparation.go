package apiservice

import (
	"math/big"

	"github.com/iotexproject/iotex-analyser-api/common/rewards"
	"github.com/pkg/errors"
)

type HermesDistributionReward map[string]map[string]*big.Int

func GetHermeOrigin(startEpoch, epochCount uint64) (HermesDistributionReward, error) {
	rewardAddress := "io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85"
	endEpoch := startEpoch + epochCount - 1
	distributePlanMap, err := distributionPlanByRewardAddress(startEpoch, endEpoch, rewardAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get reward distribution plan")
	}
	delegateMap := make(map[uint64][]string)
	for delegateName, planMap := range distributePlanMap {
		for epochNumber := range planMap {
			if _, ok := delegateMap[epochNumber]; !ok {
				delegateMap[epochNumber] = make([]string, 0)
			}
			delegateMap[epochNumber] = append(delegateMap[epochNumber], delegateName)
		}
	}
	accountRewardsMap, err := accountRewards(delegateMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account rewards")
	}

	voterVotesMap, err := rewards.WeightedVotesBySearchPairs(delegateMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get voter votes")
	}
	hermesDistributions := make(HermesDistributionReward)
	for delegate, rewardsMap := range accountRewardsMap {
		planMap := distributePlanMap[delegate]
		epochVoterMap := voterVotesMap[delegate]

		voterAddrToReward := make(map[string]*big.Int)
		for epoch, rewards := range rewardsMap {
			distributePlan := planMap[epoch]
			voterMap := epochVoterMap[epoch]

			distrReward, err := calculatedDistributedReward(distributePlan, rewards)
			if err != nil {
				return nil, errors.Wrap(err, "failed to calculate distributed reward")
			}
			for voterAddr, weightedVotes := range voterMap {
				amount := new(big.Int).Set(distrReward)
				amount = amount.Mul(amount, weightedVotes).Div(amount, distributePlan.TotalWeightedVotes)
				if _, ok := voterAddrToReward[voterAddr]; !ok {
					voterAddrToReward[voterAddr] = big.NewInt(0)
				}
				voterAddrToReward[voterAddr].Add(voterAddrToReward[voterAddr], amount)
			}
		}
		hermesDistributions[delegate] = voterAddrToReward
	}
	return hermesDistributions, nil
}

func GetHermeFixed(startEpoch, epochCount uint64) (HermesDistributionReward, error) {
	rewardAddress := "io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85"
	endEpoch := startEpoch + epochCount - 1
	distributePlanMap, err := distributionPlanByRewardAddress(startEpoch, endEpoch, rewardAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get reward distribution plan")
	}
	delegateMap := make(map[uint64][]string)
	for delegateName, planMap := range distributePlanMap {
		for epochNumber := range planMap {
			if _, ok := delegateMap[epochNumber]; !ok {
				delegateMap[epochNumber] = make([]string, 0)
			}
			delegateMap[epochNumber] = append(delegateMap[epochNumber], delegateName)
		}
	}
	accountRewardsMap, err := accountRewards(delegateMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account rewards")
	}

	voterVotesMap, err := rewards.WeightedVotesBySearchPairsFix(delegateMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get voter votes")
	}
	hermesDistributions := make(HermesDistributionReward)
	for delegate, rewardsMap := range accountRewardsMap {
		planMap := distributePlanMap[delegate]
		epochVoterMap := voterVotesMap[delegate]

		voterAddrToReward := make(map[string]*big.Int)
		for epoch, rewards := range rewardsMap {
			distributePlan := planMap[epoch]
			voterMap := epochVoterMap[epoch]

			distrReward, err := calculatedDistributedReward(distributePlan, rewards)
			if err != nil {
				return nil, errors.Wrap(err, "failed to calculate distributed reward")
			}
			for voterAddr, weightedVotes := range voterMap {
				amount := new(big.Int).Set(distrReward)
				amount = amount.Mul(amount, weightedVotes).Div(amount, distributePlan.TotalWeightedVotes)
				if _, ok := voterAddrToReward[voterAddr]; !ok {
					voterAddrToReward[voterAddr] = big.NewInt(0)
				}
				voterAddrToReward[voterAddr].Add(voterAddrToReward[voterAddr], amount)
			}
		}
		hermesDistributions[delegate] = voterAddrToReward
	}
	return hermesDistributions, nil
}
