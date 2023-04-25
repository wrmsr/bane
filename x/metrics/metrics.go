package metrics

import bt "github.com/wrmsr/bane/core/types"

//

type Tag = bt.Kv[string, string]
type Tags = []Tag

type MetricKind int8

const (
	InvalidKind MetricKind = iota
	CountKind
	RateKind
	GaugeKind
	HistogramKind
)

type Metric struct {
	Kind  MetricKind
	Name  string
	Value float64
	Tags  Tags
}

func MakeMetric(kind MetricKind, name string, value float64, tags Tags) Metric {
	return Metric{
		Kind:  kind,
		Name:  name,
		Value: value,
		Tags:  tags,
	}
}

//

type Metrics interface {
	EmitMetric(m Metric)
}

//

func Count(mtr Metrics, name string, value float64, tags Tags) {
	mtr.EmitMetric(MakeMetric(CountKind, name, value, tags))
}

func Rate(mtr Metrics, name string, value float64, tags Tags) {
	mtr.EmitMetric(MakeMetric(RateKind, name, value, tags))
}

func Gauge(mtr Metrics, name string, value float64, tags Tags) {
	mtr.EmitMetric(MakeMetric(GaugeKind, name, value, tags))
}

func Histogram(mtr Metrics, name string, value float64, tags Tags) {
	mtr.EmitMetric(MakeMetric(HistogramKind, name, value, tags))
}

//

type NopMetrics struct{}

var _ Metrics = NopMetrics{}

func (n NopMetrics) EmitMetric(m Metric) {}
