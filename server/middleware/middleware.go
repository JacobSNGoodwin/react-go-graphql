package middleware

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	"github.com/maxbrain0/react-go-graphql/server/config"
	"github.com/maxbrain0/react-go-graphql/server/logger"
	uuid "github.com/satori/go.uuid"
)

type contextKey string

const contextKeyHeader = contextKey("header")
const contextKeyDB = contextKey("db")
const contextKeyAuth = contextKey("Auth")
const contextKeyAuthProviders = contextKey("AuthProviders")
const contextKeyWriter = contextKey("writer")

var ctxLogger = logger.CtxLogger

// Config holds references that will be accessed in middleware
type Config struct {
	GQLHandler *handler.Handler
	DB         *gorm.DB
	AUTH       *config.Auth
}

// UserInfo holds authorization info sent and received in jwt custom claims
type UserInfo struct {
	ID uuid.UUID `json:"id"`
}

// UserClaims used for creating and parsing JWTs
type UserClaims struct {
	UserInfo UserInfo `json:"userInfo"`
	jwt.StandardClaims
}

// HTTPMiddleware adds the request header to a graphql handler function
func HTTPMiddleware(c Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get cookies and reconstruct token - verify token and append authorization roles to
		// the req context
		userInfo := userFromCookies(&w, r)
		ctx := context.WithValue(r.Context(), contextKeyAuth, userInfo)

		// ctxLogger.WithField("Auth", claims.UserInfo).Debugln("UserInfo on context")

		// Configure DB and auth (for verifying tokens with google/fb)
		ctx = context.WithValue(ctx, contextKeyHeader, r.Header)
		ctx = context.WithValue(ctx, contextKeyDB, c.DB)
		ctx = context.WithValue(ctx, contextKeyAuthProviders, c.AUTH)

		// Configure response
		ctx = context.WithValue(ctx, contextKeyWriter, &w)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		c.GQLHandler.ContextHandler(ctx, w, r)
	})
}

// GetHeader returns the header as a strgin
func GetHeader(ctx context.Context) http.Header {
	return ctx.Value(contextKeyHeader).(http.Header)
}

// GetDB retrieves gorm connection from context
func GetDB(ctx context.Context) *gorm.DB {
	return ctx.Value(contextKeyDB).(*gorm.DB)
}

// GetAuth returns UserInfo as parsed from jwt
func GetAuth(ctx context.Context) UserInfo {
	return ctx.Value(contextKeyAuth).(UserInfo)
}

// GetAuthProviders loads a map with key strings to the oauth2 provider and values containing oauth2.config
func GetAuthProviders(ctx context.Context) *config.Auth {
	return ctx.Value(contextKeyAuthProviders).(*config.Auth)
}

// GetWriter retrieves the http.ResponseWriter from the current context
func GetWriter(ctx context.Context) *http.ResponseWriter {
	return ctx.Value(contextKeyWriter).(*http.ResponseWriter)
}
