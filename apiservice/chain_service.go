package apiservice

import (
	"context"
	"database/sql"
	"math/big"
	"strings"
	"time"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/rewards"
	"github.com/iotexproject/iotex-analyser-api/common/votings"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/pkg/errors"
)

// isUndefinedTableErr reports whether err is a PostgreSQL "undefined_table"
// error (SQLSTATE 42P01). Some tables — like staking_record — are populated
// by external Windmill jobs in production and may simply not exist in
// local-dev. Handlers should treat that as an empty result, not a 500.
func isUndefinedTableErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "SQLSTATE 42P01")
}

// ChainService is the service to handle chain related requests
type ChainService struct {
	api.UnimplementedChainServiceServer
}

// Chain returns the chain info
func (s *ChainService) Chain(ctx context.Context, req *api.ChainRequest) (*api.ChainResponse, error) {
	resp := &api.ChainResponse{}

	epoch, height, err := common.GetCurrentEpochAndHeight()
	if err != nil {
		return nil, err
	}
	resp.MostRecentEpoch = epoch
	resp.MostRecentBlockHeight = height

	totalSupply, err := common.GetTotalSupply(height)
	if err != nil {
		return nil, err
	}
	resp.TotalSupply = totalSupply

	totalCirculatingSupply, err := common.GetTotalCirculatingSupply(height, totalSupply)
	if err != nil {
		return nil, err
	}
	resp.TotalCirculatingSupply = totalCirculatingSupply

	exactCirculatingSupply, err := common.GetExactCirculatingSupply(height, totalCirculatingSupply)
	if err != nil {
		return nil, err
	}
	resp.ExactCirculatingSupply = exactCirculatingSupply

	conn, err := common.NewDefaultGRPCConn(config.Default.RPC)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := iotexapi.NewAPIServiceClient(conn)

	availableRewards, err := rewards.GetAvailableRewards(ctx, client)
	if err != nil {
		return nil, err
	}
	totalRewards, err := rewards.GetTotalRewards(ctx, client)
	if err != nil {
		return nil, err
	}
	resp.Rewards = &api.ChainResponse_Rewards{
		TotalAvailable: availableRewards.String(),
		TotalBalance:   totalRewards.String(),
		TotalUnclaimed: new(big.Int).Sub(totalRewards, availableRewards).String(),
	}
	totalCirculatingSupplyNoRewardPool, err := common.GetTotalCirculatingSupplyNoRewardPool(availableRewards.String(), totalCirculatingSupply)
	if err != nil {
		return nil, err
	}
	resp.TotalCirculatingSupplyNoRewardPool = totalCirculatingSupplyNoRewardPool

	meta, err := votings.GetVotingMeta()
	if err != nil {
		return nil, err
	}

	resp.VotingResultMeta = &api.VotingResultMeta{
		TotalCandidates:    meta.TotalCandidates,
		TotalWeightedVotes: meta.TotalWeightedVotes,
		VotedTokens:        meta.VotedTokens,
	}

	return resp, nil
}

// MostRecentTPS gives the most recent TPS
func (s *ChainService) MostRecentTPS(ctx context.Context, req *api.MostRecentTPSRequest) (*api.MostRecentTPSResponse, error) {
	resp := &api.MostRecentTPSResponse{}

	_, height, err := common.GetCurrentEpochAndHeight()
	if err != nil {
		return nil, err
	}

	blockWindow := req.GetBlockWindow()
	if height < blockWindow {
		blockWindow = height
	}

	start := height - blockWindow + 1
	end := height
	db := db.DB()
	query := "select (select timestamp from block where block_height=?) start_time,(select timestamp from block where block_height=?) end_time,sum(num_actions) num_actions from block where block_height>=? and block_height<=?"
	var result struct {
		StartTime  time.Time
		EndTime    time.Time
		NumActions uint64
	}
	err = db.Raw(query, start, end, start, end).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	resp.MostRecentTPS = float64(result.NumActions) / float64(result.EndTime.Unix()-result.StartTime.Unix())
	return resp, nil
}

// NumberOfActions gives the number of actions within a epoch frame
func (s *ChainService) NumberOfActions(ctx context.Context, req *api.NumberOfActionsRequest) (*api.NumberOfActionsResponse, error) {
	resp := &api.NumberOfActionsResponse{}

	currentEpoch, _, err := common.GetCurrentEpochAndHeight()
	if err != nil {
		return nil, err
	}

	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	if startEpoch > currentEpoch {
		return resp, nil
	}
	endEpoch := startEpoch + epochCount - 1
	db := db.DB()
	query := "select sum(num_actions) num_actions from block b right join block_meta bm on bm.block_height=b.block_height where bm.epoch_num>=? and bm.epoch_num<=?"
	var result struct {
		NumActions uint64
	}
	err = db.Raw(query, startEpoch, endEpoch).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	resp.Exist = true
	resp.Count = result.NumActions
	return resp, nil
}

