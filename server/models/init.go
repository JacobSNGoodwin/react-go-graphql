package models

import "github.com/sirupsen/logrus"

import "github.com/maxbrain0/react-go-graphql/server/database"

// Init assures tables for provided models are available and initialized a couple of users and roles
func Init() {
	db := database.Conn

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})

	var admin Role
	var editor Role
	var user1 User

	// create two roles for first user
	db.FirstOrCreate(&admin, Role{
		Name: AdminRole,
	})

	ctxLogger.WithFields(logrus.Fields{
		"id":        admin.ID,
		"Name":      admin.Name,
		"UpdatedAt": admin.UpdatedAt,
	}).Debugln("Created or found role")

	db.FirstOrCreate(&editor, Role{
		Name: EditorRole,
	})

	ctxLogger.WithFields(logrus.Fields{
		"id":        editor.ID,
		"Name":      editor.Name,
		"UpdatedAt": editor.UpdatedAt,
	}).Debugln("Created or found role")

	// Create users
	db.FirstOrCreate(&user1, User{
		Name:     "Jacob",
		Email:    "jacob.goodwin@gmail.com",
		ImageURI: "https://lh3.googleusercontent.com/a-/AAuE7mCsAHdorySC7ttxiSQOx7xtcUHhMwX6LlJwDT65LsE=s96-c",
	})

	// seems hwe have to do it this way for back ref
	db.Model(&user1).Association("Roles").Append([]Role{admin, editor})

	ctxLogger.WithFields(logrus.Fields{
		"id":        user1.ID,
		"Name":      user1.Name,
		"UpdatedAt": user1.UpdatedAt,
		"Roles":     user1.Roles,
	}).Debugln("Created or found user")
}
