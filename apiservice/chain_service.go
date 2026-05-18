package apiservice

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"strconv"
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
	"github.com/shopspring/decimal"
)

// blockIntervalSecs is the chain's post-hardfork block interval, used as the
// TPS denominator. Pre-hardfork blocks were 5 s; the homepage chart shows the
// current state, so a single value is fine.
const blockIntervalSecs = 2.5

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

// GetChainStats returns the homepage chain stats by reading three pre-computed
// keys from iotexscanv3_kv (maintained by the iotex-statistics Windmill job).
// Supply values are already in IOTX units.
func (s *ChainService) GetChainStats(ctx context.Context, req *api.GetChainStatsRequest) (*api.GetChainStatsResponse, error) {
	resp := &api.GetChainStatsResponse{}

	var rows []struct {
		Key   string
		Value sql.NullString
	}
	if err := db.DB().WithContext(ctx).Raw(
		`SELECT key, value FROM iotexscanv3_kv WHERE key IN (?, ?, ?)`,
		"actions_num", "total_supply", "total_circulating_supply",
	).Scan(&rows).Error; err != nil {
		if isUndefinedTableErr(err) {
			return resp, nil
		}
		return nil, errors.Wrap(err, "failed to read iotexscanv3_kv")
	}

	for _, r := range rows {
		if !r.Value.Valid || r.Value.String == "" {
			continue
		}
		switch r.Key {
		case "actions_num":
			n, err := strconv.ParseUint(r.Value.String, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "actions_num is not a uint64: %q", r.Value.String)
			}
			resp.ActionsNum = n
		case "total_supply":
			resp.TotalSupply = r.Value.String
		case "total_circulating_supply":
			resp.CirculatingSupply = r.Value.String
		}
	}
	return resp, nil
}

// GetTpsHistory returns daily avg/max TPS over a date range. avg_tps and
// max_tps are computed from block.num_actions divided by the block interval.
func (s *ChainService) GetTpsHistory(ctx context.Context, req *api.GetTpsHistoryRequest) (*api.GetTpsHistoryResponse, error) {
	resp := &api.GetTpsHistoryResponse{}
	gormDB := db.DB()
	q := fmt.Sprintf(`SELECT to_char(date_trunc('day', timestamp AT TIME ZONE 'UTC'), 'YYYY-MM-DD') AS date,
		ROUND(AVG(num_actions)::numeric / %g, 2)::float8 AS avg_tps,
		ROUND(MAX(num_actions)::numeric / %g, 2)::float8 AS max_tps
		FROM block
		WHERE timestamp >= ?::date AND timestamp < (?::date + INTERVAL '1 day')
		GROUP BY 1
		ORDER BY 1 ASC`, blockIntervalSecs, blockIntervalSecs)
	var rows []struct {
		Date   string
		AvgTps float64
		MaxTps float64
	}
	if err := gormDB.WithContext(ctx).Raw(q, req.GetStart(), req.GetEnd()).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get tps history")
	}
	for _, r := range rows {
		resp.Data = append(resp.Data, &api.TpsHistoryPoint{
			Date:   r.Date,
			AvgTps: r.AvgTps,
			MaxTps: r.MaxTps,
		})
	}
	return resp, nil
}

// GetGasHistory returns daily gas price stats and total gas fee, aggregated
// live from block_action joined with block. Rows with gas_price = 0 (e.g.
// system actions) are excluded.
//
// Implementation note: we resolve the block_height bounds in a cheap
// timestamp-indexed query first, then pass them as literal integers into the
// main aggregation. Without this, the planner picks block_pkey + filter and
// loses block_action partition pruning, causing a 10×+ slowdown (17 s → 2 s
// for a 7-day window).
func (s *ChainService) GetGasHistory(ctx context.Context, req *api.GetGasHistoryRequest) (*api.GetGasHistoryResponse, error) {
	resp := &api.GetGasHistoryResponse{}
	gormDB := db.DB().WithContext(ctx)

	var bounds struct {
		Lo sql.NullInt64
		Hi sql.NullInt64
	}
	// MIN/MAX with WHERE on a different indexed column tricks PG into scanning
	// block_pkey backwards filtering by timestamp — 47 s on a 7-day window.
	// ORDER BY timestamp ... LIMIT 1 forces use of idx_block_timestamp (~1 ms).
	if err := gormDB.Raw(
		`SELECT
			(SELECT block_height FROM block
			 WHERE timestamp >= ?::date AND timestamp < (?::date + INTERVAL '1 day')
			 ORDER BY timestamp ASC LIMIT 1) AS lo,
			(SELECT block_height FROM block
			 WHERE timestamp >= ?::date AND timestamp < (?::date + INTERVAL '1 day')
			 ORDER BY timestamp DESC LIMIT 1) AS hi`,
		req.GetStart(), req.GetEnd(),
		req.GetStart(), req.GetEnd(),
	).Scan(&bounds).Error; err != nil {
		return nil, errors.Wrap(err, "failed to resolve block_height bounds")
	}
	if !bounds.Lo.Valid || !bounds.Hi.Valid {
		return resp, nil // no blocks in window
	}

	var rows []struct {
		Date        string
		MaxGasPrice sql.NullString
		MinGasPrice sql.NullString
		AvgGasPrice sql.NullString
		TotalGasFee sql.NullString
	}
	// Inline the block_height bounds as literals (not parameters) so PG's
	// planner sees concrete values at plan time and can prune block_action's
	// 49 partitions. With prepared-statement parameters the planner falls
	// back to a generic plan over all partitions, costing ~14× more.
	// Safe to inline: bounds are int64 derived from a DB lookup, never user
	// input.
	q := fmt.Sprintf(`SELECT to_char(b.timestamp::date, 'YYYY-MM-DD') AS date,
		MAX(ba.gas_price)::text                       AS max_gas_price,
		MIN(ba.gas_price)::text                       AS min_gas_price,
		ROUND(AVG(ba.gas_price))::text                AS avg_gas_price,
		SUM(ba.gas_price * ba.gas_consumed)::text     AS total_gas_fee
		FROM block_action ba
		JOIN block b ON b.block_height = ba.block_height
		WHERE ba.block_height BETWEEN %d AND %d
		  AND ba.gas_price > 0
		GROUP BY 1
		ORDER BY 1 ASC`, bounds.Lo.Int64, bounds.Hi.Int64)
	if err := gormDB.Raw(q).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get gas history")
	}
	for _, r := range rows {
		resp.Data = append(resp.Data, &api.GasHistoryPoint{
			Date:        r.Date,
			MaxGasPrice: r.MaxGasPrice.String,
			MinGasPrice: r.MinGasPrice.String,
			AvgGasPrice: r.AvgGasPrice.String,
			TotalGasFee: r.TotalGasFee.String,
		})
	}
	return resp, nil
}

