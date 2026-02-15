package models

type Temperature struct {
	Name                      string   `json:"Name"`
	SensorNumber              int      `json:"SensorNumber"`
	PhysicalContext           string   `json:"PhysicalContext"`
	LowerThresholdCritical    *float64 `json:"LowerThresholdCritical"`
	LowerThresholdFatal       *float64 `json:"LowerThresholdFatal"`
	LowerThresholdNonCritical *float64 `json:"LowerThresholdNonCritical"`
	MaxReadingRangeTemp       *float64 `json:"MaxReadingRangeTemp"`
	MinReadingRangeTemp       *float64 `json:"MinReadingRangeTemp"`
	ReadingCelsius            *float64 `json:"ReadingCelsius"`
	UpperThresholdCritical    *float64 `json:"UpperThresholdCritical"`
	UpperThresholdFatal       *float64 `json:"UpperThresholdFatal"`
	UpperThresholdNonCritical *float64 `json:"UpperThresholdNonCritical"`
	Status                    Status   `json:"Status"`
}

type Status struct {
	Health *string `json:"Health"`
	State  *string `json:"State"`
}

/*Data Points Available:
- LowerThresholdCritical
- LowerThresholdFatal
- LowerThresholdNonCritical
- MaxReadingRangeTemp
- MinReadingRangeTemp
- ReadingCelsius
- PhysicalContext
- SensorNumber
- UpperThresholdCritical
- UpperThresholdFatal
- UpperThresholdNonCritical
- Status
	- Health
	- State

** Name of the sensor is taken as the label for the metric
*/
