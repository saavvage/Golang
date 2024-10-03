package main

import (
	"log"
	"net/http"
	"rapif/db"
	"rapif/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize database connections
	db.InitSQL()
	db.InitGORM()

	// Create a new router
	router := mux.NewRouter()

	// SQL-based routes
	router.HandleFunc("/users/sql", handlers.GetUsersSQL).Methods("GET")
	router.HandleFunc("/users/sql", handlers.CreateUserSQL).Methods("POST")
	router.HandleFunc("/users/sql/{id}", handlers.UpdateUserSQL).Methods("PUT")
	router.HandleFunc("/users/sql/{id}", handlers.DeleteUserSQL).Methods("DELETE")

	// GORM-based routes
	router.HandleFunc("/users/gorm", handlers.GetUsersGORM).Methods("GET")
	router.HandleFunc("/users/gorm", handlers.CreateUserGORM).Methods("POST")
	router.HandleFunc("/users/gorm/{id}", handlers.UpdateUserGORM).Methods("PUT")
	router.HandleFunc("/users/gorm/{id}", handlers.DeleteUserGORM).Methods("DELETE")

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
