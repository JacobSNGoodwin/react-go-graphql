package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/maxbrain0/react-go-graphql/server/database"
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

	// setup db - variables can be conditionally set for env or with flags
	dbHost := "localhost"
	dbPort := 5432
	dbUser := "user"
	dbName := "gql_demo"
	dbSSLMode := "disable"

	var d = database.Database{
		Host:    dbHost,
		Port:    dbPort,
		User:    dbUser,
		Name:    dbName,
		SSLMode: dbSSLMode,
	}

	// connects to db given above parameters
	d.Connect()

	ctxLogger.WithFields(logrus.Fields{
		"host":   dbHost,
		"port":   dbPort,
		"dbname": dbName,
	}).Info("Connection to Postgres DB established")

	d.Init()
	defer d.DB.Close()

	// setup handler endpoint
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	// use middleware which gets request headers and injects db
	http.Handle("/graphql", gql.HTTPMiddleware(h))

	// run server in go func, and gracefully shut down server and database connection
	srv := &http.Server{
		Addr: fmt.Sprintf(":%v", port),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			ctxLogger.WithFields(logrus.Fields{
				"port": port,
			}).Fatal("Failed to serve application on given port")
		}
	}()

	ctxLogger.WithFields(logrus.Fields{
		"port": port,
	}).Info("Server successfully listening on port")

	<-done

	// disconnect postgres
	if err := d.DB.Close(); err != nil {
		ctxLogger.Fatalf("Failed to shut down databse %v", dbPort)
	}

	ctxLogger.Info("Successfully closed connection to postgres")

	// give 5 seconds to shutdown server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		ctxLogger.Fatalf("Failed to shut down server on %v", port)
	}

	ctxLogger.Info("Successfully shut down server")

}
