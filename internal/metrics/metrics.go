package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// OpsProcessed ...
	OpsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	HandlerCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_http_requests_total",
		Help: "The total number of HTTP requests",
	}, []string{"endpoint"})
)
