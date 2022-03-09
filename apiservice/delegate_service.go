package apiservice

import (
	"context"
	"sort"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/votings"
	"github.com/pkg/errors"
)

type DelegateService struct {
	api.UnimplementedDelegateServiceServer
}

func (s *DelegateService) GetBucketInfo(ctx context.Context, req *api.DelegateRequest) (*api.DelegateResponse, error) {
	resp := &api.DelegateResponse{}
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
	resp.BucketInfo = &api.BucketInfoOutput{Exist: true, BucketInfoList: bucketInfoLists}
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
