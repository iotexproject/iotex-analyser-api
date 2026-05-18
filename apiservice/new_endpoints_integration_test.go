package apiservice

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/api/pagination"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/stretchr/testify/require"
)

// setupRealPG loads PG credentials from env vars and connects. The test is
// skipped unless ITEST_DB_HOST is set, so this never runs in CI by accident.
//
//	ITEST_DB_HOST=65.108.2.142 ITEST_DB_PORT=5432 \
//	ITEST_DB_USER=analyser_testnet ITEST_DB_PASSWORD=... \
//	ITEST_DB_NAME=testnet \
//	go test ./apiservice -run TestNewEndpoints_Integration -v
func setupRealPG(t *testing.T) {
	t.Helper()
	host := os.Getenv("ITEST_DB_HOST")
	if host == "" {
		t.Skip("ITEST_DB_HOST not set; skipping integration test")
	}
	config.Default.Database.Driver = "postgres"
	config.Default.Database.Host = host
	config.Default.Database.Port = envOr("ITEST_DB_PORT", "5432")
	config.Default.Database.User = envOr("ITEST_DB_USER", "postgres")
	config.Default.Database.Password = os.Getenv("ITEST_DB_PASSWORD")
	config.Default.Database.Name = envOr("ITEST_DB_NAME", "testnet")
	_, err := db.Connect()
	require.NoError(t, err)
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func dumpJSON(t *testing.T, label string, v any) {
	t.Helper()
	b, err := json.MarshalIndent(v, "", "  ")
	require.NoError(t, err)
	t.Logf("%s: %s", label, string(b))
}

func TestNewEndpoints_Integration_GetChainStats(t *testing.T) {
	setupRealPG(t)
	svc := &ChainService{}
	resp, err := svc.GetChainStats(context.Background(), &api.GetChainStatsRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotEmpty(t, resp.TotalSupply, "total_supply should be non-empty")
	require.NotEmpty(t, resp.CirculatingSupply, "circulating_supply should be non-empty")
	dumpJSON(t, "GetChainStats", resp)
}

func TestNewEndpoints_Integration_GetTpsHistory(t *testing.T) {
	setupRealPG(t)
	svc := &ChainService{}
	resp, err := svc.GetTpsHistory(context.Background(), &api.GetTpsHistoryRequest{
		Start: "2026-05-01",
		End:   "2026-05-07",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	dumpJSON(t, "GetTpsHistory (2026-05-01..07)", resp)
	for _, p := range resp.Data {
		require.NotEmpty(t, p.Date)
		require.GreaterOrEqual(t, p.AvgTps, 0.0)
		require.GreaterOrEqual(t, p.MaxTps, p.AvgTps, "max should be >= avg on day %s", p.Date)
	}
}

func TestNewEndpoints_Integration_GetGasHistory(t *testing.T) {
	setupRealPG(t)
	svc := &ChainService{}
	resp, err := svc.GetGasHistory(context.Background(), &api.GetGasHistoryRequest{
		Start: "2026-05-01",
		End:   "2026-05-07",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	dumpJSON(t, "GetGasHistory (2026-05-01..07)", resp)
}

func TestNewEndpoints_Integration_GetSupplyHistory(t *testing.T) {
	setupRealPG(t)
	svc := &ChainService{}
	resp, err := svc.GetSupplyHistory(context.Background(), &api.GetSupplyHistoryRequest{
		Start: "2026-05-01",
		End:   "2026-05-07",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	dumpJSON(t, "GetSupplyHistory (2026-05-01..07)", resp)

	// 2023-11 had real burn activity — exercises the delta path with non-zero values.
	resp2, err := svc.GetSupplyHistory(context.Background(), &api.GetSupplyHistoryRequest{
		Start: "2023-11-01",
		End:   "2023-11-07",
	})
	require.NoError(t, err)
	require.NotNil(t, resp2)
	dumpJSON(t, "GetSupplyHistory (2023-11-01..07, burn period)", resp2)
}

func TestNewEndpoints_Integration_GetXRC20Stats(t *testing.T) {
	setupRealPG(t)
	svc := &XRC20Service{}
	resp, err := svc.GetXRC20Stats(context.Background(), &api.GetXRC20StatsRequest{
		Pagination: &pagination.Pagination{First: 10, Skip: 0},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	dumpJSON(t, "GetXRC20Stats (top 10 by holders)", resp)
	require.NotZero(t, resp.Count, "total xrc20 contract count should be > 0")
	for _, item := range resp.Items {
		require.NotEmpty(t, item.Address)
		require.Greater(t, item.Holders, uint64(0))
	}
}
