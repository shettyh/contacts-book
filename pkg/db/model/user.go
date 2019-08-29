package model

// Users table data model
type User struct {
	Email    string    `gorm:"primary_key;" json:"email"`
	Name     string    `gorm:"not null" json:"name"`
	Phone    string    `gorm:"not null;unique" json:"phone"`
	Password string    `gorm:"not null" json:"password"`
	Contacts []Contact `gorm:"foreignkey:UserId" json:"contacts,omitempty"`
}
