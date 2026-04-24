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
	"github.com/iotexproject/iotex-analyser-api/common/actions"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-analyser-api/internal/sync/errgroup"
	"github.com/iotexproject/iotex-analyser-api/model"
	"github.com/iotexproject/iotex-core/v2/blockchain/genesis"
	"github.com/iotexproject/iotex-core/v2/ioctl/util"
	"github.com/pkg/errors"
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

// ─── Bucket list / detail helpers ───────────────────────────────────────────

var allowedBucketSortFields = map[string]bool{
	"timestamp": true, "amount": true, "create_time": true,
	"stake_start_time": true, "unstake_start_time": true,
	"staked_amount": true, "bucket_id": true,
}

func parseBucketSortParam(sort string) (field, order string) {
	field, order = "timestamp", "DESC"
	if sort == "" {
		return
	}
	parts := [2]string{"timestamp", "desc"}
	if i := len(sort) - 1; i > 0 {
		for idx, c := range sort {
			if c == ':' {
				parts[0] = sort[:idx]
				if idx+1 < len(sort) {
					parts[1] = sort[idx+1:]
				}
				break
			}
		}
	}
	if allowedBucketSortFields[parts[0]] {
		field = "staking_buckets." + parts[0]
	}
	if parts[1] == "asc" {
		order = "ASC"
	}
	return
}

func intervalToUnixSeconds(interval string) int64 {
	now := time.Now()
	switch interval {
	case "1D":
		return now.Add(-24 * time.Hour).Unix()
	case "7D":
		return now.Add(-7 * 24 * time.Hour).Unix()
	case "30D", "1M":
		return now.Add(-30 * 24 * time.Hour).Unix()
	case "1Y":
		return now.Add(-365 * 24 * time.Hour).Unix()
	}
	return 0 // "ALL" or unknown → no filter
}

type bucketExRow struct {
	BucketID          int64
	ActionHash        string
	Timestamp         sql.NullString
	CreateTime        sql.NullString
	StakeStartTime    sql.NullString
	UnstakeStartTime  sql.NullString
	Amount            sql.NullString
	StakedAmount      sql.NullString
	ActType           sql.NullString
	Sender            string
	OwnerAddress      sql.NullString
	Candidate         sql.NullString
	AutoStake         sql.NullBool
	Duration          sql.NullString
	GasPrice          sql.NullString
	GasLimit          sql.NullString
	Recipient         sql.NullString
	DelegateName      sql.NullString
}

func toBucketInfoEx(r bucketExRow) *api.BucketInfoEx {
	b := &api.BucketInfoEx{
		BucketId:   r.BucketID,
		ActionHash: r.ActionHash,
		Sender:     r.Sender,
	}
	if r.Timestamp.Valid {
		b.Timestamp = r.Timestamp.String
	}
	if r.CreateTime.Valid {
		b.CreateTime = r.CreateTime.String
	}
	if r.StakeStartTime.Valid {
		b.StakeStartTime = r.StakeStartTime.String
	}
	if r.UnstakeStartTime.Valid {
		b.UnstakeStartTime = r.UnstakeStartTime.String
	}
	if r.Amount.Valid {
		b.Amount = r.Amount.String
	}
	if r.StakedAmount.Valid {
		b.StakedAmount = r.StakedAmount.String
	}
	if r.ActType.Valid {
		b.ActType = r.ActType.String
	}
	if r.OwnerAddress.Valid {
		b.OwnerAddress = r.OwnerAddress.String
	}
	if r.Candidate.Valid {
		b.Candidate = r.Candidate.String
	}
	if r.AutoStake.Valid {
		b.AutoStake = r.AutoStake.Bool
	}
	if r.Duration.Valid {
		b.Duration = r.Duration.String
	}
	if r.GasPrice.Valid {
		b.GasPrice = r.GasPrice.String
	}
	if r.GasLimit.Valid {
		b.GasLimit = r.GasLimit.String
	}
	if r.Recipient.Valid {
		b.Recipient = r.Recipient.String
	}
	if r.DelegateName.Valid {
		b.DelegateName = r.DelegateName.String
	}
	return b
}

