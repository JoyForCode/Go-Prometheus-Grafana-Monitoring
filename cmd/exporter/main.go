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
	metrics.InitTemperatureMetrics()
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

		fetch2 := func() {
			temperature, err := redfish.RetrieveTemperatureValues()
			if err != nil {
				log.Printf("Error retrieving temperature values: %v", err)
				metrics.UpdateHealth(false)
				return
			}
			metrics.UpdateTemperatureMetrics(temperature)
		}

		fetch3 := func() {
			client := redfish.NewClient(
				redfish.BaseURL,
				"root",
				"calvin",
			)

			temps, err := client.GetTemperatures(redfish.Temperature_Endpoints)
			if err != nil {
				log.Printf("Error retrieving temperature values: %v", err)
				return
			}

			for _, t := range temps {
				metrics.UpdateTemperatureMetrics(t)
			}
		}
		fetch()
		fetch2()
		fetch3()
		for range ticker.C {
			fetch()
			fetch2()
			fetch3()
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Exporter running on :9105")
	log.Fatal(http.ListenAndServe(":9105", nil))
}
