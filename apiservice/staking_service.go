package apiservice

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common/actions"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-analyser-api/internal/sync/errgroup"
	"github.com/iotexproject/iotex-analyser-api/model"
	"github.com/iotexproject/iotex-core/v2/blockchain/genesis"
	"github.com/iotexproject/iotex-core/v2/ioctl/util"
)

type StakingService struct {
	api.UnimplementedStakingServiceServer
}

// curl -d '{"address": ["io10avlgwgxv2k22dup4q0ah998vklg4rcrgl04m8", "io1fuhhg9jgdxwpms9dsdfwjdc90nt7v67hx40cd8"], "height":11900487 }' http://127.0.0.1:8889/api.StakingService.VoteByHeight
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
		stakeAmounts, voteWeights, err := actions.GetStakedBucketByVoterAndHeight(addr, height)
		if err != nil {
			return nil, err
		}
		systemStakeAmounts, systemVoteWeights, err := actions.GetSystemStakedBucketByVoterAndHeight(addr, height)
		if err != nil {
			return nil, err
		}
		systemV2StakeAmounts, systemV2VoteWeights, err := actions.GetSystemV2StakedBucketByVoterAndHeight(addr, height)
		if err != nil {
			return nil, err
		}
		stakeAmounts = stakeAmounts.Add(stakeAmounts, systemStakeAmounts)
		voteWeights = voteWeights.Add(voteWeights, systemVoteWeights)
		stakeAmounts = stakeAmounts.Add(stakeAmounts, systemV2StakeAmounts)
		voteWeights = voteWeights.Add(voteWeights, systemV2VoteWeights)
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

func (s *StakingService) BucketByID(ctx context.Context, req *api.BucketByIDRequest) (*api.BucketByIDResponse, error) {
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
	bucketIDs := req.GetBucketId()
	resp := &api.BucketByIDResponse{Height: height}

	nativeBuckets, err := actions.GetNativeBucketsByIDsAndHeight(bucketIDs, height)
	if err != nil {
		return nil, err
	}
	for _, b := range nativeBuckets {
		resp.NativeBuckets = append(resp.NativeBuckets, nativeBucketToStakingInfo(b))
	}

	if req.GetIncludeSystem() {
		g := new(errgroup.Group)
		var systemBuckets, systemV2Buckets []*model.SystemStakingBucket
		var systemV3Buckets []*model.SystemStakingBucketV3
		g.Go(func(ctx context.Context) error {
			var err error
			systemBuckets, err = actions.GetSystemBucketsByIDsAndHeight(bucketIDs, height)
			return err
		})
		g.Go(func(ctx context.Context) error {
			var err error
			systemV2Buckets, err = actions.GetSystemV2BucketsByIDsAndHeight(bucketIDs, height)
			return err
		})
		g.Go(func(ctx context.Context) error {
			var err error
			systemV3Buckets, err = actions.GetSystemV3BucketsByIDsAndHeight(bucketIDs, height)
			return err
		})
		if err := g.Wait(); err != nil {
			return nil, err
		}
		for _, b := range systemBuckets {
			resp.SystemBuckets = append(resp.SystemBuckets, systemBucketToStakingInfo(b))
		}
		for _, b := range systemV2Buckets {
			resp.SystemV2Buckets = append(resp.SystemV2Buckets, systemBucketToStakingInfo(b))
		}
		for _, b := range systemV3Buckets {
			resp.SystemV3Buckets = append(resp.SystemV3Buckets, systemV3BucketToStakingInfo(b))
		}
	}
	return resp, nil
}

func nativeBucketToStakingInfo(b *model.StakingBucket) *api.StakingBucketInfo {
	return &api.StakingBucketInfo{
		BucketId:         b.BucketID,
		OwnerAddress:     b.OwnerAddress,
		Candidate:        b.Candidate,
		StakedAmount:     b.StakedAmount,
		VotingPower:      b.VotingPower,
		Duration:         b.Duration * 86400,
		AutoStake:        b.AutoStake,
		CreateTime:       uint32(b.CreateTime),
		StakeStartTime:   uint32(b.StakeStartTime),
		UnstakeStartTime: uint32(b.UnstakeStartTime),
		BlockHeight:      b.BlockHeight,
	}
}

func systemBucketToStakingInfo(b *model.SystemStakingBucket) *api.StakingBucketInfo {
	return &api.StakingBucketInfo{
		BucketId:         b.BucketID,
		OwnerAddress:     b.OwnerAddress,
		Candidate:        b.Candidate,
		StakedAmount:     b.StakedAmount,
		VotingPower:      b.VotingPower,
		Duration:         b.Duration * 86400, // stored in days
		AutoStake:        b.AutoStake,
		CreateTime:       uint32(b.CreateTime),
		StakeStartTime:   uint32(b.StakeStartTime),
		UnstakeStartTime: uint32(b.UnstakeStartTime),
		BlockHeight:      b.BlockHeight,
	}
}

func systemV3BucketToStakingInfo(b *model.SystemStakingBucketV3) *api.StakingBucketInfo {
	return &api.StakingBucketInfo{
		BucketId:         b.BucketID,
		OwnerAddress:     b.OwnerAddress,
		Candidate:        b.Candidate,
		StakedAmount:     b.StakedAmount,
		VotingPower:      b.VotingPower,
		Duration:         b.Duration, // stored in seconds
		AutoStake:        b.AutoStake,
		CreateTime:       uint32(b.CreateTime),
		StakeStartTime:   uint32(b.StakeStartTime),
		UnstakeStartTime: uint32(b.UnstakeStartTime),
		BlockHeight:      b.BlockHeight,
	}
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

func getCandidateStaking(height uint64, addr string) ([]*Staking, error) {
	db := db.DB()
	query := `WITH max_ids AS (
		SELECT MAX(id) AS max_id
		FROM staking_buckets
		WHERE block_height <= ?
		GROUP BY bucket_id
	)
	SELECT id,block_height,bucket_id,owner_address,candidate,staked_amount as amount,act_type,auto_stake,duration
	FROM staking_buckets t1
	RIGHT JOIN max_ids t2 ON  t1.id=t2.max_id where candidate=? order by bucket_id`
	rows, err := db.Raw(query, height, addr).Rows()
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
