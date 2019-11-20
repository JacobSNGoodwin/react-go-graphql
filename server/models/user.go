package models

// User holds user information and role
type User struct {
	Base
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);unique_index"`
	Role  string `gorm:"size:255"`
}
