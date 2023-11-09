package rewards

import (
	"context"
	"encoding/csv"
	"fmt"
	"math"
	"math/big"
	"os"
	"sync"
	"time"

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
func rewardsToSplit(startEpoch uint64, endEpoch uint64, delegateName string, epochRewardPerc, blockRewardPerc, foundationBonusPerc uint64) (map[uint64]*big.Int, error) {
	var rewards []*EpochFoundationReward
	db := db.DB()
	if err := db.Table("hermes_account_rewards").Where("epoch_number >= ?  AND epoch_number <= ? AND candidate_name= ?", startEpoch, endEpoch, delegateName).Scan(&rewards).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	distrRewardMap := make(map[uint64]*big.Int)

	for _, reward := range rewards {
		rewardToSplit, _ := new(big.Int).SetString(reward.EpochReward, 10)
		blockReward, _ := new(big.Int).SetString(reward.BlockReward, 10)
		rewardToSplit.Add(rewardToSplit, blockReward.Mul(blockReward, big.NewInt(int64(blockRewardPerc))).Div(blockReward, big.NewInt(100)))

		boundationBonus, _ := new(big.Int).SetString(reward.FoundationBonus, 10)
		rewardToSplit.Add(rewardToSplit, boundationBonus.Mul(boundationBonus, big.NewInt(int64(foundationBonusPerc))).Div(boundationBonus, big.NewInt(100)))

		distrRewardMap[reward.EpochNumber] = rewardToSplit.Mul(rewardToSplit, big.NewInt(int64(epochRewardPerc))).Div(rewardToSplit, big.NewInt(100))
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

func GetBookkeeping(ctx context.Context, startEpoch uint64, epochCount uint64, delegateName string, epochRewardPerc, blockRewardPerc, foundationBonusPerc uint64) (map[string]*big.Int, error) {
	endEpoch := startEpoch + epochCount - 1

	distrRewardMap, err := rewardsToSplit(startEpoch, endEpoch, delegateName, epochRewardPerc, blockRewardPerc, foundationBonusPerc)
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

func WeightedVotesBySearchPairsFix(delegateMap map[uint64][]string) (map[string]map[uint64]map[string]*big.Int, error) {
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
	f := func(ctx context.Context, epochNum uint64) ([]AggregateVoting, error) {
		var votes []AggregateVoting
		fileName := fmt.Sprintf("epoch_fix_%d.csv", epochNum)
		file, err := os.Open(fileName)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		csvReader := csv.NewReader(file)
		csvReader.FieldsPerRecord = -1
		votesArr, err := csvReader.ReadAll()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, row := range votesArr {
			votes = append(votes, AggregateVoting{
				EpochNumber:    epochNum,
				CandidateName:  row[0],
				VoterAddress:   row[1],
				AggregateVotes: row[2],
			})
		}
		file.Close()
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
		voters := value.([]AggregateVoting)
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
			if _, ok := epochVoterMap[epoch]; !ok {
				epochVoterMap[epoch] = make(map[string]*big.Int)
			}
			voterMap := epochVoterMap[epoch]

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

func WeightedVotesBySearchPairs(delegateMap map[uint64][]string) (map[string]map[uint64]map[string]*big.Int, error) {
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
	f := func(ctx context.Context, epochNum uint64) ([]AggregateVoting, error) {
		var votes []AggregateVoting
		fileName := fmt.Sprintf("epoch_%d.csv", epochNum)
		file, err := os.Open(fileName)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		csvReader := csv.NewReader(file)
		csvReader.FieldsPerRecord = -1
		votesArr, err := csvReader.ReadAll()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, row := range votesArr {
			votes = append(votes, AggregateVoting{
				EpochNumber:    epochNum,
				CandidateName:  row[0],
				VoterAddress:   row[1],
				AggregateVotes: row[2],
			})
		}
		file.Close()
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
		voters := value.([]AggregateVoting)
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
			if _, ok := epochVoterMap[epoch]; !ok {
				epochVoterMap[epoch] = make(map[string]*big.Int)
			}
			voterMap := epochVoterMap[epoch]

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

type HermesVoteInfo struct {
	Recipient  string
	StartEpoch uint64
	EndEpoch   uint64
	Amount     string
	ActionHash string
	Timestamp  time.Time
}

type HermesDelegateInfo struct {
	DelegateName string
	StartEpoch   uint64
	EndEpoch     uint64
	Amount       string
	ActionHash   string
	Timestamp    time.Time
}

func GetTotalHermesByDelegate(ctx context.Context, startEpoch uint64, endEpoch uint64, delegateName string) (int64, string, error) {
	db := db.DB()
	var result struct {
		Count int64
		Sum   string
	}
	query := "select count(1) as count,sum(amount) as sum from (select * from block_receipt_transactions where  sender in ('io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu','io16y9wk2xnwurvtgmd2mds2gcdfe2lmzad6dcw29')) as t1 inner join (select t5.*,t6.timestamp from hermes_distributes as t5 left join block t6 on t6.block_height=t5.block_height where epoch_number>=? and epoch_number<=? ) as t2 on t1.action_hash=t2.action_hash where delegate_name=?"
	if err := db.Raw(query, startEpoch, endEpoch, delegateName).Scan(&result).Error; err != nil {
		return 0, "", err
	}
	return result.Count, result.Sum, nil
}

func GetHermesByDelegate(ctx context.Context, startEpoch uint64, endEpoch uint64, delegateName string, skip, first uint64) ([]*HermesVoteInfo, error) {
	db := db.DB()
	var results []*HermesVoteInfo
	query := "select recipient,start_epoch,end_epoch,amount,t1.action_hash,t2.timestamp from (select * from block_receipt_transactions where  sender in ('io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu','io16y9wk2xnwurvtgmd2mds2gcdfe2lmzad6dcw29')) as t1 inner join (select t5.*,t6.timestamp from hermes_distributes as t5 left join block t6 on t6.block_height=t5.block_height where epoch_number>=? and epoch_number<=? ) as t2 on t1.action_hash=t2.action_hash where delegate_name=? limit ? offset ?"
	if err := db.Raw(query, startEpoch, endEpoch, delegateName, first, skip).Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func GetTotalHermesByVoter(ctx context.Context, startEpoch uint64, endEpoch uint64, voterAddress string) (int64, string, error) {
	db := db.DB()
	var result struct {
		Count int64
		Sum   string
	}
	query := "select count(1) as count,sum(amount) as sum from (select * from block_receipt_transactions where  sender in ('io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu','io16y9wk2xnwurvtgmd2mds2gcdfe2lmzad6dcw29')) as t1 inner join (select t5.*,t6.timestamp from hermes_distributes as t5 left join block t6 on t6.block_height=t5.block_height where epoch_number>=? and epoch_number<=? ) as t2 on t1.action_hash=t2.action_hash where recipient=?"
	if err := db.Raw(query, startEpoch, endEpoch, voterAddress).Scan(&result).Error; err != nil {
		return 0, "", err
	}
	return result.Count, result.Sum, nil
}

func GetHermesByVoter(ctx context.Context, startEpoch uint64, endEpoch uint64, voterAddress string, skip, first uint64) ([]*HermesDelegateInfo, error) {
	db := db.DB()
	var results []*HermesDelegateInfo
	query := "select delegate_name,start_epoch,end_epoch,amount,t1.action_hash,t2.timestamp from (select * from block_receipt_transactions where  sender in ('io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu','io16y9wk2xnwurvtgmd2mds2gcdfe2lmzad6dcw29')) as t1 inner join (select t5.*,t6.timestamp from hermes_distributes as t5 left join block t6 on t6.block_height=t5.block_height where epoch_number>=? and epoch_number<=? ) as t2 on t1.action_hash=t2.action_hash where recipient=? order by t1.id desc limit ? offset ?"
	if err := db.Raw(query, startEpoch, endEpoch, voterAddress, first, skip).Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

type HermesRatio struct {
	EpochNumber               uint64
	BlockRewardPercentage     float64
	EpochRewardPercentage     float64
	FoundationBonusPercentage float64
}

func GetHermesRatioByDelegate(ctx context.Context, startEpoch uint64, endEpoch uint64, delegateName string) ([]*HermesRatio, error) {
	db := db.DB()
	var results []*HermesRatio
	query := "select epoch_number,block_reward_percentage,epoch_reward_percentage,foundation_bonus_percentage from hermes_voting_results where epoch_number>=? and epoch_number<=? and delegate_name=?"
	if err := db.Raw(query, startEpoch, endEpoch, delegateName).Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

type HermesMeta struct {
	Count                   uint64
	NumberOfDelegates       uint64
	NumberOfRecipients      uint64
	TotalRewardsDistributed string
}

func GetHermesMeta(ctx context.Context, startEpoch uint64, endEpoch uint64) (*HermesMeta, error) {
	db := db.DB()
	var result *HermesMeta
	query := "SELECT count(1) as count, COUNT(DISTINCT delegate_name) as number_of_delegates, COUNT(DISTINCT recipient) as number_of_recipients, sum(amount) as total_rewards_distributed from (select * from block_receipt_transactions where  sender in ('io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu','io16y9wk2xnwurvtgmd2mds2gcdfe2lmzad6dcw29')) as t1 inner join (select t5.*,t6.timestamp from hermes_distributes as t5 left join block t6 on t6.block_height=t5.block_height where epoch_number>=? and epoch_number<=? ) as t2 on t1.action_hash=t2.action_hash"
	if err := db.WithContext(ctx).Raw(query, startEpoch, endEpoch).Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// DelegateHermesAverage defines the Hermes average stats for each delegate
type DelegateHermesAverage struct {
	DelegateName       string
	RewardDistribution string
	TotalWeightedVotes string
}
