package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// User es el usuario logueado
type User struct {
	ID          string   `json:"id"  validate:"required"`
	Name        string   `json:"name"  validate:"required"`
	Permissions []string `json:"permissions"`
	Login       string   `json:"login"  validate:"required"`
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if strings.Index(tokenString, "Bearer ") != 0 {
			ErrorResponse(w, 401, "Unauthorized")
			return
		}
		token := tokenString[7:]
			req, err := http.NewRequest("GET", "http://localhost:3000/v1/users/current", nil)
			if err != nil || req == nil {
				ErrorResponse(w, 500, "internal_server_error")
				return
			}
			req.Header.Add("Authorization", "Bearer "+token)
			resp, err := http.DefaultClient.Do(req)
			if err != nil || resp.StatusCode != 200 {
				ErrorResponse(w, 401, "Unauthorized")
				return
			}
		

		user := User{}
		json.NewDecoder(resp.Body).Decode(&user)

		ctx := context.WithValue(r.Context(), "user_logged", tokenString)
		ctx2 := context.WithValue(ctx, "user_id", user.ID)

		next.ServeHTTP(w, r.WithContext(ctx2))
	})
}