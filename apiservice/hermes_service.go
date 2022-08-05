package apiservice

import (
	"context"
	"math/big"
	"sort"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/rewards"
	"github.com/pkg/errors"
)

// HermesService provides hermes service
type HermesService struct {
	api.UnimplementedHermesServiceServer
}

//grpcurl -plaintext -d '{"startEpoch": 22416, "epochCount": 1, "rewardAddress": "io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85"}' 127.0.0.1:8888 api.AccountService.Hermes
func (s *HermesService) Hermes(ctx context.Context, req *api.HermesRequest) (*api.HermesResponse, error) {
	resp := &api.HermesResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	rewardAddress := req.GetRewardAddress()
	endEpoch := startEpoch + epochCount - 1
	waiverThreshold := 100

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
	hermesDistributions := make([]*DelegateHermesDistribution, 0, len(accountRewardsMap))
	for delegate, rewardsMap := range accountRewardsMap {
		planMap := distributePlanMap[delegate]
		epochVoterMap := voterVotesMap[delegate]

		voterAddrToReward := make(map[string]*big.Int)
		balanceAfterDistribution := big.NewInt(0)
		voterCountMap := make(map[string]bool)
		feeWaiver := true
		var stakingAddress string

		for epoch, rewards := range rewardsMap {
			distributePlan := planMap[epoch]
			voterMap := epochVoterMap[epoch]

			if stakingAddress == "" {
				stakingAddress = distributePlan.StakingAddress
			}

			totalRewards := new(big.Int).Set(rewards.BlockReward)
			totalRewards.Add(totalRewards, rewards.EpochReward).Add(totalRewards, rewards.FoundationBonus)
			balanceAfterDistribution.Add(balanceAfterDistribution, totalRewards)
			waiverThresholdF := float64(waiverThreshold)
			if distributePlan.BlockRewardPercentage < waiverThresholdF || distributePlan.EpochRewardPercentage < waiverThresholdF || distributePlan.FoundationBonusPercentage < waiverThresholdF {
				feeWaiver = false
			}
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
				balanceAfterDistribution.Sub(balanceAfterDistribution, amount)
				voterCountMap[voterAddr] = true
			}
		}
		rewardDistribution, err := convertVoterDistributionMapToList(voterAddrToReward)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert voter distribution map to list")
		}
		hermesDistributions = append(hermesDistributions, &DelegateHermesDistribution{
			DelegateName:        delegate,
			Distributions:       rewardDistribution,
			StakingIotexAddress: stakingAddress,
			VoterCount:          uint64(len(voterCountMap)),
			WaiveServiceFee:     feeWaiver,
			Refund:              balanceAfterDistribution.String(),
		})
	}

	hermesDistribution := make([]*api.HermesDistribution, 0, len(hermesDistributions))
	for _, ret := range hermesDistributions {
		rds := make([]*api.RewardDistribution, 0)
		for _, distribution := range ret.Distributions {
			v := &api.RewardDistribution{
				VoterEthAddress:   distribution.VoterEthAddress,
				VoterIotexAddress: distribution.VoterIotexAddress,
				Amount:            distribution.Amount,
			}
			rds = append(rds, v)
		}
		sort.Slice(rds, func(i, j int) bool { return rds[i].VoterEthAddress < rds[j].VoterEthAddress })

		hermesDistribution = append(hermesDistribution, &api.HermesDistribution{
			DelegateName:        ret.DelegateName,
			RewardDistribution:  rds,
			StakingIotexAddress: ret.StakingIotexAddress,
			VoterCount:          ret.VoterCount,
			WaiveServiceFee:     ret.WaiveServiceFee,
			Refund:              ret.Refund,
		})
	}
	sort.Slice(hermesDistribution, func(i, j int) bool { return hermesDistribution[i].DelegateName < hermesDistribution[j].DelegateName })
	resp.HermesDistribution = hermesDistribution
	return resp, nil
}

//HermesByVoter returns Hermes voters' receiving history
func (s *HermesService) HermesByVoter(ctx context.Context, req *api.HermesByVoterRequest) (*api.HermesByVoterResponse, error) {
	resp := &api.HermesByVoterResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	voterAddress := req.GetVoterAddress()
	endEpoch := startEpoch + epochCount - 1
	if count, sum, err := rewards.GetTotalHermesByVoter(ctx, startEpoch, endEpoch, voterAddress); err != nil {
		return nil, err
	} else {
		resp.Count = uint64(count)
		resp.TotalRewardReceived = sum
	}
	if resp.Count > 0 {
		resp.Exist = true
		skip := common.PageOffset(req.GetPagination())
		first := common.PageSize(req.GetPagination())
		result, err := rewards.GetHermesByVoter(ctx, startEpoch, endEpoch, voterAddress, skip, first)
		if err != nil {
			return nil, err
		}
		for _, v := range result {
			resp.Delegates = append(resp.Delegates, &api.HermesByVoterResponse_Delegate{
				DelegateName: v.DelegateName,
				FromEpoch:    v.StartEpoch,
				ToEpoch:      v.EndEpoch,
				Amount:       v.Amount,
				ActHash:      v.ActionHash,
				Timestamp:    uint64(v.Timestamp.Unix()),
			})
		}
	}
	return resp, nil
}

