package redfish

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	Username   string
	Password   string
	HTTPClient *http.Client
}

func NewClient(baseURL, username, password string) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return &Client{
		BaseURL:  baseURL,
		Username: username,
		Password: password,
		HTTPClient: &http.Client{
			Transport: tr,
			Timeout:   15 * time.Second,
		},
	}
}

func (c *Client) Get(endpoint string, target interface{}) (int, error) {
	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return 0, nil
	}

	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return http.StatusNotFound, nil
	}

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}

/*The client.go file serves as a centrailzed client creation for all HTTP requests, and allows using same client for different endpoint types.*/
