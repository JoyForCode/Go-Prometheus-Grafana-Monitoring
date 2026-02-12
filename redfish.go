package main

import (
	"crypto/tls"
	// "fmt"
	// "io"
	"net/http"
	"encoding/json"
)

type PowerControl struct {
	PowerConsumedWatts float64 `json:"PowerConsumedWatts"`
	PowerCapacityWatts float64 `json:"PowerCapacityWatts"`
}

func RetrievePowerValues() (*PowerControl, error){

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: tr,
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
		panic(err)
	}

	defer resp.Body.Close()

	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))

	var power PowerControl
	err = json.NewDecoder(resp.Body).Decode(&power)
	if err != nil {
		return nil, err
	}

	return &power, nil
}
