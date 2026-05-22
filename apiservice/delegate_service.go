package apiservice

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"sort"
	"time"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/rewards"
	"github.com/iotexproject/iotex-analyser-api/common/votings"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			bucketInfoList.BucketInfo = make([]*api.BucketInfo, 0)
			bucketInfoLists = append(bucketInfoLists, bucketInfoList)
			continue
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
	blockRewardPerc := req.GetBlockRewardPerc()
	foundationBonusPerc := req.GetFoundationBonusPerc()
	epochRewardPerc := req.GetEpochRewardPerc()

	if epochRewardPerc > 100 || blockRewardPerc > 100 || foundationBonusPerc > 100 {
		return nil, errors.New("percentage should be 0-100")
	}
	rewards, err := rewards.GetBookkeeping(ctx, startEpoch, epochCount, delegateName, epochRewardPerc, blockRewardPerc, foundationBonusPerc)
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
	if resp.Count == 0 {
		return resp, nil
	}
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

// Staking returns the staking info of the delegate
func (s *DelegateService) Staking(ctx context.Context, req *api.StakingRequest) (*api.StakingResponse, error) {
	resp := &api.StakingResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	endEpoch := startEpoch + epochCount - 1
	db := db.DB()
	query := "SELECT epoch_number,total_weighted_votes,self_staking FROM hermes_voting_results WHERE epoch_number >= ? AND epoch_number <= ? AND delegate_name = ?"
	var result []struct {
		EpochNumber        uint64
		TotalWeightedVotes string
		SelfStaking        string
	}
	if err := db.Raw(query, startEpoch, endEpoch, delegateName).Scan(&result).Error; err != nil {
		return nil, err
	}
	resp.Exist = len(result) > 0
	resp.StakingInfo = make([]*api.StakingResponse_StakingInfo, 0)
	for _, v := range result {
		resp.StakingInfo = append(resp.StakingInfo, &api.StakingResponse_StakingInfo{
			EpochNumber:  v.EpochNumber,
			TotalStaking: v.TotalWeightedVotes,
			SelfStaking:  v.SelfStaking,
		})
	}
	return resp, nil
}

// ProbationHistoricalRate returns the probation historical rate of the delegate
func (s *DelegateService) ProbationHistoricalRate(ctx context.Context, req *api.ProbationHistoricalRateRequest) (*api.ProbationHistoricalRateResponse, error) {
	resp := &api.ProbationHistoricalRateResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	endEpoch := startEpoch + epochCount
	db := db.DB()
	query := "SELECT count(epoch_number) FROM hermes_voting_results WHERE epoch_number >= ? AND epoch_number < ? AND delegate_name = ?"
	var count int
	if err := db.WithContext(ctx).Raw(query, startEpoch, endEpoch, delegateName).Scan(&count).Error; err != nil {
		return nil, err
	}
	probationExist := func(epochNumber uint64, address string) bool {
		query := "SELECT count(*) FROM probation WHERE epoch_number = ? AND address = ?"
		var count int
		if err := db.WithContext(ctx).Raw(query, epochNumber, address).Scan(&count).Error; err != nil {
			return false
		}
		return count > 0
	}
	probationCount := uint64(0)
	for i := startEpoch; i < startEpoch+epochCount; i++ {
		address, err := votings.GetOperatorAddress(delegateName, i)
		if err != nil {
			return nil, err
		}
		exist := probationExist(i, address)
		if exist {
			probationCount++
		}
	}
	rate := float64(probationCount) / float64(count)
	log.Printf("probationCount: %d, count: %d, rate: %f", probationCount, count, rate)
	resp.ProbationHistoricalRate = fmt.Sprintf("%.2f", rate)
	return resp, nil
}

func (s *DelegateService) PaidToDelegates(ctx context.Context, req *api.PaidToDelegatesRequest) (*api.PaidToDelegatesResponse, error) {
	resp := &api.PaidToDelegatesResponse{}
	//schedule := req.GetSchedule()
	date := req.GetDate() // 2019-01-01
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	minH, maxH, err := common.GetBlockHeightRangeByMonth(dateTime)
	if err != nil {
		return nil, err
	}
	db := db.DB()
	minEpochNum := common.GetEpochNum(minH)
	maxEpochNum := common.GetEpochNum(maxH)
	query := "SELECT candidate_name, SUM(block_reward) block_reward, SUM(epoch_reward) epoch_reward, SUM(foundation_bonus) foundation_bonus FROM hermes_account_rewards WHERE epoch_number >= ?  AND epoch_number <= ? GROUP BY candidate_name"
	var result []struct {
		CandidateName   string
		BlockReward     string
		EpochReward     string
		FoundationBonus string
	}

	if err := db.Raw(query, minEpochNum, maxEpochNum).Scan(&result).Error; err != nil {
		return nil, err
	}
	resp.DelegateInfo = make([]*api.PaidToDelegatesResponse_DelegateInfo, 0)
	for _, v := range result {
		amount := new(big.Int)
		res, err := stringToBigInt(v.BlockReward)
		if err != nil {
			return nil, err
		}
		amount.Add(amount, res)
		res, err = stringToBigInt(v.EpochReward)
		if err != nil {
			return nil, err
		}
		amount.Add(amount, res)
		res, err = stringToBigInt(v.FoundationBonus)
		if err != nil {
			return nil, err
		}

		resp.DelegateInfo = append(resp.DelegateInfo, &api.PaidToDelegatesResponse_DelegateInfo{
			DelegateName:    v.CandidateName,
			BlockReward:     v.BlockReward,
			EpochReward:     v.EpochReward,
			FoundationBonus: v.FoundationBonus,
			Amount:          amount.String(),
		})
	}
	return resp, nil
}

