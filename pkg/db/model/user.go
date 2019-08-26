package model

// Users table data model
type User struct {
	Email    string    `gorm:"primary_key;" json:"email"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Password string    `json:"password"`
	Contacts []Contact `gorm:"foreignkey:UserId" json:"contacts, omitempty"`
}
