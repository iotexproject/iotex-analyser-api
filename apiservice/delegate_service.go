package apiservice

import (
	"context"
	"sort"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/rewards"
	"github.com/iotexproject/iotex-analyser-api/common/votings"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
)

type DelegateService struct {
	api.UnimplementedDelegateServiceServer
}

func (s *DelegateService) BucketInfo(ctx context.Context, req *api.BucketInfoRequest) (*api.BucketInfoResponse, error) {
	resp := &api.BucketInfoResponse{
		BucketInfoList: make([]*api.BucketInfoList, 0),
	}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	bucketMap, err := s.getBucketInformation(startEpoch, epochCount, delegateName)
	if err != nil {
		return nil, err
	}
	bucketInfoLists := make([]*api.BucketInfoList, 0)
	for epoch, bucketList := range bucketMap {
		bucketInfoList := &api.BucketInfoList{EpochNumber: epoch, Count: uint64(len(bucketList))}
		bucketInfo := make([]*api.BucketInfo, 0)
		for _, bucket := range bucketList {
			voterIotexAddr := bucket.VoterAddress
			voterAddr, _ := address.FromString(voterIotexAddr)
			bucketInfo = append(bucketInfo, &api.BucketInfo{
				BucketID:          bucket.BucketID,
				VoterEthAddress:   voterAddr.Hex(),
				VoterIotexAddress: voterIotexAddr,
				IsNative:          bucket.IsNative,
				Votes:             bucket.Votes,
				WeightedVotes:     bucket.WeightedVotes,
				RemainingDuration: bucket.RemainingDuration,
				StartTime:         bucket.StartTime,
				Decay:             bucket.Decay,
			})
		}
		page := req.GetPagination()
		var skip, first uint64

		if page != nil {
			skip = page.GetSkip()
			first = page.GetFirst()
		}
		if skip >= uint64(len(bucketInfo)) {
			return nil, errors.New("invalid pagination skip number for bucket info")
		}
		if uint64(len(bucketInfo))-skip < first {
			first = uint64(len(bucketInfo)) - skip
		}
		bucketInfoList.BucketInfo = bucketInfo[skip : skip+first]
		bucketInfoLists = append(bucketInfoLists, bucketInfoList)
	}
	sort.Slice(bucketInfoLists, func(i, j int) bool { return bucketInfoLists[i].EpochNumber < bucketInfoLists[j].EpochNumber })
	resp.Count = uint64(len(bucketInfoLists))
	resp.Exist = resp.Count > 0
	resp.BucketInfoList = bucketInfoLists
	return resp, nil
}

func (s *DelegateService) getBucketInformation(startEpoch, epochCount uint64, delegateName string) (map[uint64][]*votings.VotingInfo, error) {
	currentEpoch, _, err := common.GetCurrentEpochAndHeight()
	if err != nil {
		return nil, errors.New("failed to get most recent epoch")
	}
	endEpoch := startEpoch + epochCount - 1
	if endEpoch > currentEpoch {
		endEpoch = currentEpoch
	}
	bucketInfoMap := make(map[uint64][]*votings.VotingInfo)
	for i := startEpoch; i <= endEpoch; i++ {
		voteInfoList, err := votings.GetBucketInfoByEpoch(i, delegateName)
		if err != nil {
			return nil, err
		}
		bucketInfoMap[i] = voteInfoList
	}
	return bucketInfoMap, nil
}

func (s *DelegateService) BookKeeping(ctx context.Context, req *api.BookKeepingRequest) (*api.BookKeepingResponse, error) {
	resp := &api.BookKeepingResponse{
		RewardDistribution: make([]*api.DelegateRewardDistribution, 0),
	}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	includeBlockReward := req.GetIncludeBlockReward()
	includeFoundationBonus := req.GetIncludeFoundationBonus()
	percentAge := req.GetPercentage()

	if percentAge > 100 {
		return nil, errors.New("percentage should be 0-100")
	}
	rewards, err := rewards.GetBookkeeping(ctx, startEpoch, epochCount, delegateName, int(percentAge), includeBlockReward, includeFoundationBonus)
	if err != nil {
		return nil, err
	}
	rds := make([]*api.DelegateRewardDistribution, 0)
	for ioAddr, amount := range rewards {
		voterAddr, _ := address.FromString(ioAddr)
		v := &api.DelegateRewardDistribution{
			VoterEthAddress:   voterAddr.Hex(),
			VoterIotexAddress: ioAddr,
			Amount:            amount.String(),
		}
		rds = append(rds, v)
	}

	sort.Slice(rds, func(i, j int) bool { return rds[i].VoterEthAddress < rds[j].VoterEthAddress })

	resp.Count = uint64(len(rds))
	resp.Exist = resp.Count > 0

	page := req.GetPagination()
	var skip, first uint64

	if page != nil {
		skip = page.GetSkip()
		first = page.GetFirst()
	}
	if skip >= uint64(resp.Count) {
		return nil, errors.New("invalid pagination skip number")
	}
	if resp.Count-skip < first {
		first = resp.Count - skip
	}
	if first == skip && first == 0 {
		resp.RewardDistribution = rds
		return resp, nil
	} else {
		resp.RewardDistribution = rds[skip : skip+first]
	}
	return resp, nil
}

func (s *DelegateService) Productivity(ctx context.Context, req *api.ProductivityRequest) (*api.ProductivityResponse, error) {
	resp := &api.ProductivityResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	endEpoch := startEpoch + epochCount - 1
	db := db.DB()
	query := "select count(case when producer_name=? then block_height end) production, count(case when expected_producer_name=? then block_height end) expected_production from block_meta where epoch_num>=? and epoch_num<=?"
	var result struct {
		Production         uint64
		ExpectedProduction uint64
	}
	err := db.Raw(query, delegateName, delegateName, startEpoch, endEpoch).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	resp.Productivity = &api.Productivity{
		Exist:              result.Production > 0 || result.ExpectedProduction > 0,
		Production:         result.Production,
		ExpectedProduction: result.ExpectedProduction,
	}
	return resp, nil
}

func (s *DelegateService) Reward(ctx context.Context, req *api.RewardRequest) (*api.RewardResponse, error) {
	resp := &api.RewardResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	endEpoch := startEpoch + epochCount - 1
	db := db.DB()
	query := "SELECT SUM(block_reward) block_reward, SUM(epoch_reward) epoch_reward, SUM(foundation_bonus) foundation_bonus FROM hermes_account_rewards WHERE epoch_number >= ?  AND epoch_number <= ? AND candidate_name=?"
	var result struct {
		BlockReward     string
		EpochReward     string
		FoundationBonus string
	}
	err := db.Raw(query, startEpoch, endEpoch, delegateName).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	resp.Reward = &api.Reward{
		Exist:           result.BlockReward != "0" || result.EpochReward != "0" || result.FoundationBonus != "0",
		BlockReward:     result.BlockReward,
		EpochReward:     result.EpochReward,
		FoundationBonus: result.FoundationBonus,
	}
	return resp, nil
}
