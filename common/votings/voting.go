package votings

import (
	"math/big"
	"time"

	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
	"github.com/pkg/errors"
)

func GetVoteBucketList(epochNum uint64) (*iotextypes.VoteBucketList, error) {
	chainClient, err := common.DefaultChainClient()
	if err != nil {
		return nil, err
	}
	return GetAllStakingBuckets(chainClient, common.GetEpochHeight(epochNum))
}

func GetBucketInfoByEpoch(epochNum uint64, delegateName string) ([]*VotingInfo, error) {
	height := common.GetEpochHeight(epochNum)
	if height >= common.FairbankEffectiveHeight() {
		return getStakingBucketInfoByEpoch(height, epochNum, delegateName)
	}
	return nil, errors.New("not supported")
}

func getStakingBucketInfoByEpoch(height uint64, epochNum uint64, delegateName string) ([]*VotingInfo, error) {
	bucketList, err := GetVoteBucketList(epochNum)
	if err != nil {
		return nil, err
	}
	candidateList, err := GetCandidateList(epochNum)
	if err != nil {
		return nil, err
	}

	var candidateAddress string
	for _, cand := range candidateList.Candidates {
		if cand.Name == delegateName {
			candidateAddress = cand.Id
			break
		}
	}

	// update weighted votes based on probation
	pblist, err := GetProbationList(epochNum)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get probation list from table")
	}

	intensityRate, probationMap := stakingProbationListToMap(candidateList, pblist)
	var votinginfoList []*VotingInfo
	selfStakeIndex := selfStakeIndexMap(candidateList)
	for _, vote := range bucketList.Buckets {
		if vote.UnstakeStartTime.AsTime().After(vote.StakeStartTime.AsTime()) {
			continue
		}
		if vote.CandidateAddress == candidateAddress {
			selfStake := false
			if _, ok := selfStakeIndex[vote.Index]; ok {
				selfStake = true
			}
			weightedVotes, err := CalculateVoteWeight(config.Default.Genesis.VoteWeightCalConsts, vote, selfStake)
			if err != nil {
				return nil, errors.Wrap(err, "calculate vote weight error")
			}
			if _, ok := probationMap[vote.CandidateAddress]; ok {
				// filter based on probation
				votingPower := new(big.Float).SetInt(weightedVotes)
				weightedVotes, _ = votingPower.Mul(votingPower, big.NewFloat(intensityRate)).Int(nil)
			}
			votinginfo := &VotingInfo{
				BucketID:          vote.Index,
				EpochNumber:       epochNum,
				VoterAddress:      vote.Owner,
				IsNative:          true,
				Votes:             vote.StakedAmount,
				WeightedVotes:     weightedVotes.Text(10),
				RemainingDuration: CalcRemainingTime(vote).String(),
				StartTime:         time.Unix(vote.StakeStartTime.Seconds, int64(vote.StakeStartTime.Nanos)).String(),
				Decay:             !vote.AutoStake,
			}
			votinginfoList = append(votinginfoList, votinginfo)
		}
	}
	return votinginfoList, nil
}

func selfStakeIndexMap(candidates *iotextypes.CandidateListV2) map[uint64]struct{} {
	ret := make(map[uint64]struct{})
	for _, can := range candidates.Candidates {
		ret[can.SelfStakeBucketIdx] = struct{}{}
	}
	return ret
}
