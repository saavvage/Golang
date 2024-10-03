package models

// User model used for both SQL and GORM
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name" gorm:"unique"`
	Age  int    `json:"age"`
}
