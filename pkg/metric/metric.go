package metric

import (
	"log"
	"net"
	"strconv"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	skipPaths = make(map[string]bool, 0)
)

// App Metrics interface
type Metrics interface {
	IncHits(status int, method, path string)
	ObserveResponseTime(status int, method, path string, observeTime float64)
	SkipPath(path string) bool
	SetSkipPath(paths []string)
}

// Prometheus Metrics struct
type PrometheusMetrics struct {
	HitsTotal prometheus.Counter
	Hits      *prometheus.CounterVec
	Times     *prometheus.HistogramVec
}

// Create metrics with address and name
func CreateMetrics(address string, name string) (Metrics, error) {
	var metr PrometheusMetrics
	metr.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: name + "_hits_total",
	})

	if err := prometheus.Register(metr.HitsTotal); err != nil {
		return nil, err
	}

	metr.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name + "_hits",
		},
		[]string{"status", "method", "path"},
	)

	if err := prometheus.Register(metr.Hits); err != nil {
		return nil, err
	}

	metr.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: name + "_times",
		},
		[]string{"status", "method", "path"},
	)

	if err := prometheus.Register(metr.Times); err != nil {
		return nil, err
	}

	if err := prometheus.Register(prometheus.NewBuildInfoCollector()); err != nil {
		return nil, err
	}

	go func() {
		router := fiber.New()
		router.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
		log.Printf("Metrics server is running on port: %s", address)
		ln, _ := net.Listen("tcp", ":"+address)
		err := router.Listener(ln)
		if err != nil {
			log.Fatalf("Error starting Server: %v", err)
		}
	}()

	return &metr, nil
}

// IncHits
func (metr *PrometheusMetrics) IncHits(status int, method, path string) {
	metr.HitsTotal.Inc()
	metr.Hits.WithLabelValues(strconv.Itoa(status), method, path).Inc()
}

// Observer response time
func (metr *PrometheusMetrics) ObserveResponseTime(status int, method, path string, observeTime float64) {
	metr.Times.WithLabelValues(strconv.Itoa(status), method, path).Observe(observeTime)
}

func (metr *PrometheusMetrics) SkipPath(path string) bool {
	return skipPaths[path]
}

func (metr *PrometheusMetrics) SetSkipPath(paths []string) {
	for _, val := range paths {
		skipPaths[val] = true
	}
}
