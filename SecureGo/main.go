package main
import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/SecureGo/middleware"
    "github.com/SecureGo/handlers"
)

func main() {
    r := mux.NewRouter()

    // Use midware
    r.Use(middleware.SecurityHeadersMiddleware)
    r.Use(middleware.RequestLogger)
	//Use handlers
    r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
    r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
    r.HandleFunc("/admin", middleware.AuthMiddleware(handlers.AdminHandler)).Methods("GET")

    // Metric's action
    r.Path("/metrics").Handler(middleware.MetricsHandler())

	//set up server
    http.ListenAndServe(":8080", r)
}
