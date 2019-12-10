package models

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/middleware"
	uuid "github.com/satori/go.uuid"
)

// User holds user information and role
type User struct {
	Base
	Name     string  `json:"name" gorm:"type:varchar(100);not null"`
	Email    string  `json:"email" gorm:"type:varchar(100);unique_index"`
	ImageURI string  `json:"imageUri" gorm:"type:text"`
	Roles    []*Role `json:"roles" gorm:"many2many:user_roles"`
}

type userClaims struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Roles []string  `json:"roles" gorm:"many2many:user_roles"`
	jwt.StandardClaims
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

	if err := createAndSendToken(p, u); err != nil {
		return err
	}

	return nil
}

func createAndSendToken(p graphql.ResolveParams, u *User) error {
	currentTime := time.Now()
	expiryTime := currentTime.Add(time.Minute * 60)

	// map []user.Role to []string with Role.Name only
	var roles []string
	for _, role := range u.Roles {
		roles = append(roles, role.Name)
	}

	// create and sign the token
	claims := userClaims{
		ID:    u.ID,
		Email: u.Email,
		Roles: roles,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "graphql.demo",
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: expiryTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return err
	}

	// Split the signed string into two parts for using the two cookie approach
	split := strings.Split(ss, ".")

	if len(split) != 3 {
		return fmt.Errorf("Unable to login User")
	}

	// send token to user httpOnlyCookie, secure if env is production
	w := middleware.GetWriter(p.Context)

	http.SetCookie(*w, &http.Cookie{
		Name:    "gqldemo_userinfo",
		Value:   split[0] + "." + split[1],
		Expires: expiryTime,
		Secure:  os.Getenv("APP_ENV") == "prod",
	})

	http.SetCookie(*w, &http.Cookie{
		Name:     "gqldemo_signature",
		Value:    split[2],
		HttpOnly: true,
		Secure:   os.Getenv("APP_ENV") == "prod",
	})

	return nil
}