// TotalTransferredTokens gives the amount of tokens transferred within a time frame
func (s *ChainService) TotalTransferredTokens(ctx context.Context, req *api.TotalTransferredTokensRequest) (*api.TotalTransferredTokensResponse, error) {
	db := db.DB()
	resp := &api.TotalTransferredTokensResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	endEpoch := startEpoch + epochCount - 1
	startHeight := common.GetEpochHeight(startEpoch)
	endHeight := common.GetEpochLastBlockHeight(endEpoch)
	query := "select SUM(amount) from block_receipt_transactions where block_height>=? and block_height<=?"

	var result string
	if err := db.WithContext(ctx).Raw(query, startHeight, endHeight).Scan(&result).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get total number of holders")
	}

	resp.TotalTransferredTokens = result
	return resp, nil
}

func (s *ChainService) BlockSizeByHeight(ctx context.Context, req *api.BlockSizeByHeightRequest) (*api.BlockSizeByHeightResponse, error) {
	resp := &api.BlockSizeByHeightResponse{}
	db := db.DB()
	query := "select block_size from block_meta where block_height=?"
	var size uint64
	if err := db.WithContext(ctx).Raw(query, req.GetHeight()).Scan(&size).Error; err != nil {
		return nil, err
	}
	resp.BlockSize = float64(size) * 0.499
	resp.ServerVersion = getHardForkVersion(req.GetHeight())
	return resp, nil
}

// https://iotexscan.io/hard-fork-history
func getHardForkVersion(blk uint64) string {
	vers := []struct {
		h uint64
		v string
	}{
		{432001, "0.6.2"},
		{864001, "0.7.2"},
		{1512001, "0.8.3"},
		{1641601, "0.9.0"},
		{1816201, "0.10.0"},
		{5165641, "1.0.0"},
		{6544441, "1.1.0"},
		{11267641, "1.2.0"},
		{12289321, "1.3.0"},
		{13685401, "1.4.0"},
		{13816441, "1.5.0"},
		{13979161, "1.6.0"},
		{16509241, "1.7.0"},
		{17662681, "1.8.0"},
	}
	for i := len(vers) - 1; i >= 0; i-- {
		if blk >= vers[i].h {
			return vers[i].v
		}
	}
	return "0.6.0"
}

// GetLatestBlockHeight returns the latest block height
func (s *ChainService) GetLatestBlockHeight(ctx context.Context, req *api.GetLatestBlockHeightRequest) (*api.GetLatestBlockHeightResponse, error) {
	resp := &api.GetLatestBlockHeightResponse{}

	_, height, err := common.GetCurrentEpochAndHeight()
	if err != nil {
		return nil, err
	}

	resp.Height = height
	return resp, nil
}

// GetBlocks returns a list of blocks with pagination
func (s *ChainService) GetBlocks(ctx context.Context, req *api.GetBlocksRequest) (*api.GetBlocksResponse, error) {
	resp := &api.GetBlocksResponse{}

	page := req.GetPage()
	limit := req.GetLimit()
	beforeHeight := req.GetBeforeHeight()

	if page <= 0 {
		page = 1
	}

	var start uint64
	if beforeHeight > 0 {
		start = beforeHeight
	} else {
		_, height, err := common.GetCurrentEpochAndHeight()
		if err != nil {
			return nil, err
		}
		start = height - (page-1)*limit
	}

	gormDB := db.DB()
	query := `SELECT
		m.base_fee,
		m.priority_bonus,
		b.block_height,
		b.block_hash,
		b.producer_address,
		b.num_actions,
		(EXTRACT(EPOCH FROM (b.timestamp AT TIME ZONE 'UTC')) * 1000)::bigint as timestamp,
		m.epoch_num,
		m.gas_consumed,
		m.producer_name,
		m.block_reward
	FROM block b
	LEFT JOIN block_meta m ON m.block_height = b.block_height
	WHERE b.block_height <= ?
	ORDER BY b.block_height DESC
	LIMIT ?`

	type BlockResult struct {
		BaseFee         sql.NullString
		PriorityBonus   sql.NullString
		BlockHeight     uint64
		BlockHash       string
		ProducerAddress string
		NumActions      uint64
		Timestamp       int64
		EpochNum        uint64
		GasConsumed     uint64
		ProducerName    sql.NullString
		BlockReward     sql.NullString
	}

	var results []BlockResult
	if err := gormDB.WithContext(ctx).Raw(query, start, limit).Scan(&results).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get blocks")
	}

	for _, r := range results {
		block := &api.BlockInfo{
			BlockHeight:     r.BlockHeight,
			BlockHash:       r.BlockHash,
			ProducerAddress: r.ProducerAddress,
			NumActions:      r.NumActions,
			Timestamp:       r.Timestamp,
			GasConsumed:     r.GasConsumed,
			EpochNum:        r.EpochNum,
		}

		if r.BaseFee.Valid {
			block.BaseFee = r.BaseFee.String
		}
		if r.PriorityBonus.Valid {
			block.PriorityBonus = r.PriorityBonus.String
		}
		if r.ProducerName.Valid {
			block.ProducerName = r.ProducerName.String
		}
		if r.BlockReward.Valid {
			block.BlockReward = r.BlockReward.String
		}

		resp.Blocks = append(resp.Blocks, block)
	}

	return resp, nil
}

