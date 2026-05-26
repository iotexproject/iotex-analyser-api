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
