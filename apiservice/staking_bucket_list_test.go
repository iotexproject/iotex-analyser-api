package apiservice

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNativeBucketListQuery(t *testing.T) {
	require := require.New(t)

	// timestamp sorts on block_height, with id breaking ties within a block; the page of
	// ids is picked in a subquery and the outer query keeps the same order.
	q, args := nativeBucketListQuery("", nil, "staking_buckets.timestamp", "DESC", 20, 40)
	require.Contains(q, "FROM (SELECT staking_buckets.id FROM staking_buckets ORDER BY staking_buckets.block_height DESC, staking_buckets.id DESC LIMIT ? OFFSET ?) page")
	require.True(strings.HasSuffix(q, "ORDER BY staking_buckets.block_height DESC, staking_buckets.id DESC"), q)
	require.Equal([]interface{}{int64(20), int64(40)}, args)

	q, args = nativeBucketListQuery("staking_buckets.block_height >= ?", []interface{}{int64(53000000)}, "staking_buckets.timestamp", "ASC", 20, 0)
	require.Contains(q, "FROM (SELECT staking_buckets.id FROM staking_buckets WHERE staking_buckets.block_height >= ? ORDER BY staking_buckets.block_height ASC, staking_buckets.id ASC LIMIT ? OFFSET ?) page")
	require.Equal([]interface{}{int64(53000000), int64(20), int64(0)}, args)

	// Other sort fields are passed through unchanged.
	q, args = nativeBucketListQuery("staking_buckets.timestamp > ?", []interface{}{int64(1761000000)}, "staking_buckets.amount", "DESC", 10, 0)
	require.Contains(q, "FROM (SELECT staking_buckets.id FROM staking_buckets WHERE staking_buckets.timestamp > ? ORDER BY staking_buckets.amount DESC LIMIT ? OFFSET ?) page")
	require.True(strings.HasSuffix(q, "ORDER BY staking_buckets.amount DESC"), q)
	require.Equal([]interface{}{int64(1761000000), int64(10), int64(0)}, args)
}

func TestParseBucketSortParam(t *testing.T) {
	for _, tc := range []struct {
		sort, field, order string
	}{
		{"", "staking_buckets.timestamp", "DESC"},
		{"timestamp:desc", "staking_buckets.timestamp", "DESC"},
		{"timestamp:asc", "staking_buckets.timestamp", "ASC"},
		{"amount:asc", "staking_buckets.amount", "ASC"},
		{"unknown:asc", "staking_buckets.timestamp", "ASC"},
		{"amount; DROP TABLE x:desc", "staking_buckets.timestamp", "DESC"},
	} {
		field, order := parseBucketSortParam(tc.sort)
		require.Equal(t, tc.field, field, "sort=%q", tc.sort)
		require.Equal(t, tc.order, order, "sort=%q", tc.sort)
	}
}

// iotexscan fetches the count and group count with an empty sort; that request must
// get the block_height order too, not a full sort on timestamp.
func TestNativeBucketListQueryDefaultSort(t *testing.T) {
	for _, sort := range []string{"", "unknown:desc"} {
		field, order := parseBucketSortParam(sort)
		q, _ := nativeBucketListQuery("", nil, field, order, 1, 0)
		require.Contains(t, q, "ORDER BY staking_buckets.block_height DESC, staking_buckets.id DESC LIMIT ? OFFSET ?", "sort=%q", sort)
		require.NotContains(t, q, "ORDER BY timestamp", "sort=%q", sort)
	}
}

func TestNativeBucketCountQueries(t *testing.T) {
	require := require.New(t)
	cond, condArgs := "staking_buckets.block_height >= ?", []interface{}{int64(53000000)}

	q, args := nativeBucketCountQuery("", nil)
	require.Equal("SELECT COUNT(1) AS count FROM staking_buckets", q)
	require.Empty(args)
	q, args = nativeBucketCountQuery(cond, condArgs)
	require.Equal("SELECT COUNT(1) AS count FROM staking_buckets WHERE staking_buckets.block_height >= ?", q)
	require.Equal(condArgs, args)

	q, args = nativeBucketGroupCountQuery("", nil)
	require.Contains(q, "WITH RECURSIVE")
	require.Empty(args)
	q, args = nativeBucketGroupCountQuery(cond, condArgs)
	require.Equal("SELECT COUNT(DISTINCT bucket_id) AS count FROM staking_buckets WHERE staking_buckets.block_height >= ?", q)
	require.Equal(condArgs, args)
}
