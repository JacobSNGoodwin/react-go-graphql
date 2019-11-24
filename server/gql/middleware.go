package gql

import (
	"context"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
)

type contextKey string

const contextKeyHeader = contextKey("header")
const contextKeyDB = contextKey("db")

// MiddlewareConfig holds references that will be accessed in middleware
type MiddlewareConfig struct {
	GQLHandler *handler.Handler
	DB         *gorm.DB
	E          *casbin.Enforcer
}

// HTTPMiddleware adds the request header to a graphql handler function
func HTTPMiddleware(c MiddlewareConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c1 := context.WithValue(r.Context(), contextKeyHeader, r.Header)
		ctx := context.WithValue(c1, contextKeyDB, c.DB)
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
