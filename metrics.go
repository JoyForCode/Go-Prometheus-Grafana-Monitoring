package main

type Power struct {
	PSU []struct {
		PowerConsumedWatts float64 `json:"PowerConsumedWatts"`
	} `json:"PSU"`
}
