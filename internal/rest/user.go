package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/leandrocepeda32/reviews/internal/utils/errors"
)


func (r *RestClient) GetUser(ctx context.Context) (User, error) {

	user := User{}
	token := ctx.Value("user_logged")
	requestUrl := "http://localhost:3000/v1/users/current"
	req, err := http.NewRequest("GET", requestUrl, nil)

	if err != nil {
		return user,err
	}

	req.Header.Set("Authorization", fmt.Sprintf("%v", token))

	res, err := r.HTTPClient.Do(req)

	if err != nil {
		return user, errors.NewRestError("rest_client_error", http.StatusServiceUnavailable)
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return user,errors.NewRestError("rest_client_error", res.StatusCode)
	}

	json.NewDecoder(res.Body).Decode(&user)

	return user, nil
}