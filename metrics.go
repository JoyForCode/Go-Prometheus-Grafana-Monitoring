package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	psuWatts = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "power_consumed_watts",
			Help: "Current PSU power consumption in watts",
		},
	)

	psuCapacityWatts = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:"power_capacity_watts",
			Help:"Maximum PSU power capacity in watts",
		},
	)
)

func InitMetrics() {
	prometheus.MustRegister(psuWatts)
	prometheus.MustRegister()
}

func UpdateMetrics(power *PowerControl) {
	psuWatts.Set(power.PowerConsumedWatts)
	psuCapacityWatts.Set(power.PowerCapacityWatts)
}
