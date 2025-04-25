package main

import (
	"log"
	"monitoring/internal/monitoring"
	"net/http"
)

func init() {
	monitoring.StartProfiling("http://127.0.0.1:4040", "my.super.application")
	// monitoring.StartMetrics("127.0.0.1:8001")
}

func main() {
	// go func() {
	// 	for {
	// 		metrics.OpsProcessed.Inc()
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }()

	// http.HandleFunc("/alert", func(w http.ResponseWriter, r *http.Request) {
	// 	data, err := io.ReadAll(r.Body)
	// 	if err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		fmt.Printf("%s\n", data)
	// 	}

	// 	w.WriteHeader(http.StatusOK)
	// })

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// metrics.HandlerCount.WithLabelValues("test").Inc()
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		// metrics.HandlerCount.WithLabelValues("help").Inc()
		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatalln(err)
	}
	// var s string

	// fmt.Scanln(&s)
}
