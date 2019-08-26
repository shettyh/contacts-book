package model

// Contacts table data model
type Contact struct {
	Email string `gorm:"primary_key" json:"email"`
	Name  string `gorm:"not null" json:"name"`
	// TODO: can have multiple phone numbers ?
	Phone  string `json:"phone"`
	UserId string `gorm:"primary_key" json:"user_id, omitempty"`
}
