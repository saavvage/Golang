package main

import (
	"fmt"
	"log"
	"os"
	"task-manager/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "artem"
	password = "12345678"
	dbname   = "postgres"
)

func Connect() {
	
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unsuccessfull connection to database:", err)
	}

	log.Println("Successfully connected to the database")
	DB.AutoMigrate(&models.Task{})
}

type Task struct {
	ID int
	Task string
	Status string
}

//122
.App {
	text-align: center;
  }
  
  .App-logo {
	height: 40vmin;
	pointer-events: none;
  }
  
  @media (prefers-reduced-motion: no-preference) {
	.App-logo {
	  animation: App-logo-spin infinite 20s linear;
	}
  }
  
  .App-header {
	background-color: #282c34;
	min-height: 100vh;
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
	font-size: calc(10px + 2vmin);
	color: white;
  }
  
  .App-link {
	color: #61dafb;
  }
  
  @keyframes App-logo-spin {
	from {
	  transform: rotate(0deg);
	}
	to {
	  transform: rotate(360deg);
	}
  }
  
  .sidebar {
	height: 100vh;
	padding: 10px;
  }
  
  .mt-4 {
	margin-top: 1.5rem;
  }
  
  .btn-primary, .btn-danger {
	width: 100%;
  }
  
//211
import React, { useState, useEffect } from 'react';
import { Container, Row, Col, ListGroup, Button, Form, Badge } from 'react-bootstrap';
import axios from 'axios';
import './App.css';
import TaskList from './components/TaskList';
import TaskForm from './components/TaskForm';
import ToDoList from './To-Do-List';

function App() {
  const [tasks, setTasks] = useState([]);
  const [filteredTasks, setFilteredTasks] = useState([]);
  const [taskTitle, setTaskTitle] = useState("");
  const [filter, setFilter] = useState("all");

  // get task from server
  useEffect(() => {
    axios.get('http://localhost:8080/tasks')
      .then(response => {
        setTasks(response.data);
        setFilteredTasks(response.data);
      })
      .catch(error => console.error('Error fetching tasks:', error));
  }, []);

  // filter tasks
  useEffect(() => {
    if (filter === "completed") {
      setFilteredTasks(tasks.filter(task => task.completed));
    } else if (filter === "pending") {
      setFilteredTasks(tasks.filter(task => !task.completed));
    } else {
      setFilteredTasks(tasks);
    }
  }, [filter, tasks]);

  // add task
  const addTask = () => {
    const newTask = { title: taskTitle, completed: false };
    axios.post('http://localhost:8080/tasks', newTask)
      .then(response => {
        setTasks([...tasks, response.data]);
        setTaskTitle("");
      })
      .catch(error => console.error('Error adding task:', error));
  };

  // delete task
  const deleteTask = (id) => {
    axios.delete(`http://localhost:8080/tasks/${id}`)
      .then(() => {
        setTasks(tasks.filter(task => task.id !== id));
      })
      .catch(error => console.error('Error deleting task:', error));
  };

  // Mark task as done
  const toggleComplete = (id) => {
    const task = tasks.find(task => task.id === id);
    const updatedTask = { ...task, completed: !task.completed };
    axios.put(`http://localhost:8080/tasks/${id}`, updatedTask)
      .then(response => {
        setTasks(tasks.map(task => task.id === id ? response.data : task));
      })
      .catch(error => console.error('Error updating task:', error));
  };

  return (
    <Container fluid>
      <Row>
        <Col xs={2} className="sidebar bg-light">
          <h4 className="mt-4">Task Manager</h4>
          <ListGroup>
            <ListGroup.Item action onClick={() => setFilter('all')}>All Tasks</ListGroup.Item>
            <ListGroup.Item action onClick={() => setFilter('completed')}>Completed</ListGroup.Item>
            <ListGroup.Item action onClick={() => setFilter('pending')}>Pending</ListGroup.Item>
          </ListGroup>
          <Form.Group className="mt-4">
            <Form.Control 
              type="text" 
              placeholder="New Task" 
              value={taskTitle} 
              onChange={(e) => setTaskTitle(e.target.value)} 
            />
            <Button variant="primary" className="mt-2" onClick={addTask}>Add Task</Button>
          </Form.Group>
        </Col>

        <Col xs={10}>
          <h2 className="mt-4">Tasks</h2>
          <ListGroup>
            {filteredTasks.map(task => (
              <ListGroup.Item key={task.id} className="d-flex justify-content-between align-items-center">
                <div>
                  <Form.Check 
                    type="checkbox" 
                    checked={task.completed} 
                    onChange={() => toggleComplete(task.id)} 
                    label={task.title} 
                  />
                  <Badge bg={task.completed ? 'success' : 'warning'} className="ms-3">
                    {task.completed ? 'Completed' : 'Pending'}
                  </Badge>
                </div>
                <Button variant="danger" onClick={() => deleteTask(task.id)}>Delete</Button>
              </ListGroup.Item>
            ))}
          </ListGroup>
        </Col>
      </Row>
    </Container>
  );
}

export default App;
 //controller
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
 //main.go
 package main

import (
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"task-manager/controllers"
	"task-manager/database"
	"task-manager/models"
)

func main() {
	// Connection to database
	database.Connect()

	// Automatic migration of Task
	database.DB.AutoMigrate(&models.Task{})

	// Add router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	// Add CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Start server with CORS
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}


npx create-react-app frontend
cd frontend
npm install axios react-router-dom