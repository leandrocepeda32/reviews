package rest

import (
	"net/http"
	"time"
)

type RestClient struct {
	HTTPClient *http.Client
}

func NewRestClient(time time.Duration) RestClient {
	return RestClient{
		HTTPClient: &http.Client{
			Timeout: time,
		},
	}
}