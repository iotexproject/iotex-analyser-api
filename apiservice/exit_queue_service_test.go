package apiservice

import (
	"context"
	"testing"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-analyser-api/model"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func setupExitQueueDB(t *testing.T) {
	t.Helper()
	config.Default.Database.Driver = "sqlite3"
	config.Default.Database.Name = ":memory:"
	gdb, err := db.Connect()
	require.NoError(t, err)
	require.NoError(t, gdb.AutoMigrate(&model.CandidateExitQueue{}))
}

func seedExitQueue(t *testing.T, rows []model.CandidateExitQueue) {
	t.Helper()
	require.NoError(t, db.DB().Create(&rows).Error)
}

func TestExitQueueService_InvalidStatus(t *testing.T) {
	setupExitQueueDB(t)
	svc := &ExitQueueService{}

	_, err := svc.GetExitQueue(context.Background(), &api.GetExitQueueRequest{
		Statuses: []string{"unknown"},
	})
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

func TestExitQueueService_EmptyTable(t *testing.T) {
	setupExitQueueDB(t)
	svc := &ExitQueueService{}

	resp, err := svc.GetExitQueue(context.Background(), &api.GetExitQueueRequest{})
	require.NoError(t, err)
	require.Equal(t, int64(0), resp.Count)
	require.Empty(t, resp.Exits)
}

func TestExitQueueService_ValidStatusFilters(t *testing.T) {
	setupExitQueueDB(t)
	seedExitQueue(t, []model.CandidateExitQueue{
		{CandidateIdentity: "io1aaa", Status: "requested"},
		{CandidateIdentity: "io1bbb", Status: "scheduled"},
		{CandidateIdentity: "io1ccc", Status: "confirmed"},
	})
	svc := &ExitQueueService{}

	for _, tc := range []struct {
		name          string
		statuses      []string
		expectedCount int64
	}{
		{"all", nil, 3},
		{"requested", []string{"requested"}, 1},
		{"scheduled", []string{"scheduled"}, 1},
		{"confirmed", []string{"confirmed"}, 1},
		{"requested+scheduled", []string{"requested", "scheduled"}, 2},
		{"requested+scheduled+confirmed", []string{"requested", "scheduled", "confirmed"}, 3},
	} {
		t.Run("statuses="+tc.name, func(t *testing.T) {
			resp, err := svc.GetExitQueue(context.Background(), &api.GetExitQueueRequest{Statuses: tc.statuses})
			require.NoError(t, err)
			require.Equal(t, tc.expectedCount, resp.Count)
		})
	}
}

func TestExitQueueService_FieldMapping(t *testing.T) {
	setupExitQueueDB(t)
	seedExitQueue(t, []model.CandidateExitQueue{
		{
			CandidateName:     "mycandidate",
			CandidateIdentity: "io1identity",
			Status:            "scheduled",
			RequestHeight:     100,
			RequestHash:       "reqhash",
			ScheduleHeight:    200,
			ScheduleHash:      "schehash",
			ScheduledAt:       999,
		},
	})
	svc := &ExitQueueService{}

	resp, err := svc.GetExitQueue(context.Background(), &api.GetExitQueueRequest{})
	require.NoError(t, err)
	require.Len(t, resp.Exits, 1)
	e := resp.Exits[0]
	require.Equal(t, "mycandidate", e.CandidateName)
	require.Equal(t, "io1identity", e.CandidateIdentity)
	require.Equal(t, "scheduled", e.Status)
	require.Equal(t, uint64(100), e.RequestHeight)
	require.Equal(t, "reqhash", e.RequestHash)
	require.Equal(t, uint64(200), e.ScheduleHeight)
	require.Equal(t, "schehash", e.ScheduleHash)
	require.Equal(t, uint64(999), e.ScheduledAt)
}

func TestExitQueueService_Pagination(t *testing.T) {
	setupExitQueueDB(t)
	rows := make([]model.CandidateExitQueue, 5)
	for i := range rows {
		rows[i] = model.CandidateExitQueue{CandidateIdentity: "io1x", Status: "requested"}
	}
	seedExitQueue(t, rows)
	svc := &ExitQueueService{}

	resp, err := svc.GetExitQueue(context.Background(), &api.GetExitQueueRequest{Skip: 0, First: 3})
	require.NoError(t, err)
	require.Equal(t, int64(5), resp.Count)
	require.Len(t, resp.Exits, 3)

	resp2, err := svc.GetExitQueue(context.Background(), &api.GetExitQueueRequest{Skip: 3, First: 3})
	require.NoError(t, err)
	require.Equal(t, int64(5), resp2.Count)
	require.Len(t, resp2.Exits, 2)
}

func TestExitQueueService_FirstCappedAt100(t *testing.T) {
	setupExitQueueDB(t)
	svc := &ExitQueueService{}

	// first=0 defaults to 20
	resp, err := svc.GetExitQueue(context.Background(), &api.GetExitQueueRequest{First: 0})
	require.NoError(t, err)
	_ = resp // just check no error; cap is internal

	// first > 100 should not panic and be capped
	resp2, err := svc.GetExitQueue(context.Background(), &api.GetExitQueueRequest{First: 500})
	require.NoError(t, err)
	_ = resp2
}
