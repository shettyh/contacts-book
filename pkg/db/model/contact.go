package model

type Contact struct {
	Email  string `gorm:"primary_key"`
	Name   string `gorm:"not null"`
	Phone  string
	UserId string `gorm:"primary_key"`
}
