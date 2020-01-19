package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/database"
	"github.com/maxbrain0/react-go-graphql/server/errors"
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

// Login takes the current user and logs them in if they exist.
// It creates the user if the user doesn't yet exist
func (u *User) Login(p graphql.ResolveParams) error {
	// Add error checking
	db := database.Conn
	w := p.Context.Value(ContextKeyWriter).(*http.ResponseWriter)

	if err := db.
		Preload("Roles").
		Where(User{Email: u.Email}).
		Attrs(User{Name: u.Name, ImageURI: u.ImageURI}).
		First(&u).Error; err != nil {
		return errors.NewInternal("User not found", err)
	}

	val, err := json.Marshal(u)
	if err != nil {
		return errors.NewInternal("Internal error logging in user", err)
	}

	// user will expire after 24 hours, same with token
	if err := inmem.Conn.Set(u.ID.String(), val, time.Hour*24).Err(); err != nil {
		return errors.NewInternal("Internal error logging user", err)
	}

	// If either token creation or redis store fails, consider login to have failed
	if err := createAndSendToken(w, u.ID); err != nil {
		return errors.NewInternal("Internal error logging in user", err)
	}
	return nil
}

// GetAll returns a list of all users
func (u *Users) GetAll(p graphql.ResolveParams) error {
	db := database.Conn
	ctxUser := p.Context.Value(ContextKeyUser).(User)

	if !hasRole(ctxUser.Roles, "admin") {
		return errors.NewForbidden("Not authorized", nil)
	}

	ctxLogger.Infoln("GetAll Users")

	if result :=
		db.
			Order("email").
			Limit(p.Args["limit"].(int)).
			Offset(p.Args["offset"].(int)).
			Preload("Roles").
			Find(&u); result.Error != nil {
		return errors.NewInternal("Error fetching users", nil)
	}

	return nil
}

// GetByID gets user from database based on the users ID
func (u *User) GetByID(p graphql.ResolveParams) error {
	db := database.Conn
	ctxUser := p.Context.Value(ContextKeyUser).(User)

	if !hasRole(ctxUser.Roles, "admin") {
		return errors.NewForbidden("Not authorized", nil)
	}

	ctxLogger.WithField("id", p.Args["id"].(string)).Infoln("GetByID Users")
	// Find by uuid or email, which should both be unique
	if err := db.
		Preload("Roles").
		Where("id = ?", uuid.FromStringOrNil(p.Args["id"].(string))).
		Find(&u).Error; err != nil {
		ctxLogger.WithError(err).Debugln("DB Error finding user by ID")
		return errors.NewInternal("Error finding user", nil)
	}

	return nil
}

// GetCurrent retrieves the current user directly from the context
// to avoid double data calls
func (u *User) GetCurrent(p graphql.ResolveParams) error {
	ctxUser := p.Context.Value(ContextKeyUser).(User)

	ctxLogger.Infoln("GetCurrent User")
	if uuid.Equal(ctxUser.ID, uuid.Nil) {
		return errors.NewAuthentication("User is not logged in", nil)
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
		return errors.NewForbidden("Not authorized", nil)
	}

	ctxLogger.WithFields(logrus.Fields{
		"Email": u.Email,
		"Roles": rs,
	}).Infoln("Creating user with roles")

	if err := db.Create(&u).Model(&u).Association("Roles").Append(rs).Error; err != nil {
		ctxLogger.WithError(err).Debugln("DB Error creating user")
		return errors.NewInternal("Error creating user", nil)
	}

	return nil
}

// Update attempts to update an existing user
func (u *User) Update(p graphql.ResolveParams, updates map[string]interface{}, updateRoles bool, rs []Role) error {
	db := database.Conn

	ctxUser := p.Context.Value(ContextKeyUser).(User)

	if !hasRole(ctxUser.Roles, "admin") {
		return errors.NewForbidden("Not authorized", nil)
	}

	ctxLogger.WithField("id", u.ID).Infoln("Update User: %v")

	if err := db.First(&u).Error; err != nil {
		ctxLogger.WithError(err).Debugln("DB Error updating user")
		return errors.NewInternal("Error updating user", nil)
	}

	if updateRoles {
		// need to clear redis cache since admin has updated user's role. That user will need to
		// login again
		db.Model(&u).Updates(updates).Association("Roles").Replace(rs)
		inmem.Conn.Del(u.ID.String())
	} else {
		db.Model(&u).Updates(updates)
	}

	return nil
}

// Delete removes user with u.ID
func (u *User) Delete(p graphql.ResolveParams) error {
	db := database.Conn
	ctxUser := p.Context.Value(ContextKeyUser).(User)

	if !hasRole(ctxUser.Roles, "admin") {
		return errors.NewForbidden("Not authorized", nil)
	}

	ctxLogger.WithField("id", u.ID).Infoln("Delete User")

	// create a transacation to make sure both roles are unassociated and user is deleted
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&u).Association("Roles").Clear().Error; err != nil {
		ctxLogger.WithError(err).Debugln("DB Error deleting user")
		tx.Rollback()
		errors.NewInternal("Error deleting user", nil)
	}

	if err := tx.Delete(&u).Error; err != nil {
		ctxLogger.WithError(err).Debugln("DB Error deleting user")
		tx.Rollback()
		return errors.NewInternal("Error deleting user", nil)
	}

	if err := tx.Commit().Error; err != nil {
		return errors.NewInternal("Error deleting user", nil)
	}

	return nil
}
