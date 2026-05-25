package common

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	MaxBalanceHistoryDays = 60
	balanceHistoryTimeFmt = "2006-01-02 15:04:05"
)

type BalanceHistoryPoint struct {
	Timestamp int64
	Balance   string
	Delta     string
}

// GetBalanceHistory returns the IOTX balance series for addr over the last
// `days` days. days is clamped to [1, MaxBalanceHistoryDays]. days==1 yields 24
// hourly buckets; days>=2 yields N daily buckets. Each point reports the
// cumulative balance at the bucket end (rau) and the net flow within the
// bucket. The bucket window is anchored to UTC wall-clock boundaries.
func GetBalanceHistory(ctx context.Context, addr string, days int32) ([]BalanceHistoryPoint, error) {
	if days <= 0 {
		days = 1
	}
	if days > MaxBalanceHistoryDays {
		days = MaxBalanceHistoryDays
	}

	normAddr, err := Address(addr)
	if err != nil {
		return nil, err
	}
	addrLower := strings.ToLower(*normAddr)

	var (
		stepInterval string
		truncUnit    string
		n            int
	)
	if days == 1 {
		stepInterval, truncUnit, n = "1 hour", "hour", 24
	} else {
		stepInterval, truncUnit, n = "1 day", "day", int(days)
	}

	now := time.Now().UTC()
	var end time.Time
	if truncUnit == "hour" {
		end = now.Truncate(time.Hour)
	} else {
		end = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	}
	step := time.Hour
	if truncUnit == "day" {
		step = 24 * time.Hour
	}
	start := end.Add(-time.Duration(n-1) * step)
	nextAfterEnd := end.Add(step)

	gdb := db.DB().WithContext(ctx)

	// Resolve block_height bounds via indexed timestamp lookup. The pattern
	// (subquery with ORDER BY timestamp ... LIMIT 1) forces use of
	// idx_block_timestamp; a plain MAX(block_height) WHERE timestamp < ?
	// would scan block_pkey backwards. See GetGasHistory for the same trick.
	var bounds struct {
		BaselineH sql.NullInt64
		EndH      sql.NullInt64
	}
	if err := gdb.Raw(
		`SELECT
			(SELECT block_height FROM block
			 WHERE (timestamp AT TIME ZONE 'UTC') < ?::timestamp
			 ORDER BY timestamp DESC LIMIT 1) AS baseline_h,
			(SELECT block_height FROM block
			 WHERE (timestamp AT TIME ZONE 'UTC') < ?::timestamp
			 ORDER BY timestamp DESC LIMIT 1) AS end_h`,
		start.Format(balanceHistoryTimeFmt),
		nextAfterEnd.Format(balanceHistoryTimeFmt),
	).Scan(&bounds).Error; err != nil {
		return nil, errors.Wrap(err, "failed to resolve block_height bounds")
	}

	var baselineH, endH int64
	if bounds.BaselineH.Valid {
		baselineH = bounds.BaselineH.Int64
	}
	if bounds.EndH.Valid {
		endH = bounds.EndH.Int64
	}

	// Baseline balance: state at end of baseline_h. We need a single number,
	// not a window-sum — so prefer the eth-archive RPC (single O(1) state
	// read, ~100ms regardless of account history). The SQL fallback summing
	// account_income works but degrades to tens of seconds for old contracts
	// with millions of historical rows.
	baseline, err := resolveBaseline(ctx, gdb, addrLower, baselineH)
	if err != nil {
		return nil, errors.Wrap(err, "failed to compute baseline balance")
	}

	startStr := start.Format(balanceHistoryTimeFmt)
	endStr := end.Format(balanceHistoryTimeFmt)

	var (
		query string
		args  []any
	)
	if endH > baselineH {
		// Inline height bounds as literals so PG's planner sees concrete
		// values and can prune account_income / block_action partitions.
		// Bounds are int64 from a DB lookup; safe to inline.
		query = fmt.Sprintf(`
WITH buckets AS (
    SELECT generate_series(?::timestamp, ?::timestamp, ?::interval) AS bucket
),
deltas AS (
    SELECT date_trunc(?, b.timestamp AT TIME ZONE 'UTC') AS bucket,
           SUM(ai.in_flow - ai.out_flow)::numeric        AS delta
    FROM account_income ai
    JOIN block b ON b.block_height = ai.block_height
    WHERE ai.address = ?
      AND ai.block_height BETWEEN %d AND %d
    GROUP BY 1
)
SELECT
    EXTRACT(EPOCH FROM (b.bucket + ?::interval))::bigint                          AS ts,
    COALESCE(d.delta, 0)::text                                                    AS delta,
    (?::numeric + SUM(COALESCE(d.delta, 0)) OVER (ORDER BY b.bucket))::text       AS balance
FROM buckets b
LEFT JOIN deltas d ON d.bucket = b.bucket
ORDER BY b.bucket`, baselineH+1, endH)
		args = []any{startStr, endStr, stepInterval, truncUnit, addrLower, stepInterval, baseline}
	} else {
		// No blocks in window → all buckets carry the baseline, zero delta.
		query = `
WITH buckets AS (
    SELECT generate_series(?::timestamp, ?::timestamp, ?::interval) AS bucket
)
SELECT
    EXTRACT(EPOCH FROM (bucket + ?::interval))::bigint AS ts,
    '0'::text                                          AS delta,
    ?::text                                            AS balance
FROM buckets
ORDER BY bucket`
		args = []any{startStr, endStr, stepInterval, stepInterval, baseline}
	}

	var rows []struct {
		Ts      int64
		Delta   string
		Balance string
	}
	if err := gdb.Raw(query, args...).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query balance history")
	}

	out := make([]BalanceHistoryPoint, 0, len(rows))
	for _, r := range rows {
		out = append(out, BalanceHistoryPoint{
			Timestamp: r.Ts,
			Balance:   trimNumericText(r.Balance),
			Delta:     trimNumericText(r.Delta),
		})
	}
	return out, nil
}

// resolveBaseline returns the addr's balance (rau, integer text) at the end
// of baselineH. If EthArchiveEndPoint is configured, it queries
// eth_getBalance against that archive node (constant-time). Otherwise it
// falls back to summing account_income — fine for fresh accounts but slow
// for old contracts.
func resolveBaseline(ctx context.Context, gdb *gorm.DB, addr string, baselineH int64) (string, error) {
	if endpoint := config.Default.EthArchiveEndPoint; endpoint != "" {
		bal, err := EthGetBalanceAtHeight(ctx, endpoint, addr, uint64(baselineH))
		if err != nil {
			return "", err
		}
		return bal.String(), nil
	}
	var s sql.NullString
	if err := gdb.Raw(
		`SELECT COALESCE(SUM(in_flow - out_flow), 0)::text
		 FROM account_income
		 WHERE address = ? AND block_height <= ?`,
		addr, baselineH,
	).Scan(&s).Error; err != nil {
		return "", err
	}
	if s.Valid && s.String != "" {
		return s.String, nil
	}
	return "0", nil
}

// trimNumericText drops a trailing ".0..." from a numeric::text result. Rau
// deltas are integer-valued, but PG numeric output occasionally carries a
// decimal tail depending on column scale.
func trimNumericText(s string) string {
	if i := strings.IndexByte(s, '.'); i >= 0 {
		s = s[:i]
	}
	if s == "" || s == "-" {
		return "0"
	}
	return s
}
