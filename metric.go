package promsarama

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rcrowley/go-metrics"
)

type Meter struct {
	gauge prometheus.Gauge
}

func (m *Meter) Count() int64            { return -1 }
func (m *Meter) Rate1() float64          { return -1.0 }
func (m *Meter) Rate5() float64          { return -1.0 }
func (m *Meter) Rate15() float64         { return -1.0 }
func (m *Meter) RateMean() float64       { return -1.0 }
func (m *Meter) Snapshot() metrics.Meter { return metrics.Meter(m) }
func (m *Meter) Stop()                   {}

func (m *Meter) Mark(v int64) {
	m.gauge.Set(float64(v))
}

var _ metrics.Meter = (*Meter)(nil)

type Counter struct {
	gauge prometheus.Gauge
}

func (c *Counter) Count() int64              { return -1 }
func (c *Counter) Clear()                    {}
func (c *Counter) Snapshot() metrics.Counter { return metrics.Counter(c) }
func (c *Counter) Dec(v int64) {
	c.gauge.Sub(float64(v))
}
func (c *Counter) Inc(v int64) {
	c.gauge.Add(float64(v))
}

var _ metrics.Counter = (*Counter)(nil)

type Histogram struct {
	histogram prometheus.Histogram
}

func (h *Histogram) Clear()                          {}
func (h *Histogram) Count() int64                    { return -1 }
func (h *Histogram) Max() int64                      { return -1 }
func (h *Histogram) Mean() float64                   { return -1 }
func (h *Histogram) Min() int64                      { return -1 }
func (h *Histogram) Percentile(float64) float64      { return -1.0 }
func (h *Histogram) Percentiles([]float64) []float64 { return nil }
func (h *Histogram) Sample() metrics.Sample          { return nil }
func (h *Histogram) Snapshot() metrics.Histogram     { return nil }
func (h *Histogram) StdDev() float64                 { return -1.0 }
func (h *Histogram) Sum() int64                      { return -1 }
func (h *Histogram) Variance() float64               { return -1.0 }

func (h *Histogram) Update(v int64) {
	h.histogram.Observe(float64(v))
}

var _ metrics.Histogram = (*Histogram)(nil)