const tsExpr = `to_char(to_timestamp(staking_buckets.timestamp), 'YYYY-MM-DD"T"HH24:MI:SS"Z"')`
const createTimeExpr = `to_char(to_timestamp(staking_buckets.create_time), 'YYYY-MM-DD"T"HH24:MI:SS"Z"')`
const stakeStartExpr = `to_char(to_timestamp(staking_buckets.stake_start_time), 'YYYY-MM-DD"T"HH24:MI:SS"Z"')`
const unstakeStartExpr = `to_char(to_timestamp(staking_buckets.unstake_start_time), 'YYYY-MM-DD"T"HH24:MI:SS"Z"')`

func nativeBucketListQuery(startTime int64, field, order string, limit, offset int64) (string, []interface{}) {
	base := `SELECT staking_buckets.bucket_id, staking_buckets.action_hash,
		` + tsExpr + ` AS timestamp,
		` + createTimeExpr + ` AS create_time,
		` + stakeStartExpr + ` AS stake_start_time,
		` + unstakeStartExpr + ` AS unstake_start_time,
		staking_buckets.amount, staking_buckets.staked_amount, staking_buckets.act_type,
		staking_buckets.sender, staking_buckets.owner_address, staking_buckets.candidate,
		staking_buckets.auto_stake, staking_buckets.duration::text AS duration,
		NULL AS gas_price, NULL AS gas_limit, NULL AS recipient,
		delegate.name AS delegate_name
	FROM staking_buckets
	LEFT JOIN delegate ON delegate.candidate = staking_buckets.candidate`
	var args []interface{}
	if startTime > 0 {
		base += " WHERE staking_buckets.timestamp > ?"
		args = append(args, startTime)
	}
	base += fmt.Sprintf(" ORDER BY %s %s LIMIT ? OFFSET ?", field, order)
	args = append(args, limit, offset)
	return base, args
}

func nftBucketListQuery(table string, requireFinal bool, startTime int64, field, order string, limit, offset int64) (string, []interface{}) {
	base := fmt.Sprintf(`SELECT staking_buckets.bucket_id, staking_buckets.act_hash AS action_hash,
		`+tsExpr+` AS timestamp,
		`+createTimeExpr+` AS create_time,
		NULL AS stake_start_time, NULL AS unstake_start_time,
		staking_buckets.amount, staking_buckets.staked_amount, staking_buckets.event_type AS act_type,
		staking_buckets.sender, staking_buckets.owner_address, d.candidate AS candidate,
		staking_buckets.auto_stake, (staking_buckets.duration / 86400.0)::text AS duration,
		NULL AS gas_price, NULL AS gas_limit, NULL AS recipient,
		d.name AS delegate_name
	FROM %s staking_buckets
	LEFT JOIN delegate d ON d.owner_address = staking_buckets.delegate_owner_address`, table)
	var args []interface{}
	var conditions []string
	if requireFinal {
		conditions = append(conditions, "staking_buckets.final = true")
	}
	if startTime > 0 {
		conditions = append(conditions, "staking_buckets.timestamp > ?")
		args = append(args, startTime)
	}
	if len(conditions) > 0 {
		base += " WHERE " + conditions[0]
		for _, c := range conditions[1:] {
			base += " AND " + c
		}
	}
	base += fmt.Sprintf(" ORDER BY %s %s LIMIT ? OFFSET ?", field, order)
	args = append(args, limit, offset)
	return base, args
}

func bucketCountQuery(table string, requireFinal bool, startTime int64) (string, []interface{}) {
	base := fmt.Sprintf("SELECT COUNT(1) AS count FROM %s staking_buckets", table)
	var args []interface{}
	var conditions []string
	if requireFinal {
		conditions = append(conditions, "staking_buckets.final = true")
	}
	if startTime > 0 {
		conditions = append(conditions, "staking_buckets.timestamp > ?")
		args = append(args, startTime)
	}
	if len(conditions) > 0 {
		base += " WHERE " + conditions[0]
		for _, c := range conditions[1:] {
			base += " AND " + c
		}
	}
	return base, args
}

