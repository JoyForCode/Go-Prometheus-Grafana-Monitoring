package metrics

import (
	"strconv"

	"goprom/internal/models"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	temperatureCelsius = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "redfish_temperature_celsius",
			Help: "Temperature sensor reading in Celsius.",
		},
		[]string{"sensor", "physical_context", "sensor_number"},
	)

	temperatureLowerCritical = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "redfish_temperature_lower_threshold_critical_celsius",
			Help: "Lower critical temperature threshold in Celsius.",
		},
		[]string{"sensor", "physical_context", "sensor_number"},
	)

	temperatureLowerFatal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "redfish_temperature_lower_threshold_fatal_celsius",
			Help: "Lower fatal temperature threshold in Celsius.",
		},
		[]string{"sensor", "physical_context", "sensor_number"},
	)

	temperatureUpperCritical = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "redfish_temperature_upper_threshold_critical_celsius",
			Help: "Upper critical temperature threshold in Celsius.",
		},
		[]string{"sensor", "physical_context", "sensor_number"},
	)

	temperatureUpperFatal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "redfish_temperature_upper_threshold_fatal_celsius",
			Help: "Upper fatal temperature threshold in Celsius.",
		},
		[]string{"sensor", "physical_context", "sensor_number"},
	)
)

func InitTemperatureMetrics() {
	prometheus.MustRegister(
		temperatureCelsius,
		temperatureLowerCritical,
		temperatureLowerFatal,
		temperatureUpperCritical,
		temperatureUpperFatal,
	)
}

func UpdateTemperatureMetrics(temp *models.Temperature) {

	if temp.ReadingCelsius == nil {
		return
	}

	labels := prometheus.Labels{
		"sensor":           temp.Name,
		"physical_context": temp.PhysicalContext,
		"sensor_number":    strconv.Itoa(temp.SensorNumber),
	}

	temperatureCelsius.With(labels).Set(*temp.ReadingCelsius)

	if temp.LowerThresholdCritical != nil {
		temperatureLowerCritical.With(labels).Set(*temp.LowerThresholdCritical)
	}

	if temp.LowerThresholdFatal != nil {
		temperatureLowerFatal.With(labels).Set(*temp.LowerThresholdFatal)
	}

	if temp.UpperThresholdCritical != nil {
		temperatureUpperCritical.With(labels).Set(*temp.UpperThresholdCritical)
	}

	if temp.UpperThresholdFatal != nil {
		temperatureUpperFatal.With(labels).Set(*temp.UpperThresholdFatal)
	}
}
