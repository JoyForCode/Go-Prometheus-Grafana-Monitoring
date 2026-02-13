package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	InitMetrics()

	go func() {
		for {
			power, err := RetrievePowerValues()
			if err != nil {
				log.Println("Error fetching power values:", err)
			} else {
				UpdateMetrics(power)
			}

			time.Sleep(20 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Exporter running on :9105")
	log.Fatal(http.ListenAndServe(":9105", nil))
}