// ─── Added for iotex-kit modules-db migration ───

// GetDelegatesByHeight returns the delegate_record snapshot at the given block
// height (or the latest snapshot when height = 0). Replaces kit
// staking.delegatesByHeight.
func (s *DelegateService) GetDelegatesByHeight(ctx context.Context, req *api.GetDelegatesByHeightRequest) (*api.GetDelegatesByHeightResponse, error) {
	resp := &api.GetDelegatesByHeightResponse{}
	gormDB := db.DB().WithContext(ctx)
	height := req.GetHeight()
	if height == 0 {
		var maxHeight sql.NullInt64
		if err := gormDB.Raw(`SELECT MAX(block_height) FROM delegate_record`).Scan(&maxHeight).Error; err != nil {
			// delegate_record is populated by an external job in production and
			// may not exist in local-dev — treat as empty rather than 500.
			if isUndefinedTableErr(err) {
				return resp, nil
			}
			return nil, errors.Wrap(err, "failed to read max delegate_record height")
		}
		if !maxHeight.Valid {
			return resp, nil
		}
		height = uint64(maxHeight.Int64)
	}
	resp.Height = height

	var rows []struct {
		BlockHeight        uint64
		CandidateAddress   sql.NullString
		CandidateName      sql.NullString
		OperatorAddress    sql.NullString
		RewardAddress      sql.NullString
		SelfStakingTokens  sql.NullString
		TotalWeightedVotes sql.NullString
		StakeAmount        sql.NullString
		Active             sql.NullBool
	}
	if err := gormDB.Raw(
		`SELECT block_height, candidate_address, candidate_name, operator_address,
			reward_address, self_staking_tokens::text, total_weighted_votes::text,
			stake_amount::text, active
		FROM delegate_record WHERE block_height = ?`,
		height,
	).Scan(&rows).Error; err != nil {
		if isUndefinedTableErr(err) {
			return resp, nil
		}
		return nil, errors.Wrap(err, "failed to query delegate_record")
	}
	for _, r := range rows {
		d := &api.DelegateRecord{BlockHeight: r.BlockHeight}
		if r.CandidateAddress.Valid {
			d.CandidateAddress = r.CandidateAddress.String
		}
		if r.CandidateName.Valid {
			d.CandidateName = r.CandidateName.String
		}
		if r.OperatorAddress.Valid {
			d.OperatorAddress = r.OperatorAddress.String
		}
		if r.RewardAddress.Valid {
			d.RewardAddress = r.RewardAddress.String
		}
		if r.SelfStakingTokens.Valid {
			d.SelfStakingTokens = r.SelfStakingTokens.String
		}
		if r.TotalWeightedVotes.Valid {
			d.TotalWeightedVotes = r.TotalWeightedVotes.String
		}
		if r.StakeAmount.Valid {
			d.StakeAmount = r.StakeAmount.String
		}
		if r.Active.Valid {
			d.Active = r.Active.Bool
		}
		resp.Delegates = append(resp.Delegates, d)
	}
	return resp, nil
}

// GetBlocksByProducer returns blocks produced by an address with pagination
// and a total count. Replaces kit staking.delegate_recent_blocks.
func (s *DelegateService) GetBlocksByProducer(ctx context.Context, req *api.GetBlocksByProducerRequest) (*api.GetBlocksByProducerResponse, error) {
	resp := &api.GetBlocksByProducerResponse{}
	producer := req.GetProducerAddress()
	if producer == "" {
		return nil, status.Errorf(codes.InvalidArgument, "producer_address is required")
	}
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())

	var count uint64
	if err := db.DB().WithContext(ctx).Raw(
		`SELECT COUNT(1) FROM block WHERE producer_address = ?`, producer,
	).Scan(&count).Error; err != nil {
		return nil, errors.Wrap(err, "failed to count blocks for producer")
	}
	resp.Count = count

	type row struct {
		BlockHeight     uint64
		BlockHash       string
		ProducerAddress string
		NumActions      uint64
		Timestamp       int64
		GasConsumed     sql.NullInt64
		ProducerName    sql.NullString
		BlockReward     sql.NullString
	}
	var rows []row
	if err := db.DB().WithContext(ctx).Raw(
		`SELECT b.block_height, b.block_hash, b.producer_address, b.num_actions,
			(EXTRACT(EPOCH FROM (b.timestamp AT TIME ZONE 'UTC')) * 1000)::bigint AS timestamp,
			m.gas_consumed, m.producer_name, m.block_reward
		FROM block b
		LEFT JOIN block_meta m ON m.block_height = b.block_height
		WHERE b.producer_address = ?
		ORDER BY b.block_height DESC
		LIMIT ? OFFSET ?`,
		producer, first, skip,
	).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to list blocks for producer")
	}
	for _, r := range rows {
		pb := &api.ProducerBlock{
			BlockHeight:     r.BlockHeight,
			BlockHash:       r.BlockHash,
			ProducerAddress: r.ProducerAddress,
			NumActions:      r.NumActions,
			Timestamp:       r.Timestamp,
		}
		if r.GasConsumed.Valid {
			pb.GasConsumed = uint64(r.GasConsumed.Int64)
		}
		if r.ProducerName.Valid {
			pb.ProducerName = r.ProducerName.String
		}
		if r.BlockReward.Valid {
			pb.BlockReward = r.BlockReward.String
		}
		resp.Blocks = append(resp.Blocks, pb)
	}
	return resp, nil
}
