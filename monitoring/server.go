package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricsSuffix = "/metrics"
	listenAddress = "localhost:9453"
)

var (
	apiEndpoints = [2]string{"/api/apple", "/api/banana"}
)

func main() {
	// register custom exporter
	metricsPrefix := "myexporter"
	exporter := NewExporter(metricsPrefix)
	prometheus.MustRegister(exporter)

	// handle API endpoints
	for _, e := range apiEndpoints {
		http.HandleFunc(e, func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[%s] Requested to %s", time.Now().String(), r.URL.Path)
			exporter.IncreCounter()
			exporter.IncreCounterWithEndpoint(r.URL.Path)
		})
	}

	// Handle Metrics endpoint
	http.Handle(metricsSuffix, promhttp.Handler())

	log.Printf("Metrics exported at http://%s%s", listenAddress, metricsSuffix)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
