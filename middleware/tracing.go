package middleware

import "net/http"

func TracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// need to integrate with openTelemetry later
		next.ServeHTTP(w, r)
	})
}
