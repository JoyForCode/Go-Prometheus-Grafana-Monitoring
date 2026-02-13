package main

import (
	"goprom/internal/metrics"
	"goprom/internal/redfish"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	metrics.InitMetrics()

	go func() {
		//time.Sleep(20 * time.Second)
		ticker := time.NewTicker(60 * time.Second)
		defer ticker.Stop()

		fetch := func() {
			power, err := redfish.RetrievePowerValues()
			if err != nil {
				log.Printf("Error retrieving power values: %v", err)
				metrics.UpdateHealth(false)
				return
			}
			metrics.UpdateHealth(true)
			metrics.UpdateMetrics(power)
		}
		fetch()

		for range ticker.C {
			fetch()
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Exporter running on :9105")
	log.Fatal(http.ListenAndServe(":9105", nil))
}
