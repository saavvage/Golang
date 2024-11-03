package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	// dsn := os.Getenv("DATABASE_URL")
	dsn := "postgres://artem:12345678@localhost:5432/taskmanager"
	log.Println("1234141441")
	log.Println(dsn)
	log.Println("1234141441")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unsuccessful connection to database:", err)
	}
	fmt.Println("Successfull connection to database!")
}
