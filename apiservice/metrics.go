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

// statusRecorder captures the response status code so the middleware can
// label the histogram observation with it. Mirrors net/http semantics:
// only the first WriteHeader call wins; subsequent calls are passed through
// to the underlying ResponseWriter (which logs "superfluous WriteHeader" and
// ignores them) but do not update the recorded status.
type statusRecorder struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func (r *statusRecorder) WriteHeader(code int) {
	if !r.wroteHeader {
		r.status = code
		r.wroteHeader = true
	}
	r.ResponseWriter.WriteHeader(code)
}

// instrument wraps h with request-duration tracking under the given handler
// name (e.g. "api", "graphql", "docs"). The label "status_class" is the
// hundreds digit of the response status code ("2xx", "4xx", ...). If the
// handler never calls WriteHeader (net/http implicitly sends 200), the
// observation is labeled "2xx".
func instrument(name string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w}
		start := time.Now()
		h.ServeHTTP(rec, r)
		status := rec.status
		if status == 0 {
			status = http.StatusOK
		}
		httpRequestDuration.
			WithLabelValues(name, r.Method, statusClass(status)).
			Observe(time.Since(start).Seconds())
	})
}

func statusClass(code int) string {
	if code < 100 || code >= 600 {
		return "unknown"
	}
	return fmt.Sprintf("%sxx", strconv.Itoa(code/100))
}
