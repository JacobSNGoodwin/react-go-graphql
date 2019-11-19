package gql

import (
	"github.com/graphql-go/graphql"
)

// UserType holds information for users
var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "A user with id and name",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
