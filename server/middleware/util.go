package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/maxbrain0/react-go-graphql/server/models"
)

func userFromCookies(w *http.ResponseWriter, r *http.Request, rc *redis.Client) models.User {
	// get cookie containing header/payload + cookie containing signature

	// if error, we will not have auth data on requests and requests will fail
	// to clarify, and parsing problems will result in empty authorization info
	c1, err := r.Cookie("userinfo")
	if err != nil {
		return models.User{}
	}
	c2, err := r.Cookie("signature")
	if err != nil {
		return models.User{}
	}

	ts := c1.Value + "." + c2.Value

	token, err := jwt.ParseWithClaims(ts, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		ctxLogger.Debugln("Unable to parse jwt from string")
		return models.User{}
	}

	claims, ok := token.Claims.(*models.UserClaims)

	if !ok && !token.Valid {
		ctxLogger.Debugln("Invalid jwt")
		return models.User{}
	}

	val, err := rc.Get(claims.ID.String()).Result()

	if err != nil {
		ctxLogger.WithField("ID", claims.ID).Debugln("models.User not found in redis")
		return models.User{}
	}

	var u models.User
	err = json.Unmarshal([]byte(val), &u)

	if err != nil {
		ctxLogger.WithField("ID", claims.ID).Debugln("Could not deserialize user")
		return models.User{}
	}

	// If all succeeeds, extend cookie for half an hour

	c1.Expires = time.Now().Add(time.Minute * 30)
	http.SetCookie(*w, c1)

	return u
}
