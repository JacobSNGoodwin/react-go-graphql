package gql

import (
	"context"
	"net/http"

	casbin "github.com/casbin/casbin/v2"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	"golang.org/x/oauth2"
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
	AUTH       map[string]*oauth2.Config
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
func GetHeader(ctx context.Context) (http.Header, bool) {
	header, ok := ctx.Value(contextKeyHeader).(http.Header)
	return header, ok
}

// GetDB retrieves gorm connection from context
func GetDB(ctx context.Context) (*gorm.DB, bool) {
	db, ok := ctx.Value(contextKeyDB).(*gorm.DB)
	return db, ok
}

// GetAuthConfigs loads a map with key strings to the oauth2 provider and values containing oauth2.config
func GetAuthConfigs(ctx context.Context) (map[string]*oauth2.Config, bool) {
	configs, ok := ctx.Value(contextKeyAuth).(map[string]*oauth2.Config)
	return configs, ok
}
