package data

import (
	"github.com/maxbrain0/react-go-graphql/server/logger"
	"github.com/sirupsen/logrus"
)

// User holds information about a user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Data holds some mock product data
type Data struct {
	Users []User
}

var startingData = []User{
	{
		ID:   1,
		Name: "Bill123",
	},
	{
		ID:   2,
		Name: "John987",
	},
	{
		ID:   3,
		Name: "GuyWithTheHair",
	},
}

var ctxLogger = logger.CtxLogger

// Init initializes DataStruct with mock data
func (d *Data) Init() {
	for _, val := range startingData {
		d.Users = append(d.Users, val)
	}

	ctxLogger.Debug("Users array has been filled")
}

// GetUserByID does exactly what is says it does, bozo!
func (d *Data) GetUserByID(id int) (interface{}, error) {
	for _, user := range d.Users {
		if user.ID == id {
			ctxLogger.WithFields(logrus.Fields{
				"id": id,
			}).Debug("User found for requested id")
			return user, nil
		}
	}
	ctxLogger.WithFields(logrus.Fields{
		"id": id,
	}).Debug("User not found for requested id")

	return nil, nil
}

// Place holder for when we have more complicated operations with db

// GetUsers retrieves all users from Data repository
func (d *Data) GetUsers() []User {
	ctxLogger.Debug("Fetching all users")
	return d.Users
}
