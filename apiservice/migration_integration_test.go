package apiservice

// Integration tests for the RPCs added for the iotex-kit ANALYZER_DATABASE_URL
// migration (IotexscanService + the delegate/staking migration handlers).
//
// Skipped unless ITEST_DB_HOST is set, same as new_endpoints_integration_test.go,
// so it never runs in CI by accident. Reuses setupRealPG/envOr/dumpJSON from
// that file (same package). These assert the SQL executes and the response
// shape is populated — they don't pin exact values (data drifts).
//
//	ITEST_DB_HOST=65.108.2.142 ITEST_DB_PORT=5432 ITEST_DB_USER=analyser \
//	ITEST_DB_PASSWORD=... ITEST_DB_NAME=mainnet \
//	go test ./apiservice -run TestMigration_Integration -v

import (
	"context"
	"testing"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/stretchr/testify/require"
)

func TestMigration_Integration_GetBlockNumberByTime(t *testing.T) {
	setupRealPG(t)
	svc := &IotexscanService{}
	before, err := svc.GetBlockNumberByTime(context.Background(), &api.GetBlockNumberByTimeRequest{
		Timestamp: 1700000000, Closest: "before",
	})
	require.NoError(t, err)
	require.True(t, before.Exist)
	require.Greater(t, before.BlockHeight, uint64(0))

	after, err := svc.GetBlockNumberByTime(context.Background(), &api.GetBlockNumberByTimeRequest{
		Timestamp: 1700000000, Closest: "after",
	})
	require.NoError(t, err)
	require.True(t, after.Exist)
	// The "after" block must be strictly greater than the "before" block.
	require.Greater(t, after.BlockHeight, before.BlockHeight)
	dumpJSON(t, "GetBlockNumberByTime before/after", map[string]uint64{"before": before.BlockHeight, "after": after.BlockHeight})
}

func TestMigration_Integration_GetDailyNewAddresses(t *testing.T) {
	setupRealPG(t)
	svc := &IotexscanService{}
	resp, err := svc.GetDailyNewAddresses(context.Background(), &api.GetDailyNewAddressesRequest{
		StartDate: "2024-01-01", EndDate: "2024-01-03", Sort: "asc",
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Data, "expected new-address rows in range")
	for _, p := range resp.Data {
		require.NotEmpty(t, p.Date)
	}
	dumpJSON(t, "GetDailyNewAddresses", resp)
}

func TestMigration_Integration_GetActionStatusByHash_NotFound(t *testing.T) {
	setupRealPG(t)
	svc := &IotexscanService{}
	// A hash that won't exist: exercises the empty path without a fixed tx.
	resp, err := svc.GetActionStatusByHash(context.Background(), &api.GetActionStatusByHashRequest{
		ActionHash: "0000000000000000000000000000000000000000000000000000000000000000",
	})
	require.NoError(t, err)
	require.False(t, resp.Exist)
}

func TestMigration_Integration_GetContractCreationBatch_Empty(t *testing.T) {
	setupRealPG(t)
	svc := &IotexscanService{}
	// Empty input short-circuits before any query.
	resp, err := svc.GetContractCreationBatch(context.Background(), &api.GetContractCreationBatchRequest{})
	require.NoError(t, err)
	require.Empty(t, resp.Items)

	// Non-empty: IN (?) must render correctly (regression for the ANY(?)
	// malformed-array-literal bug). An EOA/absent address just yields no rows.
	resp2, err := svc.GetContractCreationBatch(context.Background(), &api.GetContractCreationBatchRequest{
		Addresses: []string{"io15qr5fzpxsnp7garl4m7k355rafzqn8grrm0grz"},
	})
	require.NoError(t, err, "IN (?) with a slice must not error")
	dumpJSON(t, "GetContractCreationBatch", resp2)
}

func TestMigration_Integration_GetDelegatesStatistics(t *testing.T) {
	setupRealPG(t)
	svc := &DelegateService{}
	resp, err := svc.GetDelegatesStatistics(context.Background(), &api.GetDelegatesStatisticsRequest{})
	require.NoError(t, err)
	require.True(t, resp.Exist)
	require.Greater(t, resp.DelegateCount, uint64(0))
	require.NotEmpty(t, resp.TotalAmount)
	dumpJSON(t, "GetDelegatesStatistics", resp)
}

func TestMigration_Integration_GetStakingHistory(t *testing.T) {
	setupRealPG(t)
	svc := &StakingService{}
	resp, err := svc.GetStakingHistory(context.Background(), &api.GetStakingHistoryRequest{
		OwnerAddress: "io15qr5fzpxsnp7garl4m7k355rafzqn8grrm0grz",
		Page:         1,
		Limit:        5,
	})
	require.NoError(t, err)
	// total is a COUNT(DISTINCT); data length is bounded by the page limit.
	require.LessOrEqual(t, uint64(len(resp.Data)), uint64(5))
	dumpJSON(t, "GetStakingHistory", resp)
}

func TestMigration_Integration_GetReceivedVotesByAddress(t *testing.T) {
	setupRealPG(t)
	svc := &DelegateService{}
	resp, err := svc.GetReceivedVotesByAddress(context.Background(), &api.GetReceivedVotesByAddressRequest{
		Address: "io15qr5fzpxsnp7garl4m7k355rafzqn8grrm0grz",
	})
	require.NoError(t, err)
	dumpJSON(t, "GetReceivedVotesByAddress", resp)
}

func TestMigration_Integration_GetProductivityHistory_Validation(t *testing.T) {
	setupRealPG(t)
	svc := &DelegateService{}
	// Missing end_date -> InvalidArgument (validation path, no DB row needed).
	_, err := svc.GetProductivityHistory(context.Background(), &api.GetProductivityHistoryRequest{
		Candidate: "io15qr5fzpxsnp7garl4m7k355rafzqn8grrm0grz",
	})
	require.Error(t, err)
}