func (s *HermesService) HermesByDelegate(ctx context.Context, req *api.HermesByDelegateRequest) (*api.HermesByDelegateResponse, error) {
	resp := &api.HermesByDelegateResponse{
		VoterInfoList: make([]*api.HermesByDelegateVoterInfo, 0),
	}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	endEpoch := startEpoch + epochCount - 1
	if count, sum, err := rewards.GetTotalHermesByDelegate(ctx, startEpoch, endEpoch, delegateName); err != nil {
		return nil, err
	} else {
		resp.Count = uint64(count)
		resp.TotalRewardsDistributed = sum
	}
	if resp.Count > 0 {
		resp.Exist = true
		skip := common.PageOffset(req.GetPagination())
		first := common.PageSize(req.GetPagination())
		result, err := rewards.GetHermesByDelegate(ctx, startEpoch, endEpoch, delegateName, skip, first)
		if err != nil {
			return nil, err
		}
		for _, v := range result {
			resp.VoterInfoList = append(resp.VoterInfoList, &api.HermesByDelegateVoterInfo{
				VoterAddress: v.Recipient,
				FromEpoch:    v.StartEpoch,
				ToEpoch:      v.EndEpoch,
				Amount:       v.Amount,
				ActHash:      v.ActionHash,
				Timestamp:    uint64(v.Timestamp.Unix()),
			})
		}
	}
	if results, err := rewards.GetHermesRatioByDelegate(ctx, startEpoch, endEpoch, delegateName); err != nil {
		return nil, err
	} else {
		resp.DistributionRatio = make([]*api.HermesByDelegateDistributionRatio, 0)
		for _, v := range results {
			resp.DistributionRatio = append(resp.DistributionRatio, &api.HermesByDelegateDistributionRatio{
				EpochNumber:          v.EpochNumber,
				BlockRewardRatio:     v.BlockRewardPercentage,
				EpochRewardRatio:     v.EpochRewardPercentage,
				FoundationBonusRatio: v.FoundationBonusPercentage,
			})
		}
	}
	return resp, nil
}

// HermesMeta provides Hermes platform metadata
func (s *HermesService) HermesMeta(ctx context.Context, req *api.HermesMetaRequest) (*api.HermesMetaResponse, error) {
	resp := &api.HermesMetaResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	endEpoch := startEpoch + epochCount - 1
	meta, err := rewards.GetHermesMeta(ctx, startEpoch, endEpoch)
	if err != nil {
		return nil, err
	}
	resp.Exist = meta.Count > 0
	resp.NumberOfDelegates = meta.NumberOfDelegates
	resp.NumberOfRecipients = meta.NumberOfRecipients
	resp.TotalRewardDistributed = meta.TotalRewardsDistributed

	return resp, nil
}

// HermesAverageStats returns the Hermes average statistics
func (s *HermesService) HermesAverageStats(ctx context.Context, req *api.HermesAverageStatsRequest) (*api.HermesAverageStatsResponse, error) {
	resp := &api.HermesAverageStatsResponse{
		AveragePerEpoch: make([]*api.HermesAverageStatsResponse_AveragePerEpoch, 0),
	}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	rewardAddress := req.GetRewardAddress()
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
	delegateAverages := make([]*rewards.DelegateHermesAverage, 0, len(accountRewardsMap))
	for delegate, rewardsMap := range accountRewardsMap {
		planMap := distributePlanMap[delegate]
		distrRewardSum := big.NewInt(0)
		totalWeightedVotesSum := big.NewInt(0)

		for epoch, rewards := range rewardsMap {
			distributePlan := planMap[epoch]
			totalWeightedVotesSum.Add(totalWeightedVotesSum, distributePlan.TotalWeightedVotes)
			distrReward, err := calculatedDistributedReward(distributePlan, rewards)
			if err != nil {
				return nil, errors.Wrap(err, "failed to calculate reward distribution plan")
			}
			distrRewardSum = distrRewardSum.Add(distrRewardSum, distrReward)
		}

		avgRewardDistribution := distrRewardSum.Div(distrRewardSum, big.NewInt(int64(len(rewardsMap))))
		avgTotalWeightedVotes := totalWeightedVotesSum.Div(totalWeightedVotesSum, big.NewInt(int64(len(rewardsMap))))

		delegateAverages = append(delegateAverages, &rewards.DelegateHermesAverage{
			DelegateName:       delegate,
			RewardDistribution: avgRewardDistribution.String(),
			TotalWeightedVotes: avgTotalWeightedVotes.String(),
		})
	}
	hermesAverages := make([]*api.HermesAverageStatsResponse_AveragePerEpoch, 0, len(delegateAverages))
	for _, ret := range delegateAverages {

		hermesAverages = append(hermesAverages, &api.HermesAverageStatsResponse_AveragePerEpoch{
			DelegateName:       ret.DelegateName,
			RewardDistribution: ret.RewardDistribution,
			TotalWeightedVotes: ret.TotalWeightedVotes,
		})
	}
	sort.Slice(hermesAverages, func(i, j int) bool { return hermesAverages[i].DelegateName < hermesAverages[j].DelegateName })
	resp.Exist = len(hermesAverages) > 0
	resp.AveragePerEpoch = hermesAverages
	return resp, nil
}
