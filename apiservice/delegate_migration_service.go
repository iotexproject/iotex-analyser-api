package apiservice

// Handlers backing the iotex-kit modules-db/delegate.ts endpoints that were
// previously served by kit connecting directly to the analyzer Postgres via
// ANALYZER_DATABASE_URL (this.analysis). SQL is ported 1:1 from those kit
// methods. Methods hang off the existing DelegateService.
//
// Reuses the package-local toIo() helper (apiservice/iotexscan_service.go) for
// 0x -> io normalization.

import (
	"context"
	"database/sql"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetDelegateHeight: latest block_meta height for a producer name.
func (s *DelegateService) GetDelegateHeight(ctx context.Context, req *api.GetDelegateHeightRequest) (*api.GetDelegateHeightResponse, error) {
	resp := &api.GetDelegateHeightResponse{}
	if req.GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}
	var row struct{ BlockHeight sql.NullInt64 }
	if err := db.DB().WithContext(ctx).Raw(
		`SELECT block_height FROM block_meta WHERE producer_name = ? ORDER BY block_height DESC LIMIT 1`,
		req.GetName(),
	).Scan(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query delegate height")
	}
	if row.BlockHeight.Valid {
		resp.Exist = true
		resp.BlockHeight = uint64(row.BlockHeight.Int64)
	}
	return resp, nil
}

// GetProductivityHistory: delegate_productivity_history rows in a date range.
func (s *DelegateService) GetProductivityHistory(ctx context.Context, req *api.GetProductivityHistoryRequest) (*api.GetProductivityHistoryResponse, error) {
	resp := &api.GetProductivityHistoryResponse{}
	cand, err := toIo(req.GetCandidate())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid candidate: %v", err)
	}
	if cand == "" {
		return nil, status.Errorf(codes.InvalidArgument, "candidate is required")
	}
	if req.GetEndDate() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "end_date is required")
	}
	query := `SELECT id, productivity::numeric AS productivity, candidate AS temp_eth_address, date_time::text AS date
		FROM delegate_productivity_history
		WHERE candidate = ? AND date_time < ?`
	args := []interface{}{cand, req.GetEndDate()}
	if req.GetStartDate() != "" {
		query += ` AND date_time > ?`
		args = append(args, req.GetStartDate())
	}
	rows, err := db.DB().WithContext(ctx).Raw(query, args...).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query productivity history")
	}
	defer rows.Close()
	for rows.Next() {
		var id sql.NullInt64
		var productivity, tempEth, date sql.NullString
		if err := rows.Scan(&id, &productivity, &tempEth, &date); err != nil {
			return nil, errors.Wrap(err, "scan productivity row")
		}
		resp.Data = append(resp.Data, &api.ProductivityHistoryItem{
			Id:             id.Int64,
			Productivity:   productivity.String,
			TempEthAddress: tempEth.String,
			Date:           date.String,
		})
	}
	return resp, rows.Err()
}

// GetProbationHistory: probation days for a candidate's operator in a date range.
func (s *DelegateService) GetProbationHistory(ctx context.Context, req *api.GetProbationHistoryRequest) (*api.GetProbationHistoryResponse, error) {
	resp := &api.GetProbationHistoryResponse{}
	cand, err := toIo(req.GetCandidate())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid candidate: %v", err)
	}
	if cand == "" {
		return nil, status.Errorf(codes.InvalidArgument, "candidate is required")
	}
	if req.GetStartDate() == "" || req.GetEndDate() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "start_date and end_date are required")
	}
	// Ported 1:1 from kit delegate.getProbationHistory.
	query := `SELECT true AS probation, address, ("year" || '-' || "month" || '-' || "day") AS date
		FROM (
			SELECT p.block_height, "count", address, "year", "month", "day", "timestamp"
			FROM probation p
			LEFT JOIN block b ON p.block_height = b.block_height
		) a
		WHERE address = (SELECT operator_address FROM delegate WHERE candidate = ?)
			AND a.timestamp > ? AND a.timestamp < ?
		GROUP BY a.year, a.month, a.day, a.address`
	rows, err := db.DB().WithContext(ctx).Raw(query, cand, req.GetStartDate(), req.GetEndDate()).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query probation history")
	}
	defer rows.Close()
	for rows.Next() {
		var probation sql.NullBool
		var address, date sql.NullString
		if err := rows.Scan(&probation, &address, &date); err != nil {
			return nil, errors.Wrap(err, "scan probation row")
		}
		resp.Data = append(resp.Data, &api.ProbationHistoryItem{
			Probation: probation.Bool,
			Address:   address.String,
			Date:      date.String,
		})
	}
	return resp, rows.Err()
}