func bucketGroupCountQuery(table string, requireFinal bool, startTime int64) (string, []interface{}) {
	base := fmt.Sprintf("SELECT COUNT(DISTINCT bucket_id) AS count FROM %s staking_buckets", table)
	var args []interface{}
	var conditions []string
	if requireFinal {
		conditions = append(conditions, "staking_buckets.final = true")
	}
	if startTime > 0 {
		conditions = append(conditions, "staking_buckets.timestamp > ?")
		args = append(args, startTime)
	}
	if len(conditions) > 0 {
		base += " WHERE " + conditions[0]
		for _, c := range conditions[1:] {
			base += " AND " + c
		}
	}
	return base, args
}

type versionConfig struct {
	table        string
	requireFinal bool
	isNative     bool
}

var bucketVersionMap = map[string]versionConfig{
	"native": {"staking_buckets", false, true},
	"nft_v1": {"system_staking_buckets", false, false},
	"nft_v2": {"system_staking_buckets_v2_record", true, false},
	"nft_v3": {"system_staking_buckets_v3_record", true, false},
}

// GetBucketList returns a paginated/filtered/sorted bucket list with counts.
func (s *StakingService) GetBucketList(ctx context.Context, req *api.GetBucketListRequest) (*api.GetBucketListResponse, error) {
	resp := &api.GetBucketListResponse{}
	gormDB := db.DB()
	limit := req.GetLimit()
	if limit <= 0 {
		limit = 20
	}
	offset := req.GetOffset()
	version := req.GetVersion()
	if version == "" {
		version = "native"
	}
	vc, ok := bucketVersionMap[version]
	if !ok {
		return nil, errors.Errorf("unknown bucket version: %s", version)
	}
	startTime := intervalToUnixSeconds(req.GetInterval())
	field, order := parseBucketSortParam(req.GetSort())

	var rows []bucketExRow
	var countRow struct{ Count int64 }
	var groupCountRow struct{ Count int64 }

	if vc.isNative {
		q, args := nativeBucketListQuery(startTime, field, order, limit, offset)
		if err := gormDB.WithContext(ctx).Raw(q, args...).Scan(&rows).Error; err != nil {
			return nil, errors.Wrap(err, "failed to get native bucket list")
		}
		cq := "SELECT COUNT(1) AS count FROM staking_buckets"
		cargs := []interface{}{}
		if startTime > 0 {
			cq += " WHERE timestamp > ?"
			cargs = append(cargs, startTime)
		}
		if err := gormDB.WithContext(ctx).Raw(cq, cargs...).Scan(&countRow).Error; err != nil {
			return nil, errors.Wrap(err, "failed to get native bucket count")
		}
		gcq := "SELECT COUNT(DISTINCT bucket_id) AS count FROM staking_buckets"
		gcargs := []interface{}{}
		if startTime > 0 {
			gcq += " WHERE timestamp > ?"
			gcargs = append(gcargs, startTime)
		}
		if err := gormDB.WithContext(ctx).Raw(gcq, gcargs...).Scan(&groupCountRow).Error; err != nil {
			return nil, errors.Wrap(err, "failed to get native bucket group count")
		}
	} else {
		q, args := nftBucketListQuery(vc.table, vc.requireFinal, startTime, field, order, limit, offset)
		if err := gormDB.WithContext(ctx).Raw(q, args...).Scan(&rows).Error; err != nil {
			return nil, errors.Wrap(err, "failed to get nft bucket list")
		}
		cq, cargs := bucketCountQuery(vc.table, vc.requireFinal, startTime)
		if err := gormDB.WithContext(ctx).Raw(cq, cargs...).Scan(&countRow).Error; err != nil {
			return nil, errors.Wrap(err, "failed to get nft bucket count")
		}
		gcq, gcargs := bucketGroupCountQuery(vc.table, vc.requireFinal, startTime)
		if err := gormDB.WithContext(ctx).Raw(gcq, gcargs...).Scan(&groupCountRow).Error; err != nil {
			return nil, errors.Wrap(err, "failed to get nft bucket group count")
		}
	}
	for _, r := range rows {
		resp.Buckets = append(resp.Buckets, toBucketInfoEx(r))
	}
	resp.Count = countRow.Count
	resp.GroupCount = groupCountRow.Count
	return resp, nil
}

