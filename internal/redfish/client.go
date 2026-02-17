package redfish

import (
	"time"
	"net/http"
	"crypto/tls"
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
		BaseURL: baseURL,
		Username: username,
		Password: password,
		HTTPClient: &http.Client{
			Transport: tr,
			Timeout: 15 * time.Second,
		},
	}
}

/*The client.go file serves as a centrailzed client creation for all HTTP requests, and allows using same client for different endpoint types.*/
