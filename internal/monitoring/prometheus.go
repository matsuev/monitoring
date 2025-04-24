package monitoring

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartMetrics ...
func StartMetrics(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: promhttp.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}
