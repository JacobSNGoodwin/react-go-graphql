package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/database"
	"github.com/maxbrain0/react-go-graphql/server/inmem"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

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

// Users holds an array of users
type Users []User

// LoginOrCreate takes the current user and logs them in if they exist.
// It creates the user if the user doesn't yet exist
func (u *User) LoginOrCreate(p graphql.ResolveParams) error {
	// Add error checking
	db := database.Conn
	w := p.Context.Value(ContextKeyWriter).(*http.ResponseWriter)

	if err := db.
		Preload("Roles").
		Where(User{Email: u.Email}).
		Attrs(User{Name: u.Name, ImageURI: u.ImageURI}).
		FirstOrCreate(&u).Error; err != nil {
		return err
	}

	val, err := json.Marshal(u)
	if err != nil {
		return errFailedToAuthenticate
	}

	// user will expire after 24 hours, same with token
	if err := inmem.Conn.Set(u.ID.String(), val, time.Hour*24).Err(); err != nil {
		return errFailedToAuthenticate
	}

	// If either token creation or redis store fails, consider login to have failed
	if err := createAndSendToken(w, u.ID); err != nil {
		return errFailedToAuthenticate
	}
	return nil
}

// GetAll returns a list of all users
func (u *Users) GetAll(p graphql.ResolveParams) error {
	db := database.Conn
	ctxUser := p.Context.Value(ContextKeyUser).(User)

	if !hasRole(ctxUser.Roles, "admin") {
		return errNotAuthorized
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

// GetByID gets user from database based on the users ID
func (u *User) GetByID(p graphql.ResolveParams) error {
	db := database.Conn
	ctxUser := p.Context.Value(ContextKeyUser).(User)

	if !hasRole(ctxUser.Roles, "admin") {
		return errNotAuthorized
	}

	// Find by uuid or email, which should both be unique
	if result := db.
		Preload("Roles").
		Where("id = ?", uuid.FromStringOrNil(p.Args["id"].(string))).
		Find(&u); result.Error != nil {
		return result.Error
	}

	return nil
}

// GetCurrent retrieves the current user directly from the context
// to avoid double data calls
func (u *User) GetCurrent(p graphql.ResolveParams) error {
	ctxUser := p.Context.Value(ContextKeyUser).(User)

	if uuid.Equal(ctxUser.ID, uuid.Nil) {
		return errNotAuthorized
	}

	*u = ctxUser
	return nil
}

// Create adds a new User to the database
// If it fails, returns a Failed to create error
func (u *User) Create(p graphql.ResolveParams, rs []Role) error {
	db := database.Conn
	ctxUser := p.Context.Value(ContextKeyUser).(User)

	if !hasRole(ctxUser.Roles, "admin") {
		return errNotAuthorized
	}

	ctxLogger.WithFields(logrus.Fields{
		"Email": u.Email,
		"Roles": rs,
	}).Debugln("Creating user with roles")

	if err := db.Create(&u).Model(&u).Association("Roles").Append(rs).Error; err != nil {
		return errFailedToCreate
	}

	return nil
}
