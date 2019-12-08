package middleware

import (
	"context"
	"net/http"

	casbin "github.com/casbin/casbin/v2"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	"github.com/maxbrain0/react-go-graphql/server/config"
	"github.com/maxbrain0/react-go-graphql/server/logger"
)

type contextKey string

const contextKeyHeader = contextKey("header")
const contextKeyDB = contextKey("db")
const contextKeyAuth = contextKey("auth")
const contextKeyWriter = contextKey("writer")

var ctxLogger = logger.CtxLogger

// Config holds references that will be accessed in middleware
type Config struct {
	GQLHandler *handler.Handler
	DB         *gorm.DB
	E          *casbin.Enforcer
	AUTH       *config.Auth
}

// HTTPMiddleware adds the request header to a graphql handler function
func HTTPMiddleware(c Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), contextKeyHeader, r.Header)
		ctx = context.WithValue(ctx, contextKeyDB, c.DB)
		ctx = context.WithValue(ctx, contextKeyAuth, c.AUTH)

		wRef := &w

		ctxLogger.Debugf("The value of wRef passed into context: %v\n", wRef)
		ctx = context.WithValue(ctx, contextKeyWriter, wRef)
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

// GetAuth loads a map with key strings to the oauth2 provider and values containing oauth2.config
func GetAuth(ctx context.Context) *config.Auth {
	return ctx.Value(contextKeyAuth).(*config.Auth)
}

// GetWriter retrieves the http.ResponseWriter from the current context
func GetWriter(ctx context.Context) *http.ResponseWriter {
	return ctx.Value(contextKeyWriter).(*http.ResponseWriter)
}
