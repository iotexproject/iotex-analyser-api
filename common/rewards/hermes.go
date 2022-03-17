package rewards

import (
	"context"
	"math"
	"math/big"
	"sync"

	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-analyser-api/internal/sync/errgroup"
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

func voterVotes(ctx context.Context, startEpoch uint64, endEpoch uint64, delegateName string) (map[uint64]map[string]*big.Int, error) {

	g := errgroup.Group{}
	g.GOMAXPROCS(8)
	db := db.DB()
	f := func(ctx context.Context, epochNum uint64, delegateName string) (map[string]*big.Int, error) {
		var votes []*AggregateVoting
		if err := db.Table("hermes_aggregate_votings").Select("voter_address,aggregate_votes").Where("epoch_number = ? AND candidate_name = ?", epochNum, delegateName).Scan(&votes).Error; err != nil {
			return nil, errors.WithStack(err)
		}
		voterVotes := make(map[string]*big.Int)
		for _, vote := range votes {
			a, ok := new(big.Int).SetString(vote.AggregateVotes, 10)
			if !ok {
				return nil, errors.New("failed to covert string to big int")
			}
			if val, ok := voterVotes[vote.VoterAddress]; ok {
				voterVotes[vote.VoterAddress] = new(big.Int).Add(val, a)
			} else {
				voterVotes[vote.VoterAddress] = a
			}
		}
		return voterVotes, nil
	}
	var epochMap sync.Map
	for epoch := startEpoch; epoch <= endEpoch; epoch++ {
		epoch := epoch // fix https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func(context.Context) error {
			voters, err := f(ctx, epoch, delegateName)
			if err != nil {
				return err
			}
			epochMap.Store(epoch, voters)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	epochToVoters := make(map[uint64]map[string]*big.Int)
	epochMap.Range(func(key, value interface{}) bool {
		epochToVoters[key.(uint64)] = value.(map[string]*big.Int)
		return true
	})

	return epochToVoters, nil
}

func GetBookkeeping(ctx context.Context, startEpoch uint64, epochCount uint64, delegateName string, percentage int, includeBlockReward, includeFoundationBonus bool) (map[string]*big.Int, error) {
	endEpoch := startEpoch + epochCount - 1

	distrRewardMap, err := rewardsToSplit(startEpoch, endEpoch, delegateName, percentage, includeBlockReward, includeFoundationBonus)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get reward distribution map")
	}
	delegateTotalVotesMap, err := totalWeightedVotes(startEpoch, endEpoch, delegateName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get delegate total weighted votes")
	}
	epochToVotersMap, err := voterVotes(ctx, startEpoch, endEpoch, delegateName)
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

func WeightedVotesBySearchPairs(delegateMap map[uint64][]string) (map[string]map[uint64]map[string]*big.Int, error) {
	db := db.DB()
	g := errgroup.Group{}
	g.GOMAXPROCS(8)
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

	f := func(ctx context.Context, epochNum uint64) ([]*AggregateVoting, error) {
		var votes []*AggregateVoting
		if err := db.Table("hermes_aggregate_votings").Select("candidate_name,voter_address,aggregate_votes").Where("epoch_number = ?", epochNum).Scan(&votes).Error; err != nil {
			return nil, errors.WithStack(err)
		}
		return votes, nil
	}
	var epochMap sync.Map
	for epoch := minEpoch; epoch <= maxEpoch; epoch++ {
		epoch := epoch
		g.Go(func(ctx context.Context) error {
			voters, err := f(ctx, epoch)
			if err != nil {
				return err
			}
			epochMap.Store(epoch, voters)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	voterVotesMap := make(map[string]map[uint64]map[string]*big.Int) //map[candidateName][epoch][voterAddr]weightedVotes
	epochMap.Range(func(key, value interface{}) bool {
		epoch := key.(uint64)
		voters := value.([]*AggregateVoting)
		for _, row := range voters {
			exist := false
			for _, v := range delegateMap[epoch] {
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

			weightedVotesInt, ok := new(big.Int).SetString(row.AggregateVotes, 10)
			if !ok {
				return false
			}
			if val, ok := voterMap[row.VoterAddress]; !ok {
				voterMap[row.VoterAddress] = weightedVotesInt
			} else {
				voterMap[row.VoterAddress] = new(big.Int).Add(val, weightedVotesInt)
			}
		}
		return true
	})
	return voterVotesMap, nil
}
