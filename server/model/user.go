package model

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/logger"
	"github.com/maxbrain0/react-go-graphql/server/middleware"
	"github.com/maxbrain0/react-go-graphql/server/schema"
	uuid "github.com/satori/go.uuid"
)

var ctxLogger = logger.CtxLogger

// User holds data of type schema.User for a single user
type User schema.User

// Users is used for parsing gorm reads with multiple users
type Users []*schema.User

// UserClaims is used as data type for creating jwts
type UserClaims schema.UserClaims

// LoginOrCreate takes the current user and logs them in if they exist.
// It creates the user if the user doesn't yet exist
func (u *User) LoginOrCreate(p graphql.ResolveParams) error {
	db := middleware.GetDB(p.Context)
	// rc := middleware.GetRedis(p.Context)

	// Add error checking
	if err := db.
		Preload("Roles").
		Where(User{Email: u.Email}).
		Attrs(User{Name: u.Name, ImageURI: u.ImageURI}).
		FirstOrCreate(&u).Error; err != nil {
		return err
	}

	w := middleware.GetWriter(p.Context)

	if err := middleware.CreateAndSendToken(w, u.ID); err != nil {
		return err
	}

	return nil
}

// GetAll returns a list of all users
func (u Users) GetAll(p graphql.ResolveParams) error {
	db := middleware.GetDB(p.Context)
	// userInfo := middleware.GetAuth(p.Context)

	// if !userInfo.Roles.IsAdmin {
	// 	return fmt.Errorf("User is not authorized to view other users")
	// }

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

// GetByID gets user from database based on the users ID
func (u *User) GetByID(p graphql.ResolveParams) error {
	db := middleware.GetDB(p.Context)
	// userInfo := middleware.GetAuth(p.Context)

	// if !userInfo.Roles.IsAdmin {
	// 	return fmt.Errorf("User is not authorized to view other users")
	// }

	// Find by uuid or email, which should both be unique
	if result := db.
		Preload("Roles").
		Where("id = ?", uuid.FromStringOrNil(p.Args["id"].(string))).
		Find(&u); result.Error != nil {
		return result.Error
	}

	return nil
}

// GetCurrent gets user from database based on the users ID
func (u *User) GetCurrent(p graphql.ResolveParams) error {
	db := middleware.GetDB(p.Context)
	userInfo := middleware.GetAuth(p.Context)

	// Find by uuid or email, which should both be unique
	if result := db.
		Preload("Roles").
		Where("id = ?", userInfo.ID).
		Find(&u); result.Error != nil {
		return result.Error
	}

	return nil
}