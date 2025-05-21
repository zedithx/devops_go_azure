package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	StatusHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "status_endpoint_hits_total",
			Help: "Total number of calls to /status endpoint",
		},
	)
)

func InitMetrics() {
	prometheus.MustRegister(StatusHits)
}