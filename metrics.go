package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/common/version"
)

type ServiceMetrics struct {
	Decisions	prometheus.Counter
}

func NewServiceMetrics() ServiceMetrics {
	return ServiceMetrics{
		Decisions: promauto.NewCounter(prometheus.CounterOpts{
			Name: "decisions_total",
			Help: "The total number of decisions logged",
		}),
	}
}

func initMetrics() {
	prometheus.MustRegister(version.NewCollector("opa_audit_logging_service"))
}