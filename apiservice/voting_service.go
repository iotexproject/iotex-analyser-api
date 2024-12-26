package apiservice

import (
	"context"
	"math/big"
	"sort"
	"sync"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-analyser-api/internal/sync/errgroup"
	"github.com/iotexproject/iotex-core/v2/blockchain/genesis"
	"github.com/pkg/errors"
)

// VotingService provides voting service
type VotingService struct {
	api.UnimplementedVotingServiceServer
}

// CandidateInfo provides candidate information
func (s *VotingService) CandidateInfo(ctx context.Context, req *api.CandidateInfoRequest) (*api.CandidateInfoResponse, error) {
	resp := &api.CandidateInfoResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	epoch, _, err := common.GetCurrentEpochAndHeight()
	if err != nil {
		return nil, err
	}
	endEpoch := startEpoch + epochCount - 1
	if endEpoch > epoch {
		endEpoch = epoch
	}
	g := errgroup.Group{}
	g.GOMAXPROCS(8)
	getCandidateInfoByEpoch := func(epochNumber uint64) ([]*api.CandidateInfoResponse_Candidates, error) {
		var result []struct {
			DelegateName       string
			StakingAddress     string
			TotalWeightedVotes string
			SelfStaking        string
			OperatorAddress    string
			RewardAddress      string
		}
		db := db.DB()
		query := "SELECT delegate_name, staking_address, total_weighted_votes, self_staking, operator_address, reward_address FROM hermes_voting_results WHERE epoch_number = ?"
		if err := db.Raw(query, epochNumber).Scan(&result).Error; err != nil {
			return nil, err
		}
		var candidates []*api.CandidateInfoResponse_Candidates
		for _, candidate := range result {
			candidates = append(candidates, &api.CandidateInfoResponse_Candidates{
				Name:               candidate.DelegateName,
				Address:            candidate.StakingAddress,
				TotalWeightedVotes: candidate.TotalWeightedVotes,
				SelfStakingTokens:  candidate.SelfStaking,
				OperatorAddress:    candidate.OperatorAddress,
				RewardAddress:      candidate.RewardAddress,
			})
		}
		return candidates, nil
	}
	var epochMap sync.Map
	for epoch := startEpoch; epoch <= endEpoch; epoch++ {
		epoch := epoch
		g.Go(func(context.Context) error {
			candidates, err := getCandidateInfoByEpoch(epoch)
			if err != nil {
				return err
			}
			epochMap.Store(epoch, candidates)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	epochMap.Range(func(key, value interface{}) bool {
		resp.CandidateInfo = append(resp.CandidateInfo, &api.CandidateInfoResponse_CandidateInfo{
			EpochNumber: key.(uint64),
			Candidates:  value.([]*api.CandidateInfoResponse_Candidates),
		})
		return true
	})
	sort.Slice(resp.CandidateInfo, func(i, j int) bool {
		return resp.CandidateInfo[i].EpochNumber < resp.CandidateInfo[j].EpochNumber
	})
	return resp, nil
}

// RewardSources provides reward sources for voters
func (s *VotingService) RewardSources(ctx context.Context, req *api.RewardSourcesRequest) (*api.RewardSourcesResponse, error) {
	resp := &api.RewardSourcesResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	voterIotxAddress := req.GetVoterIotxAddress()
	endEpoch := startEpoch + epochCount - 1
	weightedVotesMap, err := weightedVotesByVoterAddress(startEpoch, endEpoch, voterIotxAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get voter's weighted votes")
	}
	delegateMap := make(map[uint64][]string)
	for delegateName, planMap := range weightedVotesMap {
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
	distributePlanMap, err := distributionPlanByDelegate(delegateMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get reward distribution plan")
	}
	delegateDistributionMap := make(map[string]*big.Int)
	for delegate, rewardsMap := range accountRewardsMap {
		planMap := distributePlanMap[delegate]
		delegateMap := weightedVotesMap[delegate]
		delegateDistributionMap[delegate] = big.NewInt(0)

		for epoch, rewards := range rewardsMap {
			distributePlan := planMap[epoch]
			distrReward, err := calculatedDistributedReward(distributePlan, rewards)
			if err != nil {
				return nil, errors.Wrap(err, "failed to calculate reward distribution plan")
			}
			weightedVotes := delegateMap[epoch]
			amount := distrReward.Mul(distrReward, weightedVotes).Div(distrReward, distributePlan.TotalWeightedVotes)

			delegateDistributionMap[delegate].Add(delegateDistributionMap[delegate], amount)
		}
	}

	for delegateName, amount := range delegateDistributionMap {
		resp.DelegateDistributions = append(resp.DelegateDistributions, &api.RewardSourcesResponse_DelegateDistributions{
			DelegateName: delegateName,
			Amount:       amount.String(),
		})
	}
	resp.Exist = len(resp.DelegateDistributions) > 0
	return resp, nil
}

// VotingMeta provides metadata of voting results
func (s *VotingService) VotingMeta(ctx context.Context, req *api.VotingMetaRequest) (*api.VotingMetaResponse, error) {
	resp := &api.VotingMetaResponse{
		CandidateMeta: make([]*api.VotingMetaResponse_CandidateMeta, 0),
	}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	endEpoch := startEpoch + epochCount - 1
	var results []struct {
		EpochNumber        uint64
		VotedToken         string
		DelegateCount      uint64
		TotalWeightedVotes string
	}
	db := db.DB()
	query := "select epoch_number,voted_token, delegate_count,total_weighted_votes from hermes_voting_meta where epoch_number>=? and epoch_number<=?"
	if err := db.Raw(query, startEpoch, endEpoch).Scan(&results).Error; err != nil {
		return nil, err
	}
	for _, result := range results {
		resp.CandidateMeta = append(resp.CandidateMeta, &api.VotingMetaResponse_CandidateMeta{
			EpochNumber:        result.EpochNumber,
			TotalCandidates:    result.DelegateCount,
			TotalWeightedVotes: result.TotalWeightedVotes,
			ConsensusDelegates: genesis.Default.NumCandidateDelegates,
			VotedTokens:        result.VotedToken,
		})
	}
	sort.Slice(resp.CandidateMeta, func(i, j int) bool {
		return resp.CandidateMeta[i].EpochNumber < resp.CandidateMeta[j].EpochNumber
	})
	resp.Exist = len(resp.CandidateMeta) > 0
	return resp, nil
}

// weightedVotesByVoterAddress gets voter's weighted votes for delegates by voter's address
func weightedVotesByVoterAddress(startEpoch uint64, endEpoch uint64, voterAddress string) (map[string]map[uint64]*big.Int, error) {
	db := db.DB()

	query := "SELECT epoch_number,candidate_name,aggregate_votes FROM hermes_aggregate_votings WHERE epoch_number >= ?  AND epoch_number <= ? AND voter_address= ?"

	var results []struct {
		EpochNumber    uint64
		CandidateName  string
		AggregateVotes string
	}
	if err := db.Raw(query, startEpoch, endEpoch, voterAddress).Scan(&results).Error; err != nil {
		return nil, err
	}

	weightedVotesMap := make(map[string]map[uint64]*big.Int)
	for _, r := range results {
		if _, ok := weightedVotesMap[r.CandidateName]; !ok {
			weightedVotesMap[r.CandidateName] = make(map[uint64]*big.Int)
		}
		weightedVotesInt, errs := stringToBigInt(r.AggregateVotes)
		if errs != nil {
			return nil, errors.Wrap(errs, "failed to convert to big int")

		}
		if val, ok := weightedVotesMap[r.CandidateName][r.EpochNumber]; !ok {
			weightedVotesMap[r.CandidateName][r.EpochNumber] = weightedVotesInt
		} else {
			val.Add(val, weightedVotesInt)
		}
	}
	return weightedVotesMap, nil
}

func distributionPlanByDelegate(delegateMap map[uint64][]string) (map[string]map[uint64]*HermesDistributionPlan, error) {

	db := db.DB()
	var conds [][]interface{}
	for k, delegates := range delegateMap {
		for _, delegate := range delegates {
			conds = append(conds, []interface{}{k, delegate})
		}
	}
	var ids []HermesVotingResult
	if err := db.Table("hermes_voting_results").Where("(epoch_number, delegate_name) IN ?", conds).Find(&ids).Error; err != nil {
		return nil, err
	}

	return parseDistributionPlanFromVotingResult(ids)
}
