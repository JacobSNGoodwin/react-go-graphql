package models

// AdminRole holds const value for creator
var AdminRole = Role{
	Name: "admin",
}

// EditorRole holds const value for editor
var EditorRole = Role{
	Name: "editor",
}

// Role struct holds database and response characteristics for a user role
type Role struct {
	Base
	Name  string  `json:"name,omitempty" gorm:"type:varchar(100);not null;unique"`
	Users []*User `json:"users" gorm:"many2many:user_roles"`
}
