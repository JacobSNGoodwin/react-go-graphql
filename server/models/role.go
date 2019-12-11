package models

// AdminRole holds const value for creator
const AdminRole = "admin"

// EditorRole holds const value for editor
const EditorRole = "editor"

// Role struct holds database and response characteristics for a user role
type Role struct {
	Base
	Name  string  `json:"name,omitempty" gorm:"type:varchar(100);not null;unique"`
	Users []*User `json:"users" gorm:"many2many:user_roles"`
}
