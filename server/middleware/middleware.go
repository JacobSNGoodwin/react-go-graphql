package middleware

import (
	"context"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/maxbrain0/react-go-graphql/server/logger"
)

var ctxLogger = logger.CtxLogger

type contextKey string

const contextKeyHeader = contextKey("header")

// HTTPMiddleware adds the request header to a graphql handler function
func HTTPMiddleware(next *handler.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), contextKeyHeader, r.Header)
		next.ContextHandler(ctx, w, r)
	})
}

// GetHeader returns the header as a strgin
func GetHeader(ctx context.Context) (http.Header, bool) {
	header, ok := ctx.Value(contextKeyHeader).(http.Header)
	return header, ok
}
