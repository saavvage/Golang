package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

// User struct represents a user in the database
type User struct {
	ID   int
	Name string
	Age  int
}

// initDB establishes the connection pool to the PostgreSQL database
func initDB() {
	var err error
	connStr := "user=artem dbname=postgres sslmode=disable password=12345678"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	fmt.Println("Successfully connected to the database")
}

// CreateUsersTable creates the users table with constraints
func CreateUsersTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		age INT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	fmt.Println("Users table created successfully")
}

// InsertUsers inserts multiple users into the users table within a transaction
func InsertUsers(users []User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO users (name, age) VALUES ($1, $2)"
	for _, user := range users {
		_, err := tx.Exec(query, user.Name, user.Age)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("Users inserted successfully")
	return nil
}

// GetUsersWithFilterAndPagination queries users with filtering and pagination
func GetUsersWithFilterAndPagination(minAge, maxAge, limit, offset int) ([]User, error) {
	query := `SELECT id, name, age FROM users WHERE age >= $1 AND age <= $2 LIMIT $3 OFFSET $4`
	rows, err := db.Query(query, minAge, maxAge, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// UpdateUser updates a user's details by ID
func UpdateUser(id int, name string, age int) error {
	query := `UPDATE users SET name=$1, age=$2 WHERE id=$3`
	_, err := db.Exec(query, name, age, id)
	if err != nil {
		return err
	}

	fmt.Println("User updated successfully")
	return nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	fmt.Println("User deleted successfully")
	return nil
}

func main() {
	initDB()
	defer db.Close()

	// Create the users table
	CreateUsersTable()

	// Example: Insert multiple users
	users := []User{
		{Name: "Artem", Age: 20},
		{Name: "Sabir", Age: 21},
	}
	err := InsertUsers(users)
	if err != nil {
		log.Fatal("Error inserting users:", err)
	}

	// Example: Query users with filtering and pagination
	filteredUsers, err := GetUsersWithFilterAndPagination(20, 30, 2, 0)
	if err != nil {
		log.Fatal("Error querying users:", err)
	}
	fmt.Println("Filtered users:", filteredUsers)

	// Example: Update a user
	err = UpdateUser(1, "Artem Updated", 20)
	if err != nil {
		log.Fatal("Error updating user:", err)
	}

	// Example: Delete a user
	err = DeleteUser(2)
	if err != nil {
		log.Fatal("Error deleting user:", err)
	}
}
