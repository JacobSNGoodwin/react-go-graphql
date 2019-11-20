package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/maxbrain0/react-go-graphql/server/gql"
	"github.com/maxbrain0/react-go-graphql/server/logger"
	"github.com/sirupsen/logrus"
)

var ctxLogger = logger.CtxLogger

func main() {
	// could receive a flag
	port := 8080

	//schema setup and serve
	// config of query and mutations setup
	schemaConfig := graphql.SchemaConfig{Query: gql.RootQuery}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", gql.HTTPMiddleware(h))

	ctxLogger.WithFields(logrus.Fields{
		"port": port,
	}).Info("Starting server")

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		ctxLogger.WithFields(logrus.Fields{
			"port": port,
		}).Fatal("Failed to serve application on given port")
	}
}
