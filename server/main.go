package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	// setup db
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=user dbname=gql_demo password=password sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to create connection to postgres database: %v", err.Error())
	}
	ctxLogger.WithFields(logrus.Fields{
		"host":   "localhost",
		"port":   5432,
		"dbname": "gql-demo",
	}).Info("Connection to Postgres DB established")

	defer db.Close()

	// setup handler endpoint
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	// use middleware which gets request headers and injects db
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
