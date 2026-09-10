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
	"fmt"

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

// receivedVoteSource describes one table that can hold buckets voting for a
// delegate. Native staking and the system-contract (NFT) staking tables are
// separate append-only action logs with different delegate columns, so each is
// reduced to its latest row per bucket_id independently and the results are
// merged. The response carries no bucket_id, so a flat merge is what the caller
// sees — bucket_id collides across native and system tables and must never be
// used to join between them.
type receivedVoteSource struct {
	table string
	// delegateCol is the column naming the delegate this bucket votes for.
	delegateCol string
	// activeCond is the SQL keeping only buckets still voting, evaluated on the
	// latest row of each bucket.
	activeCond string
	// durationExpr normalizes duration to whole days, matching native staking.
	durationExpr string
}

var receivedVoteSources = []receivedVoteSource{
	// Native staking: a bucket that has left native staking zeroes its amount on
	// the way out, which is what staked_amount > 0 tests. Across mainnet that is
	// exactly WithdrawStake (62,177 buckets) and MigrateStake (254, moved to the
	// staking contract and counted there instead); every other act_type has a
	// non-zero amount on its latest row, so nothing live is excluded. An unstaked
	// bucket keeps its amount and is recognised the way the chain-side reader does
	// it, by unstake_start_time having overtaken stake_start_time.
	{
		table:        "staking_buckets",
		delegateCol:  "candidate",
		activeCond:   "l.staked_amount > 0 AND NOT (l.unstake_start_time > l.stake_start_time)",
		durationExpr: "l.duration::text",
	},
	// System-contract staking v1/v2/v3. `final` marks a settled record and `muted`
	// one that no longer counts; Withdrawal/Unstaked rows additionally zero
	// staked_amount, which is what actually separates a live bucket from a spent one.
	{
		table:        "system_staking_buckets_record",
		delegateCol:  "delegate_owner_address",
		activeCond:   "l.final AND NOT l.muted AND l.staked_amount > 0 AND NOT (l.unstake_start_time > l.stake_start_time)",
		durationExpr: "(CASE WHEN l.duration_type = 0 THEN l.duration ELSE l.duration / 86400 END)::text",
	},
	{
		table:        "system_staking_buckets_v2_record",
		delegateCol:  "delegate_owner_address",
		activeCond:   "l.final AND NOT l.muted AND l.staked_amount > 0 AND NOT (l.unstake_start_time > l.stake_start_time)",
		durationExpr: "(CASE WHEN l.duration_type = 0 THEN l.duration ELSE l.duration / 86400 END)::text",
	},
	{
		table:        "system_staking_buckets_v3_record",
		delegateCol:  "delegate_owner_address",
		activeCond:   "l.final AND NOT l.muted AND l.staked_amount > 0 AND NOT (l.unstake_start_time > l.stake_start_time)",
		durationExpr: "(CASE WHEN l.duration_type = 0 THEN l.duration ELSE l.duration / 86400 END)::text",
	},
}

// resolveDelegateAddresses maps whatever delegate address the caller passed to
// the pair the two bucket families are keyed by. They are the same value for
// most delegates but not all (4 of 123 on mainnet as of 2026-09), so a caller
// passing either one must still get the complete list.
func resolveDelegateAddresses(ctx context.Context, addr string) (candidate, ownerAddress string) {
	var row struct {
		Candidate    sql.NullString
		OwnerAddress sql.NullString
	}
	res := db.DB().WithContext(ctx).Raw(
		`SELECT candidate, owner_address FROM delegate WHERE candidate = ? OR owner_address = ? LIMIT 1`,
		addr, addr,
	).Scan(&row)
	if res.Error != nil || res.RowsAffected == 0 {
		// Not a known delegate (or the lookup failed): query both families with
		// the address as given rather than silently returning nothing.
		return addr, addr
	}
	return row.Candidate.String, row.OwnerAddress.String
}

// GetReceivedVotesByAddress lists the buckets currently voting for a delegate,
// across native staking and the system-contract staking tables.
//
// Note both sources are action logs, one row per staking action, so a bucket
// appears many times (mainnet: 102,334 rows for 87 buckets on one address).
// Each source is therefore collapsed to the latest row per bucket_id, and a
// bucket whose newest row moved it to another delegate is dropped — on mainnet
// that is 1,158 of 1,722 otherwise-active buckets for a single delegate, so
// skipping the check overcounts roughly threefold.
func (s *DelegateService) GetReceivedVotesByAddress(ctx context.Context, req *api.GetReceivedVotesByAddressRequest) (*api.GetReceivedVotesByAddressResponse, error) {
	resp := &api.GetReceivedVotesByAddressResponse{}
	addr, err := toIo(req.GetAddress())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address: %v", err)
	}
	if addr == "" {
		return nil, status.Errorf(codes.InvalidArgument, "address is required")
	}
	candidate, ownerAddress := resolveDelegateAddresses(ctx, addr)

	for _, src := range receivedVoteSources {
		delegateAddr := candidate
		if src.delegateCol != "candidate" {
			delegateAddr = ownerAddress
		}
		items, err := queryReceivedVotes(ctx, src, delegateAddr)
		if err != nil {
			return nil, err
		}
		resp.Data = append(resp.Data, items...)
	}
	return resp, nil
}

func queryReceivedVotes(ctx context.Context, src receivedVoteSource, delegateAddr string) ([]*api.ReceivedVoteItem, error) {
	// DISTINCT ON picks the newest row per bucket among the rows naming this
	// delegate; NOT EXISTS then drops buckets whose newest row overall is newer
	// still, i.e. buckets that have since moved to a different delegate.
	query := fmt.Sprintf(`
		WITH latest AS (
			SELECT DISTINCT ON (bucket_id) *
			FROM %[1]s
			WHERE %[2]s = ?
			ORDER BY bucket_id, id DESC
		)
		SELECT l.owner_address, l.staked_amount::text, l.voting_power::text, %[3]s
		FROM latest l
		WHERE %[4]s
		  AND NOT EXISTS (SELECT 1 FROM %[1]s s WHERE s.bucket_id = l.bucket_id AND s.id > l.id)
		ORDER BY l.voting_power DESC, l.bucket_id DESC`,
		src.table, src.delegateCol, src.durationExpr, src.activeCond)

	rows, err := db.DB().WithContext(ctx).Raw(query, delegateAddr).Rows()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to query received votes from %s", src.table)
	}
	defer rows.Close()
	var items []*api.ReceivedVoteItem
	for rows.Next() {
		var staker, amount, votes, duration sql.NullString
		if err := rows.Scan(&staker, &amount, &votes, &duration); err != nil {
			return nil, errors.Wrapf(err, "scan received vote row from %s", src.table)
		}
		items = append(items, &api.ReceivedVoteItem{
			Staker:   staker.String,
			Amount:   amount.String,
			Votes:    votes.String,
			Duration: duration.String,
		})
	}
	return items, rows.Err()
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
