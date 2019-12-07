package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	casbin "github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/maxbrain0/react-go-graphql/server/config"
	"github.com/maxbrain0/react-go-graphql/server/database"
	"github.com/maxbrain0/react-go-graphql/server/gql"
	"github.com/maxbrain0/react-go-graphql/server/logger"
	"github.com/maxbrain0/react-go-graphql/server/middleware"
	"github.com/sirupsen/logrus"
)

var ctxLogger = logger.CtxLogger

func main() {
	// load env variables from .env file
	// need check to run coad only in DEV mode
	envPath, err := filepath.Abs("./config/.env")
	if err != nil {
		ctxLogger.Fatal("Failed to load development env file")
	}

	err = godotenv.Load(envPath)

	if err != nil {
		ctxLogger.Fatal("Error loading .env file")
	}

	//
	port := os.Getenv("SERVER_PORT")

	//schema setup and serve
	// config of query and mutations setup
	schemaConfig := graphql.SchemaConfig{Query: gql.RootQuery, Mutation: gql.RootMutation}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// setup db - variables can be conditionally set for env or with flags
	dbHost := os.Getenv("PG_HOST")
	dbPort := os.Getenv("PG_PORT")
	dbUser := os.Getenv("PG_USER")
	dbName := os.Getenv("PG_DB_NAME")
	dbSSLMode := os.Getenv("PG_SSL_MODE")

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

	a, err := gormadapter.NewAdapterByDB(d.DB)

	if err != nil {
		ctxLogger.Fatalf("Unable to connect gorm adapter: %v", err.Error())
	}

	e, err := casbin.NewEnforcer("config/rbac_model.conf", a)

	if err != nil {
		ctxLogger.Fatalf("Unable to setup casbin config: %v", err.Error())
	}

	d.Init(e)
	defer d.DB.Close()

	// setup auth config for login queries and mutations
	authConfig := &config.Auth{}

	authConfig.Load()

	// setup handler endpoint
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	// use middleware which gets request headers and injects db
	http.Handle("/graphql", middleware.HTTPMiddleware(middleware.Config{
		GQLHandler: h,
		DB:         d.DB,
		E:          e,
		AUTH:       authConfig,
	}))

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
