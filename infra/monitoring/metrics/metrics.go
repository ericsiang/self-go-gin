package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsCollector struct {
	// HTTP 請求指標
	HttpRequestTotal    *prometheus.CounterVec
	HttpRequestDuration *prometheus.HistogramVec
	HttpRequestSize     *prometheus.SummaryVec
	HttpResponseSize    *prometheus.SummaryVec

	// 業務指標
	BusinessEvents         *prometheus.CounterVec
	DatabaseQueryDuration  *prometheus.HistogramVec
	CacheOperationDuration *prometheus.HistogramVec

	// 系統指標
	GoroutineCount prometheus.Gauge
	MemoryUsage    prometheus.Gauge
	CPUUsage       prometheus.Gauge
}

func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{
		HttpRequestTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "path", "status", "api_version"},
		),
		HttpRequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "HTTP request duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method", "path", "status"},
		),
		BusinessEvents: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "business_events_total",
				Help: "Total number of business events",
			},
			[]string{"event_type", "domain", "status"},
		),
		DatabaseQueryDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "database_query_duration_seconds",
				Help:    "Database query duration in seconds",
				Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
			},
			[]string{"operation", "table", "success"},
		),
		GoroutineCount: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "go_goroutines",
				Help: "Number of goroutines",
			},
		),
		MemoryUsage: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "go_memory_usage_bytes",
				Help: "Memory usage in bytes",
			},
		),
	}
}

// RecordHttpRequest 記錄 HTTP 請求指標
func (m *MetricsCollector) RecordHttpRequest(method, path, status, apiVersion string, duration float64) {
	m.HttpRequestTotal.WithLabelValues(method, path, status, apiVersion).Inc()
	m.HttpRequestDuration.WithLabelValues(method, path, status).Observe(duration)
}

// RecordBusinessEvent 記錄業務事件
func (m *MetricsCollector) RecordBusinessEvent(eventType, domain, status string) {
	m.BusinessEvents.WithLabelValues(eventType, domain, status).Inc()
}

// RecordDatabaseQuery 記錄資料庫查詢
func (m *MetricsCollector) RecordDatabaseQuery(operation, table string, success bool, duration float64) {
	m.DatabaseQueryDuration.WithLabelValues(operation, table, boolToString(success)).Observe(duration)
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
