package apiservice

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	dto "github.com/prometheus/client_model/go"
)

func TestInstrumentObservesRequestDuration(t *testing.T) {
	httpRequestDuration.Reset()

	cases := []struct {
		handler   string
		method    string
		writeCode int
	}{
		{"api", "GET", 0},      // default 200
		{"graphql", "POST", 404},
		{"api", "POST", 502},
		{"docs", "GET", 301},
	}

	for _, tc := range cases {
		inner := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			if tc.writeCode != 0 {
				w.WriteHeader(tc.writeCode)
			}
		})
		req := httptest.NewRequest(tc.method, "/anything", nil)
		rr := httptest.NewRecorder()
		instrument(tc.handler, inner).ServeHTTP(rr, req)
	}

	// Four distinct (handler, method, status_class) tuples ⇒ four series in
	// the histogram.
	if got, want := testutil.CollectAndCount(httpRequestDuration), 4; got != want {
		t.Errorf("got %d series in httpRequestDuration, want %d", got, want)
	}

	// Spot-check one combination through the dto.Metric to confirm the
	// expected (handler=api, method=GET, status_class=2xx) tuple was the one
	// that recorded the default-200 case.
	m := &dto.Metric{}
	if err := httpRequestDuration.WithLabelValues("api", "GET", "2xx").(prometheus.Histogram).Write(m); err != nil {
		t.Fatalf("Write: %v", err)
	}
	if got := m.Histogram.GetSampleCount(); got != 1 {
		t.Errorf("api/GET/2xx sample count = %d, want 1", got)
	}
}

// TestInstrumentDuplicateWriteHeader pins net/http semantics: when a handler
// calls WriteHeader twice, the client only sees the first status, and the
// metric label must match. Regression guard for the original code that
// overwrote `status` on every call.
func TestInstrumentDuplicateWriteHeader(t *testing.T) {
	httpRequestDuration.Reset()

	// Inner handler writes 200 first, then "tries" to write 500.
	inner := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.WriteHeader(http.StatusInternalServerError)
	})
	req := httptest.NewRequest("GET", "/anything", nil)
	rr := httptest.NewRecorder()
	instrument("api", inner).ServeHTTP(rr, req)

	// Expect a 2xx observation, not 5xx.
	m := &dto.Metric{}
	if err := httpRequestDuration.WithLabelValues("api", "GET", "2xx").(prometheus.Histogram).Write(m); err != nil {
		t.Fatalf("Write: %v", err)
	}
	if got := m.Histogram.GetSampleCount(); got != 1 {
		t.Errorf("api/GET/2xx sample count = %d, want 1 (first WriteHeader wins)", got)
	}
	if got := testutil.CollectAndCount(httpRequestDuration); got != 1 {
		t.Errorf("got %d series, want exactly 1 (no spurious 5xx series)", got)
	}
}

func TestStatusClass(t *testing.T) {
	for _, tc := range []struct {
		code int
		want string
	}{
		{200, "2xx"}, {201, "2xx"}, {301, "3xx"},
		{404, "4xx"}, {500, "5xx"}, {0, "unknown"}, {999, "unknown"},
	} {
		if got := statusClass(tc.code); got != tc.want {
			t.Errorf("statusClass(%d) = %q, want %q", tc.code, got, tc.want)
		}
	}
}
