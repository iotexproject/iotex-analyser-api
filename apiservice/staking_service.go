package apiservice

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-analyser-api/internal/sync/errgroup"
	"github.com/iotexproject/iotex-core/blockchain/genesis"
	"github.com/iotexproject/iotex-core/ioctl/util"
)

type StakingService struct {
	api.UnimplementedStakingServiceServer
}

//curl -d '{"address": ["io10avlgwgxv2k22dup4q0ah998vklg4rcrgl04m8", "io1fuhhg9jgdxwpms9dsdfwjdc90nt7v67hx40cd8"], "height":11900487 }' http://127.0.0.1:8889/api.StakingService.VoteByHeight
func (s *StakingService) VoteByHeight(ctx context.Context, req *api.VoteByHeightRequest) (*api.VoteByHeightResponse, error) {
	resp := &api.VoteByHeightResponse{
		Height: req.GetHeight(),
	}
	height := req.GetHeight()
	for _, addr := range req.GetAddress() {
		if addr[:2] == "0x" || addr[:2] == "0X" {
			add, err := address.FromHex(addr)
			if err != nil {
				return nil, err
			}

			addr = add.String()
		}
		bucketIDs, err := getBucketIDsByAddressWithHeight(addr, height)
		if err != nil {
			return nil, err
		}
		stakeAmounts := big.NewInt(0)
		voteWeights := big.NewInt(0)
		for _, bucketID := range bucketIDs {
			stakeAmount, err := getSumStake(addr, height, bucketID)
			if err != nil {
				return nil, err
			}
			stakeAmounts = stakeAmounts.Add(stakeAmounts, stakeAmount)
			duration, autoStake, selfAutoStake := getVoteBucketParams(addr, height, bucketID)
			voteBucket := &VoteBucket{
				StakedAmount:   stakeAmount,
				AutoStake:      autoStake,
				StakedDuration: duration,
			}
			voteWeight := calculateVoteWeight(config.Default.Genesis.VoteWeightCalConsts, voteBucket, selfAutoStake)
			voteWeights = voteWeights.Add(voteWeights, voteWeight)
		}
		resp.StakeAmount = append(resp.StakeAmount, util.RauToString(stakeAmounts, util.IotxDecimalNum))
		resp.VoteWeight = append(resp.VoteWeight, util.RauToString(voteWeights, util.IotxDecimalNum))
	}
	return resp, nil
}

