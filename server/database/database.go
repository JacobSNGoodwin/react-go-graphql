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
	Port    int
	User    string
	Name    string
	SSLMode string
	DB      *gorm.DB
}

var ctxLogger = logger.CtxLogger

// Connect returns the DB connection for the connection info in the struct
func (d *Database) Connect() {
	connStr := "host=%s port=%d user=%s dbname=%s sslmode=%s"
	connStr = fmt.Sprintf(connStr, d.Host, d.Port, d.User, d.Name, d.SSLMode)

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=user dbname=gql_demo password=password sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to create connection to postgres database: %v", err.Error())
	}

	d.DB = db
}

// Init assures tables for provided models are available and initialized a couple of users and roles
func (d *Database) Init() {
	d.DB.AutoMigrate(&models.User{}, &models.Role{})

	var admin models.Role
	var editor models.Role
	var user1 models.User
	var user2 models.User

	// Create roles first, then these role ids can be assigned to users
	d.DB.Where(models.Role{Name: "Admin"}).FirstOrCreate(&admin)
	ctxLogger.WithFields(logrus.Fields{
		"id":        admin.ID,
		"Name":      admin.Name,
		"UpdatedAt": admin.UpdatedAt,
	}).Debug("Created or found role")

	d.DB.Where(models.Role{Name: "Editor"}).FirstOrCreate(&editor)
	ctxLogger.WithFields(logrus.Fields{
		"id":        editor.ID,
		"Name":      editor.Name,
		"UpdatedAt": editor.UpdatedAt,
	}).Debug("Created or found role")

	// Create users
	d.DB.FirstOrCreate(&user1, models.User{
		Name:  "Jacob",
		Email: "jacob.test.com",
	})

	// seems hwe have to do it this way for back ref
	d.DB.Model(&user1).Association("Roles").Append([]models.Role{admin, editor})
	ctxLogger.WithFields(logrus.Fields{
		"id":        user1.ID,
		"Name":      user1.Name,
		"UpdatedAt": user1.UpdatedAt,
	}).Debug("Created or found user")

	d.DB.FirstOrCreate(&user2, models.User{
		Name:  "Thea",
		Email: "thea.test.com",
	})

	// seems hwe have to do it this way for back ref
	d.DB.Model(&user2).Association("Roles").Append([]models.Role{editor})
	ctxLogger.WithFields(logrus.Fields{
		"id":        user2.ID,
		"Name":      user2.Name,
		"UpdatedAt": user2.UpdatedAt,
	}).Debug("Created or found user")

	// check if we can get back ref (users for given role)
	// var editorUsers []models.User
	// d.DB.Model(&editor).Related(&editorUsers, "Users")

	// for _, user := range editorUsers {
	// 	ctxLogger.WithFields(logrus.Fields{
	// 		"id":        user.ID,
	// 		"Name":      user.Name,
	// 		"UpdatedAt": user.UpdatedAt,
	// 	}).Debug("Found user in editor role")
	// }
}
