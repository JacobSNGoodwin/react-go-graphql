package models

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/maxbrain0/react-go-graphql/server/logger"
	uuid "github.com/satori/go.uuid"
)

// Types and vars accessible to models in general

type contextKey string

// ContextKeyUser used as a type for setting and acccessing User on the http/graphql context
const ContextKeyUser = contextKey("User")

// ContextKeyWriter used to access writer in http pipeline (for setting cookies outside of middleware)
const ContextKeyWriter = contextKey("writer")

var (
	errFailedToAuthenticate = fmt.Errorf("Failed to authenticate user")
	errFailedToCreate       = fmt.Errorf("Failed to create resource")
	errNotAuthorized        = fmt.Errorf("Not authorized")
	errNotFound             = fmt.Errorf("Not found")
)

var ctxLogger = logger.CtxLogger

// utility functions for models

// createAndSendToken is a utility function for writing tokens
func createAndSendToken(w *http.ResponseWriter, id uuid.UUID) error {
	currentTime := time.Now()
	tokenExpiryTime := currentTime.Add(time.Hour * 24)
	cookieExpiryTime := currentTime.Add(time.Minute * 30)

	// create and sign the token
	claims := UserClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "graphql.demo",
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: tokenExpiryTime.Unix(),
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
	http.SetCookie(*w, &http.Cookie{
		Name:    "userinfo",
		Value:   split[0] + "." + split[1],
		Expires: cookieExpiryTime,
		Secure:  os.Getenv("APP_ENV") == "prod",
	})

	http.SetCookie(*w, &http.Cookie{
		Name:     "signature",
		Value:    split[2],
		HttpOnly: true,
		Secure:   os.Getenv("APP_ENV") == "prod",
	})

	return nil
}

func hasRole(rs []*Role, r string) bool {
	// possible that val.Name is null? Should be populated with 0 value, at least
	for _, val := range rs {
		if val.Name == r {
			return true
		}
	}
	return false
}
