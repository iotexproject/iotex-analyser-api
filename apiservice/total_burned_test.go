package apiservice

import (
	"context"
	"testing"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/stretchr/testify/require"
)

func setupStoreDB(t *testing.T) {
	t.Helper()
	config.Default.Database.Driver = "sqlite3"
	config.Default.Database.Name = ":memory:"
	_, err := db.Connect()
	require.NoError(t, err)
	require.NoError(t, db.DB().Exec(`CREATE TABLE store (key TEXT, value TEXT)`).Error)
}

func TestGetTotalBurned_Empty(t *testing.T) {
	setupStoreDB(t)
	svc := &ChainService{}

	resp, err := svc.GetTotalBurned(context.Background(), &api.GetTotalBurnedRequest{})
	require.NoError(t, err)
	require.Equal(t, "", resp.GetTotalBurned())
	require.Equal(t, "", resp.GetTotalBurnedRau())
	require.Equal(t, uint64(0), resp.GetAsOfBlockHeight())
}

func TestGetTotalBurned_Value(t *testing.T) {
	setupStoreDB(t)
	require.NoError(t, db.DB().Exec(
		`INSERT INTO store (key, value) VALUES (?, ?)`,
		"total_burned_gasfee",
		`{"block_height":43638689,"amount":"6640079577070144894752174"}`,
	).Error)

	svc := &ChainService{}
	resp, err := svc.GetTotalBurned(context.Background(), &api.GetTotalBurnedRequest{})
	require.NoError(t, err)
	require.Equal(t, "6640079577070144894752174", resp.GetTotalBurnedRau())
	require.Equal(t, "6640079.58", resp.GetTotalBurned()) // rau / 1e18, 2dp
	require.Equal(t, uint64(43638689), resp.GetAsOfBlockHeight())
}
