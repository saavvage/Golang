package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var SQLDB *sql.DB

// InitGORM initializes GORM connection with PostgreSQL
func InitGORM() {
	dsn := "user=artem password=12345678 dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database with GORM: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error getting SQL DB from GORM: %v", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Connected to the database using GORM")
}

// InitSQL initializes connection using database/sql
func InitSQL() {
	dsn := "user=artem password=12345678 dbname=postgres sslmode=disable"
	var err error
	SQLDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database with SQL: %v", err)
	}
	SQLDB.SetMaxOpenConns(25)
	SQLDB.SetMaxIdleConns(25)
	SQLDB.SetConnMaxLifetime(5 * time.Minute)

	err = SQLDB.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	log.Println("Connected to the database using SQL")
}
