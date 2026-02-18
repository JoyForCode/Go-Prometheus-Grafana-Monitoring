package collector

import "github.com/prometheus/client_golang/prometheus"

type RedfishAPI interface {
	Get(endpoint string, target interface{}) (int, error)
}

type RedfishCollector struct {
	client RedfishAPI

	powerConsumedDesc *prometheus.Desc
	powerCapacityDesc *prometheus.Desc
	redfishUpDesc     *prometheus.Desc
}

func NewRedfishCollector(client RedfishAPI) *RedfishCollector {
	return &RedfishCollector{
		client: client,
		powerConsumedDesc: prometheus.NewDesc(
			"power_consumed_watts",
			"Current PSU power consumption in watts",
			nil,
			nil,
		),

		powerCapacityDesc: prometheus.NewDesc(
			"power_capacity_watts",
			"PSU power capacity in watts",
			nil,
			nil,
		),

		redfishUpDesc: prometheus.NewDesc(
			"redfish_up",
			"Redfish API Response Status (1 = up, 0 = down)",
			nil,
			nil,
		),
	}
}

func (c *RedfishCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.powerConsumedDesc
	ch <- c.powerCapacityDesc
	ch <- c.redfishUpDesc
}

