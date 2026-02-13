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
		//time.Sleep(20 * time.Second)
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()

		fetch := func() {
			power, err := RetrievePowerValues()
			if err != nil {
				log.Printf("Error retrieving power values: %v", err)
				UpdateHealth(false)
				return
			}
			UpdateHealth(true)
			UpdateMetrics(power)
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
