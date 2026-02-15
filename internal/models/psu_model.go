package models

type PowerControl struct {
	PowerConsumedWatts float64 `json:"PowerConsumedWatts"`
	PowerCapacityWatts float64 `json:"PowerCapacityWatts"`
}