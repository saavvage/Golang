package main

import (
	"log"
	"net/http"
	"task-manager/controllers"
	"task-manager/database"
	"task-manager/middleware"
	"task-manager/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Task{}, &models.User{})

	router := mux.NewRouter()

	// routes ro register
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	// routes for tasks
	taskRouter := router.PathPrefix("/tasks").Subrouter()
	taskRouter.Use(middleware.AuthMiddleware)
	taskRouter.HandleFunc("", controllers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("", controllers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/{id}", controllers.GetTask).Methods("GET")
	taskRouter.HandleFunc("/{id}", controllers.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/{id}", controllers.DeleteTask).Methods("DELETE")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
