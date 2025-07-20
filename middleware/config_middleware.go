package middleware

import (
	"context"
	"net/http"

	"github.com/devldm/grammar-check-go/handlers"
)

type key string

const APIConfigKey key = "api_config"

func ConfigMiddleware(config *handlers.APIConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), APIConfigKey, config)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
