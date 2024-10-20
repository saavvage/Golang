package main

import (
	"log"
	"net/http"
	"task-manager/controllers"
	"task-manager/database"
	"task-manager/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to database
	database.Connect()

	// Automatic migration to tasks
	database.DB.AutoMigrate(&models.Task{})

	// Create a router
	router := mux.NewRouter()

	// Define routes for tasks
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	// Add CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Start the project with CORS
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
