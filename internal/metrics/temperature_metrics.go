package metrics

import (
	"strconv"

	"goprom/internal/models"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	ReadingCelsius = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "current_temperature_reading_celsius",
			Help: "Current temperature reading in Celsius",
		}, []string{"name", "physical_context", "sensor_number"},
	)
)

func InitTemperatureMetrics() {
	prometheus.MustRegister(ReadingCelsius)
}

func UpdateTemperatureMetrics(temperature *models.Temperature) {
	if temperature.ReadingCelsius == nil {
		return
	}

	ReadingCelsius.With(prometheus.Labels{
		"name":             temperature.Name,
		"physical_context": temperature.PhysicalContext,
		"sensor_number":    strconv.Itoa(temperature.SensorNumber),
	}).Set(*temperature.ReadingCelsius)
}

func mapHeathToFloat(health *string) float64 {
	if health == nil {
		return -1
	}

	switch *health {
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
