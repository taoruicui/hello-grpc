package metrics

import (
	"net/http"
	"sync"
	"time"

	"fmt"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// these values prevent double prometheus registration
var (
	metricsRegLock    sync.RWMutex
	metricsRegistered bool
	PromMetrics       *PromMiddleware
)

type PromMiddleware struct {
	APILatency *prometheus.SummaryVec
	APIHits    *prometheus.CounterVec
	APIErrors  *prometheus.CounterVec
}

// Same as DefObjectives from the golang package, but adding 1.0 as an objective
var SummaryVecObjectives = map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001, 1.0: 0.0}

func init() {
	var err error
	PromMetrics, err = DefaultPromMiddleware()
	if err != nil {
		log.Fatal("failed to create service metrics")
	}
}

func (pm *PromMiddleware) Register() error {
	metricsRegLock.Lock()
	if metricsRegistered {
		metricsRegLock.Unlock()
		return nil
	}
	defer metricsRegLock.Unlock()
	if err := prometheus.Register(pm.APILatency); err != nil {
		return errors.Wrap(err, "Registering APILatency")
	}
	if err := prometheus.Register(pm.APIHits); err != nil {
		return errors.Wrap(err, "Registering APIHits")
	}
	if err := prometheus.Register(pm.APIErrors); err != nil {
		return errors.Wrap(err, "Registering APIErrors")
	}
	metricsRegistered = true
	return nil
}

// TODO need revisit here later, as current implementation is creating monitoring middleware
// and plugin into the server. This provided much convenience but we need to figure out a way
// of identifying the labels based on the uri (possibly?)
func (pm *PromMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r)

		uri := r.RequestURI
		pm.APIHits.WithLabelValues(uri).Inc()
		next.ServeHTTP(w, r)
	})
}

func DefaultPromMiddleware() (*PromMiddleware, error) {
	pm := &PromMiddleware{
		APILatency: prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Namespace:  "helloservice",
				Name:       "api_latency",
				Help:       "api latency in nanoseconds",
				Objectives: SummaryVecObjectives,
				MaxAge:     time.Minute,
			},
			[]string{"api"},
		),
		APIHits: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "helloservice",
				Name:      "api_hits",
				Help:      "how many times a particular api has been hit",
			},
			[]string{"api"},
		),
		APIErrors: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "helloservice",
				Name:      "api_errors",
				Help:      "how many times a particular api has returned with an error",
			},
			[]string{"api"},
		),
	}
	return pm, pm.Register()
}
