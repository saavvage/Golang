package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// User model with a one-to-one relationship with Profile
type User struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"not null"`
	Age     int     `gorm:"not null"`
	Profile Profile `gorm:"foreignKey:UserID"`
}

// Profile model associated with User
type Profile struct {
	ID                uint `gorm:"primaryKey"`
	UserID            uint `gorm:"not null;unique"`
	Bio               string
	ProfilePictureURL string
}

// Initialize the GORM connection with PostgreSQL
func initDB() {
	// Update the Data Source Name (DSN) with your PostgreSQL credentials
	dsn := "user=artem password=12345678 dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable detailed logging
	})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Set connection pooling settings
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error accessing raw DB:", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("Successfully connected to the database")
}

// AutoMigrate models to create tables with constraints
func AutoMigrateModels() {
	err := db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		log.Fatal("Failed to auto migrate:", err)
	}
	fmt.Println("Tables created successfully")
}

// InsertUserWithProfile inserts a user and their associated profile in a single transaction
func InsertUserWithProfile(name string, age int, bio string, profilePic string) error {
	user := User{
		Name: name,
		Age:  age,
		Profile: Profile{
			Bio:               bio,
			ProfilePictureURL: profilePic,
		},
	}

	// Use transaction to insert the user and profile
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err // Rollback on error
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error inserting user with profile:", err)
		return err
	}
	fmt.Println("User and profile inserted successfully")
	return nil
}

// GetUsersWithProfiles retrieves all users along with their profiles using eager loading
func GetUsersWithProfiles() ([]User, error) {
	var users []User
	err := db.Preload("Profile").Find(&users).Error // Eager load profiles
	if err != nil {
		return nil, err
	}

	// Print the users and their profiles
	for _, user := range users {
		fmt.Printf("User: %s, Age: %d, Bio: %s, Profile Pic: %s\n", user.Name, user.Age, user.Profile.Bio, user.Profile.ProfilePictureURL)
	}

	return users, nil
}

// UpdateUserProfile updates a user's profile by their ID
func UpdateUserProfile(userID uint, newBio string, newProfilePic string) error {
	var user User
	err := db.Preload("Profile").First(&user, userID).Error // Load the user and profile
	if err != nil {
		return err
	}

	// Update the profile fields
	user.Profile.Bio = newBio
	user.Profile.ProfilePictureURL = newProfilePic

	// Save the changes
	err = db.Save(&user.Profile).Error
	if err != nil {
		return err
	}

	fmt.Println("User profile updated successfully")
	return nil
}

// DeleteUser deletes a user and their associated profile
func DeleteUser(userID uint) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		// Find the user
		var user User
		if err := tx.Preload("Profile").First(&user, userID).Error; err != nil {
			return err
		}

		// Delete the user (which will delete the associated profile)
		if err := tx.Delete(&user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal("Error deleting user:", err)
		return err
	}

	fmt.Println("User and associated profile deleted successfully")
	return nil
}

func main() {
	// Initialize the database connection
	initDB()

	// Auto migrate to create tables
	AutoMigrateModels()

	// Insert a new user and profile
	InsertUserWithProfile("Artem", 20, "Software Engineer", "https://example.com/profile.jpg")

	// Get users and their profiles
	GetUsersWithProfiles()

	// Update a user's profile
	UpdateUserProfile(1, "Updated Bio", "https://example.com/updated.jpg")

	// Delete a user by ID
	DeleteUser(1)
}
