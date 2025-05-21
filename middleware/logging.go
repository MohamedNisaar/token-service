package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var logger = logrus.New()

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Infof("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		logger.Infof("Completed in %v", time.Since(start))
	})
}
