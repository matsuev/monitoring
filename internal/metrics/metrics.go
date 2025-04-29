package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HandlerCount ...
	HandlerCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_requests_total",
		Help: "The total number of HTTP requests",
	}, []string{"endpoint"})
)
