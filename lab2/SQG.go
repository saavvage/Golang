package main

import (
	_ "database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "artem"
	password = "12345678"
	dbname   = "postgres"
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  int
}

func main() {
	dsn := "host=localhost user=artem password=12345678 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	db.AutoMigrate(&User{})
}

//3.

func insertUserGorm(db *gorm.DB, name string, age int) {
	user := User{Name: name, Age: age}
	db.Create(&user)
	fmt.Println("User inserted successfully")
}

//4.

func getUsersGorm(db *gorm.DB) {
	var users []User
	db.Find(&users)
	for _, user := range users {
		fmt.Println("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}
