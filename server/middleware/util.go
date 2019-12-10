package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func userFromCookies(w *http.ResponseWriter, r *http.Request) UserInfo {
	// get cookie containing header/payload + cookie containing signature

	// if error, we will not have auth data on requests and requests will fail
	// to clarify, and parsing problems will result in empty authorization info
	c1, err := r.Cookie("userinfo")
	if err != nil {
		return UserInfo{}
	}
	c2, err := r.Cookie("signature")
	if err != nil {
		return UserInfo{}
	}

	ts := c1.Value + "." + c2.Value

	token, err := jwt.ParseWithClaims(ts, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		ctxLogger.Debugln("Unable to parse jwt from string")
		return UserInfo{}
	}

	claims, ok := token.Claims.(*UserClaims)

	if !ok && !token.Valid {
		ctxLogger.Debugln("Invalid jwt")
		return UserInfo{}
	}

	c1.Expires = time.Now().Add(time.Minute * 30)
	http.SetCookie(*w, c1)
	return claims.UserInfo
}

// CreateAndSendToken is a utility function for writing tokens
func CreateAndSendToken(w *http.ResponseWriter, u UserInfo) error {
	currentTime := time.Now()
	tokenExpiryTime := currentTime.Add(time.Hour * 24)
	cookieExpiryTime := currentTime.Add(time.Minute * 30)

	// create and sign the token
	claims := UserClaims{
		UserInfo: u,
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
