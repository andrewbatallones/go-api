package middleware

import (
	"net/http"
	"strings"
)

// This middleware is responsible for setting needed information for any API request.
func SetApi(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isApiPath(r.URL.Path) {
			w.Header().Set("Content-Type", "application/json")
		}

		next.ServeHTTP(w, r)
	})
}

func isApiPath(path string) bool {
	return strings.Contains(path, "api") || path == "/healthcheck"
}
