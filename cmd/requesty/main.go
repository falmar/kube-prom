package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
)

func makeHandler(counter *prometheus.CounterVec, path, code string) http.Handler {
	counter = counter.MustCurryWith(prometheus.Labels{"path": path})

	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		counter.With(prometheus.Labels{"code": code}).Inc()
	})
}

func main() {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total http requests",
		},
		[]string{"path", "code"},
	)

	prometheus.MustRegister(counter)

	http.Handle("/handler", makeHandler(counter, "/handler", "200"))
	http.Handle("/error", makeHandler(counter, "/handler", "500"))
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
