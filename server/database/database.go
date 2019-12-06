package database

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
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
func (d *Database) Init(e *casbin.Enforcer) {
	d.DB.AutoMigrate(&models.User{})

	// var admin models.Role
	// var editor models.Role
	var user1 models.User

	// Create users
	d.DB.FirstOrCreate(&user1, models.User{
		Name:     "Jacob",
		Email:    "jacob.goodwin@gmail.com",
		ImageURI: "",
	})

	// seems hwe have to do it this way for back ref
	// d.DB.Model(&user1).Association("Roles").Append([]models.Role{admin, editor})
	ctxLogger.WithFields(logrus.Fields{
		"id":        user1.ID,
		"Name":      user1.Name,
		"UpdatedAt": user1.UpdatedAt,
	}).Debug("Created or found user")

	e.AddRoleForUser(user1.Email, "Admin")
	e.AddRoleForUser(user1.Email, "Edit")

	e.AddPolicy("Admin", "path", "write")
	e.AddPolicy("Admin", "path", "read")

	e.AddPolicy("Edit", "path", "read")
}