// GetBucketsByBucketId returns all bucket actions for a given bucket ID.
func (s *StakingService) GetBucketsByBucketId(ctx context.Context, req *api.GetBucketsByBucketIdRequest) (*api.GetBucketsByBucketIdResponse, error) {
	resp := &api.GetBucketsByBucketIdResponse{}
	gormDB := db.DB()
	limit := req.GetLimit()
	if limit <= 0 {
		limit = 20
	}
	offset := req.GetOffset()
	bucketID := req.GetBucketId()
	version := req.GetVersion()
	if version == "" {
		version = "native"
	}
	vc, ok := bucketVersionMap[version]
	if !ok {
		return nil, errors.Errorf("unknown bucket version: %s", version)
	}

	var rows []bucketExRow
	var countRow struct{ Total int64 }

	if vc.isNative {
		q := `SELECT staking_buckets.action_hash, staking_buckets.bucket_id,
			` + tsExpr + ` AS timestamp,
			` + createTimeExpr + ` AS create_time,
			` + stakeStartExpr + ` AS stake_start_time,
			` + unstakeStartExpr + ` AS unstake_start_time,
			staking_buckets.amount, staking_buckets.staked_amount, staking_buckets.act_type,
			staking_buckets.sender, staking_buckets.owner_address, staking_buckets.candidate,
			staking_buckets.auto_stake, staking_buckets.duration::text AS duration,
			NULL AS gas_price, NULL AS gas_limit, NULL AS recipient,
			'' AS delegate_name
		FROM staking_buckets
		WHERE staking_buckets.bucket_id = ?
		ORDER BY staking_buckets.timestamp DESC LIMIT ? OFFSET ?`
		if err := gormDB.WithContext(ctx).Raw(q, bucketID, limit, offset).Scan(&rows).Error; err != nil {
			return nil, errors.Wrap(err, "failed to get buckets by id")
		}
		if err := gormDB.WithContext(ctx).Raw("SELECT COUNT(id) AS total FROM staking_buckets WHERE bucket_id = ?", bucketID).Scan(&countRow).Error; err != nil {
			return nil, errors.Wrap(err, "failed to count buckets by id")
		}
	} else {
		q := fmt.Sprintf(`SELECT staking_buckets.act_hash AS action_hash, staking_buckets.bucket_id,
			`+tsExpr+` AS timestamp,
			`+createTimeExpr+` AS create_time,
			NULL AS stake_start_time, NULL AS unstake_start_time,
			staking_buckets.amount, staking_buckets.staked_amount, staking_buckets.event_type AS act_type,
			staking_buckets.sender, staking_buckets.owner_address, d.candidate AS candidate,
			staking_buckets.auto_stake, (staking_buckets.duration / 86400.0)::text AS duration,
			NULL AS gas_price, NULL AS gas_limit, NULL AS recipient,
			'' AS delegate_name
		FROM %s staking_buckets
		LEFT JOIN delegate d ON staking_buckets.delegate_owner_address = d.owner_address
		WHERE staking_buckets.bucket_id = ?
		ORDER BY staking_buckets.timestamp DESC LIMIT ? OFFSET ?`, vc.table)
		if err := gormDB.WithContext(ctx).Raw(q, bucketID, limit, offset).Scan(&rows).Error; err != nil {
			return nil, errors.Wrap(err, "failed to get nft buckets by id")
		}
		if err := gormDB.WithContext(ctx).Raw(fmt.Sprintf("SELECT COUNT(id) AS total FROM %s WHERE bucket_id = ?", vc.table), bucketID).Scan(&countRow).Error; err != nil {
			return nil, errors.Wrap(err, "failed to count nft buckets by id")
		}
	}
	for _, r := range rows {
		resp.Buckets = append(resp.Buckets, toBucketInfoEx(r))
	}
	resp.Count = countRow.Total
	return resp, nil
}

