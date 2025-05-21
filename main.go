package main

import (
	"github.com/MohamedNisaar/token-service/handlers"
	"github.com/MohamedNisaar/token-service/metrics"
	"github.com/MohamedNisaar/token-service/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Middlewares
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.TracingMiddleware)
	r.Use(middleware.RateLimitMiddleware)

	r.HandleFunc("/token", handlers.GenerateTokenHandler).Methods("POST")
	r.Handle("/metrics", metrics.GetPrometheusHandler()).Methods("GET")

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
