package models

// Role holds roles for users
type Role struct {
	Base
	Name  string  `gorm:"type:varchar(100);not null;unique"`
	Users []*User `gorm:"many2many:user_roles"`
}