// GetBucketByBucketId returns the latest bucket record for a given bucket ID.
func (s *StakingService) GetBucketByBucketId(ctx context.Context, req *api.GetBucketByBucketIdRequest) (*api.GetBucketByBucketIdResponse, error) {
	resp := &api.GetBucketByBucketIdResponse{}
	gormDB := db.DB()
	bucketID := req.GetBucketId()
	version := req.GetVersion()
	if version == "" {
		version = "native"
	}
	vc, ok := bucketVersionMap[version]
	if !ok {
		return nil, errors.Errorf("unknown bucket version: %s", version)
	}

	var row bucketExRow
	var err error
	if vc.isNative {
		q := `SELECT staking_buckets.action_hash, staking_buckets.bucket_id,
			` + tsExpr + ` AS timestamp,
			` + createTimeExpr + ` AS create_time,
			` + stakeStartExpr + ` AS stake_start_time,
			` + unstakeStartExpr + ` AS unstake_start_time,
			staking_buckets.amount, staking_buckets.staked_amount, staking_buckets.act_type,
			staking_buckets.sender, staking_buckets.owner_address, staking_buckets.candidate,
			staking_buckets.auto_stake, staking_buckets.duration::text AS duration,
			NULL AS gas_price, NULL AS gas_limit, NULL AS recipient,
			'' AS delegate_name
		FROM staking_buckets
		WHERE bucket_id = ?
		ORDER BY id DESC LIMIT 1`
		err = gormDB.WithContext(ctx).Raw(q, bucketID).Scan(&row).Error
	} else {
		q := fmt.Sprintf(`SELECT staking_buckets.act_hash AS action_hash, staking_buckets.bucket_id,
			`+tsExpr+` AS timestamp,
			`+createTimeExpr+` AS create_time,
			NULL AS stake_start_time, NULL AS unstake_start_time,
			staking_buckets.amount, staking_buckets.staked_amount, staking_buckets.event_type AS act_type,
			staking_buckets.sender, staking_buckets.owner_address, d.candidate AS candidate,
			staking_buckets.auto_stake, (staking_buckets.duration / 86400.0)::text AS duration,
			NULL AS gas_price, NULL AS gas_limit, NULL AS recipient,
			'' AS delegate_name
		FROM %s staking_buckets
		LEFT JOIN delegate d ON staking_buckets.delegate_owner_address = d.owner_address
		WHERE staking_buckets.bucket_id = ?
		ORDER BY staking_buckets.id DESC LIMIT 1`, vc.table)
		err = gormDB.WithContext(ctx).Raw(q, bucketID).Scan(&row).Error
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get bucket by id")
	}
	if row.ActionHash == "" && row.BucketID == 0 {
		return resp, nil
	}
	resp.Exist = true
	resp.Bucket = toBucketInfoEx(row)
	return resp, nil
}

// GetNativeBuckets returns a DISTINCT ON (bucket_id) list of all native buckets.
func (s *StakingService) GetNativeBuckets(ctx context.Context, req *api.GetNativeBucketsRequest) (*api.GetNativeBucketsResponse, error) {
	resp := &api.GetNativeBucketsResponse{}
	gormDB := db.DB()
	limit := req.GetLimit()
	if limit <= 0 {
		limit = 20
	}
	offset := req.GetOffset()

	q := `SELECT DISTINCT ON (bucket_id)
		staking_buckets.action_hash, staking_buckets.bucket_id,
		` + tsExpr + ` AS timestamp,
		` + createTimeExpr + ` AS create_time,
		` + stakeStartExpr + ` AS stake_start_time,
		` + unstakeStartExpr + ` AS unstake_start_time,
		staking_buckets.amount, staking_buckets.staked_amount, staking_buckets.act_type,
		staking_buckets.sender, staking_buckets.owner_address, staking_buckets.candidate,
		staking_buckets.auto_stake, staking_buckets.duration::text AS duration,
		block_action_partition.gas_price, block_action_partition.gas_limit, block_action_partition.recipient
	FROM staking_buckets
	LEFT JOIN block_action_partition ON staking_buckets.action_hash = block_action_partition.action_hash
	ORDER BY bucket_id DESC, staking_buckets.timestamp DESC
	LIMIT ? OFFSET ?`
	var rows []bucketExRow
	if err := gormDB.WithContext(ctx).Raw(q, limit, offset).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get native buckets")
	}
	var countRow struct{ Count int64 }
	if err := gormDB.WithContext(ctx).Raw("SELECT COUNT(DISTINCT bucket_id) AS count FROM staking_buckets").Scan(&countRow).Error; err != nil {
		return nil, errors.Wrap(err, "failed to count native buckets")
	}
	for _, r := range rows {
		resp.Buckets = append(resp.Buckets, toBucketInfoEx(r))
	}
	resp.Count = countRow.Count
	return resp, nil
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
