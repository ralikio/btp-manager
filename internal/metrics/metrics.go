package metrics

import (
	"github.com/kyma-project/btp-manager/internal/conditions"
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type Metrics struct {
	WorkqueueSizeGauge prometheus.Gauge
	ReasonCounters     map[conditions.Reason]prometheus.Counter
}

func (m *Metrics) registerMetrics() {
	m.WorkqueueSizeGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "workqueue_size",
			Help: "Size of work queue",
		},
	)
	metrics.Registry.MustRegister(m.WorkqueueSizeGauge)

	reasonsCounters := make(map[conditions.Reason]prometheus.Counter, len(conditions.Reasons))
	for reason, metadata := range conditions.Reasons {
		counter := prometheus.NewCounter(prometheus.CounterOpts{
			Name:        string(reason),
			ConstLabels: prometheus.Labels{"state": string(metadata.State)},
		})
		reasonsCounters[reason] = counter
		metrics.Registry.MustRegister(counter)
	}
	m.ReasonCounters = reasonsCounters
}

func InitMetrics() *Metrics {
	metrics := &Metrics{}
	metrics.registerMetrics()
	return metrics
}
