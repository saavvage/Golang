package controllers

import (
	"encoding/json"
	"net/http"
	"task-manager/database"
	"task-manager/models"

	"github.com/gorilla/mux"
)

// GetTasks return each tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

// GetTask return by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var task models.Task
	if err := database.DB.First(&task, vars["id"]).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// CreateTask 
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Create(&task)
	json.NewEncoder(w).Encode(task)
}

// UpdateTask 
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var task models.Task
	if err := database.DB.First(&task, vars["id"]).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Save(&task)
	json.NewEncoder(w).Encode(task)
}

// DeleteTask 
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var task models.Task
	if err := database.DB.First(&task, vars["id"]).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}
