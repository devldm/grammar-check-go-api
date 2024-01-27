package middleware

import (
	"context"
	"net/http"

	"github.com/devldm/grammar-check-go/config"
)

func ConfigMiddleware(config *config.APIConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "api_config", config)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
