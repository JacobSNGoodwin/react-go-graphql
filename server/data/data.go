package data

import "github.com/maxbrain0/react-go-graphql/server/logger"

// User holds information about a user
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Data holds some mock product data
type Data struct {
	Users []User
}

var startingData = []User{
	{
		ID:   1,
		Name: "Wax seal",
	},
	{
		ID:   2,
		Name: "Toilet Flap",
	},
	{
		ID:   3,
		Name: "Bidet Seat",
	},
}

var ctxLogger = logger.CtxLogger

// InitData initializes DataStruct with mock data
func (d *Data) InitData() {
	for _, val := range startingData {
		d.Users = append(d.Users, val)
	}

	ctxLogger.Debug("Products array has been filled")
}

// func (d *Data) GetUserByID(id int64) User {
// 	for _, user := range d.Users {
// 		if user.ID == id {
// 			return user
// 		}
// 	}

// }

// GetUsers retrieves all users from Data repository
func (d *Data) GetUsers() []User {
	return d.Users
}
