package metrics

import (
	"goprom/internal/models"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	SystemBoardExhaustTemp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "system_board_exhaust_temperature_celsius",
			Help: "Current system board exhaust temperature in Celsius",
		},
	)
)

func InitTemperatureMetrics() {
	prometheus.MustRegister(SystemBoardExhaustTemp)
}

func UpdateTemperatureMetrics(temperature *models.Temperature) {
	SystemBoardExhaustTemp.Set(*temperature.ReadingCelsius)
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
