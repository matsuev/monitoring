package main

import (
	"log"
	"monitoring/internal/metrics"
	"monitoring/internal/monitoring"
	"net/http"
)

func init() {
	monitoring.StartProfiling("http://127.0.0.1:4040", "my.super.application")

	monitoring.StartMetrics("127.0.0.1:8001")
}

func main() {

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		metrics.HandlerCount.WithLabelValues("test").Inc()

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		metrics.HandlerCount.WithLabelValues("help").Inc()
		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatalln(err)
	}
}
