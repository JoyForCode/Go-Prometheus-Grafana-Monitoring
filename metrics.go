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

	redfishUp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "redfish_up",
			Help: "Redfish API Response Status (1 = up, 0 = down)",
		},
	)
)

func InitMetrics() {
	prometheus.MustRegister(psuWatts)
	prometheus.MustRegister(psuCapacityWatts)
	prometheus.MustRegister(redfishUp)
}

func UpdateMetrics(power *PowerControl) {
	psuWatts.Set(power.PowerConsumedWatts)
	psuCapacityWatts.Set(power.PowerCapacityWatts)
}

func UpdateHealth(up bool) {
	if up {
		redfishUp.Set(1)
	} else {
		redfishUp.Set(0)
	}
}