package data

import "github.com/maxbrain0/react-go-graphql/server/logger"

// User holds information about a user
type user struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Data holds some mock product data
type Data struct {
	Users []user
}

var startingData = []user{
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

// func (d *Data) GetUserByID(id int64) User {
// 	for _, user := range d.Users {
// 		if user.ID == id {
// 			return user
// 		}
// 	}

// }

// Place holder for when we have more complicated operations with db

// GetUsers retrieves all users from Data repository
// func (d *Data) GetUsers() []user {
// 	return d.users
// }
