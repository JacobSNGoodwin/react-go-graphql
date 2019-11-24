package models

// User holds user information and role
type User struct {
	Base
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
	Email string `json:"email" gorm:"type:varchar(100);unique_index"`
	// Roles []*Role `gorm:"many2many:user_roles;association_autoupdate:false;association_autocreate:false"`
}
