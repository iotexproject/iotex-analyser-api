package apiservice

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestParseDateRange(t *testing.T) {
	for _, tc := range []struct {
		name      string
		start     string
		end       string
		wantErr   bool
		wantCode  codes.Code
		wantField string // substring expected in error message
	}{
		{"valid same day", "2026-05-01", "2026-05-01", false, codes.OK, ""},
		{"valid range", "2026-05-01", "2026-05-07", false, codes.OK, ""},
		{"empty start", "", "2026-05-07", true, codes.InvalidArgument, "start"},
		{"empty end", "2026-05-01", "", true, codes.InvalidArgument, "end"},
		{"malformed start", "2026/05/01", "2026-05-07", true, codes.InvalidArgument, "start"},
		{"malformed end", "2026-05-01", "May 7", true, codes.InvalidArgument, "end"},
		{"start after end", "2026-05-08", "2026-05-01", true, codes.InvalidArgument, "after"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, _, err := parseDateRange(tc.start, tc.end)
			if !tc.wantErr {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			st, ok := status.FromError(err)
			require.True(t, ok, "want gRPC status error")
			require.Equal(t, tc.wantCode, st.Code())
			require.Contains(t, st.Message(), tc.wantField)
		})
	}
}