func (s *StakingService) CandidateVoteByHeight(ctx context.Context, req *api.CandidateVoteByHeightRequest) (*api.CandidateVoteByHeightResponse, error) {
	pluginHeight, err := db.GetIndexHeight("staking_actions")
	if err != nil {
		return nil, err
	}
	height := req.GetHeight()
	if height == 0 {
		height = pluginHeight
	} else if height > pluginHeight {
		return nil, fmt.Errorf("request height greater than plugin height, %d > %d", height, pluginHeight)
	}
	resp := &api.CandidateVoteByHeightResponse{
		Height: height,
	}
	g := new(errgroup.Group)
	for _, addr := range req.GetAddress() {
		if addr[:2] == "0x" || addr[:2] == "0X" {
			add, err := address.FromHex(addr)
			if err != nil {
				return nil, err
			}

			addr = add.String()
		}
		addr := addr
		g.Go(func(ctx context.Context) error {
			stakings, err := getCandidateStaking(height, addr)
			if err != nil {
				return err
			}
			stakeAmounts := big.NewInt(0)
			voteWeights := big.NewInt(0)
			for _, staking := range stakings {
				stakeAmount, _ := big.NewInt(0).SetString(staking.Amount, 0)
				stakeAmounts = stakeAmounts.Add(stakeAmounts, stakeAmount)
				voteBucket := &VoteBucket{
					StakedAmount:   stakeAmount,
					AutoStake:      staking.AutoStake,
					StakedDuration: staking.Duration,
				}
				selfAutoStake := false
				if staking.OwnerAddress == addr {
					selfAutoStake = true
				}
				voteWeight := calculateVoteWeight(config.Default.Genesis.VoteWeightCalConsts, voteBucket, selfAutoStake)
				voteWeights = voteWeights.Add(voteWeights, voteWeight)
			}
			resp.StakeAmount = append(resp.StakeAmount, util.RauToString(stakeAmounts, util.IotxDecimalNum))
			resp.VoteWeight = append(resp.VoteWeight, util.RauToString(voteWeights, util.IotxDecimalNum))
			resp.Address = append(resp.Address, addr)
			return nil
		})

	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return resp, nil
}

func getBucketIDsByAddressWithHeight(addr string, height uint64) ([]uint64, error) {
	db := db.DB()
	var ids []struct {
		BucketID uint64
	}
	if err := db.Table("staking_actions").Distinct("bucket_id").Where("block_height<=? and owner_address=?", height, addr).Find(&ids).Error; err != nil {
		return nil, err
	}
	bucketID := []uint64{}
	for _, id := range ids {
		bucketOwner, _ := getBucketOwnerWithHeight(id.BucketID, height)
		if addr != bucketOwner {
			continue
		}
		bucketID = append(bucketID, id.BucketID)
	}
	return bucketID, nil
}

func getBucketOwnerWithHeight(bucketID, height uint64) (string, error) {
	var addr sql.NullString
	db := db.DB()
	if err := db.Table("staking_actions").Select("owner_address").Where("block_height<=? and bucket_id=?", height, bucketID).Order("id desc").Limit(1).Scan(&addr).Error; err != nil {
		return "", err
	}
	return addr.String, nil
}

type VoteBucket struct {
	Index            uint64
	Candidate        string
	Owner            string
	StakedAmount     *big.Int
	StakedDuration   uint32
	CreateTime       time.Time
	StakeStartTime   time.Time
	UnstakeStartTime time.Time
	AutoStake        bool
}

func calculateVoteWeight(c genesis.VoteWeightCalConsts, v *VoteBucket, selfStake bool) *big.Int {
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

	amount := new(big.Float).SetInt(v.StakedAmount)
	weightedAmount, _ := amount.Mul(amount, big.NewFloat(weight)).Int(nil)
	return weightedAmount
}

type Staking struct {
	ID           uint64
	BlockHeight  uint64
	BucketID     uint64
	OwnerAddress string
	Candidate    string
	Amount       string
	ActType      string
	AutoStake    bool
	Duration     uint32
}

func getSumStake(addr string, height, bucketID uint64) (*big.Int, error) {
	db := db.DB()
	var amount sql.NullString
	if err := db.Table("staking_actions").Select("sum(amount)").Where("block_height<=? and bucket_id=? and owner_address=?", height, bucketID, addr).Scan(&amount).Error; err != nil {
		return nil, err
	}
	stakeAmount, _ := big.NewInt(0).SetString(amount.String, 0)
	return stakeAmount, nil
}

func getVoteBucketParams(addr string, height, bucketID uint64) (uint32, bool, bool) {
	var av Staking
	db := db.DB()
	if err := db.Table("staking_actions").Where("block_height<=? and bucket_id=? and owner_address=?", height, bucketID, addr).Order("id desc").Scan(&av).Error; err != nil {
		return 0, false, false
	}

	selfAutoStake := false
	if addr == av.Candidate {
		selfAutoStake = true
	}
	return av.Duration, av.AutoStake, selfAutoStake
}

func getCandidateStaking(height uint64, addr string) ([]*Staking, error) {
	db := db.DB()
	//TODO: fix the query, it's too slow
	query := "select id,block_height,bucket_id,owner_address,candidate,(select sum(b.amount) from staking_actions b where b.block_height<=? and b.bucket_id=a.bucket_id) as amount,act_type,auto_stake,duration from staking_actions a where id=any(array(select max(id) from staking_actions where block_height<=? and bucket_id=any(array(select distinct bucket_id from staking_actions where block_height<=? and candidate=?)) group by bucket_id))  and candidate=?"
	rows, err := db.Raw(query, height, height, height, addr, addr).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*Staking
	for rows.Next() {
		av := new(Staking)

		if err := db.ScanRows(rows, av); err != nil {
			return nil, err
		}
		results = append(results, av)
	}
	return results, nil
}
