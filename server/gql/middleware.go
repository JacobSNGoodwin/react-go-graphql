package gql

import (
	"context"
	"net/http"

	casbin "github.com/casbin/casbin/v2"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	"github.com/maxbrain0/react-go-graphql/server/config"
)

type contextKey string

const contextKeyHeader = contextKey("header")
const contextKeyDB = contextKey("db")
const contextKeyAuth = contextKey("auth")

// MiddlewareConfig holds references that will be accessed in middleware
type MiddlewareConfig struct {
	GQLHandler *handler.Handler
	DB         *gorm.DB
	E          *casbin.Enforcer
	AUTH       config.Auth
}

// HTTPMiddleware adds the request header to a graphql handler function
func HTTPMiddleware(c MiddlewareConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), contextKeyHeader, r.Header)
		ctx = context.WithValue(ctx, contextKeyDB, c.DB)
		ctx = context.WithValue(ctx, contextKeyAuth, c.AUTH)
		c.GQLHandler.ContextHandler(ctx, w, r)
	})
}

// GetHeader returns the header as a strgin
func GetHeader(ctx context.Context) http.Header {
	header, ok := ctx.Value(contextKeyHeader).(http.Header)

	if !ok {
		ctxLogger.Warningln("Unable to get Header key in HTTPMiddleware")
	}

	return header
}

// GetDB retrieves gorm connection from context
func GetDB(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(contextKeyDB).(*gorm.DB)

	if !ok {
		ctxLogger.Warningln("Unable to get DB key in HTTPMiddleware")
	}
	return db
}

// GetAuth loads a map with key strings to the oauth2 provider and values containing oauth2.config
func GetAuth(ctx context.Context) config.Auth {
	config, ok := ctx.Value(contextKeyAuth).(config.Auth)

	if !ok {
		ctxLogger.Warningln("Unable to get Auth key in HTTPMiddleware")
	}
	return config
}
