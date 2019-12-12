package schema

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

// User holds user information and role
type User struct {
	Base
	Name     string  `json:"name" gorm:"type:varchar(100);not null"`
	Email    string  `json:"email" gorm:"type:varchar(100);unique_index"`
	ImageURI string  `json:"imageUri" gorm:"type:text"`
	Roles    []*Role `json:"roles" gorm:"many2many:user_roles"`
}

// UserClaims used for creating and parsing JWTs
type UserClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}

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