// GetDelegateRewards: single delegate_rewards row for a candidate.
func (s *DelegateService) GetDelegateRewards(ctx context.Context, req *api.GetDelegateRewardsRequest) (*api.GetDelegateRewardsResponse, error) {
	resp := &api.GetDelegateRewardsResponse{}
	cand, err := toIo(req.GetCandidate())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid candidate: %v", err)
	}
	if cand == "" {
		return nil, status.Errorf(codes.InvalidArgument, "candidate is required")
	}
	var row struct {
		BlockReward     sql.NullString
		EpochReward     sql.NullString
		FoundationBonus sql.NullString
		BurnReward      sql.NullString
	}
	// Scan a single row; kit returned rewards?.[0].
	res := db.DB().WithContext(ctx).Raw(
		`SELECT block_reward, epoch_reward, foundation_bonus, burn_reward
		 FROM delegate_rewards WHERE candidate = ? LIMIT 1`, cand,
	).Scan(&row)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, "failed to query delegate rewards")
	}
	if res.RowsAffected > 0 {
		resp.Exist = true
		resp.BlockReward = row.BlockReward.String
		resp.EpochReward = row.EpochReward.String
		resp.FoundationBonus = row.FoundationBonus.String
		resp.BurnReward = row.BurnReward.String
	}
	return resp, nil
}

// GetDelegateRewardsHistory: hermes_delegate_rewards_history rows in a date range.
func (s *DelegateService) GetDelegateRewardsHistory(ctx context.Context, req *api.GetDelegateRewardsHistoryRequest) (*api.GetDelegateRewardsHistoryResponse, error) {
	resp := &api.GetDelegateRewardsHistoryResponse{}
	cand, err := toIo(req.GetCandidate())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid candidate: %v", err)
	}
	if cand == "" {
		return nil, status.Errorf(codes.InvalidArgument, "candidate is required")
	}
	if req.GetStartDate() == "" || req.GetEndDate() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "start_date and end_date are required")
	}
	query := `SELECT block_reward, epoch_reward, foundation_bonus, burn_reward, date_time::text
		FROM hermes_delegate_rewards_history
		WHERE date_time BETWEEN ?::date AND ?::date
			AND candidate_name = (SELECT name FROM delegate WHERE candidate = ?)`
	rows, err := db.DB().WithContext(ctx).Raw(query, req.GetStartDate(), req.GetEndDate(), cand).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query delegate rewards history")
	}
	defer rows.Close()
	for rows.Next() {
		var blockReward, epochReward, foundationBonus, burnReward, dateTime sql.NullString
		if err := rows.Scan(&blockReward, &epochReward, &foundationBonus, &burnReward, &dateTime); err != nil {
			return nil, errors.Wrap(err, "scan rewards history row")
		}
		resp.Data = append(resp.Data, &api.DelegateRewardsHistoryItem{
			BlockReward:     blockReward.String,
			EpochReward:     epochReward.String,
			FoundationBonus: foundationBonus.String,
			BurnReward:      burnReward.String,
			DateTime:        dateTime.String,
		})
	}
	return resp, rows.Err()
}

// GetReceivedVotesByAddress: staking_buckets rows where owner_address = addr.
func (s *DelegateService) GetReceivedVotesByAddress(ctx context.Context, req *api.GetReceivedVotesByAddressRequest) (*api.GetReceivedVotesByAddressResponse, error) {
	resp := &api.GetReceivedVotesByAddressResponse{}
	// kit passed the address through as-is (no conversion); accept 0x/io and normalize.
	addr, err := toIo(req.GetAddress())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address: %v", err)
	}
	if addr == "" {
		return nil, status.Errorf(codes.InvalidArgument, "address is required")
	}
	rows, err := db.DB().WithContext(ctx).Raw(
		`SELECT sender AS staker, staked_amount AS amount, voting_power AS votes, duration
		 FROM staking_buckets WHERE owner_address = ? ORDER BY timestamp DESC`, addr,
	).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query received votes")
	}
	defer rows.Close()
	for rows.Next() {
		var staker, amount, votes, duration sql.NullString
		if err := rows.Scan(&staker, &amount, &votes, &duration); err != nil {
			return nil, errors.Wrap(err, "scan received vote row")
		}
		resp.Data = append(resp.Data, &api.ReceivedVoteItem{
			Staker:   staker.String,
			Amount:   amount.String,
			Votes:    votes.String,
			Duration: duration.String,
		})
	}
	return resp, rows.Err()
}

// GetDelegatesStatistics: count + total stake over the delegate table.
func (s *DelegateService) GetDelegatesStatistics(ctx context.Context, req *api.GetDelegatesStatisticsRequest) (*api.GetDelegatesStatisticsResponse, error) {
	resp := &api.GetDelegatesStatisticsResponse{}
	var row struct {
		DelegateCount sql.NullInt64
		TotalAmount   sql.NullString
	}
	res := db.DB().WithContext(ctx).Raw(
		`SELECT count(1) AS delegate_count, sum(stake_amount)::text AS total_amount FROM delegate`,
	).Scan(&row)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, "failed to query delegates statistics")
	}
	if res.RowsAffected > 0 {
		resp.Exist = true
		resp.DelegateCount = uint64(row.DelegateCount.Int64)
		resp.TotalAmount = row.TotalAmount.String
	}
	return resp, nil
}
