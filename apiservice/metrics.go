package apiservice

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// httpRequestDuration is a HistogramVec keyed by (handler, method, status_class)
// instead of the full status code, to keep label cardinality bounded.
var httpRequestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "analyser_api_http_request_duration_seconds",
		Help:    "Latency of HTTP requests handled by analyser-api, by handler mount point.",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"handler", "method", "status_class"},
)

// statusRecorder captures the response status code as it is written, so the
// middleware can label the histogram observation with it.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// instrument wraps h with request-duration tracking under the given handler
// name (e.g. "api", "graphql", "docs"). The label "status_class" is the
// hundreds digit of the response status code ("2xx", "4xx", ...). Unwritten
// statuses default to 200 (net/http's implicit behavior).
func instrument(name string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		start := time.Now()
		h.ServeHTTP(rec, r)
		httpRequestDuration.
			WithLabelValues(name, r.Method, statusClass(rec.status)).
			Observe(time.Since(start).Seconds())
	})
}

func statusClass(code int) string {
	if code < 100 || code >= 600 {
		return "unknown"
	}
	return fmt.Sprintf("%sxx", strconv.Itoa(code/100))
}
