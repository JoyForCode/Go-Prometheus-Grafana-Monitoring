package main

import (
	"crypto/tls"
	"fmt"

	// "io"
	"encoding/json"
	"net/http"
	"time"
)

type PowerControl struct {
	PowerConsumedWatts float64 `json:"PowerConsumedWatts"`
	PowerCapacityWatts float64 `json:"PowerCapacityWatts"`
}

func RetrievePowerValues() (*PowerControl, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	req, err := http.NewRequest(
		"GET",
		"https://192.168.1.101/redfish/v1/Chassis/System.Embedded.1/Power/PowerControl",
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("root", "calvin")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))

	var power PowerControl
	err = json.NewDecoder(resp.Body).Decode(&power)
	if err != nil {
		return nil, err
	}

	return &power, nil
}
