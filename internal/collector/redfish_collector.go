package collector

import (
	"goprom/internal/models"
	"goprom/internal/redfish"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type RedfishAPI interface {
	Get(endpoint string, target interface{}) (int, error)
}

type RedfishCollector struct {
	client RedfishAPI

	powerConsumedDesc *prometheus.Desc
	powerCapacityDesc *prometheus.Desc

	redfishUpDesc      *prometheus.Desc
	scrapeDurationDesc *prometheus.Desc

	readingCelsiusDesc            *prometheus.Desc
	lowerThresholdCriticalDesc    *prometheus.Desc
	lowerThresholdFatalDesc       *prometheus.Desc
	upperThresholdCriticalDesc    *prometheus.Desc
	upperThresholdFatalDesc       *prometheus.Desc
	lowerThresholdNonCriticalDesc *prometheus.Desc
	upperThresholdNonCriticalDesc *prometheus.Desc
}

func NewRedfishCollector(client RedfishAPI) *RedfishCollector {
	return &RedfishCollector{
		client: client,

		powerConsumedDesc: prometheus.NewDesc(
			"redfish_power_consumed_watts",
			"Current PSU power consumption in watts",
			nil,
			nil,
		),

		powerCapacityDesc: prometheus.NewDesc(
			"redfish_power_capacity_watts",
			"PSU power capacity in watts",
			nil,
			nil,
		),

		redfishUpDesc: prometheus.NewDesc(
			"redfish_up",
			"Overall Redfish scrape status (1 = success, 0 = failure)",
			nil,
			nil,
		),

		scrapeDurationDesc: prometheus.NewDesc(
			"redfish_scrape_duration_seconds",
			"Time taken to scrape Redfish in seconds",
			nil,
			nil,
		),

		readingCelsiusDesc: prometheus.NewDesc(
			"redfish_temperature_reading_celsius",
			"Current temperature reading",
			[]string{"name", "physical_context", "sensor_number"},
			nil,
		),

		lowerThresholdCriticalDesc: prometheus.NewDesc(
			"redfish_temperature_lower_threshold_critical",
			"Lower critical temperature threshold",
			[]string{"name", "physical_context", "sensor_number"},
			nil,
		),

		lowerThresholdFatalDesc: prometheus.NewDesc(
			"redfish_temperature_lower_threshold_fatal",
			"Lower fatal temperature threshold",
			[]string{"name", "physical_context", "sensor_number"},
			nil,
		),

		upperThresholdCriticalDesc: prometheus.NewDesc(
			"redfish_temperature_upper_threshold_critical",
			"Upper critical temperature threshold",
			[]string{"name", "physical_context", "sensor_number"},
			nil,
		),

		upperThresholdFatalDesc: prometheus.NewDesc(
			"redfish_temperature_upper_threshold_fatal",
			"Upper fatal temperature threshold",
			[]string{"name", "physical_context", "sensor_number"},
			nil,
		),

		lowerThresholdNonCriticalDesc: prometheus.NewDesc(
			"redfish_temperature_lower_threshold_non_critical",
			"Lower non-critical temperature threshold",
			[]string{"name", "physical_context", "sensor_number"},
			nil,
		),

		upperThresholdNonCriticalDesc: prometheus.NewDesc(
			"redfish_temperature_upper_threshold_non_critical",
			"Upper non-critical temperature threshold",
			[]string{"name", "physical_context", "sensor_number"},
			nil,
		),
	}
}

func (c *RedfishCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- c.powerConsumedDesc
	ch <- c.powerCapacityDesc

	ch <- c.redfishUpDesc
	ch <- c.scrapeDurationDesc

	ch <- c.readingCelsiusDesc
	ch <- c.lowerThresholdCriticalDesc
	ch <- c.lowerThresholdFatalDesc
	ch <- c.upperThresholdCriticalDesc
	ch <- c.upperThresholdFatalDesc
	ch <- c.lowerThresholdNonCriticalDesc
	ch <- c.upperThresholdNonCriticalDesc
}

func (c *RedfishCollector) Collect(ch chan<- prometheus.Metric) {
	start := time.Now()
	overallSuccess := true

	var power models.PowerControl
	status, err := c.client.Get(redfish.Power_Endpoints[0], &power)

	if err != nil || (status != http.StatusOK && status != http.StatusNotFound) {
		overallSuccess = false
	} else if status == http.StatusOK {
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
	}

	for _, endpoint := range redfish.Temperature_Endpoints {
		var t models.Temperature
		status, err := c.client.Get(endpoint, &t)

		if err != nil {
			overallSuccess = false
			continue
		}
		if status == http.StatusNotFound {
			continue
		}
		if status != http.StatusOK {
			overallSuccess = false
			continue
		}

		labels := []string{t.Name, t.PhysicalContext, strconv.Itoa(t.SensorNumber)}

		if t.ReadingCelsius != nil {
			ch <- prometheus.MustNewConstMetric(
				c.readingCelsiusDesc,
				prometheus.GaugeValue,
				float64(*t.ReadingCelsius),
				labels...,
			)
		}

		if t.LowerThresholdCritical != nil {
			ch <- prometheus.MustNewConstMetric(
				c.lowerThresholdCriticalDesc,
				prometheus.GaugeValue,
				float64(*t.LowerThresholdCritical),
				labels...,
			)
		}

		if t.LowerThresholdFatal != nil {
			ch <- prometheus.MustNewConstMetric(
				c.lowerThresholdFatalDesc,
				prometheus.GaugeValue,
				float64(*t.LowerThresholdFatal),
				labels...,
			)
		}

		if t.UpperThresholdCritical != nil {
			ch <- prometheus.MustNewConstMetric(
				c.upperThresholdCriticalDesc,
				prometheus.GaugeValue,
				float64(*t.UpperThresholdCritical),
				labels...,
			)
		}

		if t.UpperThresholdFatal != nil {
			ch <- prometheus.MustNewConstMetric(
				c.upperThresholdFatalDesc,
				prometheus.GaugeValue,
				float64(*t.UpperThresholdFatal),
				labels...,
			)
		}

		if t.LowerThresholdNonCritical != nil {
			ch <- prometheus.MustNewConstMetric(
				c.lowerThresholdNonCriticalDesc,
				prometheus.GaugeValue,
				float64(*t.LowerThresholdNonCritical),
				labels...,
			)
		}

		if t.UpperThresholdNonCritical != nil {
			ch <- prometheus.MustNewConstMetric(
				c.upperThresholdNonCriticalDesc,
				prometheus.GaugeValue,
				float64(*t.UpperThresholdNonCritical),
				labels...,
			)
		}
	}

	if overallSuccess {
		ch <- prometheus.MustNewConstMetric(c.redfishUpDesc, prometheus.GaugeValue, 1)
	} else {
		ch <- prometheus.MustNewConstMetric(c.redfishUpDesc, prometheus.GaugeValue, 0)
	}

	duration := time.Since(start).Seconds()
	ch <- prometheus.MustNewConstMetric(
		c.scrapeDurationDesc,
		prometheus.GaugeValue,
		duration,
	)
}