// GetEpochInfo returns the current epoch number and its starting block height
func (s *ChainService) GetEpochInfo(ctx context.Context, req *api.GetEpochInfoRequest) (*api.GetEpochInfoResponse, error) {
	resp := &api.GetEpochInfoResponse{}
	gormDB := db.DB()
	var row struct {
		EpochHeight uint64
		EpochNum    uint64
	}
	if err := gormDB.WithContext(ctx).Raw(
		"SELECT epoch_height, epoch_num FROM block_meta ORDER BY block_height DESC LIMIT 1",
	).Scan(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get epoch info")
	}
	resp.EpochHeight = row.EpochHeight
	resp.EpochNum = row.EpochNum
	return resp, nil
}

// GetLatestStakingRecord returns the most recent staking statistics
func (s *ChainService) GetLatestStakingRecord(ctx context.Context, req *api.GetLatestStakingRecordRequest) (*api.GetLatestStakingRecordResponse, error) {
	resp := &api.GetLatestStakingRecordResponse{}
	gormDB := db.DB()
	var row struct {
		TotalSupply  sql.NullString
		AllStaking   sql.NullString
		StakingRatio sql.NullString
	}
	if err := gormDB.WithContext(ctx).Raw(
		"SELECT total_supply, all_staking, staking_ratio::text FROM staking_record ORDER BY date_time DESC LIMIT 1",
	).Scan(&row).Error; err != nil {
		if isUndefinedTableErr(err) {
			// Table is populated by an external job in production; not yet
			// present in this environment. Return empty response.
			return resp, nil
		}
		return nil, errors.Wrap(err, "failed to get staking record")
	}
	if row.TotalSupply.Valid {
		resp.TotalSupply = row.TotalSupply.String
	}
	if row.AllStaking.Valid {
		resp.AllStaking = row.AllStaking.String
	}
	if row.StakingRatio.Valid {
		resp.StakingRatio = row.StakingRatio.String
	}
	return resp, nil
}

// GetPeakTps returns the all-time peak TPS
func (s *ChainService) GetPeakTps(ctx context.Context, req *api.GetPeakTpsRequest) (*api.GetPeakTpsResponse, error) {
	resp := &api.GetPeakTpsResponse{}
	gormDB := db.DB()
	startBh := req.GetStartBlockHeight()
	var row struct {
		NumActions  sql.NullString
		BlockHeight uint64
	}
	var err error
	if startBh > 0 {
		err = gormDB.WithContext(ctx).Raw(
			"SELECT (SELECT max(block_height) FROM block) AS block_height, (SELECT round(max(num_actions)::NUMERIC/5,2)::text FROM block WHERE block_height > ?) AS num_actions",
			startBh,
		).Scan(&row).Error
	} else {
		err = gormDB.WithContext(ctx).Raw(
			"SELECT (SELECT max(block_height) FROM block) AS block_height, (SELECT round(max(num_actions)::NUMERIC/5,2)::text FROM block) AS num_actions",
		).Scan(&row).Error
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get peak TPS")
	}
	if row.NumActions.Valid {
		resp.NumActions = row.NumActions.String
	}
	resp.BlockHeight = row.BlockHeight
	return resp, nil
}

