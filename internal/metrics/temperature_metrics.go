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

	lowerThresholdNonCritical = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "redfish_temperature_lower_threshold_non_critical_celsius",
			Help: "Lower non-critical temperature threshold in Celsius.",
		},
		[]string{"sensor", "physical_context", "sensor_number"},
	)

	upperThresholdNonCritical = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "redfish_temperature_upper_threshold_non_critical_celsius",
			Help: "Upper non-critical temperature threshold in Celsius.",
		},
		[]string{"sensor", "physical_context", "sensor_number"},
	)

	TemperatureHealth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "redfish_temperature_health",
			Help: "Health status of the temperature sensor (1=OK, 0.5=Warning, 0=Critical, -1=Unknown).",
		},
		[]string{"sensor", "physical_context", "sensor_number"},
	)

	TemperatureState = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "redfish_temperature_state",
			Help: "State of the temperature sensor (1=Enabled, 0=Disabled, -1=Unknown).",
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
		TemperatureHealth,
		TemperatureState,
		lowerThresholdNonCritical,
		upperThresholdNonCritical,
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

	if temp.LowerThresholdNonCritical != nil {
		lowerThresholdNonCritical.With(labels).Set(*temp.LowerThresholdNonCritical)
	}

	if temp.UpperThresholdNonCritical != nil {
		upperThresholdNonCritical.With(labels).Set(*temp.UpperThresholdNonCritical)
	}

	TemperatureHealth.With(labels).Set(mapHealthToFloat(temp.Status.Health))
	TemperatureState.With(labels).Set(mapStateToFloat(temp.Status.State))
}

func mapHealthToFloat(h *string) float64 {
	if h == nil {
		return -1
	}

	switch *h {
	case "OK":
		return 1
	case "Warning":
		return 0.5
	case "Critical":
		return 0
	default:
		return -1
	}
}

func mapStateToFloat(s *string) float64 {
	if s == nil {
		return -1
	}

	switch *s {
	case "Enabled":
		return 1
	case "Disabled":
		return 0
	default:
		return -1
	}
}
