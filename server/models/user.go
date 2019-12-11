package models

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/middleware"
)

// User holds user information and role
type User struct {
	Base
	Name     string  `json:"name" gorm:"type:varchar(100);not null"`
	Email    string  `json:"email" gorm:"type:varchar(100);unique_index"`
	ImageURI string  `json:"imageUri" gorm:"type:text"`
	Roles    []*Role `json:"roles" gorm:"many2many:user_roles"`
}

// Users holds an array of users
type Users []User

// LoginOrCreate takes the current user and logs them in if they exist.
// It creates the user if the user doesn't yet exist
func (u *User) LoginOrCreate(p graphql.ResolveParams) error {
	db := middleware.GetDB(p.Context)

	// Add error checking
	if err := db.
		Preload("Roles").
		Where(User{Email: u.Email}).
		Attrs(User{Name: u.Name, ImageURI: u.ImageURI}).
		FirstOrCreate(&u).Error; err != nil {
		return err
	}

	w := middleware.GetWriter(p.Context)

	roles := middleware.RolesMap{
		IsAdmin:  false,
		IsEditor: false,
	}

	for _, role := range u.Roles {
		if role.Name == "admin" {
			roles.IsAdmin = true
		}
		if role.Name == "editor" {
			roles.IsEditor = true
		}
	}

	userInfo := middleware.UserInfo{
		ID:    u.ID,
		Email: u.Email,
		Roles: roles,
	}

	if err := middleware.CreateAndSendToken(w, userInfo); err != nil {
		return err
	}

	return nil
}

// GetAll returns a list of all users
func (u *Users) GetAll(p graphql.ResolveParams) error {
	db := middleware.GetDB(p.Context)
	userInfo := middleware.GetAuth(p.Context)

	if !userInfo.Roles.IsAdmin {
		return fmt.Errorf("User is not authorized to view other users")
	}

	if result :=
		db.
			Order("email").
			Limit(p.Args["limit"].(int)).
			Offset(p.Args["offset"].(int)).
			Preload("Roles").
			Find(&u); result.Error != nil {
		return nil
	}

	return nil
}
