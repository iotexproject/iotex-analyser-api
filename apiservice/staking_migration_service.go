package apiservice

// GetStakingHistory backs iotex-kit modules-db/staking.getStakingHistory, which
// previously hit the analyzer Postgres directly (this.analysis on
// staking_buckets). SQL ported 1:1. Method hangs off the existing
// StakingService. Uses the package-local toIo() helper.

import (
	"context"
	"database/sql"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *StakingService) GetStakingHistory(ctx context.Context, req *api.GetStakingHistoryRequest) (*api.GetStakingHistoryResponse, error) {
	resp := &api.GetStakingHistoryResponse{}
	addr, err := toIo(req.GetOwnerAddress())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid owner_address: %v", err)
	}
	if addr == "" {
		return nil, status.Errorf(codes.InvalidArgument, "owner_address is required")
	}
	limit := req.GetLimit()
	if limit == 0 {
		limit = 25
	}
	page := req.GetPage()
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	// Latest row per bucket_id (MAX(id) grouped) for the owner, ordered id desc.
	dataQuery := `SELECT bucket_id, candidate, create_time, unstake_start_time, staked_amount
		FROM staking_buckets
		WHERE owner_address = ?
			AND id IN (SELECT MAX(id) FROM staking_buckets WHERE owner_address = ? GROUP BY bucket_id)
		ORDER BY id DESC
		LIMIT ? OFFSET ?`
	rows, err := db.DB().WithContext(ctx).Raw(dataQuery, addr, addr, limit, offset).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query staking history")
	}
	defer rows.Close()
	for rows.Next() {
		var bucketID sql.NullInt64
		var candidate, createTime, unstakeStartTime, stakedAmount sql.NullString
		if err := rows.Scan(&bucketID, &candidate, &createTime, &unstakeStartTime, &stakedAmount); err != nil {
			return nil, errors.Wrap(err, "scan staking history row")
		}
		resp.Data = append(resp.Data, &api.StakingHistoryItem{
			BucketId:         uint64(bucketID.Int64),
			Candidate:        candidate.String,
			CreateTime:       createTime.String,
			UnstakeStartTime: unstakeStartTime.String,
			StakedAmount:     stakedAmount.String,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	var totalRow struct{ Total sql.NullInt64 }
	if err := db.DB().WithContext(ctx).Raw(
		`SELECT COUNT(DISTINCT bucket_id) AS total FROM staking_buckets WHERE owner_address = ?`, addr,
	).Scan(&totalRow).Error; err != nil {
		return nil, errors.Wrap(err, "failed to count staking history")
	}
	resp.Total = uint64(totalRow.Total.Int64)
	return resp, nil
}
