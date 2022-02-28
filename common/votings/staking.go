package votings

import (
	"math"
	"math/big"
	"time"

	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-core/blockchain/genesis"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func GetCandidateList(epochNum uint64) (*iotextypes.CandidateListV2, error) {
	candidateListAll := &iotextypes.CandidateListV2{}
	var vbl CandidateList
	if err := db.DB().Table("candidate_list").Where("epoch_number = ?", epochNum).First(&vbl).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to get candidate list in epoch %d", epochNum)
	}
	if err := proto.Unmarshal(vbl.CandidateList, candidateListAll); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal candidate list in epoch %d", epochNum)
	}
	return candidateListAll, nil
}

func stakingProbationListToMap(candidateList *iotextypes.CandidateListV2, probationList []*ProbationList) (intensityRate float64, probationMap map[string]uint64) {
	probationMap = make(map[string]uint64)
	if probationList != nil {
		for _, can := range candidateList.Candidates {
			for _, pb := range probationList {
				intensityRate = float64(uint64(100)-pb.IntensityRate) / float64(100)
				if pb.Address == can.OperatorAddress {
					probationMap[can.OwnerAddress] = pb.Count
				}
			}
		}
	}
	return
}

// CalculateVoteWeight calculates the weighted votes
func CalculateVoteWeight(c genesis.VoteWeightCalConsts, v *iotextypes.VoteBucket, selfStake bool) (*big.Int, error) {
	// TODO: calculation of remaining time is wrong
	remainingTime := float64(v.StakedDuration * 86400)
	weight := float64(1)
	var m float64
	if v.AutoStake {
		m = c.AutoStake
	}
	if remainingTime > 0 {
		weight += math.Log(math.Ceil(remainingTime/86400)*(1+m)) / math.Log(c.DurationLg) / 100
	}
	if selfStake && v.AutoStake && v.StakedDuration >= 91 {
		// self-stake extra bonus requires enable auto-stake for at least 3 months
		weight *= c.SelfStake
	}

	amountInt, ok := big.NewInt(0).SetString(v.StakedAmount, 10)
	if !ok {
		return nil, errors.New("failed to convert string to big int")
	}
	amount := new(big.Float).SetInt(amountInt)
	weightedAmount, _ := amount.Mul(amount, big.NewFloat(weight)).Int(nil)
	return weightedAmount, nil
}

// CalcRemainingTime calculates the remaining time of a bucket
func CalcRemainingTime(bucket *iotextypes.VoteBucket) time.Duration {
	now := time.Now()
	startTime := time.Unix(bucket.StakeStartTime.Seconds, int64(bucket.StakeStartTime.Nanos))
	if now.Before(startTime) {
		return 0
	}
	duration := time.Duration(bucket.StakedDuration) * 24 * time.Hour
	if !bucket.AutoStake {
		endTime := startTime.Add(duration)
		if endTime.After(now) {
			return endTime.Sub(now)
		}
		return 0
	}
	return duration
}
