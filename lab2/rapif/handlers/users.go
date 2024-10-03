package handlers
import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"rapif/db"
	"rapif/models"
)

// GetUsersSQL retrieves users using SQL with optional filtering and sorting
func GetUsersSQL(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, age FROM users"
	age := r.URL.Query().Get("age")
	if age != "" {
		query += " WHERE age = " + age
	}

	sort := r.URL.Query().Get("sort")
	if sort != "" {
		query += " ORDER BY " + sort
	}

	rows, err := db.SQLDB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

// GetUsersGORM retrieves users using GORM with optional filtering, sorting, and pagination
func GetUsersGORM(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	query := db.DB

	age := r.URL.Query().Get("age")
	if age != "" {
		query = query.Where("age = ?", age)
	}

	sort := r.URL.Query().Get("sort")
	if sort != "" {
		query = query.Order(sort)
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)

	err := query.Find(&users).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// CreateUserSQL creates a new user using SQL
func CreateUserSQL(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id"
	err = db.SQLDB.QueryRow(query, user.Name, user.Age).Scan(&user.ID)
	if err != nil {
		http.Error(w, "User with that name already exists", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// CreateUserGORM creates a new user using GORM
func CreateUserGORM(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.DB.Create(&user).Error
	if err != nil {
		http.Error(w, "User with that name already exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// UpdateUserSQL updates a user by ID using SQL
func UpdateUserSQL(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE users SET name = $1, age = $2 WHERE id = $3"
	_, err = db.SQLDB.Exec(query, user.Name, user.Age, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateUserGORM updates a user by ID using GORM
func UpdateUserGORM(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUserSQL deletes a user by ID using SQL
func DeleteUserSQL(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	query := "DELETE FROM users WHERE id = $1"
	_, err := db.SQLDB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUserGORM deletes a user by ID using GORM
func DeleteUserGORM(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := db.DB.Delete(&models.User{}, id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
