package diag

import "runtime/metrics"

var runtimeMetricByName = func() map[string]metrics.Description {
	m := make(map[string]metrics.Description)
	for _, d := range metrics.All() {
		m[d.Name] = d
	}
	return m
}()

type RuntimeMetric struct {
	Desc  metrics.Description
	Value any
}

func medianBucket(h *metrics.Float64Histogram) float64 {
	total := uint64(0)
	for _, count := range h.Counts {
		total += count
	}
	thresh := total / 2
	total = 0
	for i, count := range h.Counts {
		total += count
		if total >= thresh {
			return h.Buckets[i]
		}
	}
	panic("should not happen")
}

func ReadRuntimeMetrics() []RuntimeMetric {
	descs := metrics.All()

	samples := make([]metrics.Sample, len(descs))
	for i := range samples {
		samples[i].Name = descs[i].Name
	}

	metrics.Read(samples)

	r := make([]RuntimeMetric, 0, len(descs))
	for i, sample := range samples {
		var v any
		value := sample.Value
		switch value.Kind() {
		case metrics.KindUint64:
			v = value.Uint64()
		case metrics.KindFloat64:
			v = value.Float64()
		case metrics.KindFloat64Histogram:
			v = value.Float64Histogram()
		default:
			continue
		}

		r = append(r, RuntimeMetric{Desc: descs[i], Value: v})
	}
	return r
}
