package metric

import (
	"github.com/prometheus/client_golang/prometheus"

	prom "github.com/kyh0703/go-project-layout/pkg/prometheus"
)

var instance = new(Handler)

type Handler struct{}

func ProvideMetricHandler() Metric {
	return instance
}

func (*Handler) IncreaseRequestCount() {
	label := prometheus.Labels{}
	prom.RequestCount.With(label).Inc()
}

func (*Handler) IncreaseSuccessCount() {
	label := prometheus.Labels{}
	prom.SuccessCount.With(label).Inc()
}

func (*Handler) IncreaseFailCount() {
	label := prometheus.Labels{}
	prom.FailCount.With(label).Inc()
}

func (*Handler) IncreaseTx(tenant, txType string) {
}

func (*Handler) DecreaseTx(tenant, txType string) {
}