// GetSupplyHistory returns daily total/circulating supply (IOTX) plus daily
// burn/issue derived from the supply deltas.
//
// Source: block_supply (1 row per block) joined with block (for the timestamp).
// For each day we pick the last block of the day and read its supply snapshot.
// We fetch one extra day before [start, end] so we have a baseline to diff
// against for the first user-requested day.
//
//	burn(t)  = total_supply(t-1) - total_supply(t)             (zero address only receives)
//	issue(t) = (circ(t) - circ(t-1)) + burn(t)                 (lock address only sends)
//
// All output amounts are IOTX (rau / 1e18) with 2 decimals.
func (s *ChainService) GetSupplyHistory(ctx context.Context, req *api.GetSupplyHistoryRequest) (*api.GetSupplyHistoryResponse, error) {
	resp := &api.GetSupplyHistoryResponse{}

	q := `WITH daily AS (
		SELECT b.timestamp::date AS day, MAX(b.block_height) AS max_height
		FROM block b
		WHERE b.timestamp >= ?::date - INTERVAL '1 day'
		  AND b.timestamp <  (?::date + INTERVAL '1 day')
		GROUP BY 1
	)
	SELECT to_char(d.day, 'YYYY-MM-DD') AS date,
	       bs.total_supply::text             AS total_supply,
	       bs.total_circulating_supply::text AS circulating_supply
	FROM daily d
	JOIN block_supply bs ON bs.block_height = d.max_height
	ORDER BY d.day ASC`

	var rows []struct {
		Date              string
		TotalSupply       sql.NullString
		CirculatingSupply sql.NullString
	}
	if err := db.DB().WithContext(ctx).Raw(q, req.GetStart(), req.GetEnd()).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get supply history")
	}
	if len(rows) == 0 {
		return resp, nil
	}

	// Parse rau strings into decimals once. Skip rows with NULL/invalid values.
	type day struct {
		date   string
		ts     decimal.Decimal // total_supply (rau)
		cs     decimal.Decimal // circulating_supply (rau)
		hasTS  bool
		hasCS  bool
	}
	parsed := make([]day, 0, len(rows))
	for _, r := range rows {
		d := day{date: r.Date}
		if r.TotalSupply.Valid {
			if v, err := decimal.NewFromString(r.TotalSupply.String); err == nil {
				d.ts, d.hasTS = v, true
			}
		}
		if r.CirculatingSupply.Valid {
			if v, err := decimal.NewFromString(r.CirculatingSupply.String); err == nil {
				d.cs, d.hasCS = v, true
			}
		}
		parsed = append(parsed, d)
	}

	startDate := req.GetStart()
	for i, d := range parsed {
		if d.date < startDate {
			continue // baseline row used only for diff
		}
		point := &api.SupplyHistoryPoint{Date: d.date}
		if d.hasTS {
			point.TotalSupply = rauDecimalToIOTX(d.ts)
		}
		if d.hasCS {
			point.CirculatingSupply = rauDecimalToIOTX(d.cs)
		}
		if i == 0 {
			// No baseline available — leave burn/issue empty.
			resp.Data = append(resp.Data, point)
			continue
		}
		prev := parsed[i-1]
		if d.hasTS && prev.hasTS {
			burn := prev.ts.Sub(d.ts)
			point.Burn = rauDecimalToIOTX(burn)
			if d.hasCS && prev.hasCS {
				issue := d.cs.Sub(prev.cs).Add(burn)
				point.Issue = rauDecimalToIOTX(issue)
			}
		}
		resp.Data = append(resp.Data, point)
	}
	return resp, nil
}

// rauDecimalToIOTX divides a rau-denominated decimal by 1e18 and formats with
// 2 decimal places.
func rauDecimalToIOTX(rau decimal.Decimal) string {
	return rau.Shift(-18).StringFixed(2)
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
