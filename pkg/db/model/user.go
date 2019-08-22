package model

type User struct {
	Email string `gorm:"primary_key;"`
	Name string
	Phone string
	password string
	Contacts []Contact `gorm:"foreignkey:UserId"`
}
