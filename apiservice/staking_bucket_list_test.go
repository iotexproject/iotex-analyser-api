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