// GetActionHistory returns aggregated action counts over a time range
func (s *ChainService) GetActionHistory(ctx context.Context, req *api.GetActionHistoryRequest) (*api.GetActionHistoryResponse, error) {
	resp := &api.GetActionHistoryResponse{}
	gormDB := db.DB()
	interval := req.GetInterval()
	if interval != "minute" && interval != "hour" && interval != "day" {
		interval = "hour"
	}
	query := `SELECT sum(num_actions) as sum, to_char(date_time, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') as date
		FROM (
			SELECT num_actions, date_trunc(?, timestamp) AT TIME ZONE 'UTC' as date_time
			FROM block
			WHERE "timestamp" > ? AND "timestamp" <= ?
		) AS b
		GROUP BY date_time`
	var rows []struct {
		Sum  uint64
		Date string
	}
	if err := gormDB.WithContext(ctx).Raw(query, interval, req.GetStartTime(), req.GetEndTime()).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get action history")
	}
	for _, r := range rows {
		resp.Data = append(resp.Data, &api.ActionHistoryPoint{
			Date: r.Date,
			Sum:  r.Sum,
		})
	}
	return resp, nil
}

// GetBlockByHeight returns a single block by its height
func (s *ChainService) GetBlockByHeight(ctx context.Context, req *api.GetBlockByHeightRequest) (*api.GetBlockByHeightResponse, error) {
	resp := &api.GetBlockByHeightResponse{}

	gormDB := db.DB()
	query := `SELECT
		m.base_fee,
		m.priority_bonus,
		b.block_height,
		b.block_hash,
		b.producer_address,
		b.num_actions,
		(EXTRACT(EPOCH FROM (b.timestamp AT TIME ZONE 'UTC')) * 1000)::bigint as timestamp,
		m.epoch_num,
		m.gas_consumed,
		m.producer_name,
		m.block_reward
	FROM block b
	LEFT JOIN block_meta m ON m.block_height = b.block_height
	WHERE b.block_height = ?
	LIMIT 1`

	type BlockResult struct {
		BaseFee         sql.NullString
		PriorityBonus   sql.NullString
		BlockHeight     uint64
		BlockHash       string
		ProducerAddress string
		NumActions      uint64
		Timestamp       int64
		EpochNum        uint64
		GasConsumed     uint64
		ProducerName    sql.NullString
		BlockReward     sql.NullString
	}

	var r BlockResult
	if err := gormDB.WithContext(ctx).Raw(query, req.GetHeight()).Scan(&r).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get block by height")
	}
	if r.BlockHeight == 0 {
		return resp, nil
	}

	resp.Exist = true
	block := &api.BlockInfo{
		BlockHeight:     r.BlockHeight,
		BlockHash:       r.BlockHash,
		ProducerAddress: r.ProducerAddress,
		NumActions:      r.NumActions,
		Timestamp:       r.Timestamp,
		GasConsumed:     r.GasConsumed,
		EpochNum:        r.EpochNum,
	}
	if r.BaseFee.Valid {
		block.BaseFee = r.BaseFee.String
	}
	if r.PriorityBonus.Valid {
		block.PriorityBonus = r.PriorityBonus.String
	}
	if r.ProducerName.Valid {
		block.ProducerName = r.ProducerName.String
	}
	if r.BlockReward.Valid {
		block.BlockReward = r.BlockReward.String
	}
	resp.Block = block

	return resp, nil
}

// GetStakingRatioHistory returns staking ratio history over a time range
func (s *ChainService) GetStakingRatioHistory(ctx context.Context, req *api.GetStakingRatioHistoryRequest) (*api.GetStakingRatioHistoryResponse, error) {
	resp := &api.GetStakingRatioHistoryResponse{}
	gormDB := db.DB()
	startTime := req.GetStartTime()
	endTime := req.GetEndTime()

	var rows []struct {
		DateTime string
		Ratio    string
	}
	var err error
	if startTime != "" && endTime != "" {
		err = gormDB.WithContext(ctx).Raw(
			`SELECT to_char(date_time AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"') as date_time,
				staking_ratio::text as ratio
			FROM staking_record
			WHERE date_time::date >= ?::date AND date_time::date <= ?::date
			ORDER BY id ASC`,
			startTime, endTime,
		).Scan(&rows).Error
	} else {
		err = gormDB.WithContext(ctx).Raw(
			`SELECT to_char(date_time AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"') as date_time,
				staking_ratio::text as ratio
			FROM staking_record ORDER BY id ASC`,
		).Scan(&rows).Error
	}
	if err != nil {
		if isUndefinedTableErr(err) {
			return resp, nil
		}
		return nil, errors.Wrap(err, "failed to get staking ratio history")
	}
	for _, r := range rows {
		resp.Data = append(resp.Data, &api.StakingRatioPoint{
			DateTime: r.DateTime,
			Ratio:    r.Ratio,
		})
	}
	return resp, nil
}
