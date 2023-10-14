package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

var StarCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "star_todo_count",
	},
)

func NewMonitoryRegistry() prometheus.Gatherer {
	registry := prometheus.NewRegistry()
	registry.MustRegister(StarCounter)

	return registry
}
