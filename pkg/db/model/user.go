package model

// Users table data model
type User struct {
	Email    string `gorm:"primary_key;"`
	Name     string
	Phone    string
	Password string
	Contacts []Contact `gorm:"foreignkey:UserId"`
}
