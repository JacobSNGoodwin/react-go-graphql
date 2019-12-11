package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/maxbrain0/react-go-graphql/server/logger"
	"github.com/maxbrain0/react-go-graphql/server/models"
	"github.com/sirupsen/logrus"
)

// Database holds connection string info for that database
type Database struct {
	Host    string
	Port    string
	User    string
	Name    string
	SSLMode string
	DB      *gorm.DB
}

var ctxLogger = logger.CtxLogger

// Connect returns the DB connection for the connection info in the struct
func (d *Database) Connect() {
	connStr := "host=%s port=%s user=%s dbname=%s sslmode=%s"
	connStr = fmt.Sprintf(connStr, d.Host, d.Port, d.User, d.Name, d.SSLMode)

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=user dbname=gql_demo password=password sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to create connection to postgres database: %v", err.Error())
	}

	d.DB = db
}

// Init assures tables for provided models are available and initialized a couple of users and roles
func (d *Database) Init() {
	d.DB.AutoMigrate(&models.User{})
	d.DB.AutoMigrate(&models.Role{})

	var admin models.Role
	var editor models.Role
	var user1 models.User

	// create two roles for first user
	d.DB.FirstOrCreate(&admin, models.Role{
		Name: models.AdminRole,
	})

	ctxLogger.WithFields(logrus.Fields{
		"id":        admin.ID,
		"Name":      admin.Name,
		"UpdatedAt": admin.UpdatedAt,
	}).Debugln("Created or found role")

	d.DB.FirstOrCreate(&editor, models.Role{
		Name: models.EditorRole,
	})

	ctxLogger.WithFields(logrus.Fields{
		"id":        editor.ID,
		"Name":      editor.Name,
		"UpdatedAt": editor.UpdatedAt,
	}).Debugln("Created or found role")

	// Create users
	d.DB.FirstOrCreate(&user1, models.User{
		Name:     "Jacob",
		Email:    "jacob.goodwin@gmail.com",
		ImageURI: "https://lh3.googleusercontent.com/a-/AAuE7mCsAHdorySC7ttxiSQOx7xtcUHhMwX6LlJwDT65LsE=s96-c",
	})

	// seems hwe have to do it this way for back ref
	d.DB.Model(&user1).Association("Roles").Append([]models.Role{admin, editor})

	ctxLogger.WithFields(logrus.Fields{
		"id":        user1.ID,
		"Name":      user1.Name,
		"UpdatedAt": user1.UpdatedAt,
		"Roles":     user1.Roles,
	}).Debugln("Created or found user")
}
