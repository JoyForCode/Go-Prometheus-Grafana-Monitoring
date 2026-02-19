package collector

import (
	"goprom/internal/models"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

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

func (c *RedfishCollector) Collect(ch chan<- prometheus.Metric) {

	var power models.PowerControl
	status, err := c.client.Get("/redfish/v1/Chassis/System.Embedded.1/Power/PowerControl", &power)
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			c.redfishUpDesc,
			prometheus.GaugeValue,
			0,
		)
		return
	}

	if status == http.StatusNotFound {
		ch <- prometheus.MustNewConstMetric(
			c.redfishUpDesc,
			prometheus.GaugeValue,
			1,
		)
		return
	}

	if status != http.StatusOK {
		ch <- prometheus.MustNewConstMetric(
			c.redfishUpDesc,
			prometheus.GaugeValue,
			0,
		)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		c.powerConsumedDesc,
		prometheus.GaugeValue,
		float64(power.PowerConsumedWatts),
	)

	ch <- prometheus.MustNewConstMetric(
		c.powerCapacityDesc,
		prometheus.GaugeValue,
		float64(power.PowerCapacityWatts),
	)

	ch <- prometheus.MustNewConstMetric(
		c.redfishUpDesc,
		prometheus.GaugeValue,
		1,
	)
}
