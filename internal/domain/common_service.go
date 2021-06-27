package domain

import (
	"fmt"
	"net/http"
	"time"

	"github.com/leandrocepeda32/reviews/internal/utils/errors"
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

func (r *RestClient) GetArticle(id string) error {

	requestUrl := fmt.Sprintf("http://localhost:3002/v1/articles/%s", id)
	req, err := http.NewRequest("GET", requestUrl, nil)

	if err != nil {
		return err
	}

	res, err := r.HTTPClient.Do(req)

	if err != nil {
		return errors.NewRestError("rest_client_error", http.StatusServiceUnavailable)
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return errors.NewRestError("rest_client_error", res.StatusCode)
	}

	return nil
}