package main

import (
	"net/http"

	_ "github.com/auth0/go-jwt-middleware/v2"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv"
)

// HealthHandler is a handler for health check
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/healthcheck", HealthHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
