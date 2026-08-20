package apiservice

import (
	"testing"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/stretchr/testify/require"
)

// unifiedFetchLimit decides how much of each leg UnifiedVoterRewards pulls.
// Getting it wrong is not a rounding error: too small silently drops rows from
// the page, and unbounded reinstates the full-history fetch that defeats the
// LIMIT on the Hermes join.
func TestUnifiedFetchLimit(t *testing.T) {
	// A page needs offset+limit rows from each leg, because a row beyond that
	// position in its own source cannot be promoted into the page by the merge.
	require.Equal(t, 20, unifiedFetchLimit(0, 20))
	require.Equal(t, 120, unifiedFetchLimit(100, 20))

	// common.PageSize returns 0 when the caller sends a pagination message with
	// no `first`. gorm renders Limit(0) as no LIMIT, so this is exactly the
	// case that must fall back to the cap rather than to "everything".
	require.Equal(t, maxUnifiedFetch, unifiedFetchLimit(0, 0))
	require.Equal(t, maxUnifiedFetch, unifiedFetchLimit(50, 0))

	// Deep paging clamps to the cap instead of growing without bound.
	require.Equal(t, maxUnifiedFetch, unifiedFetchLimit(maxUnifiedFetch, 20))
	require.Equal(t, maxUnifiedFetch, unifiedFetchLimit(maxUnifiedFetch*10, 20))

	// offset+limit must not overflow into a small (or negative) limit.
	require.Equal(t, maxUnifiedFetch, unifiedFetchLimit(int(^uint(0)>>1), 20))
	require.Equal(t, maxUnifiedFetch, unifiedFetchLimit(-1, 20))

	// Never returns something unusable as a SQL LIMIT.
	for _, tc := range [][2]int{{0, 0}, {0, 20}, {100, 20}, {1 << 30, 1 << 30}, {-5, -5}} {
		require.Greater(t, unifiedFetchLimit(tc[0], tc[1]), 0,
			"offset=%d limit=%d", tc[0], tc[1])
	}
}

// TestMaxUnifiedEpochSpanPreventsOverflow guards the arithmetic in
// UnifiedVoterRewards: endEpoch = startEpoch + epochCount - 1 wraps for a large
// epochCount, and a wrapped endEpoch produces `epoch >= start AND epoch <= less`
// — an empty result that reads as "you earned nothing".
func TestMaxUnifiedEpochSpanPreventsOverflow(t *testing.T) {
	// With epochCount capped, the only way to overflow is an extreme
	// start_epoch, which the explicit endEpoch < startEpoch check rejects.
	const maxStart = ^uint64(0) - maxUnifiedEpochSpan
	require.Greater(t, maxStart+maxUnifiedEpochSpan-1, maxStart,
		"cap must leave headroom for any start_epoch below the wrap point")

	// And the cap is loose enough not to reject a realistic window: IoTeX runs
	// 24 epochs a day.
	require.Greater(t, uint64(maxUnifiedEpochSpan), uint64(24*365*10),
		"cap should comfortably exceed a decade of epochs")
}

// TestRewardSourceZeroIsUnspecified pins the enum contract. proto3 gives every
// unset enum field the zero value, so a real pipeline sitting on 0 would make
// an omitted source read as a positive claim, and an empty `sources` list
// indistinguishable from one naming that pipeline.
func TestRewardSourceZeroIsUnspecified(t *testing.T) {
	require.Equal(t, api.RewardSource(0), api.RewardSource_REWARD_SOURCE_UNSPECIFIED)
	require.NotEqual(t, api.RewardSource(0), api.RewardSource_HERMES_OFFCHAIN)
	require.NotEqual(t, api.RewardSource(0), api.RewardSource_ONCHAIN_IIP59)

	// The names are what JSON and GraphQL carry, and downstream reads them as
	// strings — they must not drift.
	require.Equal(t, "HERMES_OFFCHAIN", api.RewardSource_HERMES_OFFCHAIN.String())
	require.Equal(t, "ONCHAIN_IIP59", api.RewardSource_ONCHAIN_IIP59.String())
}
