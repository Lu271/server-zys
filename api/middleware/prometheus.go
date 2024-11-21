package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "The total number of HTTP request",
		},
		[]string{"uri"},
	)

	summary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "http_request_delay",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"uri"},
	)
)

func init() {
	prometheus.MustRegister(counter)
	prometheus.MustRegister(summary)
}

func Qps(c *gin.Context) {
	counter.WithLabelValues(c.Request.RequestURI).Inc()
}

func Delay(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	latency := time.Now().Sub(startTime)
	summary.WithLabelValues(c.Request.RequestURI).
		Observe(float64(latency.Milliseconds()))
}
