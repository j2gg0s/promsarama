package promsarama

import (
	"reflect"
	"strings"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rcrowley/go-metrics"
)

type Registry struct {
	m  sync.Map
	mu sync.Mutex

	Registerer prometheus.Registerer
	Buckets    []float64
}

func (r *Registry) Each(func(string, interface{})) {}

func (r *Registry) Get(string) interface{} {
	return nil
}

func (r *Registry) GetAll() map[string]map[string]interface{} {
	return nil
}

func (r *Registry) Register(string, interface{}) error {
	return nil
}

func (r *Registry) RunHealthchecks() {}

func (r *Registry) Unregister(string) {}

func (r *Registry) UnregisterAll() {}

func (r *Registry) GetOrRegister(name string, f interface{}) interface{} {
	name = strings.ReplaceAll(name, "-", "_")

	if v, ok := r.m.Load(name); ok {
		return v
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func || t.NumOut() != 1 {
		return nil
	}

	var v interface{}
	if t.Out(0).Implements(reflect.TypeOf((*metrics.Histogram)(nil)).Elem()) {
		h := prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    name,
			Help:    name,
			Buckets: r.Buckets,
		})
		r.Registerer.MustRegister(h)

		v = &Histogram{
			histogram: h,
		}
	} else if t.Out(0).Implements(reflect.TypeOf((*metrics.Counter)(nil)).Elem()) {
		g := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: name,
			Help: name,
		})
		r.Registerer.MustRegister(g)

		v = &Counter{
			gauge: g,
		}
	} else {
		g := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: name,
			Help: name,
		})
		r.Registerer.MustRegister(g)

		v = &Meter{
			gauge: g,
		}
	}
	r.m.Store(name, v)
	return v
}

func NewRegistry() *Registry {
	return &Registry{
		m:  sync.Map{},
		mu: sync.Mutex{},

		Registerer: prometheus.DefaultRegisterer,
		Buckets:    prometheus.ExponentialBuckets(1.0, 1.05, 512),
	}
}

var _ metrics.Registry = (*Registry)(nil)
