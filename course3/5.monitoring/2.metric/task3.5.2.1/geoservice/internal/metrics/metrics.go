package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	LoginCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Login Counter",
		Help: "Total number of logins",
	})
	RegisterCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Register Counter",
		Help: "Total number of registrations",
	})
	LoginDurationCounter = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "login duration seconds",
		Help:    "Login duration in seconds",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})
	RegisterDurationCounter = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "Register duration seconds",
		Help:    "Register duration in seconds",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})
)
