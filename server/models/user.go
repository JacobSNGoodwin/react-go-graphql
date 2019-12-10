package models

import (
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

	var roles []string
	for _, role := range u.Roles {
		roles = append(roles, role.Name)
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
