package models

import "github.com/sirupsen/logrus"

import "github.com/maxbrain0/react-go-graphql/server/database"

// RoleMap holds Role references to roles which helps gorm attach roles to user in graphql
var RoleMap map[string]*Role

// Init assures tables for provided models are available and initialized a couple of users and roles
func Init() {
	db := database.Conn

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})

	RoleMap = make(map[string]*Role)
	// create map of roles (for more easily getting role ref to use with GORM in gql mutations/queries)
	RoleMap["admin"] = &Role{
		Name: "admin",
	}
	RoleMap["editor"] = &Role{
		Name: "editor",
	}

	// Create Admin and Editor Roles
	// Can iterate over array or map if we need many roles in the future
	admin, _ := RoleMap["admin"]
	db.Where(*admin).FirstOrCreate(admin)
	ctxLogger.WithFields(logrus.Fields{
		"id":        admin.ID,
		"Name":      admin.Name,
		"UpdatedAt": admin.UpdatedAt,
	}).Debugln("Created or found role")

	editor, _ := RoleMap["editor"]
	db.Where(*editor).FirstOrCreate(editor)
	ctxLogger.WithFields(logrus.Fields{
		"id":        editor.ID,
		"Name":      editor.Name,
		"UpdatedAt": editor.UpdatedAt,
	}).Debugln("Created or found role")

	// Create users and append roles
	var user1 User
	db.FirstOrCreate(&user1, User{
		Name:     "Jacob",
		Email:    "jacob.goodwin@gmail.com",
		ImageURI: "https://lh3.googleusercontent.com/a-/AAuE7mCsAHdorySC7ttxiSQOx7xtcUHhMwX6LlJwDT65LsE=s96-c",
	}).Model(&user1).Association("Roles").Append([]Role{*RoleMap["admin"], *RoleMap["editor"]})

	ctxLogger.WithFields(logrus.Fields{
		"id":        user1.ID,
		"Name":      user1.Name,
		"UpdatedAt": user1.UpdatedAt,
		"Roles":     user1.Roles,
	}).Debugln("Created or found user")
}
