package models

import "github.com/sirupsen/logrus"

import "github.com/maxbrain0/react-go-graphql/server/database"

// Init assures tables for provided models are available and initialized a couple of users and roles
func Init() {
	db := database.Conn

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})

	var user1 User

	// Create Admin and Editor Roles
	// Can iterate over array or map if we need many roles in the future
	db.Where(AdminRole).FirstOrCreate(&AdminRole)

	ctxLogger.WithFields(logrus.Fields{
		"id":        AdminRole.ID,
		"Name":      AdminRole.Name,
		"UpdatedAt": AdminRole.UpdatedAt,
	}).Debugln("Created or found role")

	db.Where(EditorRole).FirstOrCreate(&EditorRole)

	ctxLogger.WithFields(logrus.Fields{
		"id":        EditorRole.ID,
		"Name":      EditorRole.Name,
		"UpdatedAt": EditorRole.UpdatedAt,
	}).Debugln("Created or found role")

	// Create users and append roles
	db.FirstOrCreate(&user1, User{
		Name:     "Jacob",
		Email:    "jacob.goodwin@gmail.com",
		ImageURI: "https://lh3.googleusercontent.com/a-/AAuE7mCsAHdorySC7ttxiSQOx7xtcUHhMwX6LlJwDT65LsE=s96-c",
	}).Model(&user1).Association("Roles").Append([]Role{AdminRole, EditorRole})

	ctxLogger.WithFields(logrus.Fields{
		"id":        user1.ID,
		"Name":      user1.Name,
		"UpdatedAt": user1.UpdatedAt,
		"Roles":     user1.Roles,
	}).Debugln("Created or found user")
}
