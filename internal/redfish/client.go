package redfish

import (
	"net/http"
)

type Client struct {
	BaseURL    string
	Username   string
	Password   string
	HTTPClient *http.Client
}

/*The client.go file serves as a centrailzed client creation for all HTTP requests, and allows using same client for different endpoint types.*/
